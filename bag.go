// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package bag implements some utility methods on top of a generic map in order to
// minimize the number of written type assertions in your project.
package bag

import (
	"sync"
)

// A Bag uses a map[string]interface{} to hold data. It is safe for concurrent use.
type Bag struct {
	data map[string]interface{}
	rw   sync.RWMutex
}

// Instantiates a Bag
func NewBag() *Bag {
	return &Bag{data: make(map[string]interface{})}
}

// Instantiates a Bag from an already existing map
func From(m map[string]interface{}) *Bag {
	return &Bag{data: m}
}

// Returns value associated to key.
func (b *Bag) Get(key string) (value interface{}) {
	b.rw.RLock()
	defer b.rw.RUnlock()

	return b.data[key]
}

// Has() returns true if key is present in data
func (b *Bag) Has(key string) bool {
	_, ok := b.data[key]
	return ok
}

func (b *Bag) Set(key string, value interface{}) {
	b.rw.Lock()
	defer b.rw.Unlock()

	b.data[key] = value
}

// Equivalent to Get but does a string type assertion on the value. 
// ok is the result of the type assertion.
func (b *Bag) GetString(key string) string {
	b.rw.RLock()
	defer b.rw.RUnlock()

	v, ok := b.data[key]
	if !ok {
		return ""
	}
	return v.(string)
}

func (b *Bag) GetMapStringString(key string) map[string]string {
	b.rw.RLock()
	defer b.rw.RUnlock()

	v, ok := b.data[key]
	if !ok {
		return make(map[string]string)
	}
	return v.(map[string]string)
}

// Same as GetString but with the bool type
func (b *Bag) GetBool(key string) bool {
	b.rw.RLock()
	defer b.rw.RUnlock()

	v, ok := b.data[key]
	if !ok {
		return false
	}

	return v.(bool)
}

// Same as GetString but with the int type
func (b *Bag) GetInt(key string) int {
	b.rw.RLock()
	defer b.rw.RUnlock()

	v, ok := b.data[key]
	if !ok {
		return 0
	}

	return v.(int)
}

func (b *Bag) GetByteSlice(key string) (value []byte) {
	b.rw.RLock()
	defer b.rw.RUnlock()

	v, ok := b.data[key]
	if !ok {
		return
	}
	return v.([]byte)
}

// Returns the underlying map
func (b *Bag) Map() map[string]interface{} {
	return b.data
}
