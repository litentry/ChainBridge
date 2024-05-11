// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"errors"
	"time"

	"github.com/Phala-Network/chainbridge-utils/msg"
	utils "github.com/litentry/ChainBridge/shared/ethereum"
)

// Number of blocks to wait for an finalization event
const ExecuteBlockWatchLimit = 9600

// Time between retrying a failed tx
const TxRetryInterval = time.Second * 2

// Maximum number of tx retries before exiting
const TxRetryLimit = 10

var ErrNonceTooLow = errors.New("nonce too low")
var ErrTxUnderpriced = errors.New("replacement transaction underpriced")
var ErrFatalTx = errors.New("submission of transaction failed")
var ErrFatalQuery = errors.New("query of chain state failed")

// proposalIsComplete returns true if the proposal state is either Passed, Transferred or Cancelled
func (w *writer) proposalIsComplete(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	prop, err := w.bridgeContract.GetProposal(w.conn.CallOpts(), uint8(srcId), uint64(nonce), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop.Status == PassedStatus || prop.Status == TransferredStatus || prop.Status == CancelledStatus
}

// proposalIsComplete returns true if the proposal state is Transferred or Cancelled
func (w *writer) proposalIsFinalized(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	prop, err := w.bridgeContract.GetProposal(w.conn.CallOpts(), uint8(srcId), uint64(nonce), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop.Status == TransferredStatus || prop.Status == CancelledStatus // Transferred (3)
}

func (w *writer) proposalIsPassed(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	prop, err := w.bridgeContract.GetProposal(w.conn.CallOpts(), uint8(srcId), uint64(nonce), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop.Status == PassedStatus
}

// hasVoted checks if this relayer has already voted
func (w *writer) hasVoted(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	hasVoted, err := w.bridgeContract.HasVotedOnProposal(w.conn.CallOpts(), utils.IDAndNonce(srcId, nonce), dataHash, w.conn.Opts().From)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}

	return hasVoted
}

func (w *writer) shouldVote(m msg.Message, dataHash [32]byte) bool {
	// Check if proposal has passed and skip if Passed or Transferred
	if w.proposalIsComplete(m.Source, m.DepositNonce, dataHash) {
		w.log.Info("Proposal complete, not voting", "src", m.Source, "nonce", m.DepositNonce)
		return false
	}

	// Check if relayer has previously voted
	if w.hasVoted(m.Source, m.DepositNonce, dataHash) {
		w.log.Info("Relayer has already voted, not voting", "src", m.Source, "nonce", m.DepositNonce)
		return false
	}

	return true
}

// createErc20Proposal creates an Erc20 proposal.
// Returns true if the proposal is successfully created or is complete or at least try executing done
func (w *writer) createErc20Proposal(m msg.Message) bool {
	w.log.Info("Creating erc20 proposal", "src", m.Source, "nonce", m.DepositNonce)

	data := ConstructErc20ProposalData(m.Payload[0].([]byte), m.Payload[1].([]byte))
	dataHash := utils.Hash(append(w.cfg.erc20HandlerContract.Bytes(), data...))

	if !w.shouldVote(m, dataHash) {
		if w.proposalIsPassed(m.Source, m.DepositNonce, dataHash) {
			// We should not vote for this proposal but it is ready to be executed
			w.executeProposal(m, data, dataHash)
			return true
		} else {
			return false
		}
	}

	w.voteProposal(m, dataHash, data)

	return true
}

// createErc721Proposal creates an Erc721 proposal.
// Returns true if the proposal is succesfully created or is complete or at least try executing done
func (w *writer) createErc721Proposal(m msg.Message) bool {
	w.log.Info("Creating erc721 proposal", "src", m.Source, "nonce", m.DepositNonce)

	data := ConstructErc721ProposalData(m.Payload[0].([]byte), m.Payload[1].([]byte), m.Payload[2].([]byte))
	dataHash := utils.Hash(append(w.cfg.erc721HandlerContract.Bytes(), data...))

	if !w.shouldVote(m, dataHash) {
		if w.proposalIsPassed(m.Source, m.DepositNonce, dataHash) {
			// We should not vote for this proposal but it is ready to be executed
			w.executeProposal(m, data, dataHash)
			return true
		} else {
			return false
		}
	}

	w.voteProposal(m, dataHash, data)

	return true
}

// createGenericDepositProposal creates a generic proposal
// returns true if the proposal is complete or is succesfully created
func (w *writer) createGenericDepositProposal(m msg.Message) bool {
	w.log.Info("Creating generic proposal", "src", m.Source, "nonce", m.DepositNonce)

	metadata := m.Payload[0].([]byte)
	data := ConstructGenericProposalData(metadata)
	toHash := append(w.cfg.genericHandlerContract.Bytes(), data...)
	dataHash := utils.Hash(toHash)

	if !w.shouldVote(m, dataHash) {
		if w.proposalIsPassed(m.Source, m.DepositNonce, dataHash) {
			// We should not vote for this proposal but it is ready to be executed
			w.executeProposal(m, data, dataHash)
			return true
		} else {
			return false
		}
	}

	w.voteProposal(m, dataHash, data)

	return true
}

// voteProposal submits a vote proposal
// a vote proposal will try to be submitted up to the TxRetryLimit times
func (w *writer) voteProposal(m msg.Message, dataHash [32]byte, data []byte) {
	for i := 0; i < TxRetryLimit; i++ {
		select {
		case <-w.stop:
			return
		default:
			err := w.conn.LockAndUpdateOpts()
			if err != nil {
				w.log.Error("Failed to update tx opts", "err", err)
				continue
			}

			tx, err := w.bridgeContract.VoteProposal(
				w.conn.Opts(),
				uint8(m.Source),
				uint64(m.DepositNonce),
				m.ResourceId,
				data,
			)
			w.conn.UnlockOpts()

			if err == nil {
				w.log.Info("Submitted proposal vote", "tx", tx.Hash(), "src", m.Source, "depositNonce", m.DepositNonce)
				if w.metrics != nil {
					w.metrics.VotesSubmitted.Inc()
				}
				return
			} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
				w.log.Debug("Nonce too low, will retry")
				time.Sleep(TxRetryInterval)
			} else {
				w.log.Warn("Voting failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce, "err", err)
				time.Sleep(TxRetryInterval)
			}

			// Verify proposal is still open for voting, otherwise no need to retry
			if w.proposalIsComplete(m.Source, m.DepositNonce, dataHash) {
				w.log.Info("Proposal voting complete on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
				return
			}
		}
	}
	w.log.Error("Submission of Vote transaction failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce)
	w.sysErr <- ErrFatalTx
}

// executeProposal executes the proposal
func (w *writer) executeProposal(m msg.Message, data []byte, dataHash [32]byte) {
	for i := 0; i < TxRetryLimit; i++ {
		select {
		case <-w.stop:
			return
		default:
			err := w.conn.LockAndUpdateOpts()
			if err != nil {
				w.log.Error("Failed to update nonce", "err", err)
				return
			}

			tx, err := w.bridgeContract.ExecuteProposal(
				w.conn.Opts(),
				uint8(m.Source),
				uint64(m.DepositNonce),
				data,
				m.ResourceId,
			)
			w.conn.UnlockOpts()

			if err == nil {
				w.log.Info("Submitted proposal execution", "tx", tx.Hash(), "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
				return
			} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
				w.log.Error("Nonce too low, will retry")
				time.Sleep(TxRetryInterval)
			} else {
				w.log.Warn("Execution failed, proposal may already be complete", "err", err)
				time.Sleep(TxRetryInterval)
			}

			// Verify proposal is still open for execution, tx will fail if we aren't the first to execute,
			// but there is no need to retry
			if w.proposalIsFinalized(m.Source, m.DepositNonce, dataHash) {
				w.log.Info("Proposal finalized on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
				return
			}
		}
	}
	w.log.Error("Submission of Execute transaction failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce)
	w.sysErr <- ErrFatalTx
}
