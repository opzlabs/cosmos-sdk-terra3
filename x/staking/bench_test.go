package staking_test

import (
	"testing"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/crypto/keys/ed25519"
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/staking"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/staking/teststaking"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/staking/types"
)

func BenchmarkValidateGenesis10Validators(b *testing.B) {
	benchmarkValidateGenesis(b, 10)
}

func BenchmarkValidateGenesis100Validators(b *testing.B) {
	benchmarkValidateGenesis(b, 100)
}

func BenchmarkValidateGenesis400Validators(b *testing.B) {
	benchmarkValidateGenesis(b, 400)
}

func benchmarkValidateGenesis(b *testing.B, n int) {
	b.ReportAllocs()

	validators := make([]types.Validator, 0, n)
	addressL, pubKeyL := makeRandomAddressesAndPublicKeys(n)
	for i := 0; i < n; i++ {
		addr, pubKey := addressL[i], pubKeyL[i]
		validator := teststaking.NewValidator(b, addr, pubKey)
		ni := int64(i + 1)
		validator.Tokens = sdk.NewInt(ni)
		validator.DelegatorShares = sdk.NewDec(ni)
		validators = append(validators, validator)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genesisState := types.DefaultGenesisState()
		genesisState.Validators = validators
		if err := staking.ValidateGenesis(genesisState); err != nil {
			b.Fatal(err)
		}
	}
}

func makeRandomAddressesAndPublicKeys(n int) (accL []sdk.ValAddress, pkL []*ed25519.PubKey) {
	for i := 0; i < n; i++ {
		pk := ed25519.GenPrivKey().PubKey().(*ed25519.PubKey)
		pkL = append(pkL, pk)
		accL = append(accL, sdk.ValAddress(pk.Address()))
	}
	return accL, pkL
}
