package mineservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/keeper"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
	// abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	// TODO: Define logic for when you would like to initalize a new genesis
	for _, record := range data.MineRecords {
		k.SetMine(ctx, record)
	}
	for _, record := range data.ResourceRecords {
		k.SetResource(ctx, record)
	}
		for _, record := range data.PlayerRecords {
		k.SetPlayer(ctx, record)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (data types.GenesisState) {
	// TODO: Define logic for exporting state
	var mineRecords []types.Mine
	var resourceRecords []types.Resource
	var playerRecords []types.Player

	mineIterator := k.GetMinesIterator(ctx)
	for ; mineIterator.Valid(); mineIterator.Next() {

		key := string(mineIterator.Key())
		mine, _ := k.GetMine(ctx, key)
		mineRecords = append(mineRecords, mine)

	}
	resourceIterator := k.GetResourcesIterator(ctx)
	for ; resourceIterator.Valid(); resourceIterator.Next() {

		key := string(resourceIterator.Key())
		resource, _ := k.GetResource(ctx, key)
		resourceRecords = append(resourceRecords, resource)

	}
	playerIterator := k.GetPlayersIterator(ctx)
	for ; playerIterator.Valid(); playerIterator.Next() {

		key := string(playerIterator.Key())
		player, _ := k.GetPlayer(ctx, key)
		playerRecords = append(playerRecords, player)

	}
	return types.GenesisState{MineRecords: mineRecords,PlayerRecords: playerRecords,ResourceRecords: resourceRecords}
}
