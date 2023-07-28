package http_util

import "strings"

func CheckHeader(headerToken string) bool {
	if headerToken == "" {
		return false
	}
	tokenArr := strings.Split(headerToken, " ")
	token := tokenArr[1]
	if token == "" {
		return false
	}

	return true
}
