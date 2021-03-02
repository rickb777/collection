// A simple type derived from map[uint64]uint64.
//
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=uint64 Type=uint64
// options: Comparable:true Stringer:true KeyList:Uint64List ValueList:Uint64List Mutable:always
// by runtemplate v3.10.0
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package collection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// Uint64Uint64Map is the primary type that represents a map
type Uint64Uint64Map map[uint64]uint64

// Uint64Uint64Tuple represents a key/value pair.
type Uint64Uint64Tuple struct {
	Key uint64
	Val uint64
}

// Uint64Uint64Tuples can be used as a builder for unmodifiable maps.
type Uint64Uint64Tuples []Uint64Uint64Tuple

// Append1 adds one item.
func (ts Uint64Uint64Tuples) Append1(k uint64, v uint64) Uint64Uint64Tuples {
	return append(ts, Uint64Uint64Tuple{k, v})
}

// Append2 adds two items.
func (ts Uint64Uint64Tuples) Append2(k1 uint64, v1 uint64, k2 uint64, v2 uint64) Uint64Uint64Tuples {
	return append(ts, Uint64Uint64Tuple{k1, v1}, Uint64Uint64Tuple{k2, v2})
}

// Append3 adds three items.
func (ts Uint64Uint64Tuples) Append3(k1 uint64, v1 uint64, k2 uint64, v2 uint64, k3 uint64, v3 uint64) Uint64Uint64Tuples {
	return append(ts, Uint64Uint64Tuple{k1, v1}, Uint64Uint64Tuple{k2, v2}, Uint64Uint64Tuple{k3, v3})
}

// Uint64Uint64Zip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewUint64Uint64Map
// constructor function.
func Uint64Uint64Zip(keys ...uint64) Uint64Uint64Tuples {
	ts := make(Uint64Uint64Tuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with Uint64Uint64Zip.
func (ts Uint64Uint64Tuples) Values(values ...uint64) Uint64Uint64Tuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts Uint64Uint64Tuples) ToMap() Uint64Uint64Map {
	return NewUint64Uint64Map(ts...)
}

//-------------------------------------------------------------------------------------------------

func newUint64Uint64Map() Uint64Uint64Map {
	return Uint64Uint64Map(make(map[uint64]uint64))
}

// NewUint64Uint64Map1 creates and returns a reference to a map containing one item.
func NewUint64Uint64Map1(k uint64, v uint64) Uint64Uint64Map {
	mm := newUint64Uint64Map()
	mm[k] = v
	return mm
}

// NewUint64Uint64Map creates and returns a reference to a map, optionally containing some items.
func NewUint64Uint64Map(kv ...Uint64Uint64Tuple) Uint64Uint64Map {
	mm := newUint64Uint64Map()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm Uint64Uint64Map) Keys() Uint64List {
	if mm == nil {
		return nil
	}

	s := make(Uint64List, 0, len(mm))
	for k := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm Uint64Uint64Map) Values() Uint64List {
	if mm == nil {
		return nil
	}

	s := make(Uint64List, 0, len(mm))
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm Uint64Uint64Map) slice() Uint64Uint64Tuples {
	s := make(Uint64Uint64Tuples, 0, len(mm))
	for k, v := range mm {
		s = append(s, Uint64Uint64Tuple{(k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice.
func (mm Uint64Uint64Map) ToSlice() Uint64Uint64Tuples {
	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm Uint64Uint64Map) OrderedSlice(keys Uint64List) Uint64Uint64Tuples {
	s := make(Uint64Uint64Tuples, 0, len(mm))
	for _, k := range keys {
		v, found := mm[k]
		if found {
			s = append(s, Uint64Uint64Tuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm Uint64Uint64Map) Get(k uint64) (uint64, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm Uint64Uint64Map) Put(k uint64, v uint64) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm Uint64Uint64Map) ContainsKey(k uint64) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm Uint64Uint64Map) ContainsAllKeys(kk ...uint64) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *Uint64Uint64Map) Clear() {
	*mm = make(map[uint64]uint64)
}

// Remove a single item from the map.
func (mm Uint64Uint64Map) Remove(k uint64) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm Uint64Uint64Map) Pop(k uint64) (uint64, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm Uint64Uint64Map) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm Uint64Uint64Map) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm Uint64Uint64Map) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm Uint64Uint64Map) DropWhere(fn func(uint64, uint64) bool) Uint64Uint64Tuples {
	removed := make(Uint64Uint64Tuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, Uint64Uint64Tuple{(k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm Uint64Uint64Map) Foreach(f func(uint64, uint64)) {
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
func (mm Uint64Uint64Map) Forall(p func(uint64, uint64) bool) bool {
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
func (mm Uint64Uint64Map) Exists(p func(uint64, uint64) bool) bool {
	for k, v := range mm {
		if p(k, v) {
			return true
		}
	}

	return false
}

// Find returns the first uint64 that returns true for the predicate p.
// False is returned if none match.
// The original map is not modified.
func (mm Uint64Uint64Map) Find(p func(uint64, uint64) bool) (Uint64Uint64Tuple, bool) {
	for k, v := range mm {
		if p(k, v) {
			return Uint64Uint64Tuple{(k), v}, true
		}
	}

	return Uint64Uint64Tuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm Uint64Uint64Map) Filter(p func(uint64, uint64) bool) Uint64Uint64Map {
	result := NewUint64Uint64Map()
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
func (mm Uint64Uint64Map) Partition(p func(uint64, uint64) bool) (matching Uint64Uint64Map, others Uint64Uint64Map) {
	matching = NewUint64Uint64Map()
	others = NewUint64Uint64Map()
	for k, v := range mm {
		if p(k, v) {
			matching[k] = v
		} else {
			others[k] = v
		}
	}
	return
}

// Map returns a new Uint64Map by transforming every element with the function f.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm Uint64Uint64Map) Map(f func(uint64, uint64) (uint64, uint64)) Uint64Uint64Map {
	result := NewUint64Uint64Map()

	for k1, v1 := range mm {
		k2, v2 := f(k1, v1)
		result[k2] = v2
	}

	return result
}

// FlatMap returns a new Uint64Map by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
// The original map is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm Uint64Uint64Map) FlatMap(f func(uint64, uint64) []Uint64Uint64Tuple) Uint64Uint64Map {
	result := NewUint64Uint64Map()

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
func (mm Uint64Uint64Map) Equals(other Uint64Uint64Map) bool {
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
func (mm Uint64Uint64Map) Clone() Uint64Uint64Map {
	result := NewUint64Uint64Map()
	for k, v := range mm {
		result[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm Uint64Uint64Map) String() string {
	return mm.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm Uint64Uint64Map) MkString(sep string) string {
	return mm.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm Uint64Uint64Map) MkString4(before, between, after, equals string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString4Bytes(before, between, after, equals).String()
}

func (mm Uint64Uint64Map) mkString4Bytes(before, between, after, equals string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""

	keys := make(Uint64List, 0, len(mm))
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

func (ts Uint64Uint64Tuples) String() string {
	return ts.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (ts Uint64Uint64Tuples) MkString(sep string) string {
	return ts.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (ts Uint64Uint64Tuples) MkString4(before, between, after, equals string) string {
	if ts == nil {
		return ""
	}
	return ts.mkString4Bytes(before, between, after, equals).String()
}

func (ts Uint64Uint64Tuples) mkString4Bytes(before, between, after, equals string) *strings.Builder {
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
func (t Uint64Uint64Tuple) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&t)
}

// MarshalJSON implements encoding.Marshaler interface.
func (t Uint64Uint64Tuple) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"key":"%v", "val":"%v"}`, t.Key, t.Val)), nil
}
