// An encapsulated immutable map[int]struct{} used as a set.
// Thread-safe.
//
//
// Generated from immutable/set.tpl with Type=int
// options: Comparable:always Numeric:true Ordered:true Stringer:true Mutable:disabled
// by runtemplate v3.4.2
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

// IIntSet is the primary type that represents a set.
type IIntSet struct {
	m map[int]struct{}
}

// NewIIntSet creates and returns a reference to an empty set.
func NewIIntSet(values ...int) *IIntSet {
	set := &IIntSet{
		m: make(map[int]struct{}),
	}
	for _, i := range values {
		set.m[i] = struct{}{}
	}
	return set
}

// ConvertIIntSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertIIntSet(values ...interface{}) (*IIntSet, bool) {
	set := NewIIntSet()

	for _, i := range values {
		switch j := i.(type) {
		case int:
			k := int(j)
			set.m[k] = struct{}{}
		case *int:
			k := int(*j)
			set.m[k] = struct{}{}
		case int8:
			k := int(j)
			set.m[k] = struct{}{}
		case *int8:
			k := int(*j)
			set.m[k] = struct{}{}
		case int16:
			k := int(j)
			set.m[k] = struct{}{}
		case *int16:
			k := int(*j)
			set.m[k] = struct{}{}
		case int32:
			k := int(j)
			set.m[k] = struct{}{}
		case *int32:
			k := int(*j)
			set.m[k] = struct{}{}
		case int64:
			k := int(j)
			set.m[k] = struct{}{}
		case *int64:
			k := int(*j)
			set.m[k] = struct{}{}
		case uint:
			k := int(j)
			set.m[k] = struct{}{}
		case *uint:
			k := int(*j)
			set.m[k] = struct{}{}
		case uint8:
			k := int(j)
			set.m[k] = struct{}{}
		case *uint8:
			k := int(*j)
			set.m[k] = struct{}{}
		case uint16:
			k := int(j)
			set.m[k] = struct{}{}
		case *uint16:
			k := int(*j)
			set.m[k] = struct{}{}
		case uint32:
			k := int(j)
			set.m[k] = struct{}{}
		case *uint32:
			k := int(*j)
			set.m[k] = struct{}{}
		case uint64:
			k := int(j)
			set.m[k] = struct{}{}
		case *uint64:
			k := int(*j)
			set.m[k] = struct{}{}
		case float32:
			k := int(j)
			set.m[k] = struct{}{}
		case *float32:
			k := int(*j)
			set.m[k] = struct{}{}
		case float64:
			k := int(j)
			set.m[k] = struct{}{}
		case *float64:
			k := int(*j)
			set.m[k] = struct{}{}
		}
	}

	return set, len(set.m) == len(values)
}

// BuildIIntSetFromChan constructs a new IIntSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildIIntSetFromChan(source <-chan int) *IIntSet {
	set := NewIIntSet()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set *IIntSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set *IIntSet) IsSet() bool {
	return true
}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set *IIntSet) ToList() *IIntList {
	if set == nil {
		return nil
	}

	return &IIntList{
		m: set.ToSlice(),
	}
}

// ToSet returns the set; this is an identity operation in this case.
func (set *IIntSet) ToSet() *IIntSet {
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set *IIntSet) ToSlice() []int {
	if set == nil {
		return nil
	}

	s := make([]int, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set *IIntSet) ToInterfaceSlice() []interface{} {

	s := make([]interface{}, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns the same set, which is immutable.
func (set *IIntSet) Clone() *IIntSet {
	return set
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *IIntSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *IIntSet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *IIntSet) Size() int {
	if set == nil {
		return 0
	}

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *IIntSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add returns a new set with all original items and all in `more`.
// The original set is not altered.
func (set *IIntSet) Add(more ...int) *IIntSet {
	newSet := NewIIntSet()

	for v := range set.m {
		newSet.doAdd(v)
	}

	for _, v := range more {
		newSet.doAdd(v)
	}

	return newSet
}

func (set *IIntSet) doAdd(i int) {
	set.m[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set *IIntSet) Contains(i int) bool {
	if set == nil {
		return false
	}

	_, found := set.m[i]
	return found
}

// ContainsAll determines whether a given item is already in the set, returning true if so.
func (set *IIntSet) ContainsAll(i ...int) bool {
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
func (set *IIntSet) IsSubset(other *IIntSet) bool {
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
func (set *IIntSet) IsSuperset(other *IIntSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set *IIntSet) Union(other *IIntSet) *IIntSet {
	if set == nil {
		return other
	}

	if other == nil {
		return set
	}

	unionedSet := NewIIntSet()

	for v := range set.m {
		unionedSet.doAdd(v)
	}

	for v := range other.m {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set *IIntSet) Intersect(other *IIntSet) *IIntSet {
	if set == nil || other == nil {
		return nil
	}

	intersection := NewIIntSet()

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
func (set *IIntSet) Difference(other *IIntSet) *IIntSet {
	if set == nil {
		return nil
	}

	if other == nil {
		return set
	}

	differencedSet := NewIIntSet()

	for v := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *IIntSet) SymmetricDifference(other *IIntSet) *IIntSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Remove removes a single item from the set. A new set is returned that has all the elements except the removed one.
func (set *IIntSet) Remove(i int) *IIntSet {
	if set == nil {
		return nil
	}

	clonedSet := NewIIntSet()

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
func (set *IIntSet) Send() <-chan int {
	ch := make(chan int)
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
func (set *IIntSet) Forall(p func(int) bool) bool {
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
func (set *IIntSet) Exists(p func(int) bool) bool {
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

// Foreach iterates over intSet and executes the function f against each element.
func (set *IIntSet) Foreach(f func(int)) {
	if set == nil {
		return
	}

	for v := range set.m {
		f(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first int that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set *IIntSet) Find(p func(int) bool) (int, bool) {

	for v := range set.m {
		if p(v) {
			return v, true
		}
	}

	var empty int
	return empty, false
}

// Filter returns a new IIntSet whose elements return true for the predicate p.
func (set *IIntSet) Filter(p func(int) bool) *IIntSet {
	if set == nil {
		return nil
	}

	result := NewIIntSet()

	for v := range set.m {
		if p(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new intSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set *IIntSet) Partition(p func(int) bool) (*IIntSet, *IIntSet) {
	if set == nil {
		return nil, nil
	}

	matching := NewIIntSet()
	others := NewIIntSet()

	for v := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new IIntSet by transforming every element with a function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *IIntSet) Map(f func(int) int) *IIntSet {
	if set == nil {
		return nil
	}

	result := NewIIntSet()

	for v := range set.m {
		result.m[f(v)] = struct{}{}
	}

	return result
}

// FlatMap returns a new IIntSet by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *IIntSet) FlatMap(f func(int) []int) *IIntSet {
	if set == nil {
		return nil
	}

	result := NewIIntSet()

	for v := range set.m {
		for _, x := range f(v) {
			result.m[x] = struct{}{}
		}
	}

	return result
}

// CountBy gives the number elements of IIntSet that return true for the predicate p.
func (set *IIntSet) CountBy(p func(int) bool) (result int) {

	for v := range set.m {
		if p(v) {
			result++
		}
	}
	return
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (set *IIntSet) Min() int {

	var m int
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
func (set *IIntSet) Max() (result int) {

	var m int
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

// MinBy returns an element of IIntSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *IIntSet) MinBy(less func(int, int) bool) int {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m int
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

// MaxBy returns an element of IIntSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set *IIntSet) MaxBy(less func(int, int) bool) int {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m int
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
// These methods are included when int is numeric.

// Sum returns the sum of all the elements in the set.
func (set *IIntSet) Sum() int {

	sum := int(0)
	for v := range set.m {
		sum = sum + v
	}
	return sum
}

//-------------------------------------------------------------------------------------------------

// Equals determines whether two sets are equal to each other, returning true if so.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set *IIntSet) Equals(other *IIntSet) bool {
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
func (set *IIntSet) StringList() []string {

	strings := make([]string, len(set.m))
	i := 0
	for v := range set.m {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set *IIntSet) String() string {
	return set.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set *IIntSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set *IIntSet) MkString3(before, between, after string) string {
	if set == nil {
		return ""
	}
	return set.mkString3Bytes(before, between, after).String()
}

func (set *IIntSet) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
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
func (set *IIntSet) UnmarshalJSON(b []byte) error {

	values := make([]int, 0)
	err := json.Unmarshal(b, &values)
	if err != nil {
		return err
	}

	s2 := NewIIntSet(values...)
	*set = *s2
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set *IIntSet) MarshalJSON() ([]byte, error) {

	buf, err := json.Marshal(set.ToSlice())
	return buf, err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set *IIntSet) StringMap() map[string]bool {
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
// You must register int with the 'gob' package before this method is used.
func (set *IIntSet) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&set.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register int with the 'gob' package before this method is used.
func (set IIntSet) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(set.m)
	return buf.Bytes(), err
}
