package utils

import (
	"fmt"

	"github.com/spf13/cobra"
)

func OPRemoteURL(cmd *cobra.Command, opURLFlagName string) (string, error) {
	opHostFlag := cmd.Flag(opURLFlagName)
	if opHostFlag == nil {
		return "", fmt.Errorf("ERROR getting remote flag")
	}

	return opHostFlag.Value.String(), nil
}
