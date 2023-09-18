package testutil

import (
	"testing"
	"time"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/testutil/network"
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	v1 "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/gov/types/v1"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 1
	suite.Run(t, NewIntegrationTestSuite(cfg))

	dp := v1.NewDepositParams(sdk.NewCoins(sdk.NewCoin(cfg.BondDenom, v1.DefaultMinDepositTokens)), time.Duration(15)*time.Second)
	vp := v1.NewVotingParams(time.Duration(5) * time.Second)
	genesisState := v1.DefaultGenesisState()
	genesisState.DepositParams = &dp
	genesisState.VotingParams = &vp
	bz, err := cfg.Codec.MarshalJSON(genesisState)
	require.NoError(t, err)
	cfg.GenesisState["gov"] = bz
	suite.Run(t, NewDepositTestSuite(cfg))
}
