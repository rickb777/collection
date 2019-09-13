// An encapsulated map[int]string.
//
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=int Type=string
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

// SharedIntStringMap is the primary type that represents a thread-safe map
type SharedIntStringMap struct {
	s *sync.RWMutex
	m map[int]string
}

// SharedIntStringTuple represents a key/value pair.
type SharedIntStringTuple struct {
	Key int
	Val string
}

// SharedIntStringTuples can be used as a builder for unmodifiable maps.
type SharedIntStringTuples []SharedIntStringTuple

// Append1 adds one item.
func (ts SharedIntStringTuples) Append1(k int, v string) SharedIntStringTuples {
	return append(ts, SharedIntStringTuple{k, v})
}

// Append2 adds two items.
func (ts SharedIntStringTuples) Append2(k1 int, v1 string, k2 int, v2 string) SharedIntStringTuples {
	return append(ts, SharedIntStringTuple{k1, v1}, SharedIntStringTuple{k2, v2})
}

// Append3 adds three items.
func (ts SharedIntStringTuples) Append3(k1 int, v1 string, k2 int, v2 string, k3 int, v3 string) SharedIntStringTuples {
	return append(ts, SharedIntStringTuple{k1, v1}, SharedIntStringTuple{k2, v2}, SharedIntStringTuple{k3, v3})
}

// SharedIntStringZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewSharedIntStringMap
// constructor function.
func SharedIntStringZip(keys ...int) SharedIntStringTuples {
	ts := make(SharedIntStringTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with SharedIntStringZip.
func (ts SharedIntStringTuples) Values(values ...string) SharedIntStringTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newSharedIntStringMap() *SharedIntStringMap {
	return &SharedIntStringMap{
		s: &sync.RWMutex{},
		m: make(map[int]string),
	}
}

// NewSharedIntStringMap1 creates and returns a reference to a map containing one item.
func NewSharedIntStringMap1(k int, v string) *SharedIntStringMap {
	mm := newSharedIntStringMap()
	mm.m[k] = v
	return mm
}

// NewSharedIntStringMap creates and returns a reference to a map, optionally containing some items.
func NewSharedIntStringMap(kv ...SharedIntStringTuple) *SharedIntStringMap {
	mm := newSharedIntStringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *SharedIntStringMap) Keys() []int {
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
func (mm *SharedIntStringMap) Values() []string {
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
func (mm *SharedIntStringMap) slice() []SharedIntStringTuple {
	if mm == nil {
		return nil
	}

	s := make([]SharedIntStringTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, SharedIntStringTuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *SharedIntStringMap) ToSlice() []SharedIntStringTuple {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *SharedIntStringMap) Get(k int) (string, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *SharedIntStringMap) Put(k int, v string) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *SharedIntStringMap) ContainsKey(k int) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *SharedIntStringMap) ContainsAllKeys(kk ...int) bool {
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
func (mm *SharedIntStringMap) Clear() {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		mm.m = make(map[int]string)
	}
}

// Remove a single item from the map.
func (mm *SharedIntStringMap) Remove(k int) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		delete(mm.m, k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *SharedIntStringMap) Pop(k int) (string, bool) {
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
func (mm *SharedIntStringMap) Size() int {
	if mm == nil {
		return 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *SharedIntStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *SharedIntStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *SharedIntStringMap) DropWhere(fn func(int, string) bool) SharedIntStringTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(SharedIntStringTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, SharedIntStringTuple{(k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *SharedIntStringMap) Foreach(f func(int, string)) {
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
func (mm *SharedIntStringMap) Forall(p func(int, string) bool) bool {
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
func (mm *SharedIntStringMap) Exists(p func(int, string) bool) bool {
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
func (mm *SharedIntStringMap) Find(p func(int, string) bool) (SharedIntStringTuple, bool) {
	if mm == nil {
		return SharedIntStringTuple{}, false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			return SharedIntStringTuple{(k), v}, true
		}
	}

	return SharedIntStringTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *SharedIntStringMap) Filter(p func(int, string) bool) *SharedIntStringMap {
	if mm == nil {
		return nil
	}

	result := NewSharedIntStringMap()
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
func (mm *SharedIntStringMap) Partition(p func(int, string) bool) (matching *SharedIntStringMap, others *SharedIntStringMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewSharedIntStringMap()
	others = NewSharedIntStringMap()
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
func (mm *SharedIntStringMap) Map(f func(int, string) (int, string)) *SharedIntStringMap {
	if mm == nil {
		return nil
	}

	result := NewSharedIntStringMap()
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
func (mm *SharedIntStringMap) FlatMap(f func(int, string) []SharedIntStringTuple) *SharedIntStringMap {
	if mm == nil {
		return nil
	}

	result := NewSharedIntStringMap()
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
func (mm *SharedIntStringMap) Equals(other *SharedIntStringMap) bool {
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
func (mm *SharedIntStringMap) Clone() *SharedIntStringMap {
	if mm == nil {
		return nil
	}

	result := NewSharedIntStringMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

func (mm *SharedIntStringMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *SharedIntStringMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *SharedIntStringMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *SharedIntStringMap) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *SharedIntStringMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
func (mm *SharedIntStringMap) GobDecode(b []byte) error {
	mm.s.Lock()
	defer mm.s.Unlock()

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register string with the 'gob' package before this method is used.
func (mm *SharedIntStringMap) GobEncode() ([]byte, error) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}
