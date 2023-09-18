package keeper

import (
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/gov/types"
)

// Implements GovHooks interface
var _ types.GovHooks = Keeper{}

// AfterProposalSubmission - call hook if registered
func (keeper Keeper) AfterProposalSubmission(ctx sdk.Context, proposalID uint64) {
	if keeper.hooks != nil {
		keeper.hooks.AfterProposalSubmission(ctx, proposalID)
	}
}

// AfterProposalDeposit - call hook if registered
func (keeper Keeper) AfterProposalDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress) {
	if keeper.hooks != nil {
		keeper.hooks.AfterProposalDeposit(ctx, proposalID, depositorAddr)
	}
}

// AfterProposalVote - call hook if registered
func (keeper Keeper) AfterProposalVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) {
	if keeper.hooks != nil {
		keeper.hooks.AfterProposalVote(ctx, proposalID, voterAddr)
	}
}

// AfterProposalFailedMinDeposit - call hook if registered
func (keeper Keeper) AfterProposalFailedMinDeposit(ctx sdk.Context, proposalID uint64) {
	if keeper.hooks != nil {
		keeper.hooks.AfterProposalFailedMinDeposit(ctx, proposalID)
	}
}

// AfterProposalVotingPeriodEnded - call hook if registered
func (keeper Keeper) AfterProposalVotingPeriodEnded(ctx sdk.Context, proposalID uint64) {
	if keeper.hooks != nil {
		keeper.hooks.AfterProposalVotingPeriodEnded(ctx, proposalID)
	}
}
