package validateCmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Name clients will use on the command line.
var jsonCommandName = "json"

// InterfacesCmd represents the Interfaces command
var jsonCommand = &cobra.Command{
	Use:   jsonCommandName,
	Short: "Validate that input is valid JSON",
	Long:  `Reads in a string of characters and checks if it's valid JSON.`,
	Run: func(cmd *cobra.Command, args []string) {

		input := openAndReadInput()

		validJSON := json.Valid(input)

		if !validJSON {
			if !silentOutput {
				fmt.Println("Invalid JSON")
			}
			os.Exit(invalidInput)
		}

		if !silentOutput {
			fmt.Println("JSON Valid")
		}

	},
}

func init() {
	validateCmd.AddCommand(jsonCommand)
}
