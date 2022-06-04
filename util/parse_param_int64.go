package util

import "strconv"

const (
	_base    = 10
	_bitSize = 64
)

// ParseParamInt64 parse string value into int64
func ParseParamInt64(input string) (int64, error) {

	num, err := strconv.ParseInt(input, _base, _bitSize)
	if err != nil {
		return 0, err
	}

	return num, nil
}
