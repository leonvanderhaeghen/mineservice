package keeper

import (
  // this line is used by starport scaffolding # 1
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
		
	
		
	
		
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for mineservice clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
    // this line is used by starport scaffolding # 2
		case types.QueryListPlayer:
			return listPlayer(ctx, k)
		case types.QueryGetPlayer:
			return getPlayer(ctx, path[1:], k)
		case types.QueryListResource:
			return listResource(ctx, k)
		case types.QueryGetResource:
			return getResource(ctx, path[1:], k)
		case types.QueryListMine:
			return listMine(ctx, k)
		case types.QueryGetMine:
			return getMine(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown mineservice query endpoint")
		}
	}
}
