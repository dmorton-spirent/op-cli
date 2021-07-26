package dumpCmd

import (
	"errors"
	"fmt"

	"github.com/dmorton-spirent/op-cli/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spirent/openperf/api/client/golang/client/ports"
)

// Name clients will use on the command line.
var dumpPortCommandName = "port"

var dumpPortCmd = &cobra.Command{
	Use:   dumpPortCommandName,
	Short: "Output a port in JSON format, by ID",
	Long: `Retrieves a given port resource from OpenPerf
and outputs as JSON to standard out.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires exactly one port ID")
		}

		ok := utils.ValidateID(args[0])

		return ok
	},
	Run: func(cmd *cobra.Command, args []string) {
		opclient := utils.OPClientConnection(opURL)

		getPortParams := ports.NewGetPortParams()

		getPortParams.SetID(args[0])

		port, ok := opclient.Ports.GetPort(getPortParams)
		if ok != nil {
			utils.ClassifyAPIError(ok, "port", args[0])
			return
		}

		buf, _ := port.GetPayload().MarshalBinary()
		fmt.Println(string(buf))
	},
}

func init() {
	dumpCmd.AddCommand(dumpPortCmd)
}
