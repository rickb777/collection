// An encapsulated map[int]int.
//
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=int Type=int
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

// SharedIntIntMap is the primary type that represents a thread-safe map
type SharedIntIntMap struct {
	s *sync.RWMutex
	m map[int]int
}

// SharedIntIntTuple represents a key/value pair.
type SharedIntIntTuple struct {
	Key int
	Val int
}

// SharedIntIntTuples can be used as a builder for unmodifiable maps.
type SharedIntIntTuples []SharedIntIntTuple

// Append1 adds one item.
func (ts SharedIntIntTuples) Append1(k int, v int) SharedIntIntTuples {
	return append(ts, SharedIntIntTuple{k, v})
}

// Append2 adds two items.
func (ts SharedIntIntTuples) Append2(k1 int, v1 int, k2 int, v2 int) SharedIntIntTuples {
	return append(ts, SharedIntIntTuple{k1, v1}, SharedIntIntTuple{k2, v2})
}

// Append3 adds three items.
func (ts SharedIntIntTuples) Append3(k1 int, v1 int, k2 int, v2 int, k3 int, v3 int) SharedIntIntTuples {
	return append(ts, SharedIntIntTuple{k1, v1}, SharedIntIntTuple{k2, v2}, SharedIntIntTuple{k3, v3})
}

// SharedIntIntZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewSharedIntIntMap
// constructor function.
func SharedIntIntZip(keys ...int) SharedIntIntTuples {
	ts := make(SharedIntIntTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with SharedIntIntZip.
func (ts SharedIntIntTuples) Values(values ...int) SharedIntIntTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newSharedIntIntMap() *SharedIntIntMap {
	return &SharedIntIntMap{
		s: &sync.RWMutex{},
		m: make(map[int]int),
	}
}

// NewSharedIntIntMap1 creates and returns a reference to a map containing one item.
func NewSharedIntIntMap1(k int, v int) *SharedIntIntMap {
	mm := newSharedIntIntMap()
	mm.m[k] = v
	return mm
}

// NewSharedIntIntMap creates and returns a reference to a map, optionally containing some items.
func NewSharedIntIntMap(kv ...SharedIntIntTuple) *SharedIntIntMap {
	mm := newSharedIntIntMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *SharedIntIntMap) Keys() []int {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]int, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *SharedIntIntMap) Values() []int {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make([]int, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *SharedIntIntMap) slice() []SharedIntIntTuple {
	if mm == nil {
		return nil
	}

	s := make([]SharedIntIntTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, SharedIntIntTuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *SharedIntIntMap) ToSlice() []SharedIntIntTuple {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *SharedIntIntMap) Get(k int) (int, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *SharedIntIntMap) Put(k int, v int) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *SharedIntIntMap) ContainsKey(k int) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *SharedIntIntMap) ContainsAllKeys(kk ...int) bool {
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
func (mm *SharedIntIntMap) Clear() {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		mm.m = make(map[int]int)
	}
}

// Remove a single item from the map.
func (mm *SharedIntIntMap) Remove(k int) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		delete(mm.m, k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *SharedIntIntMap) Pop(k int) (int, bool) {
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
func (mm *SharedIntIntMap) Size() int {
	if mm == nil {
		return 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *SharedIntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *SharedIntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *SharedIntIntMap) DropWhere(fn func(int, int) bool) SharedIntIntTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(SharedIntIntTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, SharedIntIntTuple{(k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *SharedIntIntMap) Foreach(f func(int, int)) {
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
func (mm *SharedIntIntMap) Forall(p func(int, int) bool) bool {
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
func (mm *SharedIntIntMap) Exists(p func(int, int) bool) bool {
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

// Find returns the first int that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm *SharedIntIntMap) Find(p func(int, int) bool) (SharedIntIntTuple, bool) {
	if mm == nil {
		return SharedIntIntTuple{}, false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			return SharedIntIntTuple{(k), v}, true
		}
	}

	return SharedIntIntTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *SharedIntIntMap) Filter(p func(int, int) bool) *SharedIntIntMap {
	if mm == nil {
		return nil
	}

	result := NewSharedIntIntMap()
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
func (mm *SharedIntIntMap) Partition(p func(int, int) bool) (matching *SharedIntIntMap, others *SharedIntIntMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewSharedIntIntMap()
	others = NewSharedIntIntMap()
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

// Map returns a new SharedIntMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *SharedIntIntMap) Map(f func(int, int) (int, int)) *SharedIntIntMap {
	if mm == nil {
		return nil
	}

	result := NewSharedIntIntMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new SharedIntMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *SharedIntIntMap) FlatMap(f func(int, int) []SharedIntIntTuple) *SharedIntIntMap {
	if mm == nil {
		return nil
	}

	result := NewSharedIntIntMap()
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
func (mm *SharedIntIntMap) Equals(other *SharedIntIntMap) bool {
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
func (mm *SharedIntIntMap) Clone() *SharedIntIntMap {
	if mm == nil {
		return nil
	}

	result := NewSharedIntIntMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

func (mm *SharedIntIntMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *SharedIntIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *SharedIntIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *SharedIntIntMap) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *SharedIntIntMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
// You must register int with the 'gob' package before this method is used.
func (mm *SharedIntIntMap) GobDecode(b []byte) error {
	mm.s.Lock()
	defer mm.s.Unlock()

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register int with the 'gob' package before this method is used.
func (mm *SharedIntIntMap) GobEncode() ([]byte, error) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}
