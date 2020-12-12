package mineservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/keeper"
)

func handleMsgBuyMine(ctx sdk.Context, k keeper.Keeper, msg types.MsgBuyMine) (*sdk.Result, error) {
	if k.GetPrice(ctx,msg.ID).IsAllGT(msg.Bid) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds,"Bid not high enough")
	}
	if k.HasOwner(ctx,msg.ID) {
		err := k.CoinKeeper.SendCoins(ctx,msg.Buyer,k.GetMineOwner(ctx,msg.ID),msg.Bid)
		if err != nil {
			return nil,err
		}
	}else{
		_, err := k.CoinKeeper.SubtractCoins(ctx,msg.Buyer,msg.Bid);
		if err != nil {
			return nil, err
		}
	}	
	k.SetOwner(ctx,msg.ID,msg.Buyer)
	k.SetPrice(ctx,msg.ID,msg.Bid)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
