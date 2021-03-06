// A simple type derived from map[uint64]string.
//
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=uint64 Type=string
// options: Comparable:true Stringer:true KeyList:Uint64List ValueList:StringList Mutable:always
// by runtemplate v3.10.1
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package collection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// Uint64StringMap is the primary type that represents a map
type Uint64StringMap map[uint64]string

// Uint64StringTuple represents a key/value pair.
type Uint64StringTuple struct {
	Key uint64
	Val string
}

// Uint64StringTuples can be used as a builder for unmodifiable maps.
type Uint64StringTuples []Uint64StringTuple

// Append1 adds one item.
func (ts Uint64StringTuples) Append1(k uint64, v string) Uint64StringTuples {
	return append(ts, Uint64StringTuple{k, v})
}

// Append2 adds two items.
func (ts Uint64StringTuples) Append2(k1 uint64, v1 string, k2 uint64, v2 string) Uint64StringTuples {
	return append(ts, Uint64StringTuple{k1, v1}, Uint64StringTuple{k2, v2})
}

// Append3 adds three items.
func (ts Uint64StringTuples) Append3(k1 uint64, v1 string, k2 uint64, v2 string, k3 uint64, v3 string) Uint64StringTuples {
	return append(ts, Uint64StringTuple{k1, v1}, Uint64StringTuple{k2, v2}, Uint64StringTuple{k3, v3})
}

// Uint64StringZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewUint64StringMap
// constructor function.
func Uint64StringZip(keys ...uint64) Uint64StringTuples {
	ts := make(Uint64StringTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with Uint64StringZip.
func (ts Uint64StringTuples) Values(values ...string) Uint64StringTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts Uint64StringTuples) ToMap() Uint64StringMap {
	return NewUint64StringMap(ts...)
}

//-------------------------------------------------------------------------------------------------

func newUint64StringMap() Uint64StringMap {
	return Uint64StringMap(make(map[uint64]string))
}

// NewUint64StringMap1 creates and returns a reference to a map containing one item.
func NewUint64StringMap1(k uint64, v string) Uint64StringMap {
	mm := newUint64StringMap()
	mm[k] = v
	return mm
}

// NewUint64StringMap creates and returns a reference to a map, optionally containing some items.
func NewUint64StringMap(kv ...Uint64StringTuple) Uint64StringMap {
	mm := newUint64StringMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm Uint64StringMap) Keys() Uint64List {
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
func (mm Uint64StringMap) Values() StringList {
	if mm == nil {
		return nil
	}

	s := make(StringList, 0, len(mm))
	for _, v := range mm {
		s = append(s, v)
	}
	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm Uint64StringMap) slice() Uint64StringTuples {
	s := make(Uint64StringTuples, 0, len(mm))
	for k, v := range mm {
		s = append(s, Uint64StringTuple{(k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice.
func (mm Uint64StringMap) ToSlice() Uint64StringTuples {
	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm Uint64StringMap) OrderedSlice(keys Uint64List) Uint64StringTuples {
	s := make(Uint64StringTuples, 0, len(mm))
	for _, k := range keys {
		v, found := mm[k]
		if found {
			s = append(s, Uint64StringTuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm Uint64StringMap) Get(k uint64) (string, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm Uint64StringMap) Put(k uint64, v string) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm Uint64StringMap) ContainsKey(k uint64) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm Uint64StringMap) ContainsAllKeys(kk ...uint64) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *Uint64StringMap) Clear() {
	*mm = make(map[uint64]string)
}

// Remove a single item from the map.
func (mm Uint64StringMap) Remove(k uint64) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm Uint64StringMap) Pop(k uint64) (string, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm Uint64StringMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm Uint64StringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm Uint64StringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm Uint64StringMap) DropWhere(fn func(uint64, string) bool) Uint64StringTuples {
	removed := make(Uint64StringTuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, Uint64StringTuple{(k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm Uint64StringMap) Foreach(f func(uint64, string)) {
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
func (mm Uint64StringMap) Forall(p func(uint64, string) bool) bool {
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
func (mm Uint64StringMap) Exists(p func(uint64, string) bool) bool {
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
func (mm Uint64StringMap) Find(p func(uint64, string) bool) (Uint64StringTuple, bool) {
	for k, v := range mm {
		if p(k, v) {
			return Uint64StringTuple{(k), v}, true
		}
	}

	return Uint64StringTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm Uint64StringMap) Filter(p func(uint64, string) bool) Uint64StringMap {
	result := NewUint64StringMap()
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
func (mm Uint64StringMap) Partition(p func(uint64, string) bool) (matching Uint64StringMap, others Uint64StringMap) {
	matching = NewUint64StringMap()
	others = NewUint64StringMap()
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
func (mm Uint64StringMap) Map(f func(uint64, string) (uint64, string)) Uint64StringMap {
	result := NewUint64StringMap()

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
func (mm Uint64StringMap) FlatMap(f func(uint64, string) []Uint64StringTuple) Uint64StringMap {
	result := NewUint64StringMap()

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
func (mm Uint64StringMap) Equals(other Uint64StringMap) bool {
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
func (mm Uint64StringMap) Clone() Uint64StringMap {
	result := NewUint64StringMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm Uint64StringMap) String() string {
	return mm.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm Uint64StringMap) MkString(sep string) string {
	return mm.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm Uint64StringMap) MkString4(before, between, after, equals string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString4Bytes(before, between, after, equals).String()
}

func (mm Uint64StringMap) mkString4Bytes(before, between, after, equals string) *strings.Builder {
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

func (ts Uint64StringTuples) String() string {
	return ts.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (ts Uint64StringTuples) MkString(sep string) string {
	return ts.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (ts Uint64StringTuples) MkString4(before, between, after, equals string) string {
	if ts == nil {
		return ""
	}
	return ts.mkString4Bytes(before, between, after, equals).String()
}

func (ts Uint64StringTuples) mkString4Bytes(before, between, after, equals string) *strings.Builder {
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
func (t Uint64StringTuple) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&t)
}

// MarshalJSON implements encoding.Marshaler interface.
func (t Uint64StringTuple) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"key":"%v", "val":"%v"}`, t.Key, t.Val)), nil
}
