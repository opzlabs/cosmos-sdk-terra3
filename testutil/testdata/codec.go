package testdata

import (
	amino "github.com/tendermint/go-amino"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec/types"
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types/msgservice"
	tx "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types/tx"
)

func NewTestInterfaceRegistry() types.InterfaceRegistry {
	registry := types.NewInterfaceRegistry()
	RegisterInterfaces(registry)
	return registry
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &TestMsg{})

	registry.RegisterInterface("Animal", (*Animal)(nil))
	registry.RegisterImplementations(
		(*Animal)(nil),
		&Dog{},
		&Cat{},
	)
	registry.RegisterImplementations(
		(*HasAnimalI)(nil),
		&HasAnimal{},
	)
	registry.RegisterImplementations(
		(*HasHasAnimalI)(nil),
		&HasHasAnimal{},
	)
	registry.RegisterImplementations(
		(*tx.TxExtensionOptionI)(nil),
		&Cat{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func NewTestAmino() *amino.Codec {
	cdc := amino.NewCodec()
	cdc.RegisterInterface((*Animal)(nil), nil)
	cdc.RegisterConcrete(&Dog{}, "testdata/Dog", nil)
	cdc.RegisterConcrete(&Cat{}, "testdata/Cat", nil)

	cdc.RegisterInterface((*HasAnimalI)(nil), nil)
	cdc.RegisterConcrete(&HasAnimal{}, "testdata/HasAnimal", nil)

	cdc.RegisterInterface((*HasHasAnimalI)(nil), nil)
	cdc.RegisterConcrete(&HasHasAnimal{}, "testdata/HasHasAnimal", nil)

	return cdc
}
