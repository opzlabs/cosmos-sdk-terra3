package std

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec/types"
	cryptocodec "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/crypto/codec"
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	txtypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types/tx"
)

// RegisterLegacyAminoCodec registers types with the Amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)
}

// RegisterInterfaces registers Interfaces from sdk/types, vesting, crypto, tx.
func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	sdk.RegisterInterfaces(interfaceRegistry)
	txtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
}
