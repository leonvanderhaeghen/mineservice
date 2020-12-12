package cli

import (
	"bufio"
    
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
)

func GetCmdCreateMine(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-mine [name] [price] [owner] [selling] [efficiency] [invetory] [resources] [uraniumCost]",
		Short: "Creates a new mine",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsName := string(args[0] )
			argsPrice := string(args[1] )
			argsOwner := string(args[2] )
			argsSelling := string(args[3] )
			argsEfficiency := string(args[4] )
			argsInvetory := string(args[5] )
			argsResources := string(args[6] )
			argsUraniumCost := string(args[7] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateMine(cliCtx.GetFromAddress(), string(argsName), string(argsPrice), string(argsOwner), string(argsSelling), string(argsEfficiency), string(argsInvetory), string(argsResources), string(argsUraniumCost))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetMine(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-mine [id]  [name] [price] [owner] [selling] [efficiency] [invetory] [resources] [uraniumCost]",
		Short: "Set a new mine",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsName := string(args[1])
			argsPrice := string(args[2])
			argsOwner := string(args[3])
			argsSelling := string(args[4])
			argsEfficiency := string(args[5])
			argsInvetory := string(args[6])
			argsResources := string(args[7])
			argsUraniumCost := string(args[8])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetMine(cliCtx.GetFromAddress(), id, string(argsName), string(argsPrice), string(argsOwner), string(argsSelling), string(argsEfficiency), string(argsInvetory), string(argsResources), string(argsUraniumCost))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteMine(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-mine [id]",
		Short: "Delete a new mine by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteMine(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
