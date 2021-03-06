package mineservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/keeper"
)

// Handle a message to delete name
func handleMsgMoveResource(ctx sdk.Context, k keeper.Keeper, msg types.MsgMoveResource) (*sdk.Result, error) {
	if !k.ResourceExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetResourceOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	var resource = types.Resource{
		ID:      msg.NewID,
    	Owner: msg.Creator,
		Selling: false,
		Amount: msg.Amount,
	}

	k.MoveResourceFromMine(ctx,msg.ID,resource,msg.PlayerID)
	
	return &sdk.Result{}, nil
}
