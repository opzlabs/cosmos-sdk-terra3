package tmservice

import (
	"context"

	coretypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/client"
)

func getNodeStatus(ctx context.Context, clientCtx client.Context) (*coretypes.ResultStatus, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return &coretypes.ResultStatus{}, err
	}
	return node.Status(ctx)
}
