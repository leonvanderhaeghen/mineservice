package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	mineserviceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	mineserviceTxCmd.AddCommand(flags.PostCommands(
	// this line is used by starport scaffolding # 1
		GetCmdCreatePlayer(cdc),
		GetCmdSetPlayer(cdc),
		GetCmdDeletePlayer(cdc),
		GetCmdCreateResource(cdc),
		GetCmdSetResource(cdc),
		GetCmdDeleteResource(cdc),
		GetCmdSellMine(cdc),
		GetCmdCreateMine(cdc),
		GetCmdBuyMine(cdc),
		GetCmdSetMine(cdc),
		GetCmdDeleteMine(cdc),
		GetCmdMoveResource(cdc),
	)...)

	return mineserviceTxCmd
}
