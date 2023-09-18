package v046

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec"
	storetypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/store/types"
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/auth/types"
)

func mapAccountAddressToAccountID(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AddressStoreKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var acc types.AccountI
		if err := cdc.UnmarshalInterface(iterator.Value(), &acc); err != nil {
			return err
		}
		store.Set(types.AccountNumberStoreKey(acc.GetAccountNumber()), acc.GetAddress().Bytes())
	}

	return nil
}

// MigrateStore performs in-place store migrations from v0.45 to v0.46. The
// migration includes:
// - Add an Account number as an index to get the account address
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	return mapAccountAddressToAccountID(ctx, storeKey, cdc)
}
