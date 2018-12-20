package filter

import (
	"math"
)

//使用数组压缩法实现bit set, 五重校验
type BloomFilter struct {
	layer uint8
	sets  [][]uint64
}

func (filter *BloomFilter) hash(index int, url []byte) uint32 {
	var hash uint32 = 0

	switch index {
	case 0:
		hash = RSHash(url)
	case 1:
		hash = JSHash(url)
	case 2:
		hash = PJWHash(url)
	case 3:
		hash = BKDRHash(url)
	case 4:
		hash = SDBMHash(url)
	case 5:
		hash = DJBHash(url)
	case 6:
		hash = DEKHash(url)
	case 7:
		hash = APHash(url)
	default:
		panic("index out of range")
	}

	return hash
}

func (filter *BloomFilter) Reset() {
	if filter.layer > 8 {
		panic("layer more than 8")
	}

	var sets = make([][]uint64, filter.layer, filter.layer)

	for i := 0; i < len(sets); i++ {
		size := (math.MaxUint32 / 64) + 1
		sets[i] = make([]uint64, size, size)
	}

	filter.sets = sets
}

func (filter *BloomFilter) Set(url []byte) {
	for index := range filter.sets {
		hash := filter.hash(index, url)
		set := filter.sets[index]
		setArrayIndex := hash / 64
		setEleIndex := hash % 64
		temp := uint64(1 << setEleIndex)
		set[setArrayIndex] = set[setArrayIndex] | temp
	}
}

func (filter *BloomFilter) Has(url []byte) bool {
	var result = true

	for index := range filter.sets {
		hash := filter.hash(index, url)
		set := filter.sets[index]
		setArrayIndex := hash / 64
		setEleIndex := hash % 64
		temp := uint64(1 << setEleIndex)
		result = result && (set[setArrayIndex]&temp == temp)
		if !result {
			break
		}
	}

	return result
}

func NewBloomFilter(layer uint8) *BloomFilter {
	filter := &BloomFilter{}
	filter.Reset()
	return filter
}
