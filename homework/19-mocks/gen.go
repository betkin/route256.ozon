package mocks_learning

import (
	"crypto/rand"
	"math/big"
)

type genRand interface {
	Random(n int) int64
	Doer()
}

type genStruct struct{}

func (r *genStruct) Doer() {}

func (r *genStruct) Random(n int) int64 {
	r.Doer()
	v, _ := rand.Int(rand.Reader, big.NewInt(int64(n)+1))
	return v.Int64()
}

func myRandom(a, b int, r genRand) int64 {
	return int64(a) + r.Random(b+1)
}
