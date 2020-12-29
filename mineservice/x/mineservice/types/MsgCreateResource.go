package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateResource{}

type MsgCreateResource struct {
	ID string 
  Owner sdk.AccAddress `json:"owner" yaml:"owner"`
  Name string `json:"name" yaml:"name"`
  MineID string `json:"mineID" yaml:"mineID"`
  PlayerID string `json:"playerID" yaml:"playerID"`

  Amount int `json:"amount" yaml:"amount"`
}

func NewMsgCreateResource(owner sdk.AccAddress, name string,mineID string,playerID string,amount int) MsgCreateResource {
  return MsgCreateResource{
	  ID: uuid.New().String(),
		Owner: owner,
		Name: name,
		MineID: mineID,
		PlayerID: playerID,
		Amount: amount,
	}
}

func (msg MsgCreateResource) Route() string {
  return RouterKey
}

func (msg MsgCreateResource) Type() string {
  return "CreateResource"
}

func (msg MsgCreateResource) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg MsgCreateResource) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateResource) ValidateBasic() error {
  if msg.Owner.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
  }
  return nil
}