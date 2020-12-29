package cli

import (
	"bufio"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
)

func GetCmdCreateResource(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-resource [name] [mineID] [amount]",
		Short: "Creates a new resource",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsName := string(args[0] )
			argsMineID := string(args[1])
			argsAmount,_ := strconv.Atoi(args[2])
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateResource(cliCtx.GetFromAddress(), string(argsName), string(argsMineID), argsAmount)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetResource(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-resource [id]  [name] [amount] [price] [selling] [minPrice]",
		Short: "Set a new resource",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsName := string(args[1])
			argsAmount,_ := strconv.Atoi(args[2])
			argsPrice := string(args[3])
			argsSelling := string(args[4])
			argsMinPrice := string(args[5])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetResource(cliCtx.GetFromAddress(), id, string(argsName),argsAmount, string(argsPrice), string(argsSelling), string(argsMinPrice))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteResource(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-resource [id]",
		Short: "Delete a new resource by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteResource(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdMoveResource(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "move-resource [id] [playerid] [amount]",
		Short: "Move a resource by ID from mine to player inv",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			id := args[0]
			argsPlayerID := string(args[1])
			argsAmount,_ := strconv.Atoi(args[2])

			msg := types.NewMsgMoveResource(id,argsPlayerID,argsAmount, cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}