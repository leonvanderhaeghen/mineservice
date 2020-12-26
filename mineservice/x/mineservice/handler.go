package mineservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/keeper"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding # 1
		case types.MsgCreatePlayer:
			return handleMsgCreatePlayer(ctx, k, msg)
		case types.MsgSetPlayer:
			return handleMsgSetPlayer(ctx, k, msg)
		case types.MsgDeletePlayer:
			return handleMsgDeletePlayer(ctx, k, msg)
		case types.MsgCreateResource:
			return handleMsgCreateResource(ctx, k, msg)
		case types.MsgSetResource:
			return handleMsgSetResource(ctx, k, msg)
		case types.MsgDeleteResource:
			return handleMsgDeleteResource(ctx, k, msg)
		case types.MsgCreateMine:
			return handleMsgCreateMine(ctx, k, msg)
		case types.MsgSetMine:
			return handleMsgSetMine(ctx, k, msg)
		case types.MsgDeleteMine:
			return handleMsgDeleteMine(ctx, k, msg)
		case types.MsgBuyMine:
			return handleMsgBuyMine(ctx,k,msg)
		case types.MsgSellMine:
			return handleMsgSellMine(ctx,k,msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
