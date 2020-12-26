package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeletePlayer{}

type MsgDeletePlayer struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeletePlayer(id string, creator sdk.AccAddress) MsgDeletePlayer {
  return MsgDeletePlayer{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeletePlayer) Route() string {
  return RouterKey
}

func (msg MsgDeletePlayer) Type() string {
  return "DeletePlayer"
}

func (msg MsgDeletePlayer) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeletePlayer) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeletePlayer) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}