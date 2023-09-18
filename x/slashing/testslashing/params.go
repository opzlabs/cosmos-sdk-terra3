package testslashing

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/slashing/types"
)

// TestParams construct default slashing params for tests.
// Have to change these parameters for tests
// lest the tests take forever
func TestParams() types.Params {
	params := types.DefaultParams()
	params.SignedBlocksWindow = 1000
	params.DowntimeJailDuration = 60 * 60

	return params
}
