package service

import "strconv"

func StringToUint64(data string) (uint64, error) {
	res, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}
