package cider

import (
	"reflect"
	"testing"
)

func Test_CheckFormat(t *testing.T) {
	type result struct {
		addr     []int
		block    int
		notError bool
	}
	type testdata struct {
		param string
		res   result
	}
	params := []testdata{
		{"0.0.0.0/0", result{[]int{0, 0, 0, 0}, 0, true}},
		{"000.000.000.000/00", result{[]int{0, 0, 0, 0}, 0, true}},
		{"255.255.255.255/32", result{[]int{255, 255, 255, 255}, 32, true}},
		{"255.255.255.255/33", result{nil, 0, false}},
		{"X,X,X,X/X", result{nil, 0, false}},
	}

	for ii, p := range params {
		addr, block, err := CheckFormat(p.param)

		for index := 0; index < len(addr); index++ {
			if p.res.addr[index] != addr[index] {
				t.Errorf("address mismatch. testindex:%d, expect:%d, test:%d", ii, p.res.addr[index], addr[index])
			}
		}
		if p.res.block != block {
			t.Errorf("block mismatch. testindex:%d, expect:%d, test:%d", ii, p.res.block, block)
		}

		if p.res.notError {
			if err != nil {
				t.Errorf("error mismatch. testindex:%d", ii)
			}
		} else {
			if err == nil {
				t.Errorf("error mismatch. testindex:%d", ii)
			}
		}
	}
}

func Test_getSubnetmask(t *testing.T) {
	type testdata struct {
		param  int
		expect []int
	}

	tests := []testdata{
		{32, []int{255, 255, 255, 255}},
		{0, []int{0, 0, 0, 0}},
		{15, []int{255, 254, 0, 0}},
	}

	for index, test := range tests {
		res := GetSubnetmask(test.param)

		if false == reflect.DeepEqual(res, test.expect) {
			t.Errorf("error subnet mismatch. testindex:%d, expect:%d, test:%d", index, test.expect, res)
		}
	}
}

func Test_GetNetworkAddress(t *testing.T) {
	type testdata struct {
		addr   []int
		subnet []int
		expect []int
	}

	tests := []testdata{
		{
			[]int{192, 168, 10, 111},
			[]int{255, 255, 255, 0},
			[]int{192, 168, 10, 0},
		},
		{
			[]int{192, 168, 255, 111},
			[]int{255, 255, 128, 0},
			[]int{192, 168, 128, 0},
		},
	}

	for index, test := range tests {
		naddr := GetNetworkAddress(test.addr, test.subnet)

		if false == reflect.DeepEqual(naddr, test.expect) {
			t.Errorf("error subnet mismatch. testindex:%d, expect:%d, test:%d", index, test.expect, naddr)
		}
	}
}

func Test_GetBroadcastAddress(t *testing.T) {
	type testdata struct {
		addr   []int
		subnet []int
		expect []int
	}

	tests := []testdata{
		{
			[]int{192, 168, 10, 111},
			[]int{255, 255, 255, 0},
			[]int{192, 168, 10, 255},
		},
		{
			[]int{192, 168, 255, 111},
			[]int{255, 255, 128, 0},
			[]int{192, 168, 255, 255},
		},
	}

	for index, test := range tests {
		naddr := GetBroadcastAddress(test.addr, test.subnet)

		if false == reflect.DeepEqual(naddr, test.expect) {
			t.Errorf("error subnet mismatch. testindex:%d, expect:%d, test:%d", index, test.expect, naddr)
		}
	}
}
