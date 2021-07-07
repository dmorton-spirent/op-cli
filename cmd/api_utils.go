package cmd

import (
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/go-openapi/runtime"
	"github.com/niemeyer/pretty"
	client "github.com/spirent/openperf/api/client/golang/client"
)

func ClassifyAPIError(err error, resourceType string, resourceID string) {

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

	fmt.Println("Error occurred while communicating with OpenPerf")
}

func OPClientConnection() *client.OpenPerfAPI {
	tcConfig := client.TransportConfig{
		Host:    OPHost,
		Schemes: []string{"http"},
	}
	//tcConfig.Host = url
	//pretty.Println(tcConfig)
	return client.NewHTTPClientWithConfig(nil, &tcConfig)
}
