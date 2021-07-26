package dumpCmd

import (
	"errors"
	"fmt"

	"github.com/dmorton-spirent/op-cli/cmd/utils"
	"github.com/spf13/cobra"
	intf "github.com/spirent/openperf/api/client/golang/client/interfaces"
)

// Name clients will use on the command line.
var dumpInterfaceCommandName = "interface"

// InterfacesCmd represents the Interfaces command
var dumpInterfaceCmd = &cobra.Command{
	Use:   dumpInterfaceCommandName,
	Short: "Output an interface in JSON format, by ID",
	Long: `Retrieves a given interface resource from OpenPerf
and outputs as JSON to standard out.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires exactly one interface ID")
		}

		ok := utils.ValidateID(args[0])

		return ok
	},
	Run: func(cmd *cobra.Command, args []string) {
		opclient := utils.OPClientConnection(opURL)

		getIntfParams := intf.NewGetInterfaceParams()

		getIntfParams.SetID(args[0])

		port, ok := opclient.Interfaces.GetInterface(getIntfParams)
		if ok != nil {
			utils.ClassifyAPIError(ok, "interface", args[0])
			return
		}

		buf, _ := port.GetPayload().MarshalBinary()
		fmt.Println(string(buf))

	},
}

func init() {
	dumpCmd.AddCommand(dumpInterfaceCmd)
}
