package cider

import (
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
