package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteResource{}

type MsgDeleteResource struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteResource(id string, creator sdk.AccAddress) MsgDeleteResource {
  return MsgDeleteResource{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteResource) Route() string {
  return RouterKey
}

func (msg MsgDeleteResource) Type() string {
  return "DeleteResource"
}

func (msg MsgDeleteResource) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteResource) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteResource) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}