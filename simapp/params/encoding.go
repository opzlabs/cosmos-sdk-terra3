package params

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/client"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec/types"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}
