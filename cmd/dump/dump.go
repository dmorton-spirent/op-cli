package dumpCmd

import (
	"github.com/dmorton-spirent/op-cli/cmd/utils"
	"github.com/spf13/cobra"
)

// Name clients will use on the command line.
var dumpCommandName = "dump"

var opURL string
var opURLFlagName string

var dumpCmd = &cobra.Command{
	Use:   dumpCommandName,
	Short: "Output a resource in JSON format",
	Long: `Output a specified resource (by type and ID) in JSON format.
Useful as a "cheat sheet" for proper JSON formatting for OpenPerf resources.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		opURL, err = utils.OPRemoteURL(cmd, opURLFlagName)
		if err != nil {
			return err
		}
		return nil
	},
}

func Register(rootCommand *cobra.Command, opFlagName string) {
	opURLFlagName = opFlagName
	rootCommand.AddCommand(dumpCmd)

}
