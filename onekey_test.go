package onekey

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	chain := Keychain()
	_ = chain.generate()
	assert.Equal(t, chain.size(), 1, "KeyChain should properly generate keys")
}

func TestConsume(t *testing.T) {
	chain := Keychain()
	key := chain.generate()
	ok := chain.consume(key)
	assert.Equal(t, ok, true, "KeyChain should be able to consume generated keys")
	assert.Equal(t, chain.size(), 0, "KeyChain should properly remove consumed keys")
}