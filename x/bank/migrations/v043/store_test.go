package v043_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/simapp"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/store/prefix"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/testutil"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/testutil/testdata"
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	v042bank "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/bank/migrations/v042"
	v043bank "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/bank/migrations/v043"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/bank/types"
)

func TestSupplyMigration(t *testing.T) {
	encCfg := simapp.MakeTestEncodingConfig()
	bankKey := sdk.NewKVStoreKey("bank")
	ctx := testutil.DefaultContext(bankKey, sdk.NewTransientStoreKey("transient_test"))
	store := ctx.KVStore(bankKey)

	oldFooCoin := sdk.NewCoin("foo", sdk.NewInt(100))
	oldBarCoin := sdk.NewCoin("bar", sdk.NewInt(200))
	oldFooBarCoin := sdk.NewCoin("foobar", sdk.NewInt(0)) // to ensure the zero denom coins pruned.

	// Old supply was stored as a single blob under the `SupplyKey`.
	var oldSupply v042bank.SupplyI
	oldSupply = &types.Supply{Total: sdk.Coins{oldFooCoin, oldBarCoin, oldFooBarCoin}}
	oldSupplyBz, err := encCfg.Codec.MarshalInterface(oldSupply)
	require.NoError(t, err)
	store.Set(v042bank.SupplyKey, oldSupplyBz)

	// Run migration.
	err = v043bank.MigrateStore(ctx, bankKey, encCfg.Codec)
	require.NoError(t, err)

	// New supply is indexed by denom.
	supplyStore := prefix.NewStore(store, types.SupplyKey)
	bz := supplyStore.Get([]byte("foo"))
	var amount math.Int
	err = amount.Unmarshal(bz)
	require.NoError(t, err)

	newFooCoin := sdk.Coin{
		Denom:  "foo",
		Amount: amount,
	}
	require.Equal(t, oldFooCoin, newFooCoin)

	bz = supplyStore.Get([]byte("bar"))
	err = amount.Unmarshal(bz)
	require.NoError(t, err)

	newBarCoin := sdk.Coin{
		Denom:  "bar",
		Amount: amount,
	}
	require.Equal(t, oldBarCoin, newBarCoin)

	// foobar shouldn't be existed in the store.
	bz = supplyStore.Get([]byte("foobar"))
	require.Nil(t, bz)
}

func TestBalanceKeysMigration(t *testing.T) {
	encCfg := simapp.MakeTestEncodingConfig()
	bankKey := sdk.NewKVStoreKey("bank")
	ctx := testutil.DefaultContext(bankKey, sdk.NewTransientStoreKey("transient_test"))
	store := ctx.KVStore(bankKey)

	_, _, addr := testdata.KeyTestPubAddr()

	// set 10 foo coin
	fooCoin := sdk.NewCoin("foo", sdk.NewInt(10))
	oldFooKey := append(append(v042bank.BalancesPrefix, addr...), []byte(fooCoin.Denom)...)
	fooBz, err := encCfg.Codec.Marshal(&fooCoin)
	require.NoError(t, err)
	store.Set(oldFooKey, fooBz)

	// set 0 foobar coin
	fooBarCoin := sdk.NewCoin("foobar", sdk.NewInt(0))
	oldKeyFooBar := append(append(v042bank.BalancesPrefix, addr...), []byte(fooBarCoin.Denom)...)
	fooBarBz, err := encCfg.Codec.Marshal(&fooBarCoin)
	require.NoError(t, err)
	store.Set(oldKeyFooBar, fooBarBz)
	require.NotNil(t, store.Get(oldKeyFooBar)) // before store migation zero values can also exist in store.

	err = v043bank.MigrateStore(ctx, bankKey, encCfg.Codec)
	require.NoError(t, err)

	newKey := types.CreatePrefixedAccountStoreKey(addr, []byte(fooCoin.Denom))
	// -7 because we replaced "balances" with 0x02,
	// +1 because we added length-prefix to address.
	require.Equal(t, len(oldFooKey)-7+1, len(newKey))
	require.Nil(t, store.Get(oldFooKey))
	require.Equal(t, fooBz, store.Get(newKey))

	newKeyFooBar := types.CreatePrefixedAccountStoreKey(addr, []byte(fooBarCoin.Denom))
	require.Nil(t, store.Get(newKeyFooBar)) // after migration zero balances pruned from store.
}
