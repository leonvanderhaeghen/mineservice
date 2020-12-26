package mineservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/keeper"
)

func handleMsgSetPlayer(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetPlayer) (*sdk.Result, error) {
	var player = types.Player{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Name: msg.Name,
	}
	if !msg.Creator.Equals(k.GetPlayerOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetPlayer(ctx, player)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
