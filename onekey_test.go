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