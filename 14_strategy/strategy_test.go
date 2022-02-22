package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStorageStrategy(t *testing.T) {
	strategyType := "file"
	storage, err := NewStorageStrategy(strategyType)
	assert.Nil(t, err)
	assert.Nil(t, storage.Save("./test.txt", []byte("123")))
}