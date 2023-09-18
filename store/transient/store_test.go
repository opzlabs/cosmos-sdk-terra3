package transient_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	pruningtypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/pruning/types"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/store/transient"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/store/types"
)

var k, v = []byte("hello"), []byte("world")

func TestTransientStore(t *testing.T) {
	tstore := transient.NewStore()

	require.Nil(t, tstore.Get(k))

	tstore.Set(k, v)

	require.Equal(t, v, tstore.Get(k))

	tstore.Commit()

	require.Nil(t, tstore.Get(k))

	// no-op
	tstore.SetPruning(pruningtypes.NewPruningOptions(pruningtypes.PruningUndefined))

	emptyCommitID := tstore.LastCommitID()
	require.Equal(t, emptyCommitID.Version, int64(0))
	require.True(t, bytes.Equal(emptyCommitID.Hash, nil))
	require.Equal(t, types.StoreTypeTransient, tstore.GetStoreType())
}
