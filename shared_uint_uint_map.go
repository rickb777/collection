// An encapsulated map[uint]uint.
//
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=uint Type=uint
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.5.4
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
)

// SharedUintUintMap is the primary type that represents a thread-safe map
type SharedUintUintMap struct {
	s *sync.RWMutex
	m map[uint]uint
}

// SharedUintUintTuple represents a key/value pair.
type SharedUintUintTuple struct {
	Key uint
	Val uint
}

// SharedUintUintTuples can be used as a builder for unmodifiable maps.
type SharedUintUintTuples []SharedUintUintTuple

// Append1 adds one item.
func (ts SharedUintUintTuples) Append1(k uint, v uint) SharedUintUintTuples {
	return append(ts, SharedUintUintTuple{k, v})
}

// Append2 adds two items.
func (ts SharedUintUintTuples) Append2(k1 uint, v1 uint, k2 uint, v2 uint) SharedUintUintTuples {
	return append(ts, SharedUintUintTuple{k1, v1}, SharedUintUintTuple{k2, v2})
}

// Append3 adds three items.
func (ts SharedUintUintTuples) Append3(k1 uint, v1 uint, k2 uint, v2 uint, k3 uint, v3 uint) SharedUintUintTuples {
	return append(ts, SharedUintUintTuple{k1, v1}, SharedUintUintTuple{k2, v2}, SharedUintUintTuple{k3, v3})
}

// SharedUintUintZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewSharedUintUintMap
// constructor function.
func SharedUintUintZip(keys ...uint) SharedUintUintTuples {
	ts := make(SharedUintUintTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with SharedUintUintZip.
func (ts SharedUintUintTuples) Values(values ...uint) SharedUintUintTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newSharedUintUintMap() *SharedUintUintMap {
	return &SharedUintUintMap{
		s: &sync.RWMutex{},
		m: make(map[uint]uint),
	}
}

// NewSharedUintUintMap1 creates and returns a reference to a map containing one item.
func NewSharedUintUintMap1(k uint, v uint) *SharedUintUintMap {
	mm := newSharedUintUintMap()
	mm.m[k] = v
	return mm
}

// NewSharedUintUintMap creates and returns a reference to a map, optionally containing some items.
func NewSharedUintUintMap(kv ...SharedUintUintTuple) *SharedUintUintMap {
	mm := newSharedUintUintMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *SharedUintUintMap) Keys() []uint {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]uint, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *SharedUintUintMap) Values() []uint {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]uint, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *SharedUintUintMap) slice() []SharedUintUintTuple {
	if mm == nil {
		return nil
	}

	s := make([]SharedUintUintTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, SharedUintUintTuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *SharedUintUintMap) ToSlice() []SharedUintUintTuple {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *SharedUintUintMap) Get(k uint) (uint, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *SharedUintUintMap) Put(k uint, v uint) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *SharedUintUintMap) ContainsKey(k uint) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *SharedUintUintMap) ContainsAllKeys(kk ...uint) bool {
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
func (mm *SharedUintUintMap) Clear() {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		mm.m = make(map[uint]uint)
	}
}

// Remove a single item from the map.
func (mm *SharedUintUintMap) Remove(k uint) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		delete(mm.m, k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *SharedUintUintMap) Pop(k uint) (uint, bool) {
	if mm == nil {
		return 0, false
	}

	mm.s.Lock()
	defer mm.s.Unlock()

	v, found := mm.m[k]
	delete(mm.m, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *SharedUintUintMap) Size() int {
	if mm == nil {
		return 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *SharedUintUintMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *SharedUintUintMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *SharedUintUintMap) DropWhere(fn func(uint, uint) bool) SharedUintUintTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(SharedUintUintTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, SharedUintUintTuple{(k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *SharedUintUintMap) Foreach(f func(uint, uint)) {
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
func (mm *SharedUintUintMap) Forall(p func(uint, uint) bool) bool {
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
func (mm *SharedUintUintMap) Exists(p func(uint, uint) bool) bool {
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

// Find returns the first uint that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm *SharedUintUintMap) Find(p func(uint, uint) bool) (SharedUintUintTuple, bool) {
	if mm == nil {
		return SharedUintUintTuple{}, false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			return SharedUintUintTuple{(k), v}, true
		}
	}

	return SharedUintUintTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *SharedUintUintMap) Filter(p func(uint, uint) bool) *SharedUintUintMap {
	if mm == nil {
		return nil
	}

	result := NewSharedUintUintMap()
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
func (mm *SharedUintUintMap) Partition(p func(uint, uint) bool) (matching *SharedUintUintMap, others *SharedUintUintMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewSharedUintUintMap()
	others = NewSharedUintUintMap()
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

// Map returns a new SharedUintMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *SharedUintUintMap) Map(f func(uint, uint) (uint, uint)) *SharedUintUintMap {
	if mm == nil {
		return nil
	}

	result := NewSharedUintUintMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new SharedUintMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *SharedUintUintMap) FlatMap(f func(uint, uint) []SharedUintUintTuple) *SharedUintUintMap {
	if mm == nil {
		return nil
	}

	result := NewSharedUintUintMap()
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

// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm *SharedUintUintMap) Equals(other *SharedUintUintMap) bool {
	if mm == nil || other == nil {
		return mm.IsEmpty() && other.IsEmpty()
	}

	mm.s.RLock()
	other.s.RLock()
	defer mm.s.RUnlock()
	defer other.s.RUnlock()

	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm.m {
		v2, found := other.m[k]
		if !found || v1 != v2 {
			return false
		}
	}
	return true
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm *SharedUintUintMap) Clone() *SharedUintUintMap {
	if mm == nil {
		return nil
	}

	result := NewSharedUintUintMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

func (mm *SharedUintUintMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *SharedUintUintMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *SharedUintUintMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *SharedUintUintMap) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *SharedUintUintMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}

	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this map type.
// You must register uint with the 'gob' package before this method is used.
func (mm *SharedUintUintMap) GobDecode(b []byte) error {
	mm.s.Lock()
	defer mm.s.Unlock()

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register uint with the 'gob' package before this method is used.
func (mm *SharedUintUintMap) GobEncode() ([]byte, error) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}
