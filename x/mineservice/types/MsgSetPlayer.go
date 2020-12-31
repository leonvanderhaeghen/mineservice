package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetPlayer{}

type MsgSetPlayer struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Name string `json:"name" yaml:"name"`
}

func NewMsgSetPlayer(creator sdk.AccAddress, id string, name string) MsgSetPlayer {
  return MsgSetPlayer{
    ID: id,
	Creator: creator,
    Name: name,

	}
}

func (msg MsgSetPlayer) Route() string {
  return RouterKey
}

func (msg MsgSetPlayer) Type() string {
  return "SetPlayer"
}

func (msg MsgSetPlayer) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetPlayer) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetPlayer) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}