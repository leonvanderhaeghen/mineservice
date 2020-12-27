package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMoveResource{}

type MsgMoveResource struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func MsgMoveResource(id string, creator sdk.AccAddress) MsgDeleteResource {
  return MsgDeleteResource{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgMoveResource) Route() string {
  return RouterKey
}

func (msg MsgMoveResource) Type() string {
  return "MoveResource"
}

func (msg MsgMoveResource) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgMoveResource) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgMoveResource) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}