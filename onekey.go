// Generate and consume unique strings
package onekey

import (
	"time"
	"code.google.com/p/go-uuid/uuid"
)

var zero = time.Time{}

type Keychain struct {
	keys map[string]time.Time
}

func NewKeychain() Keychain {
	return Keychain{
		keys: make(map[string]time.Time),
	}
}

func (k *Keychain) generate() string {
	key := uuid.New()
	k.keys[key] = time.Now()
	return key
}

func (k *Keychain) consume(key string) bool {
	t := k.keys[key]
	delete(k.keys, key)
	return (t != zero)
}

func (k *Keychain) size() int {
	return len(k.keys)
}