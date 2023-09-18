package cli

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/client"
	simappparams "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/simapp/params"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/testutil"
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	authtypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/auth/types"
)

func TestGetCommandEncode(t *testing.T) {
	encodingConfig := simappparams.MakeTestEncodingConfig()

	cmd := GetEncodeCommand()
	_ = testutil.ApplyMockIODiscardOutErr(cmd)

	authtypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	sdk.RegisterLegacyAminoCodec(encodingConfig.Amino)

	txCfg := encodingConfig.TxConfig

	// Build a test transaction
	builder := txCfg.NewTxBuilder()
	builder.SetGasLimit(50000)
	builder.SetFeeAmount(sdk.Coins{sdk.NewInt64Coin("atom", 150)})
	builder.SetMemo("foomemo")
	jsonEncoded, err := txCfg.TxJSONEncoder()(builder.GetTx())
	require.NoError(t, err)

	txFile := testutil.WriteToNewTempFile(t, string(jsonEncoded))
	txFileName := txFile.Name()

	ctx := context.Background()
	clientCtx := client.Context{}.
		WithTxConfig(encodingConfig.TxConfig).
		WithCodec(encodingConfig.Codec)
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)

	cmd.SetArgs([]string{txFileName})
	err = cmd.ExecuteContext(ctx)
	require.NoError(t, err)
}

func TestGetCommandDecode(t *testing.T) {
	encodingConfig := simappparams.MakeTestEncodingConfig()

	clientCtx := client.Context{}.
		WithTxConfig(encodingConfig.TxConfig).
		WithCodec(encodingConfig.Codec)

	cmd := GetDecodeCommand()
	_ = testutil.ApplyMockIODiscardOutErr(cmd)

	sdk.RegisterLegacyAminoCodec(encodingConfig.Amino)

	txCfg := encodingConfig.TxConfig
	clientCtx = clientCtx.WithTxConfig(txCfg)

	// Build a test transaction
	builder := txCfg.NewTxBuilder()
	builder.SetGasLimit(50000)
	builder.SetFeeAmount(sdk.Coins{sdk.NewInt64Coin("atom", 150)})
	builder.SetMemo("foomemo")

	// Encode transaction
	txBytes, err := clientCtx.TxConfig.TxEncoder()(builder.GetTx())
	require.NoError(t, err)

	// Convert the transaction into base64 encoded string
	base64Encoded := base64.StdEncoding.EncodeToString(txBytes)

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)

	// Execute the command
	cmd.SetArgs([]string{base64Encoded})
	require.NoError(t, cmd.ExecuteContext(ctx))
}
