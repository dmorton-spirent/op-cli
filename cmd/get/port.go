package getCmd

import (
	"errors"
	"fmt"

	"github.com/dmorton-spirent/op-cli/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spirent/openperf/api/client/golang/client/ports"
	"github.com/spirent/openperf/api/client/golang/models"
)

var printPortConfigFlag, printPortStatsFlag, printPortStatusFlag bool

// Name clients will use on the command line.
var commandName = "port"

var getPortCmd = &cobra.Command{
	Use:   commandName,
	Short: "Get port details, by ID",
	Long:  `Get additional details about a port, by ID.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires exactly one interface ID")
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

		// Did the user request any output?
		if !printPortConfigFlag && !printPortStatsFlag && !printPortStatusFlag {
			// No explicit request; fall back to printing stats by default.
			printPortStats(port.Payload.Stats)
			return
		}

		if printPortStatsFlag {
			printPortStats(port.Payload.Stats)
		}

		if printPortStatusFlag {
			printPortStatus(port.Payload.Status)
		}

		if printPortConfigFlag {
			printPortConfig(port.Payload.Config)
		}

	},
}

func printPortStats(stats *models.PortStats) {
	fmt.Printf("RX packets %d bytes %d errors %d\n", *stats.RxPackets, *stats.RxBytes, *stats.RxErrors)
	fmt.Printf("TX packets %d bytes %d errors %d\n", *stats.TxPackets, *stats.TxBytes, *stats.TxErrors)
}

func printPortStatus(status *models.PortStatus) {
	fmt.Printf("Link State: %s\n", *status.Link)
	fmt.Printf("Link Speed: %d\n", *status.Speed)
	fmt.Printf("Link Duplex: %s\n", *status.Duplex)
}

func printPortConfig(config *models.PortConfig) {
	maybePrintBondPort(config.Bond)
	maybePrintDpdkPort(config.Dpdk)
}

func maybePrintBondPort(cfg *models.PortConfigBond) {
	if cfg == nil {
		return
	}

	fmt.Printf("Mode: %s\n", *cfg.Mode)
	fmt.Printf("Member ports: ")
	for _, port := range cfg.Ports {
		fmt.Printf("%s", port)
	}
	fmt.Println("")
}

func maybePrintDpdkPort(cfg *models.PortConfigDpdk) {
	if cfg == nil {
		return
	}

	fmt.Printf("Device: %s\n", cfg.Device)
	fmt.Printf("Driver: %s\n", cfg.Driver)
	if cfg.Interface != "" {
		fmt.Printf("Interface: %s\n", cfg.Interface)
	}
	fmt.Printf("MAC Address: %s\n", cfg.MacAddress)
	fmt.Printf("Link Configuration: AutoNegotiation: %t, Duplex: %s, Speed: %dMbps\n", *cfg.Link.AutoNegotiation, cfg.Link.Duplex, cfg.Link.Speed)
}

func init() {
	getCmd.AddCommand(getPortCmd)

	getPortCmd.Flags().BoolVarP(&printPortStatsFlag, "stats", "s", false, "Get port statistics counters")
	getPortCmd.Flags().BoolVarP(&printPortConfigFlag, "cfg", "c", false, "Get port configuration parameters")
	getPortCmd.Flags().BoolVarP(&printPortStatusFlag, "status", "t", false, "Get port current status")
}
