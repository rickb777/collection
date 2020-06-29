// A simple type derived from map[int]int.
//
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=int Type=int
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.5.4
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

import (
	"bytes"
	"fmt"
)

// IntIntMap is the primary type that represents a map
type IntIntMap map[int]int

// IntIntTuple represents a key/value pair.
type IntIntTuple struct {
	Key int
	Val int
}

// IntIntTuples can be used as a builder for unmodifiable maps.
type IntIntTuples []IntIntTuple

// Append1 adds one item.
func (ts IntIntTuples) Append1(k int, v int) IntIntTuples {
	return append(ts, IntIntTuple{k, v})
}

// Append2 adds two items.
func (ts IntIntTuples) Append2(k1 int, v1 int, k2 int, v2 int) IntIntTuples {
	return append(ts, IntIntTuple{k1, v1}, IntIntTuple{k2, v2})
}

// Append3 adds three items.
func (ts IntIntTuples) Append3(k1 int, v1 int, k2 int, v2 int, k3 int, v3 int) IntIntTuples {
	return append(ts, IntIntTuple{k1, v1}, IntIntTuple{k2, v2}, IntIntTuple{k3, v3})
}

// IntIntZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewIntIntMap
// constructor function.
func IntIntZip(keys ...int) IntIntTuples {
	ts := make(IntIntTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with IntIntZip.
func (ts IntIntTuples) Values(values ...int) IntIntTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newIntIntMap() IntIntMap {
	return IntIntMap(make(map[int]int))
}

// NewIntIntMap1 creates and returns a reference to a map containing one item.
func NewIntIntMap1(k int, v int) IntIntMap {
	mm := newIntIntMap()
	mm[k] = v
	return mm
}

// NewIntIntMap creates and returns a reference to a map, optionally containing some items.
func NewIntIntMap(kv ...IntIntTuple) IntIntMap {
	mm := newIntIntMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm IntIntMap) Keys() []int {
	s := make([]int, 0, len(mm))
	for k := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm IntIntMap) Values() []int {
	s := make([]int, 0, len(mm))
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm IntIntMap) slice() []IntIntTuple {
	s := make([]IntIntTuple, 0, len(mm))
	for k, v := range mm {
		s = append(s, IntIntTuple{(k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm IntIntMap) ToSlice() []IntIntTuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm IntIntMap) Get(k int) (int, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm IntIntMap) Put(k int, v int) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm IntIntMap) ContainsKey(k int) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm IntIntMap) ContainsAllKeys(kk ...int) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *IntIntMap) Clear() {
	*mm = make(map[int]int)
}

// Remove a single item from the map.
func (mm IntIntMap) Remove(k int) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm IntIntMap) Pop(k int) (int, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm IntIntMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm IntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm IntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm IntIntMap) DropWhere(fn func(int, int) bool) IntIntTuples {
	removed := make(IntIntTuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, IntIntTuple{(k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm IntIntMap) Foreach(f func(int, int)) {
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
func (mm IntIntMap) Forall(p func(int, int) bool) bool {
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
func (mm IntIntMap) Exists(p func(int, int) bool) bool {
	for k, v := range mm {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first int that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm IntIntMap) Find(p func(int, int) bool) (IntIntTuple, bool) {
	for k, v := range mm {
		if p(k, v) {
			return IntIntTuple{(k), v}, true
		}
	}

	return IntIntTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm IntIntMap) Filter(p func(int, int) bool) IntIntMap {
	result := NewIntIntMap()
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
func (mm IntIntMap) Partition(p func(int, int) bool) (matching IntIntMap, others IntIntMap) {
	matching = NewIntIntMap()
	others = NewIntIntMap()
	for k, v := range mm {
		if p(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new IntMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm IntIntMap) Map(f func(int, int) (int, int)) IntIntMap {
	result := NewIntIntMap()

	for k1, v1 := range mm {
		k2, v2 := f(k1, v1)
		result[k2] = v2
	}

	return result
}

// FlatMap returns a new IntMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm IntIntMap) FlatMap(f func(int, int) []IntIntTuple) IntIntMap {
	result := NewIntIntMap()

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
func (mm IntIntMap) Equals(other IntIntMap) bool {
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
func (mm IntIntMap) Clone() IntIntMap {
	result := NewIntIntMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm IntIntMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm IntIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm IntIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm IntIntMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm IntIntMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
