package internal

import (
	"fmt"
	"strings"
)

func SimpleString(value string) string {
	return fmt.Sprintf("+%s\r\n", value)
}

func Error(message string) string {
	return fmt.Sprintf("-ERR %s\r\n", message)
}

func Integer(value int) string {
	return fmt.Sprintf(":%d\n", value)
}

func BulkString(value string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)
}

func BulkStringNil() string {
	return "$-1\r\n"
}

func Array(values []string) string {
	bulkValues := make([]string, len(values))
	for i, v := range values {
		bulkValues[i] = BulkString(v)
	}

	return strings.Join(bulkValues, "")
}
