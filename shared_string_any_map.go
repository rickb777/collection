// An encapsulated map[string]interface{}.
//
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=string Type=interface{}
// options: Comparable:<no value> Stringer:<no value> KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.5.2
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

import (
	"fmt"
	"sync"
)

// SharedStringAnyMap is the primary type that represents a thread-safe map
type SharedStringAnyMap struct {
	s *sync.RWMutex
	m map[string]interface{}
}

// SharedStringAnyTuple represents a key/value pair.
type SharedStringAnyTuple struct {
	Key string
	Val interface{}
}

// SharedStringAnyTuples can be used as a builder for unmodifiable maps.
type SharedStringAnyTuples []SharedStringAnyTuple

// Append1 adds one item.
func (ts SharedStringAnyTuples) Append1(k string, v interface{}) SharedStringAnyTuples {
	return append(ts, SharedStringAnyTuple{k, v})
}

// Append2 adds two items.
func (ts SharedStringAnyTuples) Append2(k1 string, v1 interface{}, k2 string, v2 interface{}) SharedStringAnyTuples {
	return append(ts, SharedStringAnyTuple{k1, v1}, SharedStringAnyTuple{k2, v2})
}

// Append3 adds three items.
func (ts SharedStringAnyTuples) Append3(k1 string, v1 interface{}, k2 string, v2 interface{}, k3 string, v3 interface{}) SharedStringAnyTuples {
	return append(ts, SharedStringAnyTuple{k1, v1}, SharedStringAnyTuple{k2, v2}, SharedStringAnyTuple{k3, v3})
}

// SharedStringAnyZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewSharedStringAnyMap
// constructor function.
func SharedStringAnyZip(keys ...string) SharedStringAnyTuples {
	ts := make(SharedStringAnyTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with SharedStringAnyZip.
func (ts SharedStringAnyTuples) Values(values ...interface{}) SharedStringAnyTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newSharedStringAnyMap() *SharedStringAnyMap {
	return &SharedStringAnyMap{
		s: &sync.RWMutex{},
		m: make(map[string]interface{}),
	}
}

// NewSharedStringAnyMap1 creates and returns a reference to a map containing one item.
func NewSharedStringAnyMap1(k string, v interface{}) *SharedStringAnyMap {
	mm := newSharedStringAnyMap()
	mm.m[k] = v
	return mm
}

// NewSharedStringAnyMap creates and returns a reference to a map, optionally containing some items.
func NewSharedStringAnyMap(kv ...SharedStringAnyTuple) *SharedStringAnyMap {
	mm := newSharedStringAnyMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *SharedStringAnyMap) Keys() []string {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]string, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *SharedStringAnyMap) Values() []interface{} {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]interface{}, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *SharedStringAnyMap) slice() []SharedStringAnyTuple {
	if mm == nil {
		return nil
	}

	s := make([]SharedStringAnyTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, SharedStringAnyTuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *SharedStringAnyMap) ToSlice() []SharedStringAnyTuple {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *SharedStringAnyMap) Get(k string) (interface{}, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *SharedStringAnyMap) Put(k string, v interface{}) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *SharedStringAnyMap) ContainsKey(k string) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *SharedStringAnyMap) ContainsAllKeys(kk ...string) bool {
	if mm == nil {
		return len(kk) == 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *SharedStringAnyMap) Clear() {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		mm.m = make(map[string]interface{})
	}
}

// Remove a single item from the map.
func (mm *SharedStringAnyMap) Remove(k string) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		delete(mm.m, k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *SharedStringAnyMap) Pop(k string) (interface{}, bool) {
	if mm == nil {
		return nil, false
	}

	mm.s.Lock()
	defer mm.s.Unlock()

	v, found := mm.m[k]
	delete(mm.m, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *SharedStringAnyMap) Size() int {
	if mm == nil {
		return 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *SharedStringAnyMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *SharedStringAnyMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *SharedStringAnyMap) DropWhere(fn func(string, interface{}) bool) SharedStringAnyTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(SharedStringAnyTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, SharedStringAnyTuple{(k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *SharedStringAnyMap) Foreach(f func(string, interface{})) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		for k, v := range mm.m {
			f(k, v)
		}
	}
}

// Forall applies the predicate p to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm *SharedStringAnyMap) Forall(p func(string, interface{}) bool) bool {
	if mm == nil {
		return true
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if !p(k, v) {
			return false
		}
	}

	return true
}

// Exists applies the predicate p to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm *SharedStringAnyMap) Exists(p func(string, interface{}) bool) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first interface{} that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm *SharedStringAnyMap) Find(p func(string, interface{}) bool) (SharedStringAnyTuple, bool) {
	if mm == nil {
		return SharedStringAnyTuple{}, false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			return SharedStringAnyTuple{(k), v}, true
		}
	}

	return SharedStringAnyTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *SharedStringAnyMap) Filter(p func(string, interface{}) bool) *SharedStringAnyMap {
	if mm == nil {
		return nil
	}

	result := NewSharedStringAnyMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			result.m[k] = v
		}
	}

	return result
}

// Partition applies the predicate p to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
// The original map is not modified.
func (mm *SharedStringAnyMap) Partition(p func(string, interface{}) bool) (matching *SharedStringAnyMap, others *SharedStringAnyMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewSharedStringAnyMap()
	others = NewSharedStringAnyMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new SharedAnyMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *SharedStringAnyMap) Map(f func(string, interface{}) (string, interface{})) *SharedStringAnyMap {
	if mm == nil {
		return nil
	}

	result := NewSharedStringAnyMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new SharedAnyMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *SharedStringAnyMap) FlatMap(f func(string, interface{}) []SharedStringAnyTuple) *SharedStringAnyMap {
	if mm == nil {
		return nil
	}

	result := NewSharedStringAnyMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		ts := f(k1, v1)
		for _, t := range ts {
			result.m[t.Key] = t.Val
		}
	}

	return result
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm *SharedStringAnyMap) Clone() *SharedStringAnyMap {
	if mm == nil {
		return nil
	}

	result := NewSharedStringAnyMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}
