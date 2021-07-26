package listCmd

import (
	"fmt"

	"github.com/dmorton-spirent/op-cli/cmd/utils"
	"github.com/spf13/cobra"
)

var listPortsCmdName = "ports"

// listPortsCmd represents the command to list /ports resources.
var listPortsCmd = &cobra.Command{
	Use:   listPortsCmdName,
	Short: "List IDs of all port resources",
	Long:  `List IDs of all resources from /ports endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
		opclient := utils.OPClientConnection(opURL)

		portList, ok := opclient.Ports.ListPorts(nil)
		if ok != nil {
			utils.ClassifyAPIError(ok, "interfaces", "")
			return
		}

		for _, intf := range portList.Payload {
			fmt.Println(*intf.ID)
		}

	},
}

func init() {
	listCmd.AddCommand(listPortsCmd)
}
