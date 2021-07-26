package getCmd

import (
	"errors"
	"fmt"

	"github.com/dmorton-spirent/op-cli/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spirent/openperf/api/client/golang/client/interfaces"
	"github.com/spirent/openperf/api/client/golang/models"
)

var printInterfaceConfigFlag, printInterfaceStatsFlag bool

var getInterfaceCmdName = "interface"

var getInterfaceCmd = &cobra.Command{
	Use:   getInterfaceCmdName,
	Short: "Get interface details, by ID",
	Long:  `Get additional details about an interface, by ID.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires exactly one interface ID")
		}

		ok := utils.ValidateID(args[0])

		return ok
	},
	Run: func(cmd *cobra.Command, args []string) {
		opclient := utils.OPClientConnection(opURL)

		getIntfParams := interfaces.NewGetInterfaceParams()

		getIntfParams.SetID(args[0])

		intf, ok := opclient.Interfaces.GetInterface(getIntfParams)
		if ok != nil {
			utils.ClassifyAPIError(ok, "interface", args[0])
			return
		}

		// Default to printing stats information if no
		// explicit output request.
		if !printInterfaceConfigFlag &&
			!printInterfaceStatsFlag {
			printInterfaceStats(intf.Payload.Stats)
		}

		if printInterfaceStatsFlag {
			printInterfaceStats(intf.Payload.Stats)
		}

		if printInterfaceConfigFlag {
			printInterfaceConfig(intf.Payload.Config)
		}

	},
}

func printInterfaceStats(stats *models.InterfaceStats) {
	fmt.Printf("RX packets %d bytes %d errors %d\n", *stats.RxPackets, *stats.RxBytes, *stats.RxErrors)
	fmt.Printf("TX packets %d bytes %d errors %d\n", *stats.TxPackets, *stats.TxBytes, *stats.TxErrors)
}

func printInterfaceConfig(config *models.InterfaceConfig) {
	for _, protocolConfig := range config.Protocols {

		printEthInterface(protocolConfig.Eth)
		printIpv4Interface(protocolConfig.IPV4)
		printIpv6Interface(protocolConfig.IPV6)
	}
}

func printEthInterface(intf *models.InterfaceProtocolConfigEth) {
	if intf == nil {
		return
	}

	fmt.Printf("Eth: %s\n", *intf.MacAddress)
}

func printIpv4Interface(intf *models.InterfaceProtocolConfigIPV4) {
	if intf == nil {
		return
	}

	if *intf.Method == "static" {
		printIpv4InterfaceStatic(intf.Static)
	}
}

func printIpv4InterfaceStatic(intf *models.InterfaceProtocolConfigIPV4Static) {
	if intf == nil {
		return
	}
}

func printIpv6Interface(intf *models.InterfaceProtocolConfigIPV6) {
	if intf == nil {
		return
	}

	fmt.Printf("Link Local: %s\n", intf.LinkLocalAddress)

	if *intf.Method == "static" {
		printIpv6InterfaceStatic(intf.Static)
	}
}

func printIpv6InterfaceStatic(intf *models.InterfaceProtocolConfigIPV6Static) {
	if intf == nil {
		return
	}

	fmt.Printf("IPv6 Address: %s\n", *intf.Address)
	fmt.Printf("IPv6 Gateway: %s\n", intf.Gateway)
	fmt.Printf("Prefix Length: %d\n", *intf.PrefixLength)
}

func init() {
	getCmd.AddCommand(getInterfaceCmd)

	getInterfaceCmd.Flags().BoolVarP(&printInterfaceStatsFlag, "stats", "s", false, "Get interface statistics counters")
	getInterfaceCmd.Flags().BoolVarP(&printInterfaceConfigFlag, "cfg", "c", false, "Get interface configuration parameters")
}
