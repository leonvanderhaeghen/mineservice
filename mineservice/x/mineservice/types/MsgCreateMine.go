package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateMine{}

type MsgCreateMine struct {
  ID      string
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

func NewMsgCreateMine(creator sdk.AccAddress, name string, price sdk.Coins, owner sdk.AccAddress, selling bool, efficiency int, invetory string, resources []string, uraniumCost int) MsgCreateMine {
  return MsgCreateMine{
    ID: uuid.New().String(),
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

func (msg MsgCreateMine) Route() string {
  return RouterKey
}

func (msg MsgCreateMine) Type() string {
  return "CreateMine"
}

func (msg MsgCreateMine) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateMine) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateMine) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}