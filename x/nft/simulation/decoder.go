package simulation

import (
	"bytes"
	"fmt"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec"
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types/kv"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/nft"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/nft/keeper"
)

// NewDecodeStore returns a decoder function closure that umarshals the KVPair's
// Value to the corresponding nft type.
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], keeper.ClassKey):
			var classA, classB nft.Class
			cdc.MustUnmarshal(kvA.Value, &classA)
			cdc.MustUnmarshal(kvB.Value, &classB)
			return fmt.Sprintf("%v\n%v", classA, classB)
		case bytes.Equal(kvA.Key[:1], keeper.NFTKey):
			var nftA, nftB nft.NFT
			cdc.MustUnmarshal(kvA.Value, &nftA)
			cdc.MustUnmarshal(kvB.Value, &nftB)
			return fmt.Sprintf("%v\n%v", nftA, nftB)
		case bytes.Equal(kvA.Key[:1], keeper.NFTOfClassByOwnerKey):
			return fmt.Sprintf("%v\n%v", kvA.Value, kvB.Value)
		case bytes.Equal(kvA.Key[:1], keeper.OwnerKey):
			var ownerA, ownerB sdk.AccAddress
			ownerA = sdk.AccAddress(kvA.Value)
			ownerB = sdk.AccAddress(kvB.Value)
			return fmt.Sprintf("%v\n%v", ownerA, ownerB)
		case bytes.Equal(kvA.Key[:1], keeper.ClassTotalSupply):
			var supplyA, supplyB uint64
			supplyA = sdk.BigEndianToUint64(kvA.Value)
			supplyB = sdk.BigEndianToUint64(kvB.Value)
			return fmt.Sprintf("%v\n%v", supplyA, supplyB)
		default:
			panic(fmt.Sprintf("invalid nft key %X", kvA.Key))
		}
	}
}
