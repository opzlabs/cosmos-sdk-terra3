package baseapp

import (
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
)

type state struct {
	ms  sdk.CacheMultiStore
	ctx sdk.Context
}

// CacheMultiStore calls and returns a CacheMultiStore on the state's underling
// CacheMultiStore.
func (st *state) CacheMultiStore() sdk.CacheMultiStore {
	return st.ms.CacheMultiStore()
}

// Context returns the Context of the state.
func (st *state) Context() sdk.Context {
	return st.ctx
}
