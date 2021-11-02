package cider

import (
	"errors"
	"strconv"
	"strings"
)

// CheckFormat checks whether the format of
// the input value matches the format of CIDER.
func CheckFormat(cider string) ([]int, int, error) {
	addresses := strings.Split(cider, ".")
	if len(addresses) != 4 {
		return nil, 0, errors.New("error: format error")
	}

	buf := strings.Split(addresses[len(addresses)-1], "/")
	if len(buf) != 2 {
		return nil, 0, errors.New("error: format error")
	}
	addresses[len(addresses)-1] = buf[0]
	block := buf[1]

	addr := make([]int, 4)
	for index := 0; index < len(addresses); index++ {
		v, err := strconv.Atoi(addresses[index])
		if err != nil {
			return nil, 0, errors.New("error: not number")
		}
		if v < 0 || 255 < v {
			return nil, 0, errors.New("error: out of range")
		}
		addr[index] = v
	}
	b, err := strconv.Atoi(block)
	if err != nil {
		return nil, 0, errors.New("error: not number")
	}
	if b < 0 || 32 < b {
		return nil, 0, errors.New("error: out of range")
	}
	return addr, b, nil
}
