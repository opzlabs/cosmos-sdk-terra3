package keeper

import "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/gov/types"

// UnsafeSetHooks updates the gov keeper's hooks, overriding any potential
// pre-existing hooks.
// WARNING: this function should only be used in tests.
func UnsafeSetHooks(k *Keeper, h types.GovHooks) {
	k.hooks = h
}
