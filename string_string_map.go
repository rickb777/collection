// A simple type derived from map[string]string.
//
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=string Type=string
// options: Comparable:true Stringer:true KeyList:StringList ValueList:StringList Mutable:always
// by runtemplate v3.6.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// StringStringMap is the primary type that represents a map
type StringStringMap map[string]string

// StringStringTuple represents a key/value pair.
type StringStringTuple struct {
	Key string
	Val string
}

// StringStringTuples can be used as a builder for unmodifiable maps.
type StringStringTuples []StringStringTuple

// Append1 adds one item.
func (ts StringStringTuples) Append1(k string, v string) StringStringTuples {
	return append(ts, StringStringTuple{k, v})
}

// Append2 adds two items.
func (ts StringStringTuples) Append2(k1 string, v1 string, k2 string, v2 string) StringStringTuples {
	return append(ts, StringStringTuple{k1, v1}, StringStringTuple{k2, v2})
}

// Append3 adds three items.
func (ts StringStringTuples) Append3(k1 string, v1 string, k2 string, v2 string, k3 string, v3 string) StringStringTuples {
	return append(ts, StringStringTuple{k1, v1}, StringStringTuple{k2, v2}, StringStringTuple{k3, v3})
}

// StringStringZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewStringStringMap
// constructor function.
func StringStringZip(keys ...string) StringStringTuples {
	ts := make(StringStringTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with StringStringZip.
func (ts StringStringTuples) Values(values ...string) StringStringTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts StringStringTuples) ToMap() StringStringMap {
	return NewStringStringMap(ts...)
}

//-------------------------------------------------------------------------------------------------

func newStringStringMap() StringStringMap {
	return StringStringMap(make(map[string]string))
}

// NewStringStringMap1 creates and returns a reference to a map containing one item.
func NewStringStringMap1(k string, v string) StringStringMap {
	mm := newStringStringMap()
	mm[k] = v
	return mm
}

// NewStringStringMap creates and returns a reference to a map, optionally containing some items.
func NewStringStringMap(kv ...StringStringTuple) StringStringMap {
	mm := newStringStringMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm StringStringMap) Keys() StringList {
	if mm == nil {
		return nil
	}

	s := make(StringList, 0, len(mm))
	for k := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm StringStringMap) Values() StringList {
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
func (mm StringStringMap) slice() StringStringTuples {
	s := make(StringStringTuples, 0, len(mm))
	for k, v := range mm {
		s = append(s, StringStringTuple{(k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice.
func (mm StringStringMap) ToSlice() StringStringTuples {
	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm StringStringMap) OrderedSlice(keys StringList) StringStringTuples {
	s := make(StringStringTuples, 0, len(mm))
	for _, k := range keys {
		v, found := mm[k]
		if found {
			s = append(s, StringStringTuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm StringStringMap) Get(k string) (string, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm StringStringMap) Put(k string, v string) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm StringStringMap) ContainsKey(k string) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm StringStringMap) ContainsAllKeys(kk ...string) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *StringStringMap) Clear() {
	*mm = make(map[string]string)
}

// Remove a single item from the map.
func (mm StringStringMap) Remove(k string) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm StringStringMap) Pop(k string) (string, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm StringStringMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm StringStringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm StringStringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm StringStringMap) DropWhere(fn func(string, string) bool) StringStringTuples {
	removed := make(StringStringTuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, StringStringTuple{(k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm StringStringMap) Foreach(f func(string, string)) {
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
func (mm StringStringMap) Forall(p func(string, string) bool) bool {
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
func (mm StringStringMap) Exists(p func(string, string) bool) bool {
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
func (mm StringStringMap) Find(p func(string, string) bool) (StringStringTuple, bool) {
	for k, v := range mm {
		if p(k, v) {
			return StringStringTuple{(k), v}, true
		}
	}

	return StringStringTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm StringStringMap) Filter(p func(string, string) bool) StringStringMap {
	result := NewStringStringMap()
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
func (mm StringStringMap) Partition(p func(string, string) bool) (matching StringStringMap, others StringStringMap) {
	matching = NewStringStringMap()
	others = NewStringStringMap()
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
func (mm StringStringMap) Map(f func(string, string) (string, string)) StringStringMap {
	result := NewStringStringMap()

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
func (mm StringStringMap) FlatMap(f func(string, string) []StringStringTuple) StringStringMap {
	result := NewStringStringMap()

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
func (mm StringStringMap) Equals(other StringStringMap) bool {
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
func (mm StringStringMap) Clone() StringStringMap {
	result := NewStringStringMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm StringStringMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm StringStringMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm StringStringMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm StringStringMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm StringStringMap) mkString3Bytes(before, between, after string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""

	keys := make(StringList, 0, len(mm))
	for k, _ := range mm {
		keys = append(keys, k)
	}
	keys.Sorted()

	for _, k := range keys {
		v := mm[k]
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v:%v", k, v))
		sep = between
	}

	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

func (ts StringStringTuples) String() string {
	return ts.MkString3("[", ", ", "]")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (ts StringStringTuples) MkString(sep string) string {
	return ts.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (ts StringStringTuples) MkString3(before, between, after string) string {
	if ts == nil {
		return ""
	}
	return ts.mkString3Bytes(before, between, after).String()
}

func (ts StringStringTuples) mkString3Bytes(before, between, after string) *strings.Builder {
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
func (t StringStringTuple) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&t)
}

// MarshalJSON implements encoding.Marshaler interface.
func (t StringStringTuple) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"key":"%v", "val":"%v"}`, t.Key, t.Val)), nil
}
