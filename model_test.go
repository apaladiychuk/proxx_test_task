package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		b, err := NewBoard(7, 7, 10)
		assert.NoError(t, err)
		assert.NotNil(t, b.Map)
	})
	t.Run("failed", func(t *testing.T) {
		b, err := NewBoard(7, 7, 50)
		assert.Error(t, err)
		assert.Nil(t, b)
	})

}
