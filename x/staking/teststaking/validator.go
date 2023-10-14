package teststaking

import (
	"testing"

	"github.com/stretchr/testify/require"

	cryptotypes "github.com/opzlabs/cosmos-sdk/crypto/types"
	sdk "github.com/opzlabs/cosmos-sdk/types"
	"github.com/opzlabs/cosmos-sdk/x/staking/types"
)

// NewValidator is a testing helper method to create validators in tests
func NewValidator(t testing.TB, operator sdk.ValAddress, pubKey cryptotypes.PubKey) types.Validator {
	v, err := types.NewValidator(operator, pubKey, types.Description{})
	require.NoError(t, err)
	return v
}
