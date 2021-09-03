// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	events "github.com/Phala-Network/chainbridge-substrate-events"
	"github.com/Phala-Network/go-substrate-rpc-client/v3/types"
)

type EventErc721Minted struct {
	Phase   types.Phase
	Owner   types.AccountID
	TokenId types.U256
	Topics  []types.Hash
}

type EventErc721Transferred struct {
	Phase   types.Phase
	From    types.AccountID
	To      types.AccountID
	TokenId types.U256
	Topics  []types.Hash
}

type EventErc721Burned struct {
	Phase   types.Phase
	TokenId types.AccountID
	Topics  []types.Hash
}

type EventExampleRemark struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

// EventNFTDeposited is emitted when NFT is ready to be deposited to other chain.
type EventNFTDeposited struct {
	Phase  types.Phase
	Asset  types.Hash
	Topics []types.Hash
}

// EventFeeChanged is emitted when a fee for a given key is changed.
type EventFeeChanged struct {
	Phase    types.Phase
	Key      types.Hash
	NewPrice types.U128
	Topics   []types.Hash
}

// EventNewMultiAccount is emitted when a multi account has been created.
// First param is the account that created it, second is the multisig account.
type EventNewMultiAccount struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// EventMultiAccountUpdated is emitted when a multi account has been updated. First param is the multisig account.
type EventMultiAccountUpdated struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventMultiAccountRemoved is emitted when a multi account has been removed. First param is the multisig account.
type EventMultiAccountRemoved struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventNewMultisig is emitted when a new multisig operation has begun.
// First param is the account that is approving, second is the multisig account.
type EventNewMultisig struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// TimePoint contains height and index
type TimePoint struct {
	Height types.U32
	Index  types.U32
}

// EventMultisigApproval is emitted when a multisig operation has been approved by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigApproval struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

// EventMultisigExecuted is emitted when a multisig operation has been executed by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigExecuted struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Result    types.DispatchResult
	Topics    []types.Hash
}

// EventMultisigCancelled is emitted when a multisig operation has been cancelled by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigCancelled struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

type EventTreasuryMinting struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventRadClaimsClaimed is emitted when RAD Tokens have been claimed
type EventRadClaimsClaimed struct {
	Phase  types.Phase
	Who    types.AccountID
	Value  types.U128
	Topics []types.Hash
}

// EventRadClaimsRootHashStored is emitted when RootHash has been stored for the correspondent RAD Claims batch
type EventRadClaimsRootHashStored struct {
	Phase    types.Phase
	RootHash types.Hash
	Topics   []types.Hash
}

// EventNftTransferred is emitted when the ownership of the asset has been transferred to the account
type EventNftTransferred struct {
	Phase      types.Phase
	RegistryId RegistryId
	AssetId    AssetId
	Who        types.AccountID
	Topics     []types.Hash
}

// EventRegistryMint is emitted when successfully minting an NFT
type EventRegistryMint struct {
	Phase      types.Phase
	RegistryId RegistryId
	TokenId    TokenId
	Topics     []types.Hash
}

// EventRegistryRegistryCreated is emitted when successfully creating a NFT registry
type EventRegistryRegistryCreated struct {
	Phase      types.Phase
	RegistryId RegistryId
	Topics     []types.Hash
}

// EventRegistryTmp is emitted only for testing
type EventRegistryTmp struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

type EventPhragmenElectionNewTerm struct {
	Phase types.Phase
	Term  []struct {
		Account types.AccountID
		Balance types.U128
	}
	Topics []types.Hash
}

type EventPhragmenElectionEmptyTerm struct {
	Phase  types.Phase
	Topics []types.Hash
}

type EventPhragmenElectionElectionError struct {
	Phase  types.Phase
	Topics []types.Hash
}

type EventPhragmenElectionMemberKicked struct {
	Phase  types.Phase
	Member types.AccountID
	Topics []types.Hash
}

type EventPhragmenElectionRenounced struct {
	Phase  types.Phase
	Member types.AccountID
	Topics []types.Hash
}

type EventPhragmenElectionCandidateSlashed struct {
	Phase  types.Phase
	Member types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventPhragmenElectionSeatHolderSlashed struct {
	Phase  types.Phase
	Member types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type Events struct {
	types.EventRecords
	events.ChainBridgeEvents
	events.BridgeTransferEvents
	events.PhalaGateKeeperEvents
	events.PhalaMiningEvents
	events.PhalaStakepoolEvents
	events.KittiesEvents
	PhragmenElection_NewTerm           []EventPhragmenElectionNewTerm           //nolint:stylecheck,golint
	PhragmenElection_EmptyTerm         []EventPhragmenElectionEmptyTerm         //nolint:stylecheck,golint
	PhragmenElection_ElectionError     []EventPhragmenElectionElectionError     //nolint:stylecheck,golint
	PhragmenElection_MemberKicked      []EventPhragmenElectionMemberKicked      //nolint:stylecheck,golint
	PhragmenElection_Renounced         []EventPhragmenElectionRenounced         //nolint:stylecheck,golint
	PhragmenElection_CandidateSlashed  []EventPhragmenElectionCandidateSlashed  //nolint:stylecheck,golint
	PhragmenElection_SeatHolderSlashed []EventPhragmenElectionSeatHolderSlashed //nolint:stylecheck,golint
}
