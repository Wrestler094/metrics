package utils

import "strings"

func ValidateFlags(flagPollInterval *int64, flagReportInterval *int64, flagServerAddress *string) {
	if *flagPollInterval < 1 {
		*flagPollInterval = 2
	}

	if *flagReportInterval < 1 {
		*flagReportInterval = 10
	}

	if !(strings.HasPrefix(*flagServerAddress, "http://")) {
		*flagServerAddress = "http://" + *flagServerAddress
	}
}
