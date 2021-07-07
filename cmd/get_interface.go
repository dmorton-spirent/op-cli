package cmd

import (
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/go-openapi/runtime"
	"github.com/niemeyer/pretty"
	"github.com/spf13/cobra"
	client "github.com/spirent/openperf/api/client/golang/client"
	"github.com/spirent/openperf/api/client/golang/client/interfaces"
	"github.com/spirent/openperf/api/client/golang/models"
)

// InterfacesCmd represents the Interfaces command
var getInterfaceCmd = &cobra.Command{
	Use:   "interface", // this is the name clients will use from the cli!
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires exactly one interface ID")
		}

		// FIXME: check arg against ID regex

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetInterface called")

		//pretty.Println(args)

		//pf := cmd.Parent().Flag("url")
		//urlStr := pf.Value.String()
		//fmt.Printf("urll: %s\n", url)

		//client.Interfaces.ListInterfaces()
		//tc_config := client.TransportConfig
		//tc_config := client.TransportConfig.WithHost(url)
		tcConfig := client.TransportConfig{
			Host:    OPHost,
			Schemes: []string{"http"},
		}
		//tcConfig.Host = url
		//pretty.Println(tcConfig)
		opclient := client.NewHTTPClientWithConfig(nil, &tcConfig)
		//pretty.Println(opclient)

		getIntfParams := interfaces.NewGetInterfaceParams()
		//getIntfParams := client.Interfaces.NewGetInterfaceParams()
		//getIntfParams := client.OpenPerfAPI.Interfaces.NewGetInterfaceParams()

		//getIntfParams := models.NewGetInterfaceParams()

		// Framework takes care of validating args for us.
		// If we get this far we have the correct number and format of them.
		getIntfParams.SetID(args[0])

		intf, ok := opclient.Interfaces.GetInterface(getIntfParams)
		if ok != nil {
			fmt.Println("Got error connecting to OP!")
			//pretty.Println(ok)
			var apiError *runtime.APIError
			if errors.As(ok, &apiError) {
				if apiError.Code == 404 {
					fmt.Printf("Cannot find interface: %s\n", args[0])
					return
				}
			}

			//fmt.Println(ok.Error())

			// FIXME: from here on down should be a generic
			// error handling function that's common to all commands.
			// So yeah move this to another file.
			var dnsError *net.DNSError
			if errors.As(ok, &dnsError) {
				pretty.Println(dnsError)
				fmt.Println(dnsError.Error())
				fmt.Println("Cannot resolve hostname.")
				return
			}

			var syscallErr *os.SyscallError
			if errors.As(ok, &syscallErr) {
				fmt.Println(syscallErr.Error())
				return
			}

			fmt.Println("Error occurred while communicating with OpenPerf")
			return
		}

		//pretty.Println(intf)
		// for _, intf := range interfaceList.Payload {
		// 	fmt.Println(*intf.ID)
		// }

		if cmd.Flag("stats").Value.String() == "true" {
			//pretty.Println(intf.Payload.Stats)
			printStats(intf.Payload.Stats)
		}

		if cmd.Flag("config").Value.String() == "true" {
			//pretty.Println(intf.Payload.Config)
			printConfig(intf.Payload.Config)
		}

		// Default to printing stats information if no
		// explicit output request.
		if cmd.Flag("stats").Value.String() == "false" &&
			cmd.Flag("config").Value.String() == "false" {
			printStats(intf.Payload.Stats)
		}

	},
}

func printStats(stats *models.InterfaceStats) {
	//pretty.Println(stats)
	fmt.Printf("RX packets %d bytes %d errors %d\n", *stats.RxPackets, *stats.RxBytes, *stats.RxErrors)
	fmt.Printf("TX packets %d bytes %d errors %d\n", *stats.TxPackets, *stats.TxBytes, *stats.TxErrors)
}

func printConfig(config *models.InterfaceConfig) {
	for _, protocolConfig := range config.Protocols {
		//pretty.Println(protocolConfig)
		//fmt.Println("")

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

// ipv4:
//  mode: {static, dhcp}
//  address
//  gateway
//  netmask

// ipv6:
//  mode: {static, dhcp, auto}
//  address
//  link-local
//  gateway
//  prefix length

// eth:
//  mac

// &models.InterfaceProtocolConfig{
//     Eth: &models.InterfaceProtocolConfigEth{
//         MacAddress: &"00:10:94:ae:d6:aa",
//     },
//     IPV4: (*models.InterfaceProtocolConfigIPV4)(nil),
//     IPV6: (*models.InterfaceProtocolConfigIPV6)(nil),
// }

// &models.InterfaceProtocolConfig{
//     Eth:  (*models.InterfaceProtocolConfigEth)(nil),
//     IPV4: (*models.InterfaceProtocolConfigIPV4)(nil),
//     IPV6: &models.InterfaceProtocolConfigIPV6{
//         Auto:             (*models.InterfaceProtocolConfigIPV6Auto)(nil),
//         Dhcp6:            (*models.InterfaceProtocolConfigIPV6Dhcp6)(nil),
//         LinkLocalAddress: "fe80:1::202",
//         Method:           &"static",
//         Static:           &models.InterfaceProtocolConfigIPV6Static{
//             Address:      &"2001:2::202",
//             Gateway:      "2001:2::1",
//             PrefixLength: &int32(64),
//         },
//     },
// }

// func handleURLlError(urlErr *url.Error) {

// 	if netOpError, ok := urlErr.Err.(*net.OpError); ok {
// 		switch errType := netOpError.Err
// 	}

// 	return
// }

func init() {
	getCmd.AddCommand(getInterfaceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// InterfacesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// InterfacesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getInterfaceCmd.Flags().BoolP("stats", "s", false, "Get interface statistics counters")
	getInterfaceCmd.Flags().BoolP("config", "c", false, "Get interface configuration parameters")
}
