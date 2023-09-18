package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec"
	cryptoAmino "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/crypto/codec"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/testutil/testdata"
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/auth/migrations/legacytx"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/auth/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "cosmos-sdk/Test", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, testutil.NewTxConfigTestSuite(txGen))
}
