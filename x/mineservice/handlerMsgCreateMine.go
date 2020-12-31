package mineservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/keeper"
)

func handleMsgCreateMine(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateMine) (*sdk.Result, error) {
	var mine = types.Mine{
		ID:      msg.ID,
    	Name: msg.Name,
    	Price: msg.Price,
    	Owner: msg.Owner,
    	Selling: msg.Selling,
    	Efficiency: msg.Efficiency,
    	Resources: msg.Resources,
		UraniumCost: msg.UraniumCost,
		ResourceCounter: 5,
	}
	k.CreateMine(ctx, mine)
	k.AddMineToPlayer(ctx,msg.PlayerID,mine.ID)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
