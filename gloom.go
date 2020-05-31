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

func (b *BloomFilter) Add(thing []byte) {
	indices := b.getIndices(thing)

	for _, index := range indices {
		b.bitArr.SetBit(index)
	}

	b.n++
}

func (b *BloomFilter) AddString(thing string) {
	b.Add([]byte(thing))
}

func (b *BloomFilter) Contains(thing []byte) bool {
	indices := b.getIndices(thing)

	for _, index := range indices {
		// ignore err returned by GetBit since err != nil returned only
		// when index is out of range, but we make sure index is in range
		if bit, _ := b.bitArr.GetBit(index); !bit {
			return false
		}
	}
	return true
}

func (b *BloomFilter) ContainsString(thing string) bool {
	return b.Contains([]byte(thing))
}

func (b *BloomFilter) FalsePosProb() float64 {
	p := math.Pow(1-math.Pow(1-(1/float64(b.m)), float64(b.k*b.n)), float64(b.k))
	return p
}

func (b *BloomFilter) Clear() {
	b.n = 0
	b.bitArr.Reset()
}
