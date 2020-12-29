// An encapsulated immutable map[string]interface{}.
// Thread-safe.
//
//
// Generated from immutable/map.tpl with Key=string Type=interface{}
// options: Comparable:true Stringer:true KeyList:collection.StringList ValueList:collection.AnyList Mutable:disabled
// by runtemplate v3.7.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/rickb777/collection"
	"strings"
)

// StringAnyMap is the primary type that represents a thread-safe map
type StringAnyMap struct {
	m map[string]interface{}
}

// StringAnyTuple represents a key/value pair.
type StringAnyTuple struct {
	Key string
	Val interface{}
}

// StringAnyTuples can be used as a builder for unmodifiable maps.
type StringAnyTuples []StringAnyTuple

// Append1 adds one item.
func (ts StringAnyTuples) Append1(k string, v interface{}) StringAnyTuples {
	return append(ts, StringAnyTuple{k, v})
}

// Append2 adds two items.
func (ts StringAnyTuples) Append2(k1 string, v1 interface{}, k2 string, v2 interface{}) StringAnyTuples {
	return append(ts, StringAnyTuple{k1, v1}, StringAnyTuple{k2, v2})
}

// Append3 adds three items.
func (ts StringAnyTuples) Append3(k1 string, v1 interface{}, k2 string, v2 interface{}, k3 string, v3 interface{}) StringAnyTuples {
	return append(ts, StringAnyTuple{k1, v1}, StringAnyTuple{k2, v2}, StringAnyTuple{k3, v3})
}

// StringAnyZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewStringAnyMap
// constructor function.
func StringAnyZip(keys ...string) StringAnyTuples {
	ts := make(StringAnyTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with StringAnyZip.
func (ts StringAnyTuples) Values(values ...interface{}) StringAnyTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts StringAnyTuples) ToMap() *StringAnyMap {
	return NewStringAnyMap(ts...)
}

//-------------------------------------------------------------------------------------------------

func newStringAnyMap() *StringAnyMap {
	return &StringAnyMap{
		m: make(map[string]interface{}),
	}
}

// NewStringAnyMap1 creates and returns a reference to a map containing one item.
func NewStringAnyMap1(k string, v interface{}) *StringAnyMap {
	mm := newStringAnyMap()
	mm.m[k] = v
	return mm
}

// NewStringAnyMap creates and returns a reference to a map, optionally containing some items.
func NewStringAnyMap(kv ...StringAnyTuple) *StringAnyMap {
	mm := newStringAnyMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *StringAnyMap) Keys() collection.StringList {
	if mm == nil {
		return nil
	}

	s := make(collection.StringList, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *StringAnyMap) Values() collection.AnyList {
	if mm == nil {
		return nil
	}

	s := make(collection.AnyList, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *StringAnyMap) slice() StringAnyTuples {
	if mm == nil {
		return nil
	}

	s := make(StringAnyTuples, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, StringAnyTuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *StringAnyMap) ToSlice() StringAnyTuples {
	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm *StringAnyMap) OrderedSlice(keys collection.StringList) StringAnyTuples {
	if mm == nil {
		return nil
	}

	s := make(StringAnyTuples, 0, len(mm.m))
	for _, k := range keys {
		v, found := mm.m[k]
		if found {
			s = append(s, StringAnyTuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *StringAnyMap) Get(k string) (interface{}, bool) {
	v, found := mm.m[k]
	return v, found
}

// Put adds an item to a clone of the map, replacing any prior value and returning the cloned map.
func (mm *StringAnyMap) Put(k string, v interface{}) *StringAnyMap {
	if mm == nil {
		return NewStringAnyMap1(k, v)
	}

	result := NewStringAnyMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	result.m[k] = v

	return result
}

// ContainsKey determines if a given item is already in the map.
func (mm *StringAnyMap) ContainsKey(k string) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *StringAnyMap) ContainsAllKeys(kk ...string) bool {
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
func (mm *StringAnyMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *StringAnyMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *StringAnyMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *StringAnyMap) Foreach(f func(string, interface{})) {
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
func (mm *StringAnyMap) Forall(f func(string, interface{}) bool) bool {
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
func (mm *StringAnyMap) Exists(p func(string, interface{}) bool) bool {
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

// Find returns the first interface{} that returns true for the predicate p.
// False is returned if none match.
func (mm *StringAnyMap) Find(p func(string, interface{}) bool) (StringAnyTuple, bool) {
	if mm == nil {
		return StringAnyTuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return StringAnyTuple{k, v}, true
		}
	}

	return StringAnyTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *StringAnyMap) Filter(p func(string, interface{}) bool) *StringAnyMap {
	if mm == nil {
		return nil
	}

	result := NewStringAnyMap()

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
func (mm *StringAnyMap) Partition(p func(string, interface{}) bool) (matching *StringAnyMap, others *StringAnyMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewStringAnyMap()
	others = NewStringAnyMap()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new AnyMap by transforming every element with the function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *StringAnyMap) Map(f func(string, interface{}) (string, interface{})) *StringAnyMap {
	if mm == nil {
		return nil
	}

	result := NewStringAnyMap()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new AnyMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *StringAnyMap) FlatMap(f func(string, interface{}) []StringAnyTuple) *StringAnyMap {
	if mm == nil {
		return nil
	}

	result := NewStringAnyMap()

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
func (mm *StringAnyMap) Equals(other *StringAnyMap) bool {
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
func (mm *StringAnyMap) Clone() *StringAnyMap {
	return mm
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm *StringAnyMap) String() string {
	return mm.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *StringAnyMap) MkString(sep string) string {
	return mm.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *StringAnyMap) MkString4(before, between, after, equals string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString4Bytes(before, between, after, equals).String()
}

func (mm *StringAnyMap) mkString4Bytes(before, between, after, equals string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""

	keys := make(collection.StringList, 0, len(mm.m))
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

// UnmarshalJSON implements JSON decoding for this map type.
func (mm *StringAnyMap) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&mm.m)
}

// MarshalJSON implements JSON encoding for this map type.
func (mm *StringAnyMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(mm.m)
}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this map type.
// You must register interface{} with the 'gob' package before this method is used.
func (mm *StringAnyMap) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this map type.
// You must register interface{} with the 'gob' package before this method is used.
func (mm *StringAnyMap) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}

//-------------------------------------------------------------------------------------------------

func (ts StringAnyTuples) String() string {
	return ts.MkString4("[", ", ", "]", ":")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (ts StringAnyTuples) MkString(sep string) string {
	return ts.MkString4("", sep, "", ":")
}

// MkString4 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (ts StringAnyTuples) MkString4(before, between, after, equals string) string {
	if ts == nil {
		return ""
	}
	return ts.mkString4Bytes(before, between, after, equals).String()
}

func (ts StringAnyTuples) mkString4Bytes(before, between, after, equals string) *strings.Builder {
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
func (t StringAnyTuple) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&t)
}

// MarshalJSON implements encoding.Marshaler interface.
func (t StringAnyTuple) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"key":"%v", "val":"%v"}`, t.Key, t.Val)), nil
}
