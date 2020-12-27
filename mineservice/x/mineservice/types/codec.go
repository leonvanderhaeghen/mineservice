package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
  // this line is used by starport scaffolding # 1
		cdc.RegisterConcrete(MsgCreatePlayer{}, "mineservice/CreatePlayer", nil)
		cdc.RegisterConcrete(MsgSetPlayer{}, "mineservice/SetPlayer", nil)
		cdc.RegisterConcrete(MsgDeletePlayer{}, "mineservice/DeletePlayer", nil)
		cdc.RegisterConcrete(MsgCreateResource{}, "mineservice/CreateResource", nil)
		cdc.RegisterConcrete(MsgSetResource{}, "mineservice/SetResource", nil)
		cdc.RegisterConcrete(MsgDeleteResource{}, "mineservice/DeleteResource", nil)
		cdc.RegisterConcrete(MsgMoveResource{}, "mineservice/MoveResource", nil)
		cdc.RegisterConcrete(MsgCreateMine{}, "mineservice/CreateMine", nil)
		cdc.RegisterConcrete(MsgBuyMine{}, "mineservice/BuyMine", nil)
		cdc.RegisterConcrete(MsgSetMine{}, "mineservice/SetMine", nil)
		cdc.RegisterConcrete(MsgDeleteMine{}, "mineservice/DeleteMine", nil)
		cdc.RegisterConcrete(MsgSellMine{},"mineservice/SellMine",nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
