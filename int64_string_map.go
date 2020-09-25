// A simple type derived from map[int64]string.
//
// Not thread-safe.
//
// Generated from simple/map.tpl with Key=int64 Type=string
// options: Comparable:true Stringer:true KeyList:Int64List ValueList:StringList Mutable:always
// by runtemplate v3.6.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// Int64StringMap is the primary type that represents a map
type Int64StringMap map[int64]string

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
func (ts Int64StringTuples) ToMap() Int64StringMap {
	return NewInt64StringMap(ts...)
}

//-------------------------------------------------------------------------------------------------

func newInt64StringMap() Int64StringMap {
	return Int64StringMap(make(map[int64]string))
}

// NewInt64StringMap1 creates and returns a reference to a map containing one item.
func NewInt64StringMap1(k int64, v string) Int64StringMap {
	mm := newInt64StringMap()
	mm[k] = v
	return mm
}

// NewInt64StringMap creates and returns a reference to a map, optionally containing some items.
func NewInt64StringMap(kv ...Int64StringTuple) Int64StringMap {
	mm := newInt64StringMap()
	for _, t := range kv {
		mm[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm Int64StringMap) Keys() Int64List {
	if mm == nil {
		return nil
	}

	s := make(Int64List, 0, len(mm))
	for k := range mm {
		s = append(s, k)
	}
	return s
}

// Values returns the values of the current map as a slice.
func (mm Int64StringMap) Values() StringList {
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
func (mm Int64StringMap) slice() Int64StringTuples {
	s := make(Int64StringTuples, 0, len(mm))
	for k, v := range mm {
		s = append(s, Int64StringTuple{(k), v})
	}
	return s
}

// ToSlice returns the key/value pairs as a slice.
func (mm Int64StringMap) ToSlice() Int64StringTuples {
	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm Int64StringMap) OrderedSlice(keys Int64List) Int64StringTuples {
	s := make(Int64StringTuples, 0, len(mm))
	for _, k := range keys {
		v, found := mm[k]
		if found {
			s = append(s, Int64StringTuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm Int64StringMap) Get(k int64) (string, bool) {
	v, found := mm[k]
	return v, found
}

// Put adds an item to the current map, replacing any prior value.
func (mm Int64StringMap) Put(k int64, v string) bool {
	_, found := mm[k]
	mm[k] = v
	return !found //False if it existed already
}

// ContainsKey determines if a given item is already in the map.
func (mm Int64StringMap) ContainsKey(k int64) bool {
	_, found := mm[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm Int64StringMap) ContainsAllKeys(kk ...int64) bool {
	for _, k := range kk {
		if !mm.ContainsKey(k) {
			return false
		}
	}
	return true
}

// Clear clears the entire map.
func (mm *Int64StringMap) Clear() {
	*mm = make(map[int64]string)
}

// Remove a single item from the map.
func (mm Int64StringMap) Remove(k int64) {
	delete(mm, k)
}

// Pop removes a single item from the map, returning the value present prior to removal.
func (mm Int64StringMap) Pop(k int64) (string, bool) {
	v, found := mm[k]
	delete(mm, k)
	return v, found
}

// Size returns how many items are currently in the map. This is a synonym for Len.
func (mm Int64StringMap) Size() int {
	return len(mm)
}

// IsEmpty returns true if the map is empty.
func (mm Int64StringMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm Int64StringMap) NonEmpty() bool {
	return mm.Size() > 0
}

// DropWhere applies a predicate function to every element in the map. If the function returns true,
// the element is dropped from the map.
// This is similar to Filter except that the map is modified.
func (mm Int64StringMap) DropWhere(fn func(int64, string) bool) Int64StringTuples {
	removed := make(Int64StringTuples, 0)
	for k, v := range mm {
		if fn(k, v) {
			removed = append(removed, Int64StringTuple{(k), v})
			delete(mm, k)
		}
	}
	return removed
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm Int64StringMap) Foreach(f func(int64, string)) {
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
func (mm Int64StringMap) Forall(p func(int64, string) bool) bool {
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
func (mm Int64StringMap) Exists(p func(int64, string) bool) bool {
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
func (mm Int64StringMap) Find(p func(int64, string) bool) (Int64StringTuple, bool) {
	for k, v := range mm {
		if p(k, v) {
			return Int64StringTuple{(k), v}, true
		}
	}

	return Int64StringTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
// The original map is not modified.
func (mm Int64StringMap) Filter(p func(int64, string) bool) Int64StringMap {
	result := NewInt64StringMap()
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
func (mm Int64StringMap) Partition(p func(int64, string) bool) (matching Int64StringMap, others Int64StringMap) {
	matching = NewInt64StringMap()
	others = NewInt64StringMap()
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
func (mm Int64StringMap) Map(f func(int64, string) (int64, string)) Int64StringMap {
	result := NewInt64StringMap()

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
func (mm Int64StringMap) FlatMap(f func(int64, string) []Int64StringTuple) Int64StringMap {
	result := NewInt64StringMap()

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
func (mm Int64StringMap) Equals(other Int64StringMap) bool {
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
func (mm Int64StringMap) Clone() Int64StringMap {
	result := NewInt64StringMap()
	for k, v := range mm {
		result[k] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (mm Int64StringMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm Int64StringMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm Int64StringMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm Int64StringMap) MkString3(before, between, after string) string {
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm Int64StringMap) mkString3Bytes(before, between, after string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""

	keys := make(Int64List, 0, len(mm))
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
