// An encapsulated map[int64]string.
//
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=int64 Type=string
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

// SharedInt64StringMap is the primary type that represents a thread-safe map
type SharedInt64StringMap struct {
	s *sync.RWMutex
	m map[int64]string
}

// SharedInt64StringTuple represents a key/value pair.
type SharedInt64StringTuple struct {
	Key int64
	Val string
}

// SharedInt64StringTuples can be used as a builder for unmodifiable maps.
type SharedInt64StringTuples []SharedInt64StringTuple

// Append1 adds one item.
func (ts SharedInt64StringTuples) Append1(k int64, v string) SharedInt64StringTuples {
	return append(ts, SharedInt64StringTuple{k, v})
}

// Append2 adds two items.
func (ts SharedInt64StringTuples) Append2(k1 int64, v1 string, k2 int64, v2 string) SharedInt64StringTuples {
	return append(ts, SharedInt64StringTuple{k1, v1}, SharedInt64StringTuple{k2, v2})
}

// Append3 adds three items.
func (ts SharedInt64StringTuples) Append3(k1 int64, v1 string, k2 int64, v2 string, k3 int64, v3 string) SharedInt64StringTuples {
	return append(ts, SharedInt64StringTuple{k1, v1}, SharedInt64StringTuple{k2, v2}, SharedInt64StringTuple{k3, v3})
}

// SharedInt64StringZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewSharedInt64StringMap
// constructor function.
func SharedInt64StringZip(keys ...int64) SharedInt64StringTuples {
	ts := make(SharedInt64StringTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with SharedInt64StringZip.
func (ts SharedInt64StringTuples) Values(values ...string) SharedInt64StringTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newSharedInt64StringMap() *SharedInt64StringMap {
	return &SharedInt64StringMap{
		s: &sync.RWMutex{},
		m: make(map[int64]string),
	}
}

// NewSharedInt64StringMap1 creates and returns a reference to a map containing one item.
func NewSharedInt64StringMap1(k int64, v string) *SharedInt64StringMap {
	mm := newSharedInt64StringMap()
	mm.m[k] = v
	return mm
}

// NewSharedInt64StringMap creates and returns a reference to a map, optionally containing some items.
func NewSharedInt64StringMap(kv ...SharedInt64StringTuple) *SharedInt64StringMap {
	mm := newSharedInt64StringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *SharedInt64StringMap) Keys() []int64 {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]int64, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *SharedInt64StringMap) Values() []string {
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
func (mm *SharedInt64StringMap) slice() []SharedInt64StringTuple {
	if mm == nil {
		return nil
	}

	s := make([]SharedInt64StringTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, SharedInt64StringTuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *SharedInt64StringMap) ToSlice() []SharedInt64StringTuple {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *SharedInt64StringMap) Get(k int64) (string, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *SharedInt64StringMap) Put(k int64, v string) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *SharedInt64StringMap) ContainsKey(k int64) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *SharedInt64StringMap) ContainsAllKeys(kk ...int64) bool {
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
func (mm *SharedInt64StringMap) Clear() {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		mm.m = make(map[int64]string)
	}
}

// Remove a single item from the map.
func (mm *SharedInt64StringMap) Remove(k int64) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		delete(mm.m, k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *SharedInt64StringMap) Pop(k int64) (string, bool) {
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
func (mm *SharedInt64StringMap) Size() int {
	if mm == nil {
		return 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *SharedInt64StringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *SharedInt64StringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *SharedInt64StringMap) DropWhere(fn func(int64, string) bool) SharedInt64StringTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(SharedInt64StringTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, SharedInt64StringTuple{(k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *SharedInt64StringMap) Foreach(f func(int64, string)) {
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
func (mm *SharedInt64StringMap) Forall(p func(int64, string) bool) bool {
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
func (mm *SharedInt64StringMap) Exists(p func(int64, string) bool) bool {
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
func (mm *SharedInt64StringMap) Find(p func(int64, string) bool) (SharedInt64StringTuple, bool) {
	if mm == nil {
		return SharedInt64StringTuple{}, false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			return SharedInt64StringTuple{(k), v}, true
		}
	}

	return SharedInt64StringTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *SharedInt64StringMap) Filter(p func(int64, string) bool) *SharedInt64StringMap {
	if mm == nil {
		return nil
	}

	result := NewSharedInt64StringMap()
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
func (mm *SharedInt64StringMap) Partition(p func(int64, string) bool) (matching *SharedInt64StringMap, others *SharedInt64StringMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewSharedInt64StringMap()
	others = NewSharedInt64StringMap()
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
func (mm *SharedInt64StringMap) Map(f func(int64, string) (int64, string)) *SharedInt64StringMap {
	if mm == nil {
		return nil
	}

	result := NewSharedInt64StringMap()
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
func (mm *SharedInt64StringMap) FlatMap(f func(int64, string) []SharedInt64StringTuple) *SharedInt64StringMap {
	if mm == nil {
		return nil
	}

	result := NewSharedInt64StringMap()
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
func (mm *SharedInt64StringMap) Equals(other *SharedInt64StringMap) bool {
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
func (mm *SharedInt64StringMap) Clone() *SharedInt64StringMap {
	if mm == nil {
		return nil
	}

	result := NewSharedInt64StringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

func (mm *SharedInt64StringMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *SharedInt64StringMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *SharedInt64StringMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *SharedInt64StringMap) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *SharedInt64StringMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
func (mm *SharedInt64StringMap) GobDecode(b []byte) error {
	mm.s.Lock()
	defer mm.s.Unlock()

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register string with the 'gob' package before this method is used.
func (mm *SharedInt64StringMap) GobEncode() ([]byte, error) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}
