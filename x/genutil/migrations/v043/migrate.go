package v043

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/client"
	v042bank "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/bank/migrations/v042"
	v043bank "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/bank/migrations/v043"
	bank "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/bank/types"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/genutil/types"
	v042gov "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/gov/migrations/v042"
	v043gov "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/gov/migrations/v043"
	gov "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/gov/types/v1beta1"
)

// Migrate migrates exported state from v0.40 to a v0.43 genesis state.
func Migrate(appState types.AppMap, clientCtx client.Context) types.AppMap {
	// Migrate x/gov.
	if appState[v042gov.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var oldGovState gov.GenesisState
		clientCtx.Codec.MustUnmarshalJSON(appState[v042gov.ModuleName], &oldGovState)

		// delete deprecated x/gov genesis state
		delete(appState, v042gov.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[v043gov.ModuleName] = clientCtx.Codec.MustMarshalJSON(v043gov.MigrateJSON(&oldGovState))
	}

	if appState[v042bank.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var oldBankState bank.GenesisState
		clientCtx.Codec.MustUnmarshalJSON(appState[v042bank.ModuleName], &oldBankState)

		// delete deprecated x/bank genesis state
		delete(appState, v042bank.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[v043bank.ModuleName] = clientCtx.Codec.MustMarshalJSON(v043bank.MigrateJSON(&oldBankState))
	}

	return appState
}
