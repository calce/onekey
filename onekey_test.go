package onekey

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	chain := Keychain()
	for i := 0; i<33; i++ {
		chain.generate()
	}
	assert.Equal(t, chain.size(), 33, "KeyChain should have the right size")
}

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