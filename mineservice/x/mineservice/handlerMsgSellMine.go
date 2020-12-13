package mineservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/keeper"
)

func handleMsgSellMine(ctx sdk.Context, k keeper.Keeper, msg types.MsgSellMine) (*sdk.Result, error) {
	if !msg.Owner.Equals(k.GetMineOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}	
	k.SetSelling(ctx,msg.ID,true)
	k.SetPrice(ctx,msg.ID,msg.AskingPrice)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
