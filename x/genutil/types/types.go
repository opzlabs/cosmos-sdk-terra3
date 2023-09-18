package types

import (
	"encoding/json"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/client"
	cryptotypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/crypto/types"
)

// DONTCOVER

type (
	// AppMap map modules names with their json raw representation.
	AppMap map[string]json.RawMessage

	// MigrationCallback converts a genesis map from the previous version to the
	// targeted one.
	//
	// TODO: MigrationCallback should also return an error upon failure.
	MigrationCallback func(AppMap, client.Context) AppMap

	// MigrationMap defines a mapping from a version to a MigrationCallback.
	MigrationMap map[string]MigrationCallback
)

// ModuleName is genutil
const ModuleName = "genutil"

// InitConfig common config options for init
type InitConfig struct {
	ChainID   string
	GenTxsDir string
	NodeID    string
	ValPubKey cryptotypes.PubKey
}

// NewInitConfig creates a new InitConfig object
func NewInitConfig(chainID, genTxsDir, nodeID string, valPubKey cryptotypes.PubKey) InitConfig {
	return InitConfig{
		ChainID:   chainID,
		GenTxsDir: genTxsDir,
		NodeID:    nodeID,
		ValPubKey: valPubKey,
	}
}
