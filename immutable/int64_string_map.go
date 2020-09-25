// An encapsulated immutable map[int64]string.
// Thread-safe.
//
//
// Generated from immutable/map.tpl with Key=int64 Type=string
// options: Comparable:true Stringer:true KeyList:collection.Int64List ValueList:collection.StringList Mutable:disabled
// by runtemplate v3.6.1
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

// Int64StringMap is the primary type that represents a thread-safe map
type Int64StringMap struct {
	m map[int64]string
}

// Int64StringTuple represents a key/value pair.
type Int64StringTuple struct {
	Key int64
	Val string
}

// Int64StringTuples can be used as a builder for unmodifiable maps.
type Int64StringTuples []Int64StringTuple

// Append1 adds one item.
func (ts Int64StringTuples) Append1(k int64, v string) Int64StringTuples {
	return append(ts, Int64StringTuple{k, v})
}

// Append2 adds two items.
func (ts Int64StringTuples) Append2(k1 int64, v1 string, k2 int64, v2 string) Int64StringTuples {
	return append(ts, Int64StringTuple{k1, v1}, Int64StringTuple{k2, v2})
}

// Append3 adds three items.
func (ts Int64StringTuples) Append3(k1 int64, v1 string, k2 int64, v2 string, k3 int64, v3 string) Int64StringTuples {
	return append(ts, Int64StringTuple{k1, v1}, Int64StringTuple{k2, v2}, Int64StringTuple{k3, v3})
}

// Int64StringZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewInt64StringMap
// constructor function.
func Int64StringZip(keys ...int64) Int64StringTuples {
	ts := make(Int64StringTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with Int64StringZip.
func (ts Int64StringTuples) Values(values ...string) Int64StringTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts Int64StringTuples) ToMap() *Int64StringMap {
	return NewInt64StringMap(ts...)
}

//-------------------------------------------------------------------------------------------------

func newInt64StringMap() *Int64StringMap {
	return &Int64StringMap{
		m: make(map[int64]string),
	}
}

// NewInt64StringMap1 creates and returns a reference to a map containing one item.
func NewInt64StringMap1(k int64, v string) *Int64StringMap {
	mm := newInt64StringMap()
	mm.m[k] = v
	return mm
}

// NewInt64StringMap creates and returns a reference to a map, optionally containing some items.
func NewInt64StringMap(kv ...Int64StringTuple) *Int64StringMap {
	mm := newInt64StringMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *Int64StringMap) Keys() collection.Int64List {
	if mm == nil {
		return nil
	}

	s := make(collection.Int64List, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *Int64StringMap) Values() collection.StringList {
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
func (mm *Int64StringMap) slice() Int64StringTuples {
	if mm == nil {
		return nil
	}

	s := make(Int64StringTuples, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, Int64StringTuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *Int64StringMap) ToSlice() Int64StringTuples {
	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm *Int64StringMap) OrderedSlice(keys collection.Int64List) Int64StringTuples {
	if mm == nil {
		return nil
	}

	s := make(Int64StringTuples, 0, len(mm.m))
	for _, k := range keys {
		v, found := mm.m[k]
		if found {
			s = append(s, Int64StringTuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *Int64StringMap) Get(k int64) (string, bool) {
	v, found := mm.m[k]
	return v, found
}

// Put adds an item to a clone of the map, replacing any prior value and returning the cloned map.
func (mm *Int64StringMap) Put(k int64, v string) *Int64StringMap {
	if mm == nil {
		return NewInt64StringMap1(k, v)
	}

	result := NewInt64StringMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	result.m[k] = v

	return result
}

// ContainsKey determines if a given item is already in the map.
func (mm *Int64StringMap) ContainsKey(k int64) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *Int64StringMap) ContainsAllKeys(kk ...int64) bool {
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
func (mm *Int64StringMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *Int64StringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *Int64StringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *Int64StringMap) Foreach(f func(int64, string)) {
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
func (mm *Int64StringMap) Forall(f func(int64, string) bool) bool {
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
func (mm *Int64StringMap) Exists(p func(int64, string) bool) bool {
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
func (mm *Int64StringMap) Find(p func(int64, string) bool) (Int64StringTuple, bool) {
	if mm == nil {
		return Int64StringTuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return Int64StringTuple{k, v}, true
		}
	}

	return Int64StringTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *Int64StringMap) Filter(p func(int64, string) bool) *Int64StringMap {
	if mm == nil {
		return nil
	}

	result := NewInt64StringMap()

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
func (mm *Int64StringMap) Partition(p func(int64, string) bool) (matching *Int64StringMap, others *Int64StringMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewInt64StringMap()
	others = NewInt64StringMap()

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
func (mm *Int64StringMap) Map(f func(int64, string) (int64, string)) *Int64StringMap {
	if mm == nil {
		return nil
	}

	result := NewInt64StringMap()

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
func (mm *Int64StringMap) FlatMap(f func(int64, string) []Int64StringTuple) *Int64StringMap {
	if mm == nil {
		return nil
	}

	result := NewInt64StringMap()

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
func (mm *Int64StringMap) Equals(other *Int64StringMap) bool {
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
func (mm *Int64StringMap) Clone() *Int64StringMap {
	return mm
}

//-------------------------------------------------------------------------------------------------

func (mm *Int64StringMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *Int64StringMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *Int64StringMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *Int64StringMap) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *Int64StringMap) mkString3Bytes(before, between, after string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""

	keys := make(collection.Int64List, 0, len(mm.m))
	for k, _ := range mm.m {
		keys = append(keys, k)
	}
	keys.Sorted()

	for _, k := range keys {
		v := mm.m[k]
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}

	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this map type.
// You must register string with the 'gob' package before this method is used.
func (mm *Int64StringMap) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register string with the 'gob' package before this method is used.
func (mm *Int64StringMap) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}

//-------------------------------------------------------------------------------------------------

func (ts Int64StringTuples) String() string {
	return ts.MkString3("[", ", ", "]")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (ts Int64StringTuples) MkString(sep string) string {
	return ts.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (ts Int64StringTuples) MkString3(before, between, after string) string {
	if ts == nil {
		return ""
	}
	return ts.mkString3Bytes(before, between, after).String()
}

func (ts Int64StringTuples) mkString3Bytes(before, between, after string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""
	for _, t := range ts {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", t.Key, t.Val))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this tuple type.
func (t Int64StringTuple) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&t)
}

// MarshalJSON implements encoding.Marshaler interface.
func (t Int64StringTuple) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"key":"%v", "val":"%v"}`, t.Key, t.Val)), nil
}
