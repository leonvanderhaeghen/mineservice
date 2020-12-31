package mineservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/leonvanderhaeghen/mineservice/x/mineservice/keeper"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
)

func handleMsgCreateResource(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateResource) (*sdk.Result, error) {
	var resource = types.Resource{
		ID:      msg.ID,
		MineID: msg.MineID,
    	Name: msg.Name,
    	Owner: msg.Owner,
		Selling: false,
		Amount: msg.Amount,
	}
		
	k.CreateResource(ctx, resource)
	k.AddResourceMine(ctx,resource)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
