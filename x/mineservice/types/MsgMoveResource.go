package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgMoveResource{}

type MsgMoveResource struct {
  ID      string         `json:"id" yaml:"id"`
  NewID      string
  Amount      int         `json:"amount" yaml:"amount"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  PlayerID string `json:"playerid" yaml:"playerid"`

}

func NewMsgMoveResource(id string,playerID string,amount int, creator sdk.AccAddress) MsgMoveResource {
  return MsgMoveResource{
		ID: id,
		NewID:  uuid.New().String(),
		Amount: amount,
		Creator: creator,
		PlayerID: playerID,
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