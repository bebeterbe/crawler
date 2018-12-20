package filter

import (
	"strconv"
	"testing"
)

var filter = NewBloomFilter(5)

func TestSet(t *testing.T) {
	for i := 0; i < 10000000; i++ {
		filter.Set([]byte(strconv.Itoa(i)))
	}
}

func TestHas(t *testing.T) {
	filter.Reset()
	for i := 0; i < 10000000; i++ {
		str := strconv.Itoa(i)
		filter.Set([]byte(str))
		if filter.Has([]byte(str)) == false {
			t.Fail()
		}
	}
}
