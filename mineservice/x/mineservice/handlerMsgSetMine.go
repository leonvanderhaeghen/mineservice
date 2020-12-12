package mineservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/keeper"
)

func handleMsgSetMine(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetMine) (*sdk.Result, error) {
	var mine = types.Mine{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Name: msg.Name,
    	Price: msg.Price,
    	Owner: msg.Owner,
    	Selling: msg.Selling,
    	Efficiency: msg.Efficiency,
    	Invetory: msg.Invetory,
    	Resources: msg.Resources,
    	UraniumCost: msg.UraniumCost,
	}
	if !msg.Creator.Equals(k.GetMineOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetMine(ctx, mine)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
