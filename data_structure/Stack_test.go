package data_structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	var (
		s Stack
		err error
	)

	// スタックを初期化
	s.Initialize()
	assert.Equal(t, 0, top)

	// スタックが空のときに IsEmpty, IsFull
	assert.True(t, s.IsEmpty())
	assert.False(t, s.IsFull())

	// Push の動作確認
	for i := 0; i < MAX-1; i++ {
		err := s.Push(i+1)
		assert.Equal(t, nil, err)
	}
	// スタックが一杯の状態で Push すると error を返すか
	err = s.Push(10)
	assert.NotNil(t, err)

	// スタックが一杯の状態で IsEmpty, IsFull
	assert.False(t, s.IsEmpty())
	assert.True(t, s.IsFull())

	// Pop の動作確認
	for i := 0; i < MAX-1; i++ {
		x, err := s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, MAX-1-i, x)
	}
	// スタックが空の状態で Pop すると error を返すか
	_, err = s.Pop()
	assert.NotNil(t, err)
}