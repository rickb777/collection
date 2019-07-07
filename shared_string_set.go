// An encapsulated map[string]struct{} used as a set.
//
// Thread-safe.
//
// Generated from threadsafe/set.tpl with Type=string
// options: Comparable:always Numeric:<no value> Ordered:<no value> Stringer:true ToList:true
// by runtemplate v3.4.2
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"sync"
)

// SharedStringSet is the primary type that represents a set.
type SharedStringSet struct {
	s *sync.RWMutex
	m map[string]struct{}
}

// NewSharedStringSet creates and returns a reference to an empty set.
func NewSharedStringSet(values ...string) *SharedStringSet {
	set := &SharedStringSet{
		s: &sync.RWMutex{},
		m: make(map[string]struct{}),
	}
	for _, i := range values {
		set.m[i] = struct{}{}
	}
	return set
}

// ConvertSharedStringSet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertSharedStringSet(values ...interface{}) (*SharedStringSet, bool) {
	set := NewSharedStringSet()

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

// BuildSharedStringSetFromChan constructs a new SharedStringSet from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildSharedStringSetFromChan(source <-chan string) *SharedStringSet {
	set := NewSharedStringSet()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set *SharedStringSet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set *SharedStringSet) IsSet() bool {
	return true
}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set *SharedStringSet) ToList() *SharedStringList {
	if set == nil {
		return nil
	}

	set.s.RLock()
	defer set.s.RUnlock()

	return &SharedStringList{
		s: &sync.RWMutex{},
		m: set.slice(),
	}
}

// ToSet returns the set; this is an identity operation in this case.
func (set *SharedStringSet) ToSet() *SharedStringSet {
	return set
}

// slice returns the internal elements of the current set. This is a seam for testing etc.
func (set *SharedStringSet) slice() []string {
	if set == nil {
		return nil
	}

	s := make([]string, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// ToSlice returns the elements of the current set as a slice.
func (set *SharedStringSet) ToSlice() []string {
	set.s.RLock()
	defer set.s.RUnlock()

	return set.slice()
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set *SharedStringSet) ToInterfaceSlice() []interface{} {
	set.s.RLock()
	defer set.s.RUnlock()

	s := make([]interface{}, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set *SharedStringSet) Clone() *SharedStringSet {
	if set == nil {
		return nil
	}

	clonedSet := NewSharedStringSet()

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *SharedStringSet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *SharedStringSet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *SharedStringSet) Size() int {
	if set == nil {
		return 0
	}

	set.s.RLock()
	defer set.s.RUnlock()

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *SharedStringSet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set.
func (set *SharedStringSet) Add(more ...string) {
	set.s.Lock()
	defer set.s.Unlock()

	for _, v := range more {
		set.doAdd(v)
	}
}

func (set *SharedStringSet) doAdd(i string) {
	set.m[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set *SharedStringSet) Contains(i string) bool {
	if set == nil {
		return false
	}

	set.s.RLock()
	defer set.s.RUnlock()

	_, found := set.m[i]
	return found
}

// ContainsAll determines whether the given items are all in the set, returning true if so.
func (set *SharedStringSet) ContainsAll(i ...string) bool {
	if set == nil {
		return false
	}

	set.s.RLock()
	defer set.s.RUnlock()

	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines whether every item in the other set is in this set, returning true if so.
func (set *SharedStringSet) IsSubset(other *SharedStringSet) bool {
	if set.IsEmpty() {
		return !other.IsEmpty()
	}

	if other.IsEmpty() {
		return false
	}

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v := range set.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set *SharedStringSet) IsSuperset(other *SharedStringSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set *SharedStringSet) Union(other *SharedStringSet) *SharedStringSet {
	if set == nil {
		return other
	}

	if other == nil {
		return set
	}

	unionedSet := set.Clone()

	other.s.RLock()
	defer other.s.RUnlock()

	for v := range other.m {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set *SharedStringSet) Intersect(other *SharedStringSet) *SharedStringSet {
	if set == nil || other == nil {
		return nil
	}

	intersection := NewSharedStringSet()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

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
func (set *SharedStringSet) Difference(other *SharedStringSet) *SharedStringSet {
	if set == nil {
		return nil
	}

	if other == nil {
		return set
	}

	differencedSet := NewSharedStringSet()

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

	for v := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *SharedStringSet) SymmetricDifference(other *SharedStringSet) *SharedStringSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear the entire set. Aterwards, it will be an empty set.
func (set *SharedStringSet) Clear() {
	if set != nil {
		set.s.Lock()
		defer set.s.Unlock()

		set.m = make(map[string]struct{})
	}
}

// Remove a single item from the set.
func (set *SharedStringSet) Remove(i string) {
	set.s.Lock()
	defer set.s.Unlock()

	delete(set.m, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set *SharedStringSet) Send() <-chan string {
	ch := make(chan string)
	go func() {
		if set != nil {
			set.s.RLock()
			defer set.s.RUnlock()

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
func (set *SharedStringSet) Forall(p func(string) bool) bool {
	if set == nil {
		return true
	}

	set.s.RLock()
	defer set.s.RUnlock()

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
func (set *SharedStringSet) Exists(p func(string) bool) bool {
	if set == nil {
		return false
	}

	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over the set and executes the function f against each element.
// The function can safely alter the values via side-effects.
func (set *SharedStringSet) Foreach(f func(string)) {
	if set == nil {
		return
	}

	set.s.Lock()
	defer set.s.Unlock()

	for v := range set.m {
		f(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first string that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set *SharedStringSet) Find(p func(string) bool) (string, bool) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			return v, true
		}
	}

	var empty string
	return empty, false
}

// Filter returns a new SharedStringSet whose elements return true for the predicate p.
//
// The original set is not modified
func (set *SharedStringSet) Filter(p func(string) bool) *SharedStringSet {
	if set == nil {
		return nil
	}

	result := NewSharedStringSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new SharedStringSets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original set is not modified
func (set *SharedStringSet) Partition(p func(string) bool) (*SharedStringSet, *SharedStringSet) {
	if set == nil {
		return nil, nil
	}

	matching := NewSharedStringSet()
	others := NewSharedStringSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new SharedStringSet by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *SharedStringSet) Map(f func(string) string) *SharedStringSet {
	if set == nil {
		return nil
	}

	result := NewSharedStringSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		k := f(v)
		result.m[k] = struct{}{}
	}

	return result
}

// MapToInt returns a new []int by transforming every element with function f.
// The resulting slice is the same size as the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *SharedStringSet) MapToInt(f func(string) int) []int {
	if set == nil {
		return nil
	}

	result := make([]int, 0, len(set.m))
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		result = append(result, f(v))
	}

	return result
}

// FlatMap returns a new SharedStringSet by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *SharedStringSet) FlatMap(f func(string) []string) *SharedStringSet {
	if set == nil {
		return nil
	}

	result := NewSharedStringSet()
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		for _, x := range f(v) {
			result.m[x] = struct{}{}
		}
	}

	return result
}

// FlatMapToInt returns a new []int by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *SharedStringSet) FlatMapToInt(f func(string) []int) []int {
	if set == nil {
		return nil
	}

	result := make([]int, 0, len(set.m))
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		result = append(result, f(v)...)
	}

	return result
}

// CountBy gives the number elements of SharedStringSet that return true for the predicate p.
func (set *SharedStringSet) CountBy(p func(string) bool) (result int) {
	set.s.RLock()
	defer set.s.RUnlock()

	for v := range set.m {
		if p(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of SharedStringSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *SharedStringSet) MinBy(less func(string, string) bool) string {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

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

// MaxBy returns an element of SharedStringSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set *SharedStringSet) MaxBy(less func(string, string) bool) string {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	set.s.RLock()
	defer set.s.RUnlock()

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
func (set *SharedStringSet) Equals(other *SharedStringSet) bool {
	if set == nil {
		return other == nil || other.IsEmpty()
	}

	if other == nil {
		return set.IsEmpty()
	}

	set.s.RLock()
	other.s.RLock()
	defer set.s.RUnlock()
	defer other.s.RUnlock()

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

// StringSet gets a list of strings that depicts all the elements.
func (set *SharedStringSet) StringList() []string {
	set.s.RLock()
	defer set.s.RUnlock()

	strings := make([]string, len(set.m))
	i := 0
	for v := range set.m {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set *SharedStringSet) String() string {
	return set.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set *SharedStringSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set *SharedStringSet) MkString3(before, between, after string) string {
	if set == nil {
		return ""
	}
	return set.mkString3Bytes(before, between, after).String()
}

func (set *SharedStringSet) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	set.s.RLock()
	defer set.s.RUnlock()

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
func (set *SharedStringSet) UnmarshalJSON(b []byte) error {
	set.s.Lock()
	defer set.s.Unlock()

	values := make([]string, 0)
	err := json.Unmarshal(b, &values)
	if err != nil {
		return err
	}

	s2 := NewSharedStringSet(values...)
	*set = *s2
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set *SharedStringSet) MarshalJSON() ([]byte, error) {
	set.s.RLock()
	defer set.s.RUnlock()

	buf, err := json.Marshal(set.ToSlice())
	return buf, err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set *SharedStringSet) StringMap() map[string]bool {
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
func (set *SharedStringSet) GobDecode(b []byte) error {
	set.s.Lock()
	defer set.s.Unlock()

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&set.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register string with the 'gob' package before this method is used.
func (set SharedStringSet) GobEncode() ([]byte, error) {
	set.s.RLock()
	defer set.s.RUnlock()

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(set.m)
	return buf.Bytes(), err
}
