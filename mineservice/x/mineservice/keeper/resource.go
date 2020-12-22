package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// GetResourceCount get the total number of resource
func (k Keeper) GetResourceCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.ResourceCountPrefix)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetResourceCount set the total number of resource
func (k Keeper) SetResourceCount(ctx sdk.Context, count int64)  {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.ResourceCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateResource creates a resource
func (k Keeper) CreateResource(ctx sdk.Context, resource types.Resource) {
	// Create the resource
	count := k.GetResourceCount(ctx)
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ResourcePrefix + resource.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(resource)
	store.Set(key, value)

	// Update resource count
    k.SetResourceCount(ctx, count+1)
}

// GetResource returns the resource information
func (k Keeper) GetResource(ctx sdk.Context, key string) (types.Resource, error) {
	store := ctx.KVStore(k.storeKey)
	var resource types.Resource
	byteKey := []byte(types.ResourcePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &resource)
	if err != nil {
		return resource, err
	}
	return resource, nil
}

// SetResource sets a resource
func (k Keeper) SetResource(ctx sdk.Context, resource types.Resource) {
	resourceKey := resource.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(resource)
	key := []byte(types.ResourcePrefix + resourceKey)
	store.Set(key, bz)
}

// DeleteResource deletes a resource
func (k Keeper) DeleteResource(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.ResourcePrefix + key))
}

//
// Functions used by querier
//

func listResource(ctx sdk.Context, k Keeper) ([]byte, error) {
	var resourceList []types.Resource
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ResourcePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var resource types.Resource
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &resource)
		resourceList = append(resourceList, resource)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, resourceList)
	return res, nil
}

func getResource(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	resource, err := k.GetResource(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, resource)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetResourceOwner(ctx sdk.Context, key string) sdk.AccAddress {
	resource, err := k.GetResource(ctx, key)
	if err != nil {
		return nil
	}
	return resource.Owner
}


// Check if the key exists in the store
func (k Keeper) ResourceExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.ResourcePrefix + key))
}
