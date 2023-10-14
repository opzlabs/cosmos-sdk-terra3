package types

import (
	"github.com/opzlabs/cosmos-sdk/codec"
	"github.com/opzlabs/cosmos-sdk/codec/legacy"
	"github.com/opzlabs/cosmos-sdk/codec/types"
	cryptocodec "github.com/opzlabs/cosmos-sdk/crypto/codec"
	sdk "github.com/opzlabs/cosmos-sdk/types"
	"github.com/opzlabs/cosmos-sdk/types/msgservice"
	authzcodec "github.com/opzlabs/cosmos-sdk/x/authz/codec"
)

// RegisterLegacyAminoCodec registers concrete types on LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgUnjail{}, "cosmos-sdk/MsgUnjail")
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnjail{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
