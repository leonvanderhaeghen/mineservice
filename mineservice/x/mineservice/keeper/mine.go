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

func (k Keeper) GetMinePrice(ctx sdk.Context,key string) sdk.Coins{
	mine,_ := k.GetMine(ctx,key)
	return mine.Price
}
func (k Keeper) HasMineOwner(ctx sdk.Context,key string) bool{
	mine,_ := k.GetMine(ctx,key)
	return !mine.Owner.Empty()
}

func (k Keeper) SetMinePrice(ctx sdk.Context,key string,price sdk.Coins){
	mine,_ := k.GetMine(ctx,key)
	mine.Price = price
	k.SetMine(ctx,mine)
}

func (k Keeper) SetMineOwner(ctx sdk.Context,key string,owner sdk.AccAddress){
	mine,_ := k.GetMine(ctx,key)
	mine.Owner = owner
	k.SetMine(ctx,mine)
}
func(k Keeper) SetMineSelling(ctx sdk.Context,key string,selling bool){
	mine,_ := k.GetMine(ctx,key)
	mine.Selling = selling
	k.SetMine(ctx,mine)
}

func(k Keeper) IsMineSelling(ctx sdk.Context,key string) bool{
	mine,_ := k.GetMine(ctx,key)
	return mine.Selling
}
func(k Keeper) MoveResourceFromMine(ctx sdk.Context,key string,newResource types.Resource,playerID string){
	resource,_ := k.GetResource(ctx,key)
	if k.checkResourceAmount(ctx,resource.MineID,resource.Name,newResource.Amount) {
		k.updateMineResourceAmountByName(ctx,resource.MineID,resource.Name,-newResource.Amount)
		newResource.Name = resource.Name
		k.AddResourcePlayer(ctx,newResource,playerID)
	}
}
func(k Keeper) AddResourceMine(ctx sdk.Context,resource types.Resource){
	mine,_ := k.GetMine(ctx,resource.MineID)
	if k.resourceExistsInMine(ctx,resource.MineID,resource.Name) {
		k.updateMineResourceAmountByName(ctx,resource.MineID,resource.Name,resource.Amount)
	}else{
		mine.Invetory = append(mine.Invetory, resource)
		k.SetMine(ctx,mine)
	}
}

func(k Keeper) resourceExistsInMine(ctx sdk.Context,key string,resourceName string)bool{
	mine,_ := k.GetMine(ctx,key)
	resources := mine.Invetory
	exists := false
	for i := 0; i < len(resources); i++ {
		if resources[i].Name == resourceName {
			exists = true
		}
	}
	return exists
}

func(k Keeper) updateMineResourceAmountByName(ctx sdk.Context,key string,resourceName string,resourceAmount int){
	mine,_ := k.GetMine(ctx,key)
	resources := mine.Invetory
	for i := 0; i < len(resources); i++ {
		if resources[i].Name == resourceName {
				resources[i].Amount += resourceAmount
		}
	}
	k.SetMine(ctx,mine)
}

func(k Keeper) checkResourceAmount(ctx sdk.Context,key string,resourceName string,resourceAmount int)bool{
	mine,_ := k.GetMine(ctx,key)
	resources := mine.Invetory
	for i := 0; i < len(resources); i++ {
		if resources[i].Name == resourceName {
				return resources[i].Amount >= resourceAmount
		}
	}
	return false
} 