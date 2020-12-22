package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetResource{}

type MsgSetResource struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Name string `json:"name" yaml:"name"`
  Price string `json:"price" yaml:"price"`
  Selling string `json:"selling" yaml:"selling"`
  MinPrice string `json:"minPrice" yaml:"minPrice"`
  Amount int `json:"amount" yaml:"amount"`

}

func NewMsgSetResource(creator sdk.AccAddress, id string, name string,amount int, price string, selling string, minPrice string) MsgSetResource {
  return MsgSetResource{
    ID: id,
	Creator: creator,
    Name: name,
    Price: price,
    Selling: selling,
	MinPrice: minPrice,
	  Amount: amount,

	}
}

func (msg MsgSetResource) Route() string {
  return RouterKey
}

func (msg MsgSetResource) Type() string {
  return "SetResource"
}

func (msg MsgSetResource) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetResource) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetResource) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}