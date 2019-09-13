// A simple type derived from map[int64]int64.
//
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=int64 Type=int64
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.5.2
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

import (
	"bytes"
	"fmt"
)

// Int64Int64Map is the primary type that represents a map
type Int64Int64Map map[int64]int64

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

func newInt64Int64Map() Int64Int64Map {
	return Int64Int64Map(make(map[int64]int64))
}

// NewInt64Int64Map1 creates and returns a reference to a map containing one item.
func NewInt64Int64Map1(k int64, v int64) Int64Int64Map {
	mm := newInt64Int64Map()
	mm[k] = v
	return mm
}

// NewInt64Int64Map creates and returns a reference to a map, optionally containing some items.
func NewInt64Int64Map(kv ...Int64Int64Tuple) Int64Int64Map {
	mm := newInt64Int64Map()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm Int64Int64Map) Keys() []int64 {
	s := make([]int64, 0, len(mm))
	for k := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm Int64Int64Map) Values() []int64 {
	s := make([]int64, 0, len(mm))
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm Int64Int64Map) slice() []Int64Int64Tuple {
	s := make([]Int64Int64Tuple, 0, len(mm))
	for k, v := range mm {
		s = append(s, Int64Int64Tuple{(k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm Int64Int64Map) ToSlice() []Int64Int64Tuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm Int64Int64Map) Get(k int64) (int64, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm Int64Int64Map) Put(k int64, v int64) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm Int64Int64Map) ContainsKey(k int64) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm Int64Int64Map) ContainsAllKeys(kk ...int64) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *Int64Int64Map) Clear() {
	*mm = make(map[int64]int64)
}

// Remove a single item from the map.
func (mm Int64Int64Map) Remove(k int64) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm Int64Int64Map) Pop(k int64) (int64, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm Int64Int64Map) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm Int64Int64Map) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm Int64Int64Map) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm Int64Int64Map) DropWhere(fn func(int64, int64) bool) Int64Int64Tuples {
	removed := make(Int64Int64Tuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, Int64Int64Tuple{(k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm Int64Int64Map) Foreach(f func(int64, int64)) {
	for k, v := range mm {
		f(k, v)
	}
}

// Forall applies the predicate p to every element in the map. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (mm Int64Int64Map) Forall(p func(int64, int64) bool) bool {
	for k, v := range mm {
		if !p(k, v) {
			return false
		}
	}

	return true
}

// Exists applies the predicate p to every element in the map. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (mm Int64Int64Map) Exists(p func(int64, int64) bool) bool {
	for k, v := range mm {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first int64 that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm Int64Int64Map) Find(p func(int64, int64) bool) (Int64Int64Tuple, bool) {
	for k, v := range mm {
		if p(k, v) {
			return Int64Int64Tuple{(k), v}, true
		}
	}

	return Int64Int64Tuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm Int64Int64Map) Filter(p func(int64, int64) bool) Int64Int64Map {
	result := NewInt64Int64Map()
	for k, v := range mm {
		if p(k, v) {
			result[k] = v
		}
	}

	return result
}

// Partition applies the predicate p to every element in the map. It divides the map into two copied maps,
// the first containing all the elements for which the predicate returned true, and the second containing all
// the others.
// The original map is not modified.
func (mm Int64Int64Map) Partition(p func(int64, int64) bool) (matching Int64Int64Map, others Int64Int64Map) {
	matching = NewInt64Int64Map()
	others = NewInt64Int64Map()
	for k, v := range mm {
		if p(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new Int64Map by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm Int64Int64Map) Map(f func(int64, int64) (int64, int64)) Int64Int64Map {
	result := NewInt64Int64Map()

	for k1, v1 := range mm {
		k2, v2 := f(k1, v1)
		result[k2] = v2
	}

	return result
}

// FlatMap returns a new Int64Map by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm Int64Int64Map) FlatMap(f func(int64, int64) []Int64Int64Tuple) Int64Int64Map {
	result := NewInt64Int64Map()

	for k1, v1 := range mm {
		ts := f(k1, v1)
		for _, t := range ts {
			result[t.Key] = t.Val
		}
	}

	return result
}

// Equals determines if two maps are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for maps to be equal.
func (mm Int64Int64Map) Equals(other Int64Int64Map) bool {
	if mm.Size() != other.Size() {
		return false
	}
	for k, v1 := range mm {
		v2, found := other[k]
		if !found || v1 != v2 {
			return false
		}
	}
	return true
}

// Clone returns a shallow copy of the map. It does not clone the underlying elements.
func (mm Int64Int64Map) Clone() Int64Int64Map {
	result := NewInt64Int64Map()
	for k, v := range mm {
		result[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm Int64Int64Map) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm Int64Int64Map) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm Int64Int64Map) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm Int64Int64Map) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm Int64Int64Map) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	for k, v := range mm {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}

	b.WriteString(after)
	return b
}
