package validateCmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// Name clients will use on the command line.
var yamlCommandName = "yaml"

var yamlCommand = &cobra.Command{
	Use:   yamlCommandName,
	Short: "Validate that input is valid YAML",
	Long:  `Reads in a string of characters and checks if it's valid YAML.`,
	Run: func(cmd *cobra.Command, args []string) {

		// get input from requested source.
		input := openAndReadInput()

		// YAML doesn't support a .Valid() function like JSON does.
		// So unmarshal it into a generic map and check the return value from
		// yaml.Unmarshal().
		var unmarshaledYAML = make(map[interface{}]interface{})
		ok := yaml.Unmarshal(input, &unmarshaledYAML)

		if ok != nil {
			if !silentOutput {
				fmt.Println("Invalid YAML")
			}
			os.Exit(invalidInput)
		}

		if !silentOutput {
			fmt.Println("YAML Valid")
		}

	},
}

func init() {
	validateCmd.AddCommand(yamlCommand)
}
