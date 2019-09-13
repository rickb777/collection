// An encapsulated map[uint64]string.
//
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=uint64 Type=string
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.5.2
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
)

// SharedUint64StringMap is the primary type that represents a thread-safe map
type SharedUint64StringMap struct {
	s *sync.RWMutex
	m map[uint64]string
}

// SharedUint64StringTuple represents a key/value pair.
type SharedUint64StringTuple struct {
	Key uint64
	Val string
}

// SharedUint64StringTuples can be used as a builder for unmodifiable maps.
type SharedUint64StringTuples []SharedUint64StringTuple

// Append1 adds one item.
func (ts SharedUint64StringTuples) Append1(k uint64, v string) SharedUint64StringTuples {
	return append(ts, SharedUint64StringTuple{k, v})
}

// Append2 adds two items.
func (ts SharedUint64StringTuples) Append2(k1 uint64, v1 string, k2 uint64, v2 string) SharedUint64StringTuples {
	return append(ts, SharedUint64StringTuple{k1, v1}, SharedUint64StringTuple{k2, v2})
}

// Append3 adds three items.
func (ts SharedUint64StringTuples) Append3(k1 uint64, v1 string, k2 uint64, v2 string, k3 uint64, v3 string) SharedUint64StringTuples {
	return append(ts, SharedUint64StringTuple{k1, v1}, SharedUint64StringTuple{k2, v2}, SharedUint64StringTuple{k3, v3})
}

// SharedUint64StringZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewSharedUint64StringMap
// constructor function.
func SharedUint64StringZip(keys ...uint64) SharedUint64StringTuples {
	ts := make(SharedUint64StringTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with SharedUint64StringZip.
func (ts SharedUint64StringTuples) Values(values ...string) SharedUint64StringTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newSharedUint64StringMap() *SharedUint64StringMap {
	return &SharedUint64StringMap{
		s: &sync.RWMutex{},
		m: make(map[uint64]string),
	}
}

// NewSharedUint64StringMap1 creates and returns a reference to a map containing one item.
func NewSharedUint64StringMap1(k uint64, v string) *SharedUint64StringMap {
	mm := newSharedUint64StringMap()
	mm.m[k] = v
	return mm
}

// NewSharedUint64StringMap creates and returns a reference to a map, optionally containing some items.
func NewSharedUint64StringMap(kv ...SharedUint64StringTuple) *SharedUint64StringMap {
	mm := newSharedUint64StringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *SharedUint64StringMap) Keys() []uint64 {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]uint64, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *SharedUint64StringMap) Values() []string {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]string, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *SharedUint64StringMap) slice() []SharedUint64StringTuple {
	if mm == nil {
		return nil
	}

	s := make([]SharedUint64StringTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, SharedUint64StringTuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *SharedUint64StringMap) ToSlice() []SharedUint64StringTuple {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *SharedUint64StringMap) Get(k uint64) (string, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *SharedUint64StringMap) Put(k uint64, v string) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *SharedUint64StringMap) ContainsKey(k uint64) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *SharedUint64StringMap) ContainsAllKeys(kk ...uint64) bool {
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
func (mm *SharedUint64StringMap) Clear() {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		mm.m = make(map[uint64]string)
	}
}

// Remove a single item from the map.
func (mm *SharedUint64StringMap) Remove(k uint64) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		delete(mm.m, k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *SharedUint64StringMap) Pop(k uint64) (string, bool) {
	if mm == nil {
		return "", false
	}

	mm.s.Lock()
	defer mm.s.Unlock()

	v, found := mm.m[k]
	delete(mm.m, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *SharedUint64StringMap) Size() int {
	if mm == nil {
		return 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *SharedUint64StringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *SharedUint64StringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *SharedUint64StringMap) DropWhere(fn func(uint64, string) bool) SharedUint64StringTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(SharedUint64StringTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, SharedUint64StringTuple{(k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *SharedUint64StringMap) Foreach(f func(uint64, string)) {
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
func (mm *SharedUint64StringMap) Forall(p func(uint64, string) bool) bool {
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
func (mm *SharedUint64StringMap) Exists(p func(uint64, string) bool) bool {
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

// Find returns the first string that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm *SharedUint64StringMap) Find(p func(uint64, string) bool) (SharedUint64StringTuple, bool) {
	if mm == nil {
		return SharedUint64StringTuple{}, false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			return SharedUint64StringTuple{(k), v}, true
		}
	}

	return SharedUint64StringTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *SharedUint64StringMap) Filter(p func(uint64, string) bool) *SharedUint64StringMap {
	if mm == nil {
		return nil
	}

	result := NewSharedUint64StringMap()
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
func (mm *SharedUint64StringMap) Partition(p func(uint64, string) bool) (matching *SharedUint64StringMap, others *SharedUint64StringMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewSharedUint64StringMap()
	others = NewSharedUint64StringMap()
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

// Map returns a new SharedStringMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *SharedUint64StringMap) Map(f func(uint64, string) (uint64, string)) *SharedUint64StringMap {
	if mm == nil {
		return nil
	}

	result := NewSharedUint64StringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new SharedStringMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *SharedUint64StringMap) FlatMap(f func(uint64, string) []SharedUint64StringTuple) *SharedUint64StringMap {
	if mm == nil {
		return nil
	}

	result := NewSharedUint64StringMap()
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
func (mm *SharedUint64StringMap) Equals(other *SharedUint64StringMap) bool {
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
func (mm *SharedUint64StringMap) Clone() *SharedUint64StringMap {
	if mm == nil {
		return nil
	}

	result := NewSharedUint64StringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

func (mm *SharedUint64StringMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *SharedUint64StringMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *SharedUint64StringMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *SharedUint64StringMap) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *SharedUint64StringMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
// You must register string with the 'gob' package before this method is used.
func (mm *SharedUint64StringMap) GobDecode(b []byte) error {
	mm.s.Lock()
	defer mm.s.Unlock()

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register string with the 'gob' package before this method is used.
func (mm *SharedUint64StringMap) GobEncode() ([]byte, error) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}
