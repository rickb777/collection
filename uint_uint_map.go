// A simple type derived from map[uint]uint.
//
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=uint Type=uint
// options: Comparable:true Stringer:true KeyList:UintList ValueList:UintList Mutable:always
// by runtemplate v3.10.1
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package collection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// UintUintMap is the primary type that represents a map
type UintUintMap map[uint]uint

// UintUintTuple represents a key/value pair.
type UintUintTuple struct {
	Key uint
	Val uint
}

// UintUintTuples can be used as a builder for unmodifiable maps.
type UintUintTuples []UintUintTuple

// Append1 adds one item.
func (ts UintUintTuples) Append1(k uint, v uint) UintUintTuples {
	return append(ts, UintUintTuple{k, v})
}

// Append2 adds two items.
func (ts UintUintTuples) Append2(k1 uint, v1 uint, k2 uint, v2 uint) UintUintTuples {
	return append(ts, UintUintTuple{k1, v1}, UintUintTuple{k2, v2})
}

// Append3 adds three items.
func (ts UintUintTuples) Append3(k1 uint, v1 uint, k2 uint, v2 uint, k3 uint, v3 uint) UintUintTuples {
	return append(ts, UintUintTuple{k1, v1}, UintUintTuple{k2, v2}, UintUintTuple{k3, v3})
}

// UintUintZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewUintUintMap
// constructor function.
func UintUintZip(keys ...uint) UintUintTuples {
	ts := make(UintUintTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with UintUintZip.
func (ts UintUintTuples) Values(values ...uint) UintUintTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts UintUintTuples) ToMap() UintUintMap {
	return NewUintUintMap(ts...)
}

//-------------------------------------------------------------------------------------------------

func newUintUintMap() UintUintMap {
	return UintUintMap(make(map[uint]uint))
}

// NewUintUintMap1 creates and returns a reference to a map containing one item.
func NewUintUintMap1(k uint, v uint) UintUintMap {
	mm := newUintUintMap()
	mm[k] = v
	return mm
}

// NewUintUintMap creates and returns a reference to a map, optionally containing some items.
func NewUintUintMap(kv ...UintUintTuple) UintUintMap {
	mm := newUintUintMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm UintUintMap) Keys() UintList {
	if mm == nil {
		return nil
	}

	s := make(UintList, 0, len(mm))
	for k := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm UintUintMap) Values() UintList {
	if mm == nil {
		return nil
	}

	s := make(UintList, 0, len(mm))
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm UintUintMap) slice() UintUintTuples {
	s := make(UintUintTuples, 0, len(mm))
	for k, v := range mm {
		s = append(s, UintUintTuple{(k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice.
func (mm UintUintMap) ToSlice() UintUintTuples {
	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm UintUintMap) OrderedSlice(keys UintList) UintUintTuples {
	s := make(UintUintTuples, 0, len(mm))
	for _, k := range keys {
		v, found := mm[k]
		if found {
			s = append(s, UintUintTuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm UintUintMap) Get(k uint) (uint, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm UintUintMap) Put(k uint, v uint) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm UintUintMap) ContainsKey(k uint) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm UintUintMap) ContainsAllKeys(kk ...uint) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *UintUintMap) Clear() {
	*mm = make(map[uint]uint)
}

// Remove a single item from the map.
func (mm UintUintMap) Remove(k uint) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm UintUintMap) Pop(k uint) (uint, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm UintUintMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm UintUintMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm UintUintMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm UintUintMap) DropWhere(fn func(uint, uint) bool) UintUintTuples {
	removed := make(UintUintTuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, UintUintTuple{(k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm UintUintMap) Foreach(f func(uint, uint)) {
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
func (mm UintUintMap) Forall(p func(uint, uint) bool) bool {
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
func (mm UintUintMap) Exists(p func(uint, uint) bool) bool {
	for k, v := range mm {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first uint that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm UintUintMap) Find(p func(uint, uint) bool) (UintUintTuple, bool) {
	for k, v := range mm {
		if p(k, v) {
			return UintUintTuple{(k), v}, true
		}
	}

	return UintUintTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm UintUintMap) Filter(p func(uint, uint) bool) UintUintMap {
	result := NewUintUintMap()
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
func (mm UintUintMap) Partition(p func(uint, uint) bool) (matching UintUintMap, others UintUintMap) {
	matching = NewUintUintMap()
	others = NewUintUintMap()
	for k, v := range mm {
		if p(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new UintMap by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm UintUintMap) Map(f func(uint, uint) (uint, uint)) UintUintMap {
	result := NewUintUintMap()

	for k1, v1 := range mm {
		k2, v2 := f(k1, v1)
		result[k2] = v2
	}

	return result
}

// FlatMap returns a new UintMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm UintUintMap) FlatMap(f func(uint, uint) []UintUintTuple) UintUintMap {
	result := NewUintUintMap()

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
func (mm UintUintMap) Equals(other UintUintMap) bool {
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
func (mm UintUintMap) Clone() UintUintMap {
	result := NewUintUintMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm UintUintMap) String() string {
	return mm.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm UintUintMap) MkString(sep string) string {
	return mm.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm UintUintMap) MkString4(before, between, after, equals string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString4Bytes(before, between, after, equals).String()
}

func (mm UintUintMap) mkString4Bytes(before, between, after, equals string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""

	keys := make(UintList, 0, len(mm))
	for k, _ := range mm {
		keys = append(keys, k)
	}
	keys.Sorted()

	for _, k := range keys {
		v := mm[k]
		b.WriteString(sep)
		fmt.Fprintf(b, "%v%s%v", k, equals, v)
		sep = between
	}

	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

func (ts UintUintTuples) String() string {
	return ts.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (ts UintUintTuples) MkString(sep string) string {
	return ts.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (ts UintUintTuples) MkString4(before, between, after, equals string) string {
	if ts == nil {
		return ""
	}
	return ts.mkString4Bytes(before, between, after, equals).String()
}

func (ts UintUintTuples) mkString4Bytes(before, between, after, equals string) *strings.Builder {
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
func (t UintUintTuple) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&t)
}

// MarshalJSON implements encoding.Marshaler interface.
func (t UintUintTuple) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"key":"%v", "val":"%v"}`, t.Key, t.Val)), nil
}
