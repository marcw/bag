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

// Returns value associated to key. if key does not exist, ok is false.
func (b *Bag) Get(key string) (value interface{}, ok bool) {
	b.rw.RLock()
	defer b.rw.RUnlock()

	value, ok = b.data[key]
	return
}

func (b *Bag) Set(key string, value interface{}) {
	b.rw.Lock()
	defer b.rw.Unlock()

	b.data[key] = value
}

// Equivalent to Get but does a string type assertion on the value. 
// ok is the result of the type assertion.
func (b *Bag) GetString(key string) (value string, ok bool) {
	b.rw.RLock()
	defer b.rw.RUnlock()

	v, ok := b.data[key]
	if !ok {
		return
	}

	value, ok = v.(string)
	return
}

// Same as GetString but with the bool type
func (b *Bag) GetBool(key string) (value, ok bool) {
	b.rw.RLock()
	defer b.rw.RUnlock()

	v, ok := b.data[key]
	if !ok {
		return
	}

	value, ok = v.(bool)
	return
}

// Same as GetString but with the int type
func (b *Bag) GetInt(key string) (value int, ok bool) {
	b.rw.RLock()
	defer b.rw.RUnlock()

	v, ok := b.data[key]
	if !ok {
		return
	}

	value, ok = v.(int)
	return
}
