package authz

import (
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
)

var _ Authorization = &GenericAuthorization{}

// NewGenericAuthorization creates a new GenericAuthorization object.
func NewGenericAuthorization(msgTypeURL string) *GenericAuthorization {
	return &GenericAuthorization{
		Msg: msgTypeURL,
	}
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (a GenericAuthorization) MsgTypeURL() string {
	return a.Msg
}

// Accept implements Authorization.Accept.
func (a GenericAuthorization) Accept(ctx sdk.Context, msg sdk.Msg) (AcceptResponse, error) {
	return AcceptResponse{Accept: true}, nil
}

// ValidateBasic implements Authorization.ValidateBasic.
func (a GenericAuthorization) ValidateBasic() error {
	return nil
}
