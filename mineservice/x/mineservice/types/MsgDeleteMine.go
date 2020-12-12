package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteMine{}

type MsgDeleteMine struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteMine(id string, creator sdk.AccAddress) MsgDeleteMine {
  return MsgDeleteMine{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteMine) Route() string {
  return RouterKey
}

func (msg MsgDeleteMine) Type() string {
  return "DeleteMine"
}

func (msg MsgDeleteMine) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteMine) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteMine) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}