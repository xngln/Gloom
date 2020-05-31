package gloom

import (
	"math"

	"github.com/cespare/xxhash"
	"github.com/xngln/go-datastructures/bitarray"
)

type BloomFilter struct {
	m      uint64
	k      uint64
	n      uint64
	bitArr bitarray.BitArray
}
