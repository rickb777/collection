// An encapsulated immutable map[uint64]uint64.
// Thread-safe.
//
//
// Generated from immutable/map.tpl with Key=uint64 Type=uint64
// options: Comparable:true Stringer:true KeyList:collection.Uint64List ValueList:collection.Uint64List Mutable:disabled
// by runtemplate v3.6.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/rickb777/collection"
)

// Uint64Uint64Map is the primary type that represents a thread-safe map
type Uint64Uint64Map struct {
	m map[uint64]uint64
}

// Uint64Uint64Tuple represents a key/value pair.
type Uint64Uint64Tuple struct {
	Key uint64
	Val uint64
}

// Uint64Uint64Tuples can be used as a builder for unmodifiable maps.
type Uint64Uint64Tuples []Uint64Uint64Tuple

// Append1 adds one item.
func (ts Uint64Uint64Tuples) Append1(k uint64, v uint64) Uint64Uint64Tuples {
	return append(ts, Uint64Uint64Tuple{k, v})
}

// Append2 adds two items.
func (ts Uint64Uint64Tuples) Append2(k1 uint64, v1 uint64, k2 uint64, v2 uint64) Uint64Uint64Tuples {
	return append(ts, Uint64Uint64Tuple{k1, v1}, Uint64Uint64Tuple{k2, v2})
}

// Append3 adds three items.
func (ts Uint64Uint64Tuples) Append3(k1 uint64, v1 uint64, k2 uint64, v2 uint64, k3 uint64, v3 uint64) Uint64Uint64Tuples {
	return append(ts, Uint64Uint64Tuple{k1, v1}, Uint64Uint64Tuple{k2, v2}, Uint64Uint64Tuple{k3, v3})
}

// Uint64Uint64Zip is used with the Values method to zip (i.e. interleave) a slice of
// keys with a slice of values. These can then be passed in to the NewUint64Uint64Map
// constructor function.
func Uint64Uint64Zip(keys ...uint64) Uint64Uint64Tuples {
	ts := make(Uint64Uint64Tuples, len(keys))
	for i, k := range keys {
		ts[i].Key = k
	}
	return ts
}

// Values sets the values in a tuple slice. Use this with Uint64Uint64Zip.
func (ts Uint64Uint64Tuples) Values(values ...uint64) Uint64Uint64Tuples {
	if len(ts) != len(values) {
		panic(fmt.Errorf("Mismatched %d keys and %d values", len(ts), len(values)))
	}
	for i, v := range values {
		ts[i].Val = v
	}
	return ts
}

// ToMap converts the tuples to a map.
func (ts Uint64Uint64Tuples) ToMap() *Uint64Uint64Map {
	return NewUint64Uint64Map(ts...)
}

//-------------------------------------------------------------------------------------------------

func newUint64Uint64Map() *Uint64Uint64Map {
	return &Uint64Uint64Map{
		m: make(map[uint64]uint64),
	}
}

// NewUint64Uint64Map1 creates and returns a reference to a map containing one item.
func NewUint64Uint64Map1(k uint64, v uint64) *Uint64Uint64Map {
	mm := newUint64Uint64Map()
	mm.m[k] = v
	return mm
}

// NewUint64Uint64Map creates and returns a reference to a map, optionally containing some items.
func NewUint64Uint64Map(kv ...Uint64Uint64Tuple) *Uint64Uint64Map {
	mm := newUint64Uint64Map()
	for _, t := range kv {
		mm.m[t.Key] = t.Val
	}
	return mm
}

// Keys returns the keys of the current map as a slice.
func (mm *Uint64Uint64Map) Keys() collection.Uint64List {
	if mm == nil {
		return nil
	}

	s := make(collection.Uint64List, 0, len(mm.m))
	for k := range mm.m {
		s = append(s, k)
	}

	return s
}

// Values returns the values of the current map as a slice.
func (mm *Uint64Uint64Map) Values() collection.Uint64List {
	if mm == nil {
		return nil
	}

	s := make(collection.Uint64List, 0, len(mm.m))
	for _, v := range mm.m {
		s = append(s, v)
	}

	return s
}

// slice returns the internal elements of the map. This is a seam for testing etc.
func (mm *Uint64Uint64Map) slice() Uint64Uint64Tuples {
	if mm == nil {
		return nil
	}

	s := make(Uint64Uint64Tuples, 0, len(mm.m))
	for k, v := range mm.m {
		s = append(s, Uint64Uint64Tuple{(k), v})
	}

	return s
}

// ToSlice returns the key/value pairs as a slice
func (mm *Uint64Uint64Map) ToSlice() Uint64Uint64Tuples {
	return mm.slice()
}

// OrderedSlice returns the key/value pairs as a slice in the order specified by keys.
func (mm *Uint64Uint64Map) OrderedSlice(keys collection.Uint64List) Uint64Uint64Tuples {
	if mm == nil {
		return nil
	}

	s := make(Uint64Uint64Tuples, 0, len(mm.m))
	for _, k := range keys {
		v, found := mm.m[k]
		if found {
			s = append(s, Uint64Uint64Tuple{k, v})
		}
	}
	return s
}

// Get returns one of the items in the map, if present.
func (mm *Uint64Uint64Map) Get(k uint64) (uint64, bool) {
	v, found := mm.m[k]
	return v, found
}

// Put adds an item to a clone of the map, replacing any prior value and returning the cloned map.
func (mm *Uint64Uint64Map) Put(k uint64, v uint64) *Uint64Uint64Map {
	if mm == nil {
		return NewUint64Uint64Map1(k, v)
	}

	result := NewUint64Uint64Map()

	for k, v := range mm.m {
		result.m[k] = v
	}
	result.m[k] = v

	return result
}

// ContainsKey determines if a given item is already in the map.
func (mm *Uint64Uint64Map) ContainsKey(k uint64) bool {
	if mm == nil {
		return false
	}

	_, found := mm.m[k]
	return found
}

// ContainsAllKeys determines if the given items are all in the map.
func (mm *Uint64Uint64Map) ContainsAllKeys(kk ...uint64) bool {
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
func (mm *Uint64Uint64Map) Size() int {
	if mm == nil {
		return 0
	}

	return len(mm.m)
}

// IsEmpty returns true if the map is empty.
func (mm *Uint64Uint64Map) IsEmpty() bool {
	return mm.Size() == 0
}

// NonEmpty returns true if the map is not empty.
func (mm *Uint64Uint64Map) NonEmpty() bool {
	return mm.Size() > 0
}

// Foreach applies the function f to every element in the map.
// The function can safely alter the values via side-effects.
func (mm *Uint64Uint64Map) Foreach(f func(uint64, uint64)) {
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
func (mm *Uint64Uint64Map) Forall(f func(uint64, uint64) bool) bool {
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
func (mm *Uint64Uint64Map) Exists(p func(uint64, uint64) bool) bool {
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

// Find returns the first uint64 that returns true for the predicate p.
// False is returned if none match.
func (mm *Uint64Uint64Map) Find(p func(uint64, uint64) bool) (Uint64Uint64Tuple, bool) {
	if mm == nil {
		return Uint64Uint64Tuple{}, false
	}

	for k, v := range mm.m {
		if p(k, v) {
			return Uint64Uint64Tuple{k, v}, true
		}
	}

	return Uint64Uint64Tuple{}, false
}

// Filter applies the predicate p to every element in the map and returns a copied map containing
// only the elements for which the predicate returned true.
func (mm *Uint64Uint64Map) Filter(p func(uint64, uint64) bool) *Uint64Uint64Map {
	if mm == nil {
		return nil
	}

	result := NewUint64Uint64Map()

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
func (mm *Uint64Uint64Map) Partition(p func(uint64, uint64) bool) (matching *Uint64Uint64Map, others *Uint64Uint64Map) {
	if mm == nil {
		return nil, nil
	}

	matching = NewUint64Uint64Map()
	others = NewUint64Uint64Map()

	for k, v := range mm.m {
		if p(k, v) {
			matching.m[k] = v
		} else {
			others.m[k] = v
		}
	}
	return
}

// Map returns a new Uint64Map by transforming every element with the function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *Uint64Uint64Map) Map(f func(uint64, uint64) (uint64, uint64)) *Uint64Uint64Map {
	if mm == nil {
		return nil
	}

	result := NewUint64Uint64Map()

	for k1, v1 := range mm.m {
		k2, v2 := f(k1, v1)
		result.m[k2] = v2
	}

	return result
}

// FlatMap returns a new Uint64Map by transforming every element with the function f that
// returns zero or more items in a slice. The resulting map may have a different size to the original map.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (mm *Uint64Uint64Map) FlatMap(f func(uint64, uint64) []Uint64Uint64Tuple) *Uint64Uint64Map {
	if mm == nil {
		return nil
	}

	result := NewUint64Uint64Map()

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
func (mm *Uint64Uint64Map) Equals(other *Uint64Uint64Map) bool {
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
func (mm *Uint64Uint64Map) Clone() *Uint64Uint64Map {
	return mm
}

//-------------------------------------------------------------------------------------------------

func (mm *Uint64Uint64Map) String() string {
	return mm.MkString3("[", ", ", "]")
}

// implements encoding.Marshaler interface {
//func (mm *Uint64Uint64Map) MarshalJSON() ([]byte, error) {
//	return mm.mkString3Bytes("{\"", "\", \"", "\"}").Bytes(), nil
//}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (mm *Uint64Uint64Map) MkString(sep string) string {
	return mm.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (mm *Uint64Uint64Map) MkString3(before, between, after string) string {
	if mm == nil {
		return ""
	}
	return mm.mkString3Bytes(before, between, after).String()
}

func (mm *Uint64Uint64Map) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	keys := make(collection.Uint64List, 0, len(mm.m))
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
// You must register uint64 with the 'gob' package before this method is used.
func (mm *Uint64Uint64Map) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&mm.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register uint64 with the 'gob' package before this method is used.
func (mm *Uint64Uint64Map) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(mm.m)
	return buf.Bytes(), err
}

//-------------------------------------------------------------------------------------------------

func (ts Uint64Uint64Tuples) String() string {
	return ts.MkString3("[", ", ", "]")
}

// MkString concatenates the map key/values as a string using a supplied separator. No enclosing marks are added.
func (ts Uint64Uint64Tuples) MkString(sep string) string {
	return ts.MkString3("", sep, "")
}

// MkString3 concatenates the map key/values as a string, using the prefix, separator and suffix supplied.
func (ts Uint64Uint64Tuples) MkString3(before, between, after string) string {
	if ts == nil {
		return ""
	}
	return ts.mkString3Bytes(before, between, after).String()
}

func (ts Uint64Uint64Tuples) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
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
func (t Uint64Uint64Tuple) UnmarshalJSON(b []byte) error {
	buf := bytes.NewBuffer(b)
	return json.NewDecoder(buf).Decode(&t)
}

// MarshalJSON implements encoding.Marshaler interface.
func (t Uint64Uint64Tuple) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"key":"%v", "val":"%v"}`, t.Key, t.Val)), nil
}
