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
	m := buf[1]

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
	mask, err := strconv.Atoi(m)
	if err != nil {
		return nil, 0, errors.New("error: not number")
	}
	if mask < 0 || 32 < mask {
		return nil, 0, errors.New("error: out of range")
	}
	return addr, mask, nil
}

// GetNetworkAddress retrieves a network address
// from an IP address and subnet mask
func GetNetworkAddress(addr []int, subnet []int) []int {

	naddr := make([]int, 4)

	for index := 0; index < len(naddr); index++ {
		noctet := addr[index]
		noctet &= subnet[index]
		naddr[index] = noctet
	}
	return naddr
}

// GetBroadcastAddress retrieves a broadcast address
// from an IP address and subnet mask
func GetBroadcastAddress(addr []int, subnet []int) []int {
	baddr := make([]int, 4)

	for index := 0; index < len(baddr); index++ {
		boctet := addr[index]
		for bit := 0; bit < 8; bit++ {
			sbit := subnet[index] & (0x01 << bit)
			if sbit == 0 {
				boctet |= (0x01 << bit)
			}
		}
		baddr[index] = boctet
	}
	return baddr
}

// GetSubnetmask returns a slice of the subnet mask
func GetSubnetmask(mask int) []int {
	maskbit := make([]int, 4)
	index := 0
	bitofset := 0
	for ofset := 0; ofset < mask; ofset++ {
		maskbit[index] |= (0x01 << (7 - bitofset))
		if bitofset < 7 {
			bitofset++
		} else {
			bitofset = 0
			index++
		}
	}
	return maskbit
}
