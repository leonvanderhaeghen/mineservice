package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
)

// GetPlayerCount get the total number of player
func (k Keeper) GetPlayerCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.PlayerCountPrefix)
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

// SetPlayerCount set the total number of player
func (k Keeper) SetPlayerCount(ctx sdk.Context, count int64)  {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.PlayerCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreatePlayer creates a player
func (k Keeper) CreatePlayer(ctx sdk.Context, msg types.MsgCreatePlayer) {
	// Create the player
	count := k.GetPlayerCount(ctx)
    var player = types.Player{
        Creator: msg.Creator,
        ID:      strconv.FormatInt(count, 10),
        Name: msg.Name,
    }

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.PlayerPrefix + player.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(player)
	store.Set(key, value)

	// Update player count
    k.SetPlayerCount(ctx, count+1)
}

// GetPlayer returns the player information
func (k Keeper) GetPlayer(ctx sdk.Context, key string) (types.Player, error) {
	store := ctx.KVStore(k.storeKey)
	var player types.Player
	byteKey := []byte(types.PlayerPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &player)
	if err != nil {
		return player, err
	}
	return player, nil
}

// SetPlayer sets a player
func (k Keeper) SetPlayer(ctx sdk.Context, player types.Player) {
	playerKey := player.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(player)
	key := []byte(types.PlayerPrefix + playerKey)
	store.Set(key, bz)
}

// DeletePlayer deletes a player
func (k Keeper) DeletePlayer(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.PlayerPrefix + key))
}

//
// Functions used by querier
//

func listPlayer(ctx sdk.Context, k Keeper) ([]byte, error) {
	var playerList []types.Player
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PlayerPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var player types.Player
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &player)
		playerList = append(playerList, player)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, playerList)
	return res, nil
}

func getPlayer(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	player, err := k.GetPlayer(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, player)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetPlayerOwner(ctx sdk.Context, key string) sdk.AccAddress {
	player, err := k.GetPlayer(ctx, key)
	if err != nil {
		return nil
	}
	return player.Creator
}

// Check if the key exists in the store
func (k Keeper) PlayerExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.PlayerPrefix + key))
}


func (k Keeper) AddMineToPlayer(ctx sdk.Context, key string,mineID string){
	player,_ := k.GetPlayer(ctx, key)
	mine,_ :=k.GetMine(ctx,mineID)
	player.Mines = append(player.Mines,mine)
	k.SetPlayer(ctx,player)
}
func (k Keeper) RemoveMineFromPlayer(ctx sdk.Context, key string,mineID string){
	player,_ := k.GetPlayer(ctx, key)
	mineIndex := k.GetMineIndexFromPlayer(ctx,key,mineID)
	if mineIndex >=0 {
			player.Mines = removeMine(player.Mines,mineIndex)
			k.SetPlayer(ctx,player)
	}
	
}
func removeMine(slice []types.Mine, s int) []types.Mine {
    return append(slice[:s], slice[s+1:]...)
}

func (k Keeper) GetMineIndexFromPlayer(ctx sdk.Context, key string,mineID string)int{
	player,_ := k.GetPlayer(ctx, key)
	mine,_ :=k.GetMine(ctx,mineID)
	for i := 0; i < len(player.Mines); i++ {
		if player.Mines[i].ID == mine.ID {
			return i
		}
	}
	return -1
} 
func(k Keeper) addResourcePlayer(ctx sdk.Context,resource types.Resource,playerID string){
	player,_ := k.GetPlayer(ctx,playerID)
	if k.resourceExistsInPlayer(ctx,resource.MineID,resource.Name) {
		k.updatePlayerResourceAmountByName(ctx,playerID,resource.Name,resource.Amount)
	}else{
		resource.MineID = ""
		player.Invetory = append(player.Invetory, resource)
		k.SetPlayer(ctx,player)
	}
}

func(k Keeper) resourceExistsInPlayer(ctx sdk.Context,key string,resourceName string)bool{
	player,_ := k.GetPlayer(ctx,key)
	resources := player.Invetory
	exists := false
	for i := 0; i < len(resources); i++ {
		if resources[i].Name == resourceName {
			exists = true
		}
	}
	return exists
}

func(k Keeper) updatePlayerResourceAmountByName(ctx sdk.Context,key string,resourceName string,resourceAmount int){
	player,_ := k.GetPlayer(ctx,key)
	resources := player.Invetory
	for i := 0; i < len(resources); i++ {
		if resources[i].Name == resourceName {
				resources[i].Amount += resourceAmount
		}
	}
	k.SetPlayer(ctx,player)
}