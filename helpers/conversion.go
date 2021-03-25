package helpers

import (
	"strconv"
)

func StringToInt(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}
