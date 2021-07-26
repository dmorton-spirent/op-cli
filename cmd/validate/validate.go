package validateCmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// Name clients will use on the command line.
var validateCommandName = "validate"

var silentOutput bool
var inputFileName string

const (
	ok = iota
	noFile
	ioError
	invalidInput
)

var validateCmd = &cobra.Command{
	Use:   validateCommandName,
	Short: "Validate encoding of input based on user-specified type",
	Long: `Validate reads a string from stdin and determines if it is
properly encoded. User must specify which format to validate against.
Note: the validation is not against the OpenPerf API. Rather, op-cli leverages Go's
built-in decoding support for different encoding schemes.

Use this command as a quick sanity check for correctness before POSTing anything to OpenPerf.

Returns:
0 - format correct
1 - no input file
2 - IO error
3 - Invalid Input`,
}

func Register(rootCommand *cobra.Command, opFlagName string) {
	rootCommand.AddCommand(validateCmd)
}

func init() {
	validateCmd.PersistentFlags().BoolVarP(&silentOutput, "silent", "s", false, "Silent mode. No output other than return code.")
	validateCmd.PersistentFlags().StringVarP(&inputFileName, "file", "f", "", "Read from file instead of stdin")
}

// openAndReadInput determines where to get input from, reads it and returns
// it as an array of bytes. Calls os.Exit on errors.
// Calling os.Exit allows for fine-grain error reporting as well as
// different return codes to the shell based on what sort of error occurred.
// Common to all validate subcommands; putting here for "ownership".
func openAndReadInput() []byte {

	// Make sure there's input waiting for us on stdin or an input file to read.
	stat, _ := os.Stdin.Stat()
	if os.ModeCharDevice&stat.Mode() != 0 && inputFileName == "" {
		// No input is not considered an error.
		os.Exit(ok)
	}

	// Open our input source. -f flag takes precedence over stdin.
	var inFile *os.File
	if inputFileName != "" {
		var err error
		inFile, err = os.Open(inputFileName)
		if err != nil {
			if !silentOutput {
				fmt.Printf("error while opening file %s: %s\n", inputFileName, err)
			}
			os.Exit(noFile)
		}
	} else {
		inFile = os.Stdin
	}

	// Read in file.
	input, err := io.ReadAll(inFile)
	if err != nil {
		if !silentOutput {
			fmt.Printf("error reading input: %s\n", err)
		}
		os.Exit(ioError)
	}

	return input
}
