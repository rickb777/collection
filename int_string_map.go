// A simple type derived from map[int]string.
//
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=int Type=string
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.5.4
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

import (
	"bytes"
	"fmt"
)

// IntStringMap is the primary type that represents a map
type IntStringMap map[int]string

// IntStringTuple represents a key/value pair.
type IntStringTuple struct {
	Key int
	Val string
}

// IntStringTuples can be used as a builder for unmodifiable maps.
type IntStringTuples []IntStringTuple

// Append1 adds one item.
func (ts IntStringTuples) Append1(k int, v string) IntStringTuples {
	return append(ts, IntStringTuple{k, v})
}

// Append2 adds two items.
func (ts IntStringTuples) Append2(k1 int, v1 string, k2 int, v2 string) IntStringTuples {
	return append(ts, IntStringTuple{k1, v1}, IntStringTuple{k2, v2})
}

// Append3 adds three items.
func (ts IntStringTuples) Append3(k1 int, v1 string, k2 int, v2 string, k3 int, v3 string) IntStringTuples {
	return append(ts, IntStringTuple{k1, v1}, IntStringTuple{k2, v2}, IntStringTuple{k3, v3})
}

// IntStringZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewIntStringMap
// constructor function.
func IntStringZip(keys ...int) IntStringTuples {
	ts := make(IntStringTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with IntStringZip.
func (ts IntStringTuples) Values(values ...string) IntStringTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newIntStringMap() IntStringMap {
	return IntStringMap(make(map[int]string))
}

// NewIntStringMap1 creates and returns a reference to a map containing one item.
func NewIntStringMap1(k int, v string) IntStringMap {
	mm := newIntStringMap()
	mm[k] = v
	return mm
}

// NewIntStringMap creates and returns a reference to a map, optionally containing some items.
func NewIntStringMap(kv ...IntStringTuple) IntStringMap {
	mm := newIntStringMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm IntStringMap) Keys() []int {
	s := make([]int, 0, len(mm))
	for k := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm IntStringMap) Values() []string {
	s := make([]string, 0, len(mm))
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm IntStringMap) slice() []IntStringTuple {
	s := make([]IntStringTuple, 0, len(mm))
	for k, v := range mm {
		s = append(s, IntStringTuple{(k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm IntStringMap) ToSlice() []IntStringTuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm IntStringMap) Get(k int) (string, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm IntStringMap) Put(k int, v string) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm IntStringMap) ContainsKey(k int) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm IntStringMap) ContainsAllKeys(kk ...int) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *IntStringMap) Clear() {
	*mm = make(map[int]string)
}

// Remove a single item from the map.
func (mm IntStringMap) Remove(k int) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm IntStringMap) Pop(k int) (string, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm IntStringMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm IntStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm IntStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm IntStringMap) DropWhere(fn func(int, string) bool) IntStringTuples {
	removed := make(IntStringTuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, IntStringTuple{(k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm IntStringMap) Foreach(f func(int, string)) {
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
func (mm IntStringMap) Forall(p func(int, string) bool) bool {
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
func (mm IntStringMap) Exists(p func(int, string) bool) bool {
	for k, v := range mm {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first string that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm IntStringMap) Find(p func(int, string) bool) (IntStringTuple, bool) {
	for k, v := range mm {
		if p(k, v) {
			return IntStringTuple{(k), v}, true
		}
	}

	return IntStringTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm IntStringMap) Filter(p func(int, string) bool) IntStringMap {
	result := NewIntStringMap()
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
func (mm IntStringMap) Partition(p func(int, string) bool) (matching IntStringMap, others IntStringMap) {
	matching = NewIntStringMap()
	others = NewIntStringMap()
	for k, v := range mm {
		if p(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new StringMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm IntStringMap) Map(f func(int, string) (int, string)) IntStringMap {
	result := NewIntStringMap()

	for k1, v1 := range mm {
		k2, v2 := f(k1, v1)
		result[k2] = v2
	}

	return result
}

// FlatMap returns a new StringMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm IntStringMap) FlatMap(f func(int, string) []IntStringTuple) IntStringMap {
	result := NewIntStringMap()

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
func (mm IntStringMap) Equals(other IntStringMap) bool {
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
func (mm IntStringMap) Clone() IntStringMap {
	result := NewIntStringMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm IntStringMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm IntStringMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm IntStringMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm IntStringMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm IntStringMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
