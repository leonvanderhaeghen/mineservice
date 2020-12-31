package mineservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/keeper"
)

func handleMsgSetResource(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetResource) (*sdk.Result, error) {
	var resource = types.Resource{
		Owner: msg.Creator,
		ID:      msg.ID,
		Name: msg.Name,
		Amount: msg.Amount,
    	//Price: msg.Price,
    	//Selling: msg.Selling,
    	//MinPrice: msg.MinPrice,
	}
	if !msg.Creator.Equals(k.GetResourceOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetResource(ctx, resource)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
