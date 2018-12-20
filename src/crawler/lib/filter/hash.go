package filter

func RSHash(str []byte) uint32 {
	b := 378551
	a := 63689
	hash := uint32(0)
	for i := 0; i < len(str); i++ {
		hash = hash*uint32(a) + uint32(str[i])
		a = a * b
	}
	return hash
}

func JSHash(str []byte) uint32 {
	hash := uint32(1315423911)
	for i := 0; i < len(str); i++ {
		hash ^= (hash << 5) + uint32(str[i]) + (hash >> 2)
	}
	return hash
}

func PJWHash(str []byte) uint32 {
	bitsInUnsignedInt := (uint32)(4 * 8)
	threeQuarters := (uint32)((bitsInUnsignedInt * 3) / 4)
	oneEighth := (uint32)(bitsInUnsignedInt / 8)
	highBits := (uint32)(0xFFFFFFFF) << (bitsInUnsignedInt - oneEighth)
	hash := uint32(0)
	test := uint32(0)
	for i := 0; i < len(str); i++ {
		hash = (hash << oneEighth) + uint32(str[i])
		if test = hash & highBits; test != 0 {
			hash = (hash ^ (test >> threeQuarters)) & (^highBits)
		}
	}
	return hash
}

func BKDRHash(str []byte) uint32 {
	seed := uint32(131)
	hash := uint32(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + uint32(str[i])
	}
	return hash
}

func SDBMHash(str []byte) uint32 {
	hash := uint32(0)
	for i := 0; i < len(str); i++ {
		hash = uint32(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}

func DJBHash(str []byte) uint32 {
	hash := uint32(0)
	for i := 0; i < len(str); i++ {
		hash = ((hash << 5) + hash) + uint32(str[i])
	}
	return hash
}

func DEKHash(str []byte) uint32 {
	hash := uint32(len(str))
	for i := 0; i < len(str); i++ {
		hash = ((hash << 5) ^ (hash >> 27)) ^ uint32(str[i])
	}
	return hash
}

func APHash(str []byte) uint32 {
	hash := uint32(0xAAAAAAAA)
	for i := 0; i < len(str); i++ {
		if (i & 1) == 0 {
			hash ^= (hash << 7) ^ uint32(str[i])*(hash>>3)
		} else {
			hash ^= ^((hash << 11) + uint32(str[i]) ^ (hash >> 5))
		}
	}
	return hash
}
