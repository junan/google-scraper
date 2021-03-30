package helpers

import (
	"strconv"
)

func StringToInt(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func IntToString(int int64) string {
	return strconv.FormatInt(int, 10)
}
