package mocks_learning

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type stubRandom struct{}

func (r *stubRandom) Random(n int) int64 {
	return int64(15)
}

func (r *stubRandom) Doer() {}

type spyRandom struct{ target int }

func (r *spyRandom) Random(n int) int64 {
	r.target = n
	return int64(15)
}

func (r *spyRandom) Doer() {}

func TestGen(t *testing.T) {
	// Stub
	a1 := myRandom(5, 20, new(stubRandom))
	assert.Equal(t, int64(20), a1)
	// Spy
	spy := new(spyRandom)
	a2 := myRandom(5, 20, spy)
	assert.Equal(t, int64(20), a2)
	assert.Equal(t, 21, spy.target)
	// Mock
	ctrl := gomock.NewController(t)
	mockRandom := NewMockgenRand(ctrl)
	call := mockRandom.EXPECT().Doer().AnyTimes()
	mockRandom.EXPECT().Random(21).Return(int64(15)).After(call).AnyTimes()
	a3 := myRandom(5, 20, mockRandom)
	assert.Equal(t, int64(20), a3)
}
