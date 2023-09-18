// DONTCOVER
package types

import (
	sdkerrors "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types/errors"
)

// x/evidence module sentinel errors
var (
	ErrNoEvidenceHandlerExists = sdkerrors.Register(ModuleName, 2, "unregistered handler for evidence type")
	ErrInvalidEvidence         = sdkerrors.Register(ModuleName, 3, "invalid evidence")
	ErrNoEvidenceExists        = sdkerrors.Register(ModuleName, 4, "evidence does not exist")
	ErrEvidenceExists          = sdkerrors.Register(ModuleName, 5, "evidence already exists")
)
