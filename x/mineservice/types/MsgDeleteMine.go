package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteMine{}

type MsgDeleteMine struct {
  ID      string         `json:"id" yaml:"id"`
  Owner sdk.AccAddress `json:"owner" yaml:"owner"`
}

func NewMsgDeleteMine(id string, owner sdk.AccAddress) MsgDeleteMine {
  return MsgDeleteMine{
    ID: id,
		Owner: owner,
	}
}

func (msg MsgDeleteMine) Route() string {
  return RouterKey
}

func (msg MsgDeleteMine) Type() string {
  return "DeleteMine"
}

func (msg MsgDeleteMine) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg MsgDeleteMine) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteMine) ValidateBasic() error {
  if msg.Owner.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}