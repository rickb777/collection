// An encapsulated immutable map[int]int.
// Thread-safe.
//
//
// Generated from immutable/map.tpl with Key=int Type=int
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:disabled
// by runtemplate v3.5.4
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// IntIntMap is the primary type that represents a thread-safe map
type IntIntMap struct {
	m map[int]int
}

// IntIntTuple represents a key/value pair.
type IntIntTuple struct {
	Key int
	Val int
}

// IntIntTuples can be used as a builder for unmodifiable maps.
type IntIntTuples []IntIntTuple

// Append1 adds one item.
func (ts IntIntTuples) Append1(k int, v int) IntIntTuples {
	return append(ts, IntIntTuple{k, v})
}

// Append2 adds two items.
func (ts IntIntTuples) Append2(k1 int, v1 int, k2 int, v2 int) IntIntTuples {
	return append(ts, IntIntTuple{k1, v1}, IntIntTuple{k2, v2})
}

// Append3 adds three items.
func (ts IntIntTuples) Append3(k1 int, v1 int, k2 int, v2 int, k3 int, v3 int) IntIntTuples {
	return append(ts, IntIntTuple{k1, v1}, IntIntTuple{k2, v2}, IntIntTuple{k3, v3})
}

// IntIntZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewIntIntMap
// constructor function.
func IntIntZip(keys ...int) IntIntTuples {
	ts := make(IntIntTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with IntIntZip.
func (ts IntIntTuples) Values(values ...int) IntIntTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newIntIntMap() *IntIntMap {
	return &IntIntMap{
		m: make(map[int]int),
	}
}

// NewIntIntMap1 creates and returns a reference to a map containing one item.
func NewIntIntMap1(k int, v int) *IntIntMap {
	mm := newIntIntMap()
	mm.m[k] = v
	return mm
}

// NewIntIntMap creates and returns a reference to a map, optionally containing some items.
func NewIntIntMap(kv ...IntIntTuple) *IntIntMap {
	mm := newIntIntMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *IntIntMap) Keys() []int {
	if mm == nil {
		return nil
	}

	s := make([]int, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *IntIntMap) Values() []int {
	if mm == nil {
		return nil
	}

	s := make([]int, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *IntIntMap) slice() []IntIntTuple {
	if mm == nil {
		return nil
	}

	s := make([]IntIntTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, IntIntTuple{k, v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *IntIntMap) ToSlice() []IntIntTuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *IntIntMap) Get(k int) (int, bool) {
	v, found := mm.m[k]
	return v, found
}

// Put adds an item to a clone of the map, replacing any prior value and returning the cloned map.
func (mm *IntIntMap) Put(k int, v int) *IntIntMap {
	if mm == nil {
		return NewIntIntMap1(k, v)
	}

	result := NewIntIntMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	result.m[k] = v

	return result
}

// ContainsKey determines if a given item is already in the map.
func (mm *IntIntMap) ContainsKey(k int) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *IntIntMap) ContainsAllKeys(kk ...int) bool {
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
func (mm *IntIntMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *IntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *IntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *IntIntMap) Foreach(f func(int, int)) {
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
func (mm *IntIntMap) Forall(f func(int, int) bool) bool {
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
func (mm *IntIntMap) Exists(p func(int, int) bool) bool {
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

// Find returns the first int that returns true for the predicate p.
// False is returned if none match.
func (mm *IntIntMap) Find(p func(int, int) bool) (IntIntTuple, bool) {
	if mm == nil {
		return IntIntTuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return IntIntTuple{k, v}, true
		}
	}

	return IntIntTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *IntIntMap) Filter(p func(int, int) bool) *IntIntMap {
	if mm == nil {
		return nil
	}

	result := NewIntIntMap()

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
func (mm *IntIntMap) Partition(p func(int, int) bool) (matching *IntIntMap, others *IntIntMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewIntIntMap()
	others = NewIntIntMap()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new IntMap by transforming every element with the function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *IntIntMap) Map(f func(int, int) (int, int)) *IntIntMap {
	if mm == nil {
		return nil
	}

	result := NewIntIntMap()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new IntMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *IntIntMap) FlatMap(f func(int, int) []IntIntTuple) *IntIntMap {
	if mm == nil {
		return nil
	}

	result := NewIntIntMap()

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
func (mm *IntIntMap) Equals(other *IntIntMap) bool {
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
func (mm *IntIntMap) Clone() *IntIntMap {
	return mm
}

//-------------------------------------------------------------------------------------------------

func (mm *IntIntMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *IntIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *IntIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *IntIntMap) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *IntIntMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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

// GobDecode implements 'gob' decoding for this map type.
// You must register int with the 'gob' package before this method is used.
func (mm *IntIntMap) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register int with the 'gob' package before this method is used.
func (mm *IntIntMap) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}
