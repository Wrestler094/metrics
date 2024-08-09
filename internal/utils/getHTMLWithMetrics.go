package utils

import (
	"strconv"
	"strings"
)

const begin = `<html>
	<head>
		<title>Metrics list</title>
	</head>
	<body>
		<ul>`

const end = `</ul>
	</body>
</html>`

func GetHTMLWithMetrics(gaugeMetrics map[string]float64, counterMetrics map[string]int64) string {
	var middle = ""

	for k, v := range gaugeMetrics {
		middle += "<li>" + k + " - " + strconv.FormatFloat(v, 'f', 2, 64) + "</li>"
	}

	for k, v := range counterMetrics {
		middle += "<li>" + k + " - " + strconv.FormatInt(v, 10) + "</li>"
	}

	return begin + strings.TrimSuffix(middle, ",") + end
}
