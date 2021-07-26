package utils

import (
	"errors"
	"fmt"
	"net"
	"os"
	"regexp"

	"github.com/go-openapi/runtime"
	"github.com/niemeyer/pretty"
	client "github.com/spirent/openperf/api/client/golang/client"
)

// IDRegEx regular expression that matches valid OpenPerf IDs
var IDRegEx = "^[a-z0-9-]+$"

func ClassifyAPIError(err error, resourceType string, resourceID string) {

	fmt.Println("Error occurred while communicating with OpenPerf:")

	var apiError *runtime.APIError
	if errors.As(err, &apiError) {
		if apiError.Code == 404 {
			fmt.Printf("Cannot find %s: %s\n", resourceType, resourceID)
			return
		}

		fmt.Printf("API error: %s", apiError.String())
		return
	}

	var dnsError *net.DNSError
	if errors.As(err, &dnsError) {
		pretty.Println(dnsError)
		fmt.Println(dnsError.Error())
		fmt.Println("Cannot resolve hostname.")
		return
	}

	var syscallErr *os.SyscallError
	if errors.As(err, &syscallErr) {
		fmt.Println(syscallErr.Error())
		return
	}
}

func OPClientConnection(opURL string) *client.OpenPerfAPI {
	tcConfig := client.TransportConfig{
		Host:    opURL,
		Schemes: []string{"http"},
	}

	return client.NewHTTPClientWithConfig(nil, &tcConfig)
}

func ValidateID(id string) error {
	matched, err := regexp.MatchString(IDRegEx, id)

	if err != nil {
		return err
	}

	if !matched {
		return errors.New("invalid ID format")
	}

	return nil
}
