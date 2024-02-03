package service

import "strconv"

func StringToUint64(data string) (uint64, error) {
	res, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func IsInteger(data interface{}) bool {
	switch data.(type) {
	case int, int8, int16, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case string:
		return true
	default:
		return false
	}
}
