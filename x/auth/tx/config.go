package tx

import (
	"fmt"

	signingtypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types/tx/signing"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/client"
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/auth/signing"
)

type config struct {
	handler     signing.SignModeHandler
	decoder     sdk.TxDecoder
	encoder     sdk.TxEncoder
	jsonDecoder sdk.TxDecoder
	jsonEncoder sdk.TxEncoder
	protoCodec  codec.ProtoCodecMarshaler
}

// NewTxConfig returns a new protobuf TxConfig using the provided ProtoCodec and sign modes. The
// first enabled sign mode will become the default sign mode.
// NOTE: Use NewTxConfigWithHandler to provide a custom signing handler in case the sign mode
// is not supported by default (eg: SignMode_SIGN_MODE_EIP_191).
func NewTxConfig(protoCodec codec.ProtoCodecMarshaler, enabledSignModes []signingtypes.SignMode) client.TxConfig {
	return NewTxConfigWithHandler(protoCodec, makeSignModeHandler(enabledSignModes))
}

// NewTxConfig returns a new protobuf TxConfig using the provided ProtoCodec and signing handler.
func NewTxConfigWithHandler(protoCodec codec.ProtoCodecMarshaler, handler signing.SignModeHandler) client.TxConfig {
	return &config{
		handler:     handler,
		decoder:     DefaultTxDecoder(protoCodec),
		encoder:     DefaultTxEncoder(),
		jsonDecoder: DefaultJSONTxDecoder(protoCodec),
		jsonEncoder: DefaultJSONTxEncoder(protoCodec),
		protoCodec:  protoCodec,
	}
}

func (g config) NewTxBuilder() client.TxBuilder {
	return newBuilder(g.protoCodec)
}

// WrapTxBuilder returns a builder from provided transaction
func (g config) WrapTxBuilder(newTx sdk.Tx) (client.TxBuilder, error) {
	newBuilder, ok := newTx.(*wrapper)
	if !ok {
		return nil, fmt.Errorf("expected %T, got %T", &wrapper{}, newTx)
	}

	return newBuilder, nil
}

func (g config) SignModeHandler() signing.SignModeHandler {
	return g.handler
}

func (g config) TxEncoder() sdk.TxEncoder {
	return g.encoder
}

func (g config) TxDecoder() sdk.TxDecoder {
	return g.decoder
}

func (g config) TxJSONEncoder() sdk.TxEncoder {
	return g.jsonEncoder
}

func (g config) TxJSONDecoder() sdk.TxDecoder {
	return g.jsonDecoder
}
