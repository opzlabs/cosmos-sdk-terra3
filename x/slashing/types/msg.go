package types

import (
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	sdkerrors "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types/errors"
)

// slashing message types
const (
	TypeMsgUnjail = "unjail"
)

// verify interface at compile time
var _ sdk.Msg = &MsgUnjail{}

// NewMsgUnjail creates a new MsgUnjail instance
//
//nolint:interfacer
func NewMsgUnjail(validatorAddr sdk.ValAddress) *MsgUnjail {
	return &MsgUnjail{
		ValidatorAddr: validatorAddr.String(),
	}
}

func (msg MsgUnjail) Route() string { return RouterKey }
func (msg MsgUnjail) Type() string  { return TypeMsgUnjail }
func (msg MsgUnjail) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.ValAddressFromBech32(msg.ValidatorAddr)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgUnjail) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic does a sanity check on the provided message
func (msg MsgUnjail) ValidateBasic() error {
	if _, err := sdk.ValAddressFromBech32(msg.ValidatorAddr); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("validator input address: %s", err)
	}
	return nil
}
