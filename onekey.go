// Generate and consume unique strings
package onekey

import (
	"time"
	"code.google.com/p/go-uuid/uuid"
)

var zero = time.Time{}

type OneKey struct {
	keys map[string]time.Time
}

func Keychain() OneKey {
	return OneKey{
		keys: make(map[string]time.Time),
	}
}

func (k *OneKey) generate() string {
	key := uuid.New()
	k.keys[key] = time.Now()
	return key
}

func (k *OneKey) consume(key string) bool {
	t := k.keys[key]
	return (t != zero)
}

func (k *OneKey) size() int {
	return len(k.keys)
}