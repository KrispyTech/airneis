package helpers

import "strconv"

func ConvertStringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}
