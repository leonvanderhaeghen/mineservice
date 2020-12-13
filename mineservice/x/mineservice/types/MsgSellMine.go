package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSellMine{}

type MsgSellMine struct {
  ID      string `json:"id" yaml:"id"`
  Owner sdk.AccAddress `json:"owner" yaml:"owner"`
  AskingPrice sdk.Coins `json:"price" yaml:"price"`
}

func NewMsgSellMine(owner sdk.AccAddress,id string, askingPrice sdk.Coins) MsgSellMine {
  return MsgSellMine{
    ID: id,
	Owner: owner,
    AskingPrice: askingPrice,
	}
}

func (msg MsgSellMine) Route() string {
  return RouterKey
}

func (msg MsgSellMine) Type() string {
  return "SellMine"
}

func (msg MsgSellMine) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg MsgSellMine) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSellMine) ValidateBasic() error {
  if msg.Owner.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  if len(msg.ID) == 0 {
	  return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,"id cannot be empty")
  }
  /*if msg.AskingPrice. {
	      return sdkerrors.ErrInsufficientFunds
  }*/
  return nil
}