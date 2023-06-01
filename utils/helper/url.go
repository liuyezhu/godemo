package helper

import (
	"fmt"
	"strings"
)

func HttpBuiltQuery(params map[string]string) (parseUrl string) {
	paramsArr := make([]string, 0, len(params))
	for k, v := range params {
		paramsArr = append(paramsArr, fmt.Sprintf("%s=%s", k, v))
	}
	parseUrl = strings.Join(paramsArr, "&")
	return parseUrl
}
