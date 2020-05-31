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

func New(m uint64, k uint64) *BloomFilter {
	ba := bitarray.NewBitArray(m)
	b := BloomFilter{m, k, 0, ba}
	return &b
}
