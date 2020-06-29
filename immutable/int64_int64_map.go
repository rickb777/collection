// An encapsulated immutable map[int64]int64.
// Thread-safe.
//
//
// Generated from immutable/map.tpl with Key=int64 Type=int64
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:disabled
// by runtemplate v3.5.4
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Int64Int64Map is the primary type that represents a thread-safe map
type Int64Int64Map struct {
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

//-------------------------------------------------------------------------------------------------

func newInt64Int64Map() *Int64Int64Map {
	return &Int64Int64Map{
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
func (mm *Int64Int64Map) Keys() []int64 {
	if mm == nil {
		return nil
	}

	s := make([]int64, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *Int64Int64Map) Values() []int64 {
	if mm == nil {
		return nil
	}

	s := make([]int64, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *Int64Int64Map) slice() []Int64Int64Tuple {
	if mm == nil {
		return nil
	}

	s := make([]Int64Int64Tuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, Int64Int64Tuple{k, v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *Int64Int64Map) ToSlice() []Int64Int64Tuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *Int64Int64Map) Get(k int64) (int64, bool) {
	v, found := mm.m[k]
	return v, found
}

// Put adds an item to a clone of the map, replacing any prior value and returning the cloned map.
func (mm *Int64Int64Map) Put(k int64, v int64) *Int64Int64Map {
	if mm == nil {
		return NewInt64Int64Map1(k, v)
	}

	result := NewInt64Int64Map()

	for k, v := range mm.m {
		result.m[k] = v
	}
	result.m[k] = v

	return result
}

// ContainsKey determines if a given item is already in the map.
func (mm *Int64Int64Map) ContainsKey(k int64) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *Int64Int64Map) ContainsAllKeys(kk ...int64) bool {
	if mm == nil {
		return len(kk) == 0
	}

	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm *Int64Int64Map) Size() int {
	if mm == nil {
		return 0
	}

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

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *Int64Int64Map) Foreach(f func(int64, int64)) {
	if mm != nil {
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
func (mm *Int64Int64Map) Forall(f func(int64, int64) bool) bool {
	if mm == nil {
		return true
	}

	for k, v := range mm.m {
		if !f(k, v) {
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

	for k, v := range mm.m {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first int64 that returns true for the predicate p.
// False is returned if none match.
func (mm *Int64Int64Map) Find(p func(int64, int64) bool) (Int64Int64Tuple, bool) {
	if mm == nil {
		return Int64Int64Tuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return Int64Int64Tuple{k, v}, true
		}
	}

	return Int64Int64Tuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *Int64Int64Map) Filter(p func(int64, int64) bool) *Int64Int64Map {
	if mm == nil {
		return nil
	}

	result := NewInt64Int64Map()

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
func (mm *Int64Int64Map) Partition(p func(int64, int64) bool) (matching *Int64Int64Map, others *Int64Int64Map) {
	if mm == nil {
		return nil, nil
	}

	matching = NewInt64Int64Map()
	others = NewInt64Int64Map()

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
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *Int64Int64Map) Map(f func(int64, int64) (int64, int64)) *Int64Int64Map {
	if mm == nil {
		return nil
	}

	result := NewInt64Int64Map()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new Int64Map by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *Int64Int64Map) FlatMap(f func(int64, int64) []Int64Int64Tuple) *Int64Int64Map {
	if mm == nil {
		return nil
	}

	result := NewInt64Int64Map()

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

// Clone returns the same map, which is immutable.
func (mm *Int64Int64Map) Clone() *Int64Int64Map {
	return mm
}

//-------------------------------------------------------------------------------------------------

func (mm *Int64Int64Map) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *Int64Int64Map) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *Int64Int64Map) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *Int64Int64Map) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *Int64Int64Map) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

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
// You must register int64 with the 'gob' package before this method is used.
func (mm *Int64Int64Map) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register int64 with the 'gob' package before this method is used.
func (mm *Int64Int64Map) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}
