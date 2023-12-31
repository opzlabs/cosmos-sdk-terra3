package upgrade

import (
	sdk "github.com/opzlabs/cosmos-sdk/types"
	sdkerrors "github.com/opzlabs/cosmos-sdk/types/errors"
	govtypes "github.com/opzlabs/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/opzlabs/cosmos-sdk/x/upgrade/keeper"
	"github.com/opzlabs/cosmos-sdk/x/upgrade/types"
)

// NewSoftwareUpgradeProposalHandler creates a governance handler to manage new proposal types.
// It enables SoftwareUpgradeProposal to propose an Upgrade, and CancelSoftwareUpgradeProposal
// to abort a previously voted upgrade.
func NewSoftwareUpgradeProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.SoftwareUpgradeProposal: //nolint:staticcheck
			return handleSoftwareUpgradeProposal(ctx, k, c)

		case *types.CancelSoftwareUpgradeProposal: //nolint:staticcheck
			return handleCancelSoftwareUpgradeProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized software upgrade proposal content type: %T", c)
		}
	}
}

func handleSoftwareUpgradeProposal(ctx sdk.Context, k keeper.Keeper, p *types.SoftwareUpgradeProposal) error { //nolint:staticcheck
	return k.ScheduleUpgrade(ctx, p.Plan)
}

func handleCancelSoftwareUpgradeProposal(ctx sdk.Context, k keeper.Keeper, _ *types.CancelSoftwareUpgradeProposal) error { //nolint:staticcheck
	k.ClearUpgradePlan(ctx)
	return nil
}
