// Generate and consume unique strings
package onekey

import (
	"time"
	"code.google.com/p/go-uuid/uuid"
	"errors"
)

var zero = time.Time{}

type OneKey struct {
	keys map[string]time.Time
}

func KeyChain() OneKey {
	return Keys{
		keys: make(map[string]time.Time)
	}
}

func (k *OneKey) generate() string {
	key := uuid.New()
	k.keys[key] = time.Now()
}

func (k *OneKey) consume(key) bool {
	t := k.keys[key]
	return (t != zero)
}

func (k *OneKey) size() {
	return len(k.keys)
}