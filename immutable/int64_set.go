// An encapsulated immutable map[int64]struct{} used as a set.
// Thread-safe.
//
//
// Generated from immutable/set.tpl with Type=int64
// options: Comparable:always Numeric:<no value> Integer:true Ordered:true Stringer:true Mutable:disabled
// by runtemplate v3.10.1
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package immutable

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Int64Set is the primary type that represents a set.
type Int64Set struct {
	m map[int64]struct{}
}

// NewInt64Set creates and returns a reference to an empty set.
func NewInt64Set(values ...int64) *Int64Set {
	set := &Int64Set{
		m: make(map[int64]struct{}),
	}
	for _, i := range values {
		set.m[i] = struct{}{}
	}
	return set
}

// ConvertInt64Set constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertInt64Set(values ...interface{}) (*Int64Set, bool) {
	set := NewInt64Set()

	for _, i := range values {
		switch s := i.(type) {
		case string:
			k, e := strconv.ParseInt(s, 10, 64)
			if e == nil {
				i = k
			}
		case *string:
			k, e := strconv.ParseInt(*s, 10, 64)
			if e == nil {
				i = k
			}
		}
		switch j := i.(type) {
		case int:
			k := int64(j)
			set.m[k] = struct{}{}
		case *int:
			k := int64(*j)
			set.m[k] = struct{}{}
		case int8:
			k := int64(j)
			set.m[k] = struct{}{}
		case *int8:
			k := int64(*j)
			set.m[k] = struct{}{}
		case int16:
			k := int64(j)
			set.m[k] = struct{}{}
		case *int16:
			k := int64(*j)
			set.m[k] = struct{}{}
		case int32:
			k := int64(j)
			set.m[k] = struct{}{}
		case *int32:
			k := int64(*j)
			set.m[k] = struct{}{}
		case int64:
			k := int64(j)
			set.m[k] = struct{}{}
		case *int64:
			k := int64(*j)
			set.m[k] = struct{}{}
		case uint:
			k := int64(j)
			set.m[k] = struct{}{}
		case *uint:
			k := int64(*j)
			set.m[k] = struct{}{}
		case uint8:
			k := int64(j)
			set.m[k] = struct{}{}
		case *uint8:
			k := int64(*j)
			set.m[k] = struct{}{}
		case uint16:
			k := int64(j)
			set.m[k] = struct{}{}
		case *uint16:
			k := int64(*j)
			set.m[k] = struct{}{}
		case uint32:
			k := int64(j)
			set.m[k] = struct{}{}
		case *uint32:
			k := int64(*j)
			set.m[k] = struct{}{}
		case uint64:
			k := int64(j)
			set.m[k] = struct{}{}
		case *uint64:
			k := int64(*j)
			set.m[k] = struct{}{}
		case float32:
			k := int64(j)
			set.m[k] = struct{}{}
		case *float32:
			k := int64(*j)
			set.m[k] = struct{}{}
		case float64:
			k := int64(j)
			set.m[k] = struct{}{}
		case *float64:
			k := int64(*j)
			set.m[k] = struct{}{}
		}
	}

	return set, len(set.m) == len(values)
}

// BuildInt64SetFromChan constructs a new Int64Set from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildInt64SetFromChan(source <-chan int64) *Int64Set {
	set := NewInt64Set()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set *Int64Set) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set *Int64Set) IsSet() bool {
	return true
}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set *Int64Set) ToList() *Int64List {
	if set == nil {
		return nil
	}

	return &Int64List{
		m: set.ToSlice(),
	}
}

// ToSet returns the set; this is an identity operation in this case.
func (set *Int64Set) ToSet() *Int64Set {
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set *Int64Set) ToSlice() []int64 {
	if set == nil {
		return nil
	}

	s := make([]int64, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set *Int64Set) ToInterfaceSlice() []interface{} {

	s := make([]interface{}, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns the same set, which is immutable.
func (set *Int64Set) Clone() *Int64Set {
	return set
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *Int64Set) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *Int64Set) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *Int64Set) Size() int {
	if set == nil {
		return 0
	}

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *Int64Set) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add returns a new set with all original items and all in `more`.
// The original set is not altered.
func (set *Int64Set) Add(more ...int64) *Int64Set {
	newSet := NewInt64Set()

	for v := range set.m {
		newSet.doAdd(v)
	}

	for _, v := range more {
		newSet.doAdd(v)
	}

	return newSet
}

func (set *Int64Set) doAdd(i int64) {
	set.m[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set *Int64Set) Contains(i int64) bool {
	if set == nil {
		return false
	}

	_, found := set.m[i]
	return found
}

// ContainsAll determines whether a given item is already in the set, returning true if so.
func (set *Int64Set) ContainsAll(i ...int64) bool {
	if set == nil {
		return false
	}

	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines whether every item in the other set is in this set, returning true if so.
func (set *Int64Set) IsSubset(other *Int64Set) bool {
	if set.IsEmpty() {
		return !other.IsEmpty()
	}

	if other.IsEmpty() {
		return false
	}

	for v := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set *Int64Set) IsSuperset(other *Int64Set) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set *Int64Set) Union(other *Int64Set) *Int64Set {
	if set == nil {
		return other
	}

	if other == nil {
		return set
	}

	unionedSet := NewInt64Set()

	for v := range set.m {
		unionedSet.doAdd(v)
	}

	for v := range other.m {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set *Int64Set) Intersect(other *Int64Set) *Int64Set {
	if set == nil || other == nil {
		return nil
	}

	intersection := NewInt64Set()

	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set.m {
			if other.Contains(v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v := range other.m {
			if set.Contains(v) {
				intersection.doAdd(v)
			}
		}
	}

	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set *Int64Set) Difference(other *Int64Set) *Int64Set {
	if set == nil {
		return nil
	}

	if other == nil {
		return set
	}

	differencedSet := NewInt64Set()

	for v := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *Int64Set) SymmetricDifference(other *Int64Set) *Int64Set {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Remove removes a single item from the set. A new set is returned that has all the elements except the removed one.
func (set *Int64Set) Remove(i int64) *Int64Set {
	if set == nil {
		return nil
	}

	clonedSet := NewInt64Set()

	for v := range set.m {
		if i != v {
			clonedSet.doAdd(v)
		}
	}

	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set *Int64Set) Send() <-chan int64 {
	ch := make(chan int64)
	go func() {
		if set != nil {
			for v := range set.m {
				ch <- v
			}
		}
		close(ch)
	}()

	return ch
}

//-------------------------------------------------------------------------------------------------

// Forall applies a predicate function p to every element in the set. If the function returns false,
// the iteration terminates early. The returned value is true if all elements were visited,
// or false if an early return occurred.
//
// Note that this method can also be used simply as a way to visit every element using a function
// with some side-effects; such a function must always return true.
func (set *Int64Set) Forall(p func(int64) bool) bool {
	if set == nil {
		return true
	}

	for v := range set.m {
		if !p(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate p to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set *Int64Set) Exists(p func(int64) bool) bool {
	if set == nil {
		return false
	}

	for v := range set.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over int64Set and executes the function f against each element.
func (set *Int64Set) Foreach(f func(int64)) {
	if set == nil {
		return
	}

	for v := range set.m {
		f(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first int64 that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set *Int64Set) Find(p func(int64) bool) (int64, bool) {

	for v := range set.m {
		if p(v) {
			return v, true
		}
	}

	var empty int64
	return empty, false
}

// Filter returns a new Int64Set whose elements return true for the predicate p.
func (set *Int64Set) Filter(p func(int64) bool) *Int64Set {
	if set == nil {
		return nil
	}

	result := NewInt64Set()

	for v := range set.m {
		if p(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new int64Sets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set *Int64Set) Partition(p func(int64) bool) (*Int64Set, *Int64Set) {
	if set == nil {
		return nil, nil
	}

	matching := NewInt64Set()
	others := NewInt64Set()

	for v := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new Int64Set by transforming every element with a function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *Int64Set) Map(f func(int64) int64) *Int64Set {
	if set == nil {
		return nil
	}

	result := NewInt64Set()

	for v := range set.m {
		result.m[f(v)] = struct{}{}
	}

	return result
}

// MapToString returns a new []string by transforming every element with function f.
// The resulting slice is the same size as the set.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *Int64Set) MapToString(f func(int64) string) []string {
	if set == nil {
		return nil
	}

	result := make([]string, 0, len(set.m))

	for v := range set.m {
		result = append(result, f(v))
	}

	return result
}

// FlatMap returns a new Int64Set by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *Int64Set) FlatMap(f func(int64) []int64) *Int64Set {
	if set == nil {
		return nil
	}

	result := NewInt64Set()

	for v := range set.m {
		for _, x := range f(v) {
			result.m[x] = struct{}{}
		}
	}

	return result
}

// FlatMapToString returns a new []string by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the set.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *Int64Set) FlatMapToString(f func(int64) []string) []string {
	if set == nil {
		return nil
	}

	result := make([]string, 0, len(set.m))

	for v := range set.m {
		result = append(result, f(v)...)
	}

	return result
}

// CountBy gives the number elements of Int64Set that return true for the predicate p.
func (set *Int64Set) CountBy(p func(int64) bool) (result int) {

	for v := range set.m {
		if p(v) {
			result++
		}
	}
	return
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int64 is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (set *Int64Set) Min() int64 {

	var m int64
	first := true
	for v := range set.m {
		if first {
			m = v
			first = false
		} else if v < m {
			m = v
		}
	}
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (set *Int64Set) Max() (result int64) {

	var m int64
	first := true
	for v := range set.m {
		if first {
			m = v
			first = false
		} else if v > m {
			m = v
		}
	}
	return m
}

// Fold aggregates all the values in the set using a supplied function, starting from some initial value.
func (set *Int64Set) Fold(initial int64, fn func(int64, int64) int64) int64 {

	m := initial
	for v := range set.m {
		m = fn(m, v)
	}

	return m
}

// MinBy returns an element of Int64Set containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *Int64Set) MinBy(less func(int64, int64) bool) int64 {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m int64
	first := true
	for v := range set.m {
		if first {
			m = v
			first = false
		} else if less(v, m) {
			m = v
		}
	}
	return m
}

// MaxBy returns an element of Int64Set containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set *Int64Set) MaxBy(less func(int64, int64) bool) int64 {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m int64
	first := true
	for v := range set.m {
		if first {
			m = v
			first = false
		} else if less(m, v) {
			m = v
		}
	}
	return m
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int64 is numeric.

// Sum returns the sum of all the elements in the set.
func (set *Int64Set) Sum() int64 {

	sum := int64(0)
	for v := range set.m {
		sum = sum + v
	}
	return sum
}

//-------------------------------------------------------------------------------------------------

// Equals determines whether two sets are equal to each other, returning true if so.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set *Int64Set) Equals(other *Int64Set) bool {
	if set == nil {
		return other == nil || other.IsEmpty()
	}

	if other == nil {
		return set.IsEmpty()
	}

	if set.Size() != other.Size() {
		return false
	}

	for v := range set.m {
		if !other.Contains(v) {
			return false
		}
	}

	return true
}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (set *Int64Set) StringList() []string {

	strings := make([]string, len(set.m))
	i := 0
	for v := range set.m {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set *Int64Set) String() string {
	return set.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set *Int64Set) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set *Int64Set) MkString3(before, between, after string) string {
	if set == nil {
		return ""
	}
	return set.mkString3Bytes(before, between, after).String()
}

func (set *Int64Set) mkString3Bytes(before, between, after string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""

	for v := range set.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this set type.
func (set *Int64Set) UnmarshalJSON(b []byte) error {

	values := make([]int64, 0)
	err := json.Unmarshal(b, &values)
	if err != nil {
		return err
	}

	s2 := NewInt64Set(values...)
	*set = *s2
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set *Int64Set) MarshalJSON() ([]byte, error) {

	buf, err := json.Marshal(set.ToSlice())
	return buf, err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set *Int64Set) StringMap() map[string]bool {
	if set == nil {
		return nil
	}

	strings := make(map[string]bool)
	for v := range set.m {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this set type.
// You must register int64 with the 'gob' package before this method is used.
func (set *Int64Set) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&set.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register int64 with the 'gob' package before this method is used.
func (set Int64Set) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(set.m)
	return buf.Bytes(), err
}
