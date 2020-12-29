// An encapsulated immutable map[string]struct{} used as a set.
// Thread-safe.
//
//
// Generated from immutable/set.tpl with Type=string
// options: Comparable:always Numeric:<no value> Ordered:<no value> Stringer:true Mutable:disabled
// by runtemplate v3.7.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"strings"
)

// StringSet is the primary type that represents a set.
type StringSet struct {
	m map[string]struct{}
}

// NewStringSet creates and returns a reference to an empty set.
func NewStringSet(values ...string) *StringSet {
	set := &StringSet{
		m: make(map[string]struct{}),
	}
	for _, i := range values {
		set.m[i] = struct{}{}
	}
	return set
}

// ConvertStringSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertStringSet(values ...interface{}) (*StringSet, bool) {
	set := NewStringSet()

	for _, i := range values {
		switch j := i.(type) {
		case string:
			set.m[j] = struct{}{}
		case *string:
			set.m[*j] = struct{}{}
		}
	}

	return set, len(set.m) == len(values)
}

// BuildStringSetFromChan constructs a new StringSet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildStringSetFromChan(source <-chan string) *StringSet {
	set := NewStringSet()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set *StringSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set *StringSet) IsSet() bool {
	return true
}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set *StringSet) ToList() *StringList {
	if set == nil {
		return nil
	}

	return &StringList{
		m: set.ToSlice(),
	}
}

// ToSet returns the set; this is an identity operation in this case.
func (set *StringSet) ToSet() *StringSet {
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set *StringSet) ToSlice() []string {
	if set == nil {
		return nil
	}

	s := make([]string, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set *StringSet) ToInterfaceSlice() []interface{} {

	s := make([]interface{}, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns the same set, which is immutable.
func (set *StringSet) Clone() *StringSet {
	return set
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *StringSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *StringSet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *StringSet) Size() int {
	if set == nil {
		return 0
	}

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *StringSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add returns a new set with all original items and all in `more`.
// The original set is not altered.
func (set *StringSet) Add(more ...string) *StringSet {
	newSet := NewStringSet()

	for v := range set.m {
		newSet.doAdd(v)
	}

	for _, v := range more {
		newSet.doAdd(v)
	}

	return newSet
}

func (set *StringSet) doAdd(i string) {
	set.m[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set *StringSet) Contains(i string) bool {
	if set == nil {
		return false
	}

	_, found := set.m[i]
	return found
}

// ContainsAll determines whether a given item is already in the set, returning true if so.
func (set *StringSet) ContainsAll(i ...string) bool {
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
func (set *StringSet) IsSubset(other *StringSet) bool {
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
func (set *StringSet) IsSuperset(other *StringSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set *StringSet) Union(other *StringSet) *StringSet {
	if set == nil {
		return other
	}

	if other == nil {
		return set
	}

	unionedSet := NewStringSet()

	for v := range set.m {
		unionedSet.doAdd(v)
	}

	for v := range other.m {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set *StringSet) Intersect(other *StringSet) *StringSet {
	if set == nil || other == nil {
		return nil
	}

	intersection := NewStringSet()

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
func (set *StringSet) Difference(other *StringSet) *StringSet {
	if set == nil {
		return nil
	}

	if other == nil {
		return set
	}

	differencedSet := NewStringSet()

	for v := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *StringSet) SymmetricDifference(other *StringSet) *StringSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Remove removes a single item from the set. A new set is returned that has all the elements except the removed one.
func (set *StringSet) Remove(i string) *StringSet {
	if set == nil {
		return nil
	}

	clonedSet := NewStringSet()

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
func (set *StringSet) Send() <-chan string {
	ch := make(chan string)
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
func (set *StringSet) Forall(p func(string) bool) bool {
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
func (set *StringSet) Exists(p func(string) bool) bool {
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

// Foreach iterates over stringSet and executes the function f against each element.
func (set *StringSet) Foreach(f func(string)) {
	if set == nil {
		return
	}

	for v := range set.m {
		f(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first string that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set *StringSet) Find(p func(string) bool) (string, bool) {

	for v := range set.m {
		if p(v) {
			return v, true
		}
	}

	var empty string
	return empty, false
}

// Filter returns a new StringSet whose elements return true for the predicate p.
func (set *StringSet) Filter(p func(string) bool) *StringSet {
	if set == nil {
		return nil
	}

	result := NewStringSet()

	for v := range set.m {
		if p(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new stringSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set *StringSet) Partition(p func(string) bool) (*StringSet, *StringSet) {
	if set == nil {
		return nil, nil
	}

	matching := NewStringSet()
	others := NewStringSet()

	for v := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new StringSet by transforming every element with a function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *StringSet) Map(f func(string) string) *StringSet {
	if set == nil {
		return nil
	}

	result := NewStringSet()

	for v := range set.m {
		result.m[f(v)] = struct{}{}
	}

	return result
}

// MapToInt returns a new []int by transforming every element with function f.
// The resulting slice is the same size as the set.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *StringSet) MapToInt(f func(string) int) []int {
	if set == nil {
		return nil
	}

	result := make([]int, 0, len(set.m))

	for v := range set.m {
		result = append(result, f(v))
	}

	return result
}

// FlatMap returns a new StringSet by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *StringSet) FlatMap(f func(string) []string) *StringSet {
	if set == nil {
		return nil
	}

	result := NewStringSet()

	for v := range set.m {
		for _, x := range f(v) {
			result.m[x] = struct{}{}
		}
	}

	return result
}

// FlatMapToInt returns a new []int by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the set.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *StringSet) FlatMapToInt(f func(string) []int) []int {
	if set == nil {
		return nil
	}

	result := make([]int, 0, len(set.m))

	for v := range set.m {
		result = append(result, f(v)...)
	}

	return result
}

// CountBy gives the number elements of StringSet that return true for the predicate p.
func (set *StringSet) CountBy(p func(string) bool) (result int) {

	for v := range set.m {
		if p(v) {
			result++
		}
	}
	return
}

// Fold aggregates all the values in the set using a supplied function, starting from some initial value.
func (set *StringSet) Fold(initial string, fn func(string, string) string) string {

	m := initial
	for v := range set.m {
		m = fn(m, v)
	}

	return m
}

// MinBy returns an element of StringSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *StringSet) MinBy(less func(string, string) bool) string {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m string
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

// MaxBy returns an element of StringSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set *StringSet) MaxBy(less func(string, string) bool) string {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m string
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

// Equals determines whether two sets are equal to each other, returning true if so.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set *StringSet) Equals(other *StringSet) bool {
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
func (set *StringSet) StringList() []string {

	strings := make([]string, len(set.m))
	i := 0
	for v := range set.m {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set *StringSet) String() string {
	return set.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set *StringSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set *StringSet) MkString3(before, between, after string) string {
	if set == nil {
		return ""
	}
	return set.mkString3Bytes(before, between, after).String()
}

func (set *StringSet) mkString3Bytes(before, between, after string) *strings.Builder {
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
func (set *StringSet) UnmarshalJSON(b []byte) error {

	values := make([]string, 0)
	err := json.Unmarshal(b, &values)
	if err != nil {
		return err
	}

	s2 := NewStringSet(values...)
	*set = *s2
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set *StringSet) MarshalJSON() ([]byte, error) {

	buf, err := json.Marshal(set.ToSlice())
	return buf, err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set *StringSet) StringMap() map[string]bool {
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
// You must register string with the 'gob' package before this method is used.
func (set *StringSet) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&set.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register string with the 'gob' package before this method is used.
func (set StringSet) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(set.m)
	return buf.Bytes(), err
}
