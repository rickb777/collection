// A simple type derived from map[interface{}]struct{}
//
// Not thread-safe.
//
// Generated from simple/set.tpl with Type=interface{}
// options: Numeric:<no value> Stringer:true Mutable:always
// by runtemplate v3.6.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// AnySet is the primary type that represents a set
type AnySet map[interface{}]struct{}

// NewAnySet creates and returns a reference to an empty set.
func NewAnySet(values ...interface{}) AnySet {
	set := make(AnySet)
	for _, i := range values {
		set[i] = struct{}{}
	}
	return set
}

// ConvertAnySet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
func ConvertAnySet(values ...interface{}) (AnySet, bool) {
	set := make(AnySet)

	for _, i := range values {
		switch j := i.(type) {
		case interface{}:
			k := interface{}(j)
			set[k] = struct{}{}
		case *interface{}:
			k := interface{}(*j)
			set[k] = struct{}{}
		}
	}

	return set, len(set) == len(values)
}

// BuildAnySetFromChan constructs a new AnySet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildAnySetFromChan(source <-chan interface{}) AnySet {
	set := make(AnySet)
	for v := range source {
		set[v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set AnySet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set AnySet) IsSet() bool {
	return true
}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set AnySet) ToList() AnyList {
	if set == nil {
		return nil
	}

	return AnyList(set.ToSlice())
}

// ToSet returns the set; this is an identity operation in this case.
func (set AnySet) ToSet() AnySet {
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set AnySet) ToSlice() []interface{} {
	s := make([]interface{}, 0, len(set))
	for v := range set {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set AnySet) ToInterfaceSlice() []interface{} {
	s := make([]interface{}, 0, len(set))
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set AnySet) Clone() AnySet {
	clonedSet := NewAnySet()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set AnySet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set AnySet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set AnySet) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set AnySet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set, returning the modified set.
func (set AnySet) Add(more ...interface{}) AnySet {
	for _, v := range more {
		set.doAdd(v)
	}
	return set
}

func (set AnySet) doAdd(i interface{}) {
	set[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set AnySet) Contains(i interface{}) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines whether a given item is already in the set, returning true if so.
func (set AnySet) ContainsAll(i ...interface{}) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines whether every item in the other set is in this set, returning true if so.
func (set AnySet) IsSubset(other AnySet) bool {
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set AnySet) IsSuperset(other AnySet) bool {
	return other.IsSubset(set)
}

// Append inserts more items into a clone of the set. It returns the augmented set.
// The original set is unmodified.
func (set AnySet) Append(more ...interface{}) AnySet {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set AnySet) Union(other AnySet) AnySet {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set AnySet) Intersect(other AnySet) AnySet {
	intersection := NewAnySet()
	// loop over smaller set
	if set.Size() < other.Size() {
		for v := range set {
			if other.Contains(v) {
				intersection.doAdd(v)
			}
		}
	} else {
		for v := range other {
			if set.Contains(v) {
				intersection.doAdd(v)
			}
		}
	}

	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set AnySet) Difference(other AnySet) AnySet {
	differencedSet := NewAnySet()
	for v := range set {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set AnySet) SymmetricDifference(other AnySet) AnySet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear the entire set. Aterwards, it will be an empty set.
func (set *AnySet) Clear() {
	*set = NewAnySet()
}

// Remove a single item from the set.
func (set AnySet) Remove(i interface{}) {
	delete(set, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set AnySet) Send() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for v := range set {
			ch <- v
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
func (set AnySet) Forall(p func(interface{}) bool) bool {
	for v := range set {
		if !p(v) {
			return false
		}
	}
	return true
}

// Exists applies a predicate p to every element in the set. If the function returns true,
// the iteration terminates early. The returned value is true if an early return occurred.
// or false if all elements were visited without finding a match.
func (set AnySet) Exists(p func(interface{}) bool) bool {
	for v := range set {
		if p(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over the set and executes the function f against each element.
func (set AnySet) Foreach(f func(interface{})) {
	for v := range set {
		f(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first interface{} that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set AnySet) Find(p func(interface{}) bool) (interface{}, bool) {

	for v := range set {
		if p(v) {
			return v, true
		}
	}

	var empty interface{}
	return empty, false
}

// Filter returns a new AnySet whose elements return true for the predicate p.
//
// The original set is not modified
func (set AnySet) Filter(p func(interface{}) bool) AnySet {
	result := NewAnySet()
	for v := range set {
		if p(v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new AnySets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't.
//
// The original set is not modified
func (set AnySet) Partition(p func(interface{}) bool) (AnySet, AnySet) {
	matching := NewAnySet()
	others := NewAnySet()
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// Map returns a new AnySet by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set AnySet) Map(f func(interface{}) interface{}) AnySet {
	result := NewAnySet()

	for v := range set {
		k := f(v)
		result[k] = struct{}{}
	}

	return result
}

// FlatMap returns a new AnySet by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set AnySet) FlatMap(f func(interface{}) []interface{}) AnySet {
	result := NewAnySet()

	for v := range set {
		for _, x := range f(v) {
			result[x] = struct{}{}
		}
	}

	return result
}

// CountBy gives the number elements of AnySet that return true for the predicate p.
func (set AnySet) CountBy(p func(interface{}) bool) (result int) {
	for v := range set {
		if p(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of AnySet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set AnySet) MinBy(less func(interface{}, interface{}) bool) interface{} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}
	var m interface{}
	first := true
	for v := range set {
		if first {
			m = v
			first = false
		} else if less(v, m) {
			m = v
		}
	}
	return m
}

// MaxBy returns an element of AnySet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set AnySet) MaxBy(less func(interface{}, interface{}) bool) interface{} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}
	var m interface{}
	first := true
	for v := range set {
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
func (set AnySet) Equals(other AnySet) bool {
	if set.Size() != other.Size() {
		return false
	}

	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}

	return true
}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (set AnySet) StringList() []string {
	strings := make([]string, len(set))
	i := 0
	for v := range set {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set AnySet) String() string {
	return set.mkString3Bytes("[", ", ", "]").String()
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set AnySet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set AnySet) MkString3(before, between, after string) string {
	return set.mkString3Bytes(before, between, after).String()
}

func (set AnySet) mkString3Bytes(before, between, after string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""
	for v := range set {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this set type.
func (set AnySet) UnmarshalJSON(b []byte) error {
	values := make([]interface{}, 0)
	buf := bytes.NewBuffer(b)
	err := json.NewDecoder(buf).Decode(&values)
	if err != nil {
		return err
	}
	set.Add(values...)
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set AnySet) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(set.ToSlice())
	return buf.Bytes(), err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set AnySet) StringMap() map[string]bool {
	strings := make(map[string]bool)
	for v := range set {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}
