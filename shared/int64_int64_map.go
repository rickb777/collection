// An encapsulated map[int64]int64.
//
// Thread-safe.
//
// Generated from threadsafe/map.tpl with Key=int64 Type=int64
// options: Comparable:true Stringer:true KeyList:collection.Int64List ValueList:collection.Int64List Mutable:always
// by runtemplate v3.10.0
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package shared

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/rickb777/collection"
	"strings"
	"sync"
)

// Int64Int64Map is the primary type that represents a thread-safe map
type Int64Int64Map struct {
	s *sync.RWMutex
	m map[int64]int64
}

// Int64Int64Tuple represents a key/value pair.
type Int64Int64Tuple struct {
	Key int64
	Val int64
}

// Int64Int64Tuples can be used as a builder for unmodifiable maps.
type Int64Int64Tuples []Int64Int64Tuple

// Append1 adds one item.
func (ts Int64Int64Tuples) Append1(k int64, v int64) Int64Int64Tuples {
	return append(ts, Int64Int64Tuple{k, v})
}

// Append2 adds two items.
func (ts Int64Int64Tuples) Append2(k1 int64, v1 int64, k2 int64, v2 int64) Int64Int64Tuples {
	return append(ts, Int64Int64Tuple{k1, v1}, Int64Int64Tuple{k2, v2})
}

// Append3 adds three items.
func (ts Int64Int64Tuples) Append3(k1 int64, v1 int64, k2 int64, v2 int64, k3 int64, v3 int64) Int64Int64Tuples {
	return append(ts, Int64Int64Tuple{k1, v1}, Int64Int64Tuple{k2, v2}, Int64Int64Tuple{k3, v3})
}

// Int64Int64Zip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewInt64Int64Map
// constructor function.
func Int64Int64Zip(keys ...int64) Int64Int64Tuples {
	ts := make(Int64Int64Tuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with Int64Int64Zip.
func (ts Int64Int64Tuples) Values(values ...int64) Int64Int64Tuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts Int64Int64Tuples) ToMap() *Int64Int64Map {
	return NewInt64Int64Map(ts...)
}

//-------------------------------------------------------------------------------------------------

func newInt64Int64Map() *Int64Int64Map {
	return &Int64Int64Map{
		s: &sync.RWMutex{},
		m: make(map[int64]int64),
	}
}

// NewInt64Int64Map1 creates and returns a reference to a map containing one item.
func NewInt64Int64Map1(k int64, v int64) *Int64Int64Map {
	mm := newInt64Int64Map()
	mm.m[k] = v
	return mm
}

// NewInt64Int64Map creates and returns a reference to a map, optionally containing some items.
func NewInt64Int64Map(kv ...Int64Int64Tuple) *Int64Int64Map {
	mm := newInt64Int64Map()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *Int64Int64Map) Keys() collection.Int64List {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make(collection.Int64List, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *Int64Int64Map) Values() collection.Int64List {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make(collection.Int64List, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *Int64Int64Map) slice() Int64Int64Tuples {
	if mm == nil {
		return nil
	}

	s := make(Int64Int64Tuples, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, Int64Int64Tuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *Int64Int64Map) ToSlice() Int64Int64Tuples {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm *Int64Int64Map) OrderedSlice(keys collection.Int64List) Int64Int64Tuples {
	if mm == nil {
		return nil
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	s := make(Int64Int64Tuples, 0, len(mm.m))
	for _, k := range keys {
		v, found := mm.m[k]
		if found {
			s = append(s, Int64Int64Tuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *Int64Int64Map) Get(k int64) (int64, bool) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	v, found := mm.m[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm *Int64Int64Map) Put(k int64, v int64) bool {
	mm.s.Lock()
	defer mm.s.Unlock()

	_, found := mm.m[k]
	mm.m[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm *Int64Int64Map) ContainsKey(k int64) bool {
	if mm == nil {
		return false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *Int64Int64Map) ContainsAllKeys(kk ...int64) bool {
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
func (mm *Int64Int64Map) Clear() {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		mm.m = make(map[int64]int64)
	}
}

// Remove a single item from the map.
func (mm *Int64Int64Map) Remove(k int64) {
	if mm != nil {
		mm.s.Lock()
		defer mm.s.Unlock()

		delete(mm.m, k)
	}
}

// Pop removes a single item from the map, returning the value present prior to removal.
// The boolean result is true only if the key had been present.
func (mm *Int64Int64Map) Pop(k int64) (int64, bool) {
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
func (mm *Int64Int64Map) Size() int {
	if mm == nil {
		return 0
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *Int64Int64Map) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *Int64Int64Map) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm *Int64Int64Map) DropWhere(fn func(int64, int64) bool) Int64Int64Tuples {
	mm.s.RLock()
	defer mm.s.RUnlock()

	removed := make(Int64Int64Tuples, 0)
	for k, v := range mm.m {
		if fn(k, v) {
			removed = append(removed, Int64Int64Tuple{(k), v})
			delete(mm.m, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *Int64Int64Map) Foreach(f func(int64, int64)) {
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
func (mm *Int64Int64Map) Forall(p func(int64, int64) bool) bool {
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
func (mm *Int64Int64Map) Exists(p func(int64, int64) bool) bool {
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

// Find returns the first int64 that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm *Int64Int64Map) Find(p func(int64, int64) bool) (Int64Int64Tuple, bool) {
	if mm == nil {
		return Int64Int64Tuple{}, false
	}

	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		if p(k, v) {
			return Int64Int64Tuple{(k), v}, true
		}
	}

	return Int64Int64Tuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm *Int64Int64Map) Filter(p func(int64, int64) bool) *Int64Int64Map {
	if mm == nil {
		return nil
	}

	result := NewInt64Int64Map()
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
func (mm *Int64Int64Map) Partition(p func(int64, int64) bool) (matching *Int64Int64Map, others *Int64Int64Map) {
	if mm == nil {
		return nil, nil
	}

	matching = NewInt64Int64Map()
	others = NewInt64Int64Map()
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

// Map returns a new Int64Map by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *Int64Int64Map) Map(f func(int64, int64) (int64, int64)) *Int64Int64Map {
	if mm == nil {
		return nil
	}

	result := NewInt64Int64Map()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new Int64Map by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *Int64Int64Map) FlatMap(f func(int64, int64) []Int64Int64Tuple) *Int64Int64Map {
	if mm == nil {
		return nil
	}

	result := NewInt64Int64Map()
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
func (mm *Int64Int64Map) Equals(other *Int64Int64Map) bool {
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
func (mm *Int64Int64Map) Clone() *Int64Int64Map {
	if mm == nil {
		return nil
	}

	result := NewInt64Int64Map()
	mm.s.RLock()
	defer mm.s.RUnlock()

	for k, v := range mm.m {
		result.m[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm *Int64Int64Map) String() string {
	return mm.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *Int64Int64Map) MkString(sep string) string {
	return mm.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *Int64Int64Map) MkString4(before, between, after, equals string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString4Bytes(before, between, after, equals).String()
}

func (mm *Int64Int64Map) mkString4Bytes(before, between, after, equals string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""
	mm.s.RLock()
	defer mm.s.RUnlock()

	keys := make(collection.Int64List, 0, len(mm.m))
	for k, _ := range mm.m {
		keys = append(keys, k)
	}
	keys.Sorted()

	for _, k := range keys {
		v := mm.m[k]
		b.WriteString(sep)
		fmt.Fprintf(b, "%v%s%v", k, equals, v)
		sep = between
	}

	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this map type.
// You must register int64 with the 'gob' package before this method is used.
func (mm *Int64Int64Map) GobDecode(b []byte) error {
	mm.s.Lock()
	defer mm.s.Unlock()

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this map type.
// You must register int64 with the 'gob' package before this method is used.
func (mm *Int64Int64Map) GobEncode() ([]byte, error) {
	mm.s.RLock()
	defer mm.s.RUnlock()

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}

//-------------------------------------------------------------------------------------------------

func (ts Int64Int64Tuples) String() string {
	return ts.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (ts Int64Int64Tuples) MkString(sep string) string {
	return ts.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (ts Int64Int64Tuples) MkString4(before, between, after, equals string) string {
	if ts == nil {
		return ""
	}
	return ts.mkString4Bytes(before, between, after, equals).String()
}

func (ts Int64Int64Tuples) mkString4Bytes(before, between, after, equals string) *strings.Builder {
	b := &strings.Builder{}
	sep := before
	for _, t := range ts {
		b.WriteString(sep)
		fmt.Fprintf(b, "%v%s%v", t.Key, equals, t.Val)
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this tuple type.
func (t Int64Int64Tuple) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&t)
}

// MarshalJSON implements encoding.Marshaler interface.
func (t Int64Int64Tuple) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"key":"%v", "val":"%v"}`, t.Key, t.Val)), nil
}
