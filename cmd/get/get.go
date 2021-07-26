package getCmd

import (
	"github.com/dmorton-spirent/op-cli/cmd/utils"
	"github.com/spf13/cobra"
)

var opURL string
var opURLFlagName string

// getCmdName represents the <verb> command name for the root command.
// Making it a parameter for readability vs embedding in getCmd declaration below.
var getCmdName = "get"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   getCmdName,
	Short: "Get details about a given resource, by ID",
	Long: `Get configuration, statistics, status or other details about a given resource.
Resources are specified by their ID. Different resource types have different
flags indicating which details user wants to see.

See help output for individual resource types for a complete listing of flags.
Common flags include:
    -s  resource statistics
    -c  resource configuration
    -t  resource status`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		opURL, err = utils.OPRemoteURL(cmd, opURLFlagName)
		if err != nil {
			return err
		}

		return nil
	},
}

// Register is called by the root command to give this command a chance to
// register itself with the root command. This also passes down the name of
// the flag specifying the target OpenPerf's hostname and API port.
func Register(rootCommand *cobra.Command, opFlagName string) {
	opURLFlagName = opFlagName
	rootCommand.AddCommand(getCmd)
}
