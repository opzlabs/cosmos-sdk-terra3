package types

import (
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
)

// invariant route
type InvarRoute struct {
	ModuleName string
	Route      string
	Invar      sdk.Invariant
}

// NewInvarRoute - create an InvarRoute object
func NewInvarRoute(moduleName, route string, invar sdk.Invariant) InvarRoute {
	return InvarRoute{
		ModuleName: moduleName,
		Route:      route,
		Invar:      invar,
	}
}

// get the full invariance route
func (i InvarRoute) FullRoute() string {
	return i.ModuleName + "/" + i.Route
}
