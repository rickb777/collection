// A simple type derived from map[string]int.
//
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=string Type=int
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:always
// by runtemplate v3.5.4
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

import (
	"bytes"
	"fmt"
)

// StringIntMap is the primary type that represents a map
type StringIntMap map[string]int

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

//-------------------------------------------------------------------------------------------------

func newStringIntMap() StringIntMap {
	return StringIntMap(make(map[string]int))
}

// NewStringIntMap1 creates and returns a reference to a map containing one item.
func NewStringIntMap1(k string, v int) StringIntMap {
	mm := newStringIntMap()
	mm[k] = v
	return mm
}

// NewStringIntMap creates and returns a reference to a map, optionally containing some items.
func NewStringIntMap(kv ...StringIntTuple) StringIntMap {
	mm := newStringIntMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm StringIntMap) Keys() []string {
	s := make([]string, 0, len(mm))
	for k := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm StringIntMap) Values() []int {
	s := make([]int, 0, len(mm))
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm StringIntMap) slice() []StringIntTuple {
	s := make([]StringIntTuple, 0, len(mm))
	for k, v := range mm {
		s = append(s, StringIntTuple{(k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm StringIntMap) ToSlice() []StringIntTuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm StringIntMap) Get(k string) (int, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm StringIntMap) Put(k string, v int) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm StringIntMap) ContainsKey(k string) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm StringIntMap) ContainsAllKeys(kk ...string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *StringIntMap) Clear() {
	*mm = make(map[string]int)
}

// Remove a single item from the map.
func (mm StringIntMap) Remove(k string) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm StringIntMap) Pop(k string) (int, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm StringIntMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm StringIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm StringIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm StringIntMap) DropWhere(fn func(string, int) bool) StringIntTuples {
	removed := make(StringIntTuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, StringIntTuple{(k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm StringIntMap) Foreach(f func(string, int)) {
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
func (mm StringIntMap) Forall(p func(string, int) bool) bool {
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
func (mm StringIntMap) Exists(p func(string, int) bool) bool {
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
func (mm StringIntMap) Find(p func(string, int) bool) (StringIntTuple, bool) {
	for k, v := range mm {
		if p(k, v) {
			return StringIntTuple{(k), v}, true
		}
	}

	return StringIntTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm StringIntMap) Filter(p func(string, int) bool) StringIntMap {
	result := NewStringIntMap()
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
func (mm StringIntMap) Partition(p func(string, int) bool) (matching StringIntMap, others StringIntMap) {
	matching = NewStringIntMap()
	others = NewStringIntMap()
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
func (mm StringIntMap) Map(f func(string, int) (string, int)) StringIntMap {
	result := NewStringIntMap()

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
func (mm StringIntMap) FlatMap(f func(string, int) []StringIntTuple) StringIntMap {
	result := NewStringIntMap()

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
func (mm StringIntMap) Equals(other StringIntMap) bool {
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
func (mm StringIntMap) Clone() StringIntMap {
	result := NewStringIntMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm StringIntMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm StringIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm StringIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm StringIntMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm StringIntMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
