package listCmd

import (
	"fmt"

	"github.com/dmorton-spirent/op-cli/cmd/utils"
	"github.com/spf13/cobra"
)

var listInterfacesCmdName = "interfaces"

// InterfacesCmd represents the Interfaces command
var listInterfacesCmd = &cobra.Command{
	Use:   listInterfacesCmdName,
	Short: "List IDs of all interface resources",
	Long:  `List IDs of all resources from /interfaces endpoint.`,
	Run: func(cmd *cobra.Command, args []string) {
		opclient := utils.OPClientConnection(opURL)

		interfaceList, ok := opclient.Interfaces.ListInterfaces(nil)
		if ok != nil {
			utils.ClassifyAPIError(ok, "interfaces", "")
			return
		}

		for _, intf := range interfaceList.Payload {
			fmt.Println(*intf.ID)
		}

	},
}

func init() {
	listCmd.AddCommand(listInterfacesCmd)
}
