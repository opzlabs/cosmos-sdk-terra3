package keeper

import (
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	v043 "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/slashing/migrations/v043"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v043.MigrateStore(ctx, m.keeper.storeKey)
}
