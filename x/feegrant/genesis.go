package feegrant

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec/types"
)

var _ types.UnpackInterfacesMessage = GenesisState{}

// NewGenesisState creates new GenesisState object
func NewGenesisState(entries []Grant) *GenesisState {
	return &GenesisState{
		Allowances: entries,
	}
}

// ValidateGenesis ensures all grants in the genesis state are valid
func ValidateGenesis(data GenesisState) error {
	for _, f := range data.Allowances {
		grant, err := f.GetGrant()
		if err != nil {
			return err
		}
		err = grant.ValidateBasic()
		if err != nil {
			return err
		}
	}
	return nil
}

// DefaultGenesisState returns default state for feegrant module.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (data GenesisState) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	for _, f := range data.Allowances {
		err := f.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}

	return nil
}
