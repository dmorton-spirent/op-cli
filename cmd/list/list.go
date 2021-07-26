package listCmd

import (
	"github.com/dmorton-spirent/op-cli/cmd/utils"
	"github.com/spf13/cobra"
)

var opURL string
var opURLFlagName string

// listCmdName represents the <verb> command name for the root command.
// Making it a parameter for readability vs embedding in listCmd declaration below.
var listCmdName = "list"

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   listCmdName,
	Short: "List IDs of all resources of a given type",
	Long: `List retrieves all IDs for a given resource type.
IDs are printed one per line to stdout.`,
	Example: "op-cli list resource_type",
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
	rootCommand.AddCommand(listCmd)
}
