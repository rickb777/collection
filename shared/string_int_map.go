// An encapsulated map[string]int.
//
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=string Type=int
// options: Comparable:true Stringer:true KeyList:collection.StringList ValueList:collection.IntList Mutable:always
// by runtemplate v3.6.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package shared

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/rickb777/collection"
	"sync"
)

// StringIntMap is the primary type that represents a thread-safe map
type StringIntMap struct {
	s *sync.RWMutex
	m map[string]int
}

// StringIntTuple represents a key/value pair.
type StringIntTuple struct {
	Key string
	Val int
}

// StringIntTuples can be used as a builder for unmodifiable maps.
type StringIntTuples []StringIntTuple

// Append1 adds one item.
func (ts StringIntTuples) Append1(k string, v int) StringIntTuples {
	return append(ts, StringIntTuple{k, v})
}

// Append2 adds two items.
func (ts StringIntTuples) Append2(k1 string, v1 int, k2 string, v2 int) StringIntTuples {
	return append(ts, StringIntTuple{k1, v1}, StringIntTuple{k2, v2})
}

// Append3 adds three items.
func (ts StringIntTuples) Append3(k1 string, v1 int, k2 string, v2 int, k3 string, v3 int) StringIntTuples {
	return append(ts, StringIntTuple{k1, v1}, StringIntTuple{k2, v2}, StringIntTuple{k3, v3})
}

// StringIntZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewStringIntMap
// constructor function.
func StringIntZip(keys ...string) StringIntTuples {
	ts := make(StringIntTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with StringIntZip.
func (ts StringIntTuples) Values(values ...int) StringIntTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts StringIntTuples) ToMap() *StringIntMap {
	return NewStringIntMap(ts...)
}

//-------------------------------------------------------------------------------------------------

func newStringIntMap() *StringIntMap {
	return &StringIntMap{
		s: &sync.RWMutex{},
		m: make(map[string]int),
	}
}

// NewStringIntMap1 creates and returns a reference to a map containing one item.
func NewStringIntMap1(k string, v int) *StringIntMap {
	mm := newStringIntMap()
	mm.m[k] = v
	return mm
}

// NewStringIntMap creates and returns a reference to a map, optionally containing some items.
func NewStringIntMap(kv ...StringIntTuple) *StringIntMap {
	mm := newStringIntMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *StringIntMap) Keys() collection.StringList {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make(collection.StringList, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *StringIntMap) Values() collection.IntList {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make(collection.IntList, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *StringIntMap) slice() StringIntTuples {
	if mm == nil {
		return nil
	}

	s := make(StringIntTuples, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, StringIntTuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *StringIntMap) ToSlice() StringIntTuples {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm *StringIntMap) OrderedSlice(keys collection.StringList) StringIntTuples {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make(StringIntTuples, 0, len(mm.m))
	for _, k := range keys {
		v, found := mm.m[k]
		if found {
			s = append(s, StringIntTuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *StringIntMap) Get(k string) (int, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *StringIntMap) Put(k string, v int) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *StringIntMap) ContainsKey(k string) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *StringIntMap) ContainsAllKeys(kk ...string) bool {
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
func (mm *StringIntMap) Clear() {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		mm.m = make(map[string]int)
	}
}

// Remove a single item from the map.
func (mm *StringIntMap) Remove(k string) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		delete(mm.m, k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *StringIntMap) Pop(k string) (int, bool) {
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
func (mm *StringIntMap) Size() int {
	if mm == nil {
		return 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *StringIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *StringIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *StringIntMap) DropWhere(fn func(string, int) bool) StringIntTuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(StringIntTuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, StringIntTuple{(k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *StringIntMap) Foreach(f func(string, int)) {
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
func (mm *StringIntMap) Forall(p func(string, int) bool) bool {
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
func (mm *StringIntMap) Exists(p func(string, int) bool) bool {
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
func (mm *StringIntMap) Find(p func(string, int) bool) (StringIntTuple, bool) {
	if mm == nil {
		return StringIntTuple{}, false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			return StringIntTuple{(k), v}, true
		}
	}

	return StringIntTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *StringIntMap) Filter(p func(string, int) bool) *StringIntMap {
	if mm == nil {
		return nil
	}

	result := NewStringIntMap()
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
func (mm *StringIntMap) Partition(p func(string, int) bool) (matching *StringIntMap, others *StringIntMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewStringIntMap()
	others = NewStringIntMap()
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

// Map returns a new IntMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *StringIntMap) Map(f func(string, int) (string, int)) *StringIntMap {
	if mm == nil {
		return nil
	}

	result := NewStringIntMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new IntMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *StringIntMap) FlatMap(f func(string, int) []StringIntTuple) *StringIntMap {
	if mm == nil {
		return nil
	}

	result := NewStringIntMap()
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
func (mm *StringIntMap) Equals(other *StringIntMap) bool {
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
func (mm *StringIntMap) Clone() *StringIntMap {
	if mm == nil {
		return nil
	}

	result := NewStringIntMap()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

func (mm *StringIntMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *StringIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *StringIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *StringIntMap) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *StringIntMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""
	mm.s.RLock()
	defer mm.s.RUnlock()

	keys := make(collection.StringList, 0, len(mm.m))
	for k, _ := range mm.m {
		keys = append(keys, k)
	}
	keys.Sorted()

	for _, k := range keys {
		v := mm.m[k]
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}

	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this map type.
func (mm *StringIntMap) UnmarshalJSON(b []byte) error {
	mm.s.Lock()
	defer mm.s.Unlock()

	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&mm.m)
}

// MarshalJSON implements JSON encoding for this map type.
func (mm *StringIntMap) MarshalJSON() ([]byte, error) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	return json.Marshal(mm.m)
}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this map type.
// You must register int with the 'gob' package before this method is used.
func (mm *StringIntMap) GobDecode(b []byte) error {
	mm.s.Lock()
	defer mm.s.Unlock()

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this map type.
// You must register int with the 'gob' package before this method is used.
func (mm *StringIntMap) GobEncode() ([]byte, error) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}

//-------------------------------------------------------------------------------------------------

func (ts StringIntTuples) String() string {
	return ts.MkString3("[", ", ", "]")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (ts StringIntTuples) MkString(sep string) string {
	return ts.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (ts StringIntTuples) MkString3(before, between, after string) string {
	if ts == nil {
		return ""
	}
	return ts.mkString3Bytes(before, between, after).String()
}

func (ts StringIntTuples) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""
	for _, t := range ts {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", t.Key, t.Val))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this tuple type.
func (t StringIntTuple) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&t)
}

// MarshalJSON implements encoding.Marshaler interface.
func (t StringIntTuple) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"key":"%v", "val":"%v"}`, t.Key, t.Val)), nil
}
