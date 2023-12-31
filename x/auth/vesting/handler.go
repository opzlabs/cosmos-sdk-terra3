package vesting

import (
	sdk "github.com/opzlabs/cosmos-sdk/types"
	sdkerrors "github.com/opzlabs/cosmos-sdk/types/errors"
	"github.com/opzlabs/cosmos-sdk/x/auth/keeper"
	"github.com/opzlabs/cosmos-sdk/x/auth/vesting/types"
)

// NewHandler returns a handler for x/auth message types.
func NewHandler(
	ak keeper.AccountKeeper,
	bk types.BankKeeper,
	dk types.DistrKeeper,
	sk types.StakingKeeper,
) sdk.Handler {
	msgServer := NewMsgServerImpl(ak, bk, dk, sk)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateVestingAccount:
			res, err := msgServer.CreateVestingAccount(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgCreatePeriodicVestingAccount:
			res, err := msgServer.CreatePeriodicVestingAccount(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgDonateAllVestingTokens:
			res, err := msgServer.DonateAllVestingTokens(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}
