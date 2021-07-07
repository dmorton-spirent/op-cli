package cmd

import (
	"fmt"

	"github.com/niemeyer/pretty"
	"github.com/spf13/cobra"
	"github.com/spirent/openperf/api/client/golang/client"
)

// InterfacesCmd represents the Interfaces command
var listInterfacesCmd = &cobra.Command{
	Use:   "interfaces", // this is the name clients will use from the cli!
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListInterfaces called")

		//pf := cmd.Parent().Flag("url")
		//url := pf.Value.String()
		fmt.Printf("urll: %s\n", OPHost)

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

		interfaceList, ok := opclient.Interfaces.ListInterfaces(nil)
		if ok != nil {
			fmt.Println("Got error connecting to OP!")
			pretty.Println(ok)
			return
		}

		//pretty.Println(interfaceList)
		for _, intf := range interfaceList.Payload {
			fmt.Println(*intf.ID)
		}

	},
}

func init() {
	listCmd.AddCommand(listInterfacesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// InterfacesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// InterfacesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
