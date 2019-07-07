// An encapsulated immutable map[string]string.
// Thread-safe.
//
//
// Generated from immutable/map.tpl with Key=string Type=string
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:disabled
// by runtemplate v3.4.2
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

// IStringStringMap is the primary type that represents a thread-safe map
type IStringStringMap struct {
	m map[string]string
}

// IStringStringTuple represents a key/value pair.
type IStringStringTuple struct {
	Key string
	Val string
}

// IStringStringTuples can be used as a builder for unmodifiable maps.
type IStringStringTuples []IStringStringTuple

// Append1 adds one item.
func (ts IStringStringTuples) Append1(k string, v string) IStringStringTuples {
	return append(ts, IStringStringTuple{k, v})
}

// Append2 adds two items.
func (ts IStringStringTuples) Append2(k1 string, v1 string, k2 string, v2 string) IStringStringTuples {
	return append(ts, IStringStringTuple{k1, v1}, IStringStringTuple{k2, v2})
}

// Append3 adds three items.
func (ts IStringStringTuples) Append3(k1 string, v1 string, k2 string, v2 string, k3 string, v3 string) IStringStringTuples {
	return append(ts, IStringStringTuple{k1, v1}, IStringStringTuple{k2, v2}, IStringStringTuple{k3, v3})
}

// IStringStringZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewIStringStringMap
// constructor function.
func IStringStringZip(keys ...string) IStringStringTuples {
	ts := make(IStringStringTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with IStringStringZip.
func (ts IStringStringTuples) Values(values ...string) IStringStringTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newIStringStringMap() *IStringStringMap {
	return &IStringStringMap{
		m: make(map[string]string),
	}
}

// NewIStringStringMap1 creates and returns a reference to a map containing one item.
func NewIStringStringMap1(k string, v string) *IStringStringMap {
	mm := newIStringStringMap()
	mm.m[k] = v
	return mm
}

// NewIStringStringMap creates and returns a reference to a map, optionally containing some items.
func NewIStringStringMap(kv ...IStringStringTuple) *IStringStringMap {
	mm := newIStringStringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *IStringStringMap) Keys() []string {
	if mm == nil {
		return nil
	}

	s := make([]string, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *IStringStringMap) Values() []string {
	if mm == nil {
		return nil
	}

	s := make([]string, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *IStringStringMap) slice() []IStringStringTuple {
	if mm == nil {
		return nil
	}

	s := make([]IStringStringTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, IStringStringTuple{k, v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *IStringStringMap) ToSlice() []IStringStringTuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *IStringStringMap) Get(k string) (string, bool) {
	v, found := mm.m[k]
	return v, found
}

// Put adds an item to a clone of the map, replacing any prior value and returning the cloned map.
func (mm *IStringStringMap) Put(k string, v string) *IStringStringMap {
	if mm == nil {
		return NewIStringStringMap1(k, v)
	}

	result := NewIStringStringMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	result.m[k] = v

	return result
}

// ContainsKey determines if a given item is already in the map.
func (mm *IStringStringMap) ContainsKey(k string) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *IStringStringMap) ContainsAllKeys(kk ...string) bool {
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
func (mm *IStringStringMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *IStringStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *IStringStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *IStringStringMap) Foreach(f func(string, string)) {
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
func (mm *IStringStringMap) Forall(f func(string, string) bool) bool {
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
func (mm *IStringStringMap) Exists(p func(string, string) bool) bool {
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
func (mm *IStringStringMap) Find(p func(string, string) bool) (IStringStringTuple, bool) {
	if mm == nil {
		return IStringStringTuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return IStringStringTuple{k, v}, true
		}
	}

	return IStringStringTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *IStringStringMap) Filter(p func(string, string) bool) *IStringStringMap {
	if mm == nil {
		return nil
	}

	result := NewIStringStringMap()

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
func (mm *IStringStringMap) Partition(p func(string, string) bool) (matching *IStringStringMap, others *IStringStringMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewIStringStringMap()
	others = NewIStringStringMap()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new IStringMap by transforming every element with the function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *IStringStringMap) Map(f func(string, string) (string, string)) *IStringStringMap {
	if mm == nil {
		return nil
	}

	result := NewIStringStringMap()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new IStringMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *IStringStringMap) FlatMap(f func(string, string) []IStringStringTuple) *IStringStringMap {
	if mm == nil {
		return nil
	}

	result := NewIStringStringMap()

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
func (mm *IStringStringMap) Equals(other *IStringStringMap) bool {
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
func (mm *IStringStringMap) Clone() *IStringStringMap {
	return mm
}

//-------------------------------------------------------------------------------------------------

func (mm *IStringStringMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *IStringStringMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *IStringStringMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *IStringStringMap) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *IStringStringMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	for k, v := range mm.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}

	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this map type.
func (mm *IStringStringMap) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&mm.m)
}

// MarshalJSON implements JSON encoding for this map type.
func (mm *IStringStringMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(mm.m)
}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this map type.
// You must register string with the 'gob' package before this method is used.
func (mm *IStringStringMap) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register string with the 'gob' package before this method is used.
func (mm *IStringStringMap) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}
