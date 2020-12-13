package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"


	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
)

// CreateMine creates a mine
func (k Keeper) CreateMine(ctx sdk.Context, mine types.Mine) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.MinePrefix + mine.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(mine)
	store.Set(key, value)
}
// GetMine returns the mine information
func (k Keeper) GetMine(ctx sdk.Context, key string) (types.Mine, error) {
	store := ctx.KVStore(k.storeKey)
	var mine types.Mine
	byteKey := []byte(types.MinePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &mine)
	if err != nil {
		return mine, err
	}
	return mine, nil
}

// SetMine sets a mine
func (k Keeper) SetMine(ctx sdk.Context, mine types.Mine) {
	mineKey := mine.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(mine)
	key := []byte(types.MinePrefix + mineKey)
	store.Set(key, bz)
}

// DeleteMine deletes a mine
func (k Keeper) DeleteMine(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.MinePrefix + key))
}

//
// Functions used by querier
//

func listMine(ctx sdk.Context, k Keeper) ([]byte, error) {
	var mineList []types.Mine
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.MinePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var mine types.Mine
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &mine)
		mineList = append(mineList, mine)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, mineList)
	return res, nil
}

func getMine(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	mine, err := k.GetMine(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, mine)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the mine
func (k Keeper) GetMineOwner(ctx sdk.Context, key string) sdk.AccAddress {
	mine, err := k.GetMine(ctx, key)
	if err != nil {
		return nil
	}
	return mine.Owner
}


// Check if the key exists in the store
func (k Keeper) MineExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.MinePrefix + key))
}

func (k Keeper) GetPrice(ctx sdk.Context,key string) sdk.Coins{
	mine,_ := k.GetMine(ctx,key)
	return mine.Price
}
func (k Keeper) HasOwner(ctx sdk.Context,key string) bool{
	mine,_ := k.GetMine(ctx,key)
	return !mine.Owner.Empty()
}

func (k Keeper) SetPrice(ctx sdk.Context,key string,price sdk.Coins){
	mine,_ := k.GetMine(ctx,key)
	mine.Price = price
	k.SetMine(ctx,mine)
}

func (k Keeper) SetOwner(ctx sdk.Context,key string,owner sdk.AccAddress){
	mine,_ := k.GetMine(ctx,key)
	mine.Owner = owner
	k.SetMine(ctx,mine)
}
func(k Keeper) SetSelling(ctx sdk.Context,key string,selling bool){
	mine,_ := k.GetMine(ctx,key)
	mine.Selling = selling
	k.SetMine(ctx,mine)
}

func(k Keeper) IsSelling(ctx sdk.Context,key string) bool{
	mine,_ := k.GetMine(ctx,key)
	return mine.Selling
}