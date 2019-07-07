// An encapsulated immutable map[int]int.
// Thread-safe.
//
//
// Generated from immutable/map.tpl with Key=int Type=int
// options: Comparable:true Stringer:true KeyList:<no value> ValueList:<no value> Mutable:disabled
// by runtemplate v3.4.2
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// IIntIntMap is the primary type that represents a thread-safe map
type IIntIntMap struct {
	m map[int]int
}

// IIntIntTuple represents a key/value pair.
type IIntIntTuple struct {
	Key int
	Val int
}

// IIntIntTuples can be used as a builder for unmodifiable maps.
type IIntIntTuples []IIntIntTuple

// Append1 adds one item.
func (ts IIntIntTuples) Append1(k int, v int) IIntIntTuples {
	return append(ts, IIntIntTuple{k, v})
}

// Append2 adds two items.
func (ts IIntIntTuples) Append2(k1 int, v1 int, k2 int, v2 int) IIntIntTuples {
	return append(ts, IIntIntTuple{k1, v1}, IIntIntTuple{k2, v2})
}

// Append3 adds three items.
func (ts IIntIntTuples) Append3(k1 int, v1 int, k2 int, v2 int, k3 int, v3 int) IIntIntTuples {
	return append(ts, IIntIntTuple{k1, v1}, IIntIntTuple{k2, v2}, IIntIntTuple{k3, v3})
}

// IIntIntZip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewIIntIntMap
// constructor function.
func IIntIntZip(keys ...int) IIntIntTuples {
	ts := make(IIntIntTuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with IIntIntZip.
func (ts IIntIntTuples) Values(values ...int) IIntIntTuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

//-------------------------------------------------------------------------------------------------

func newIIntIntMap() *IIntIntMap {
	return &IIntIntMap{
		m: make(map[int]int),
	}
}

// NewIIntIntMap1 creates and returns a reference to a map containing one item.
func NewIIntIntMap1(k int, v int) *IIntIntMap {
	mm := newIIntIntMap()
	mm.m[k] = v
	return mm
}

// NewIIntIntMap creates and returns a reference to a map, optionally containing some items.
func NewIIntIntMap(kv ...IIntIntTuple) *IIntIntMap {
	mm := newIIntIntMap()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *IIntIntMap) Keys() []int {
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
func (mm *IIntIntMap) Values() []int {
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
func (mm *IIntIntMap) slice() []IIntIntTuple {
	if mm == nil {
		return nil
	}

	s := make([]IIntIntTuple, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, IIntIntTuple{k, v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *IIntIntMap) ToSlice() []IIntIntTuple {
	return mm.slice()
}

// Get returns one of the items in the map, if present.
func (mm *IIntIntMap) Get(k int) (int, bool) {
	v, found := mm.m[k]
	return v, found
}

// Put adds an item to a clone of the map, replacing any prior value and returning the cloned map.
func (mm *IIntIntMap) Put(k int, v int) *IIntIntMap {
	if mm == nil {
		return NewIIntIntMap1(k, v)
	}

	result := NewIIntIntMap()

	for k, v := range mm.m {
		result.m[k] = v
	}
	result.m[k] = v

	return result
}

// ContainsKey determines if a given item is already in the map.
func (mm *IIntIntMap) ContainsKey(k int) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *IIntIntMap) ContainsAllKeys(kk ...int) bool {
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
func (mm *IIntIntMap) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *IIntIntMap) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *IIntIntMap) NonEmpty() bool {
	return mm.Size() > 0
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *IIntIntMap) Foreach(f func(int, int)) {
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
func (mm *IIntIntMap) Forall(f func(int, int) bool) bool {
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
func (mm *IIntIntMap) Exists(p func(int, int) bool) bool {
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
func (mm *IIntIntMap) Find(p func(int, int) bool) (IIntIntTuple, bool) {
	if mm == nil {
		return IIntIntTuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return IIntIntTuple{k, v}, true
		}
	}

	return IIntIntTuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *IIntIntMap) Filter(p func(int, int) bool) *IIntIntMap {
	if mm == nil {
		return nil
	}

	result := NewIIntIntMap()

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
func (mm *IIntIntMap) Partition(p func(int, int) bool) (matching *IIntIntMap, others *IIntIntMap) {
	if mm == nil {
		return nil, nil
	}

	matching = NewIIntIntMap()
	others = NewIIntIntMap()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new IIntMap by transforming every element with the function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *IIntIntMap) Map(f func(int, int) (int, int)) *IIntIntMap {
	if mm == nil {
		return nil
	}

	result := NewIIntIntMap()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new IIntMap by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *IIntIntMap) FlatMap(f func(int, int) []IIntIntTuple) *IIntIntMap {
	if mm == nil {
		return nil
	}

	result := NewIIntIntMap()

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
func (mm *IIntIntMap) Equals(other *IIntIntMap) bool {
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
func (mm *IIntIntMap) Clone() *IIntIntMap {
	return mm
}

//-------------------------------------------------------------------------------------------------

func (mm *IIntIntMap) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *IIntIntMap) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *IIntIntMap) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *IIntIntMap) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *IIntIntMap) mkString3Bytes(before, between, after string) *bytes.Buffer {
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
func (mm *IIntIntMap) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register int with the 'gob' package before this method is used.
func (mm *IIntIntMap) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}
