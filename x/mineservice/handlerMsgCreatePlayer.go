package mineservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/keeper"
)

func handleMsgCreatePlayer(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreatePlayer) (*sdk.Result, error) {
	k.CreatePlayer(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
