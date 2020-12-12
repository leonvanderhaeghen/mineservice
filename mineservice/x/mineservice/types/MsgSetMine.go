package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetMine{}

type MsgSetMine struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Name string `json:"name" yaml:"name"`
  Price sdk.Coins `json:"price" yaml:"price"`
  Owner sdk.AccAddress `json:"owner" yaml:"owner"`
  Selling bool `json:"selling" yaml:"selling"`
  Efficiency int `json:"efficiency" yaml:"efficiency"`
  Invetory string `json:"invetory" yaml:"invetory"`
  Resources []string `json:"resources" yaml:"resources"`
  UraniumCost int `json:"uraniumCost" yaml:"uraniumCost"`
  ResourceCounter int `json:"ResourceCounter" yaml:"ResourceCounter"`
}

func NewMsgSetMine(creator sdk.AccAddress, id string, name string, price sdk.Coins, owner sdk.AccAddress, selling bool, efficiency int, invetory string, resources []string, uraniumCost int) MsgSetMine {
  return MsgSetMine{
    ID: id,
		Creator: creator,
    Name: name,
    Price: price,
    Owner: owner,
    Selling: selling,
    Efficiency: efficiency,
    Invetory: invetory,
    Resources: resources,
    UraniumCost: uraniumCost,
	}
}

func (msg MsgSetMine) Route() string {
  return RouterKey
}

func (msg MsgSetMine) Type() string {
  return "SetMine"
}

func (msg MsgSetMine) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetMine) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetMine) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}