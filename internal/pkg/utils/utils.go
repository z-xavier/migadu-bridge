package utils

import (
	"fmt"
	"strings"
)

func IsMigaduHttpErr(err error, code int, errMsg string) bool {
	errStr := err.Error()
	return strings.Contains(errStr, errMsg) &&
		strings.Contains(errStr, fmt.Sprintf("status code %d", code))
}
