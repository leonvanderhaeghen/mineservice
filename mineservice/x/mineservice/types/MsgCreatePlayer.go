package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreatePlayer{}

type MsgCreatePlayer struct {
	  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Name string `json:"name" yaml:"name"`
}

func NewMsgCreatePlayer(creator sdk.AccAddress, name string, invetory string, mines string) MsgCreatePlayer {
  return MsgCreatePlayer{
	ID: uuid.New().String(),
	Creator: creator,
    Name: name,
	}
}

func (msg MsgCreatePlayer) Route() string {
  return RouterKey
}

func (msg MsgCreatePlayer) Type() string {
  return "CreatePlayer"
}

func (msg MsgCreatePlayer) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreatePlayer) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreatePlayer) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}