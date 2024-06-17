package utils

import (
	"fmt"
	"strings"
)

func ValidateServerAddress(serverAddress *string) {
	fmt.Println(serverAddress)
}

func ValidateFlags(pollInterval *int64, reportInterval *int64, serverAddress *string) {
	if *pollInterval < 1 {
		*pollInterval = 2
	}

	if *reportInterval < 1 {
		*reportInterval = 10
	}

	if !(strings.HasPrefix(*serverAddress, "http://")) {
		*serverAddress = "http://" + *serverAddress
	}
}
