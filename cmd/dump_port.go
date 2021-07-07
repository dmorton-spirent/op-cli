package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	client "github.com/spirent/openperf/api/client/golang/client"
	"github.com/spirent/openperf/api/client/golang/client/ports"
)

//var printConfigFlag, printStatsFlag, printStatusFlag bool

// Name clients will use on the command line.
var dumpPortCommandName = "port"

// InterfacesCmd represents the Interfaces command
var dumpPortCmd = &cobra.Command{
	Use:   dumpPortCommandName,
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires exactly one port ID")
		}

		// FIXME: check arg against ID regex

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("GetPort called")

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

		getPortParams := ports.NewGetPortParams()
		//getIntfParams := client.Interfaces.NewGetInterfaceParams()
		//getIntfParams := client.OpenPerfAPI.Interfaces.NewGetInterfaceParams()

		//getIntfParams := models.NewGetInterfaceParams()

		// Framework takes care of validating args for us.
		// If we get this far we have the correct number and format of them.
		getPortParams.SetID(args[0])

		port, ok := opclient.Ports.GetPort(getPortParams)
		if ok != nil {
			// var apiError *runtime.APIError
			// if errors.As(ok, &apiError) {
			// 	if apiError.Code == 404 {
			// 		fmt.Printf("Cannot find port: %s\n", args[0])
			// 		return
			// 	}

			// 	fmt.Printf("API error: %s", apiError.String())
			// 	return
			// }

			ClassifyAPIError(ok, "port", args[0])
			return
			//fmt.Println(ok.Error())

			// FIXME: from here on down should be a generic
			// error handling function that's common to all commands.
			// So yeah move this to another file.
			// var dnsError *net.DNSError
			// if errors.As(ok, &dnsError) {
			// 	pretty.Println(dnsError)
			// 	fmt.Println(dnsError.Error())
			// 	fmt.Println("Cannot resolve hostname.")
			// 	return
			// }

			// var syscallErr *os.SyscallError
			// if errors.As(ok, &syscallErr) {
			// 	fmt.Println(syscallErr.Error())
			// 	return
			// }

			// fmt.Println("Error occurred while communicating with OpenPerf")
			// return
		}

		buf, _ := port.GetPayload().MarshalBinary()
		//buf, _ := status.MarshalBinary()
		//fmt.Printf("%s\n", string(buf))
		fmt.Println(string(buf))

		// // Did the user request any output?
		// if !printConfigFlag && !printStatsFlag && !printStatusFlag {
		// 	// No explicit request; fall back to printing stats by default.
		// 	printPortStats(port.Payload.Stats)
		// 	return
		// }

		// //if cmd.Flag("stats").Value.String() == "true" {
		// if printStatsFlag {
		// 	//pretty.Println(intf.Payload.Stats)
		// 	printPortStats(port.Payload.Stats)
		// }

		// //if cmd.Flag("status").Value.String() == "true" {
		// if printStatusFlag {
		// 	printPortStatus(port.Payload.Status)
		// }

		// //if cmd.Flag("config").Value.String() == "true" {
		// if printConfigFlag {
		// 	//pretty.Println(intf.Payload.Config)
		// 	printPortConfig(port.Payload.Config)
		// }

	},
}

// func printPortStats(stats *models.PortStats) {
// 	fmt.Printf("RX packets %d bytes %d errors %d\n", *stats.RxPackets, *stats.RxBytes, *stats.RxErrors)
// 	fmt.Printf("TX packets %d bytes %d errors %d\n", *stats.TxPackets, *stats.TxBytes, *stats.TxErrors)
// }

// func printPortStatus(status *models.PortStatus) {
// 	fmt.Printf("Link State: %s\n", *status.Link)
// 	fmt.Printf("Link Speed: %d\n", *status.Speed)
// 	fmt.Printf("Link Duplex: %s\n", *status.Duplex)

// FIXME: use this if user wants JSON output instead of text.
// Cool, BUT if getting multiple things will get blobs of JSON per thing. Not super useful.
// If you need JSON output just use curl :).
// OR implement dump as a top-level command that does what curl does. Might be useful in case the
// Yocto 3.0 version of curl doesn't support unix domain sockets...

// buf, _ := status.MarshalBinary()
// fmt.Printf("marshalled: %s\n", string(buf))
//}

// func printPortConfig(config *models.PortConfig) {
// 	maybePrintBondPort(config.Bond)
// 	maybePrintDpdkPort(config.Dpdk)
// }

// func maybePrintBondPort(cfg *models.PortConfigBond) {
// 	if cfg == nil {
// 		return
// 	}

// 	fmt.Printf("Mode: %s\n", *cfg.Mode)
// 	fmt.Printf("Member ports: ")
// 	for _, port := range cfg.Ports {
// 		fmt.Printf("%s", port)
// 	}
// 	fmt.Println("")
// }

// func maybePrintDpdkPort(cfg *models.PortConfigDpdk) {
// 	if cfg == nil {
// 		return
// 	}

// 	fmt.Printf("Device: %s\n", cfg.Device)
// 	fmt.Printf("Driver: %s\n", cfg.Driver)
// 	if cfg.Interface != "" {
// 		fmt.Printf("Interface: %s\n", cfg.Interface)
// 	}
// 	fmt.Printf("MAC Address: %s\n", cfg.MacAddress)
// 	fmt.Printf("Link Configuration: AutoNegotiation: %t, Duplex: %s, Speed: %dMbps\n", *cfg.Link.AutoNegotiation, cfg.Link.Duplex, cfg.Link.Speed)
// }

func init() {
	dumpCmd.AddCommand(dumpPortCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// InterfacesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// InterfacesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// getPortCmd.Flags().BoolVarP(&printStatsFlag, "stats", "s", false, "Get port statistics counters")
	// getPortCmd.Flags().BoolVarP(&printConfigFlag, "cfg", "c", false, "Get port configuration parameters")
	// getPortCmd.Flags().BoolVarP(&printStatusFlag, "status", "t", false, "Get port current status")
}
