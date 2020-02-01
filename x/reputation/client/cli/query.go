package cli

import (
	// "crypto/sha256"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/emag3m/reputation/x/reputation/internal/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group reputation queries under a subcommand
	reputationQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	reputationQueryCmd.AddCommand(
		client.GetCommands(
			GetCmdGetReputation(queryRoute, cdc),
		)...,
	)

	return reputationQueryCmd
}

func GetCmdGetReputation(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get [applicationID] [account]",
		Short: "Query a reputation by `ApplicationID + Account Address`",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			var storageID = []byte(args[0] + args[1])

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetReputation, storageID), nil)
			if err != nil {
				fmt.Printf("could not resolve reputation %s \n%s\n", err.Error())

				return nil
			}

			var out types.Reputation
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
