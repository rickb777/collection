// An encapsulated immutable map[uint]string.
// Thread-safe.
//
//
// Generated from immutable/map.tpl with Key=uint Type=string
// options: Comparable:true Stringer:true KeyList:collection.UintList ValueList:collection.StringList Mutable:disabled
// by runtemplate v3.10.1
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

// UintStringMap is the primary type that represents a thread-safe map
type UintStringMap struct {
	m map[uint]string
}

// UintStringTuple represents a key/value pair.
type UintStringTuple struct {
	Key uint
	Val string
}

// UintStringTuples can be used as a builder for unmodifiable maps.
type UintStringTuples []UintStringTuple

// Append1 adds one item.
func (ts UintStringTuples) Append1(k uint, v string) UintStringTuples {
	return append(ts, UintStringTuple{k, v})
}

// Append2 adds two items.
func (ts UintStringTuples) Append2(k1 uint, v1 string, k2 uint, v2 string) UintStringTuples {
	return append(ts, UintStringTuple{k1, v1}, UintStringTuple{k2, v2})
}

// Append3 adds three items.
func (ts UintStringTuples) Append3(k1 uint, v1 string, k2 uint, v2 string, k3 uint, v3 string) UintStringTuples {
	return append(ts, UintStringTuple{k1, v1}, UintStringTuple{k2, v2}, UintStringTuple{k3, v3})
}

// UintStringZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewUintStringMap
// constructor function.
func UintStringZip(keys ...uint) UintStringTuples {
	ts := make(UintStringTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with UintStringZip.
func (ts UintStringTuples) Values(values ...string) UintStringTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts UintStringTuples) ToMap() *UintStringMap {
	return NewUintStringMap(ts...)
}

//-------------------------------------------------------------------------------------------------

func newUintStringMap() *UintStringMap {
	return &UintStringMap{
		m: make(map[uint]string),
	}
}

// NewUintStringMap1 creates and returns a reference to a map containing one item.
func NewUintStringMap1(k uint, v string) *UintStringMap {
	mm := newUintStringMap()
	mm.m[k] = v
	return mm
}

// NewUintStringMap creates and returns a reference to a map, optionally containing some items.
func NewUintStringMap(kv ...UintStringTuple) *UintStringMap {
	mm := newUintStringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *UintStringMap) Keys() collection.UintList {
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
func (mm *UintStringMap) Values() collection.StringList {
	if mm == nil {
		return nil
	}

	s := make(collection.StringList, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *UintStringMap) slice() UintStringTuples {
	if mm == nil {
		return nil
	}

	s := make(UintStringTuples, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, UintStringTuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *UintStringMap) ToSlice() UintStringTuples {
	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm *UintStringMap) OrderedSlice(keys collection.UintList) UintStringTuples {
	if mm == nil {
		return nil
	}

	s := make(UintStringTuples, 0, len(mm.m))
	for _, k := range keys {
		v, found := mm.m[k]
		if found {
			s = append(s, UintStringTuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *UintStringMap) Get(k uint) (string, bool) {
	v, found := mm.m[k]
	return v, found
}

// Put adds an item to a clone of the map, replacing any prior value and returning the cloned map.
func (mm *UintStringMap) Put(k uint, v string) *UintStringMap {
	if mm == nil {
		return NewUintStringMap1(k, v)
	}

	result := NewUintStringMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	result.m[k] = v

	return result
}

// ContainsKey determines if a given item is already in the map.
func (mm *UintStringMap) ContainsKey(k uint) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *UintStringMap) ContainsAllKeys(kk ...uint) bool {
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
func (mm *UintStringMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *UintStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *UintStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *UintStringMap) Foreach(f func(uint, string)) {
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
func (mm *UintStringMap) Forall(f func(uint, string) bool) bool {
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
func (mm *UintStringMap) Exists(p func(uint, string) bool) bool {
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

// Find returns the first string that returns true for the predicate p.
// False is returned if none match.
func (mm *UintStringMap) Find(p func(uint, string) bool) (UintStringTuple, bool) {
	if mm == nil {
		return UintStringTuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return UintStringTuple{k, v}, true
		}
	}

	return UintStringTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *UintStringMap) Filter(p func(uint, string) bool) *UintStringMap {
	if mm == nil {
		return nil
	}

	result := NewUintStringMap()

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
func (mm *UintStringMap) Partition(p func(uint, string) bool) (matching *UintStringMap, others *UintStringMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewUintStringMap()
	others = NewUintStringMap()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new StringMap by transforming every element with the function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *UintStringMap) Map(f func(uint, string) (uint, string)) *UintStringMap {
	if mm == nil {
		return nil
	}

	result := NewUintStringMap()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new StringMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *UintStringMap) FlatMap(f func(uint, string) []UintStringTuple) *UintStringMap {
	if mm == nil {
		return nil
	}

	result := NewUintStringMap()

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
func (mm *UintStringMap) Equals(other *UintStringMap) bool {
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
func (mm *UintStringMap) Clone() *UintStringMap {
	return mm
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm *UintStringMap) String() string {
	return mm.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *UintStringMap) MkString(sep string) string {
	return mm.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *UintStringMap) MkString4(before, between, after, equals string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString4Bytes(before, between, after, equals).String()
}

func (mm *UintStringMap) mkString4Bytes(before, between, after, equals string) *strings.Builder {
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
// You must register string with the 'gob' package before this method is used.
func (mm *UintStringMap) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this map type.
// You must register string with the 'gob' package before this method is used.
func (mm *UintStringMap) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}

//-------------------------------------------------------------------------------------------------

func (ts UintStringTuples) String() string {
	return ts.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (ts UintStringTuples) MkString(sep string) string {
	return ts.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (ts UintStringTuples) MkString4(before, between, after, equals string) string {
	if ts == nil {
		return ""
	}
	return ts.mkString4Bytes(before, between, after, equals).String()
}

func (ts UintStringTuples) mkString4Bytes(before, between, after, equals string) *strings.Builder {
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
func (t UintStringTuple) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&t)
}

// MarshalJSON implements encoding.Marshaler interface.
func (t UintStringTuple) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"key":"%v", "val":"%v"}`, t.Key, t.Val)), nil
}
