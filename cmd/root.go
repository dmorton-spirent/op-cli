package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	dumpCmd "github.com/dmorton-spirent/op-cli/cmd/dump"
	getCmd "github.com/dmorton-spirent/op-cli/cmd/get"
	listCmd "github.com/dmorton-spirent/op-cli/cmd/list"
	validateCmd "github.com/dmorton-spirent/op-cli/cmd/validate"
)

var cfgFile string

// OPHost IP address/hostname and API port number of OpenPerf.
var OPHost string

var opHostFlagName = "remote"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "op-cli",
	Short: "A command-line interface for OpenPerf",
	Long: `OpenPerf (github.com/Spirent/openperf) is an infrastructure
and application test and analysis framework.
op-cli is a utility to interact with OpenPerf's REST API from the command line.
It aims to simplify common interactions with OpenPerf, as well as provide
utility functions to aid integrators of OpenPerf.
op-cli follows the Go and Docker model of <program> <verb> <noun> to avoid an
excess of CLI flags while maintaining a single binary for all tools.

Example Usage:
op-cli list ports                 - List all ports associated with an OpenPerf instance.
op-cli get interface Interface0   - Get statistics for OpenPerf emulated interface Interface0.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.op-cli.yaml)")

	rootCmd.PersistentFlags().StringVarP(&OPHost, opHostFlagName, "r", "localhost:9000", "host and API port for OpenPerf")

	// Register subcommands. Using the init() procedure results in import loops.
	dumpCmd.Register(rootCmd, opHostFlagName)
	getCmd.Register(rootCmd, opHostFlagName)
	listCmd.Register(rootCmd, opHostFlagName)
	validateCmd.Register(rootCmd, opHostFlagName)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".op-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".op-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
