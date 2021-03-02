// An encapsulated immutable map[uint]uint.
// Thread-safe.
//
//
// Generated from immutable/map.tpl with Key=uint Type=uint
// options: Comparable:true Stringer:true KeyList:collection.UintList ValueList:collection.UintList Mutable:disabled
// by runtemplate v3.10.0
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package immutable

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/rickb777/collection"
	"strings"
)

// UintUintMap is the primary type that represents a thread-safe map
type UintUintMap struct {
	m map[uint]uint
}

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
func (ts UintUintTuples) ToMap() *UintUintMap {
	return NewUintUintMap(ts...)
}

//-------------------------------------------------------------------------------------------------

func newUintUintMap() *UintUintMap {
	return &UintUintMap{
		m: make(map[uint]uint),
	}
}

// NewUintUintMap1 creates and returns a reference to a map containing one item.
func NewUintUintMap1(k uint, v uint) *UintUintMap {
	mm := newUintUintMap()
	mm.m[k] = v
	return mm
}

// NewUintUintMap creates and returns a reference to a map, optionally containing some items.
func NewUintUintMap(kv ...UintUintTuple) *UintUintMap {
	mm := newUintUintMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *UintUintMap) Keys() collection.UintList {
	if mm == nil {
		return nil
	}

	s := make(collection.UintList, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *UintUintMap) Values() collection.UintList {
	if mm == nil {
		return nil
	}

	s := make(collection.UintList, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *UintUintMap) slice() UintUintTuples {
	if mm == nil {
		return nil
	}

	s := make(UintUintTuples, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, UintUintTuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *UintUintMap) ToSlice() UintUintTuples {
	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm *UintUintMap) OrderedSlice(keys collection.UintList) UintUintTuples {
	if mm == nil {
		return nil
	}

	s := make(UintUintTuples, 0, len(mm.m))
	for _, k := range keys {
		v, found := mm.m[k]
		if found {
			s = append(s, UintUintTuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *UintUintMap) Get(k uint) (uint, bool) {
	v, found := mm.m[k]
	return v, found
}

// Put adds an item to a clone of the map, replacing any prior value and returning the cloned map.
func (mm *UintUintMap) Put(k uint, v uint) *UintUintMap {
	if mm == nil {
		return NewUintUintMap1(k, v)
	}

	result := NewUintUintMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	result.m[k] = v

	return result
}

// ContainsKey determines if a given item is already in the map.
func (mm *UintUintMap) ContainsKey(k uint) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *UintUintMap) ContainsAllKeys(kk ...uint) bool {
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
func (mm *UintUintMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *UintUintMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *UintUintMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *UintUintMap) Foreach(f func(uint, uint)) {
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
func (mm *UintUintMap) Forall(f func(uint, uint) bool) bool {
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
func (mm *UintUintMap) Exists(p func(uint, uint) bool) bool {
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

// Find returns the first uint that returns true for the predicate p.
// False is returned if none match.
func (mm *UintUintMap) Find(p func(uint, uint) bool) (UintUintTuple, bool) {
	if mm == nil {
		return UintUintTuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return UintUintTuple{k, v}, true
		}
	}

	return UintUintTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *UintUintMap) Filter(p func(uint, uint) bool) *UintUintMap {
	if mm == nil {
		return nil
	}

	result := NewUintUintMap()

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
func (mm *UintUintMap) Partition(p func(uint, uint) bool) (matching *UintUintMap, others *UintUintMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewUintUintMap()
	others = NewUintUintMap()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new UintMap by transforming every element with the function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *UintUintMap) Map(f func(uint, uint) (uint, uint)) *UintUintMap {
	if mm == nil {
		return nil
	}

	result := NewUintUintMap()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new UintMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *UintUintMap) FlatMap(f func(uint, uint) []UintUintTuple) *UintUintMap {
	if mm == nil {
		return nil
	}

	result := NewUintUintMap()

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
func (mm *UintUintMap) Equals(other *UintUintMap) bool {
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
func (mm *UintUintMap) Clone() *UintUintMap {
	return mm
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm *UintUintMap) String() string {
	return mm.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *UintUintMap) MkString(sep string) string {
	return mm.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *UintUintMap) MkString4(before, between, after, equals string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString4Bytes(before, between, after, equals).String()
}

func (mm *UintUintMap) mkString4Bytes(before, between, after, equals string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""

	keys := make(collection.UintList, 0, len(mm.m))
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
// You must register uint with the 'gob' package before this method is used.
func (mm *UintUintMap) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this map type.
// You must register uint with the 'gob' package before this method is used.
func (mm *UintUintMap) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
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
