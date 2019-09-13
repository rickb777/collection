// A simple type derived from map[uint64]struct{}
//
// Not thread-safe.
//
// Generated from simple/set.tpl with Type=uint64
// options: Numeric:true Stringer:true Mutable:always
// by runtemplate v3.5.2
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Uint64Set is the primary type that represents a set
type Uint64Set map[uint64]struct{}

// NewUint64Set creates and returns a reference to an empty set.
func NewUint64Set(values ...uint64) Uint64Set {
	set := make(Uint64Set)
	for _, i := range values {
		set[i] = struct{}{}
	}
	return set
}

// ConvertUint64Set constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
func ConvertUint64Set(values ...interface{}) (Uint64Set, bool) {
	set := make(Uint64Set)

	for _, i := range values {
		switch j := i.(type) {
		case int:
			k := uint64(j)
			set[k] = struct{}{}
		case *int:
			k := uint64(*j)
			set[k] = struct{}{}
		case int8:
			k := uint64(j)
			set[k] = struct{}{}
		case *int8:
			k := uint64(*j)
			set[k] = struct{}{}
		case int16:
			k := uint64(j)
			set[k] = struct{}{}
		case *int16:
			k := uint64(*j)
			set[k] = struct{}{}
		case int32:
			k := uint64(j)
			set[k] = struct{}{}
		case *int32:
			k := uint64(*j)
			set[k] = struct{}{}
		case int64:
			k := uint64(j)
			set[k] = struct{}{}
		case *int64:
			k := uint64(*j)
			set[k] = struct{}{}
		case uint:
			k := uint64(j)
			set[k] = struct{}{}
		case *uint:
			k := uint64(*j)
			set[k] = struct{}{}
		case uint8:
			k := uint64(j)
			set[k] = struct{}{}
		case *uint8:
			k := uint64(*j)
			set[k] = struct{}{}
		case uint16:
			k := uint64(j)
			set[k] = struct{}{}
		case *uint16:
			k := uint64(*j)
			set[k] = struct{}{}
		case uint32:
			k := uint64(j)
			set[k] = struct{}{}
		case *uint32:
			k := uint64(*j)
			set[k] = struct{}{}
		case uint64:
			k := uint64(j)
			set[k] = struct{}{}
		case *uint64:
			k := uint64(*j)
			set[k] = struct{}{}
		case float32:
			k := uint64(j)
			set[k] = struct{}{}
		case *float32:
			k := uint64(*j)
			set[k] = struct{}{}
		case float64:
			k := uint64(j)
			set[k] = struct{}{}
		case *float64:
			k := uint64(*j)
			set[k] = struct{}{}
		}
	}

	return set, len(set) == len(values)
}

// BuildUint64SetFromChan constructs a new Uint64Set from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildUint64SetFromChan(source <-chan uint64) Uint64Set {
	set := make(Uint64Set)
	for v := range source {
		set[v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set Uint64Set) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set Uint64Set) IsSet() bool {
	return true
}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set Uint64Set) ToList() Uint64List {
	if set == nil {
		return nil
	}

	return Uint64List(set.ToSlice())
}

// ToSet returns the set; this is an identity operation in this case.
func (set Uint64Set) ToSet() Uint64Set {
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set Uint64Set) ToSlice() []uint64 {
	s := make([]uint64, 0, len(set))
	for v := range set {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set Uint64Set) ToInterfaceSlice() []interface{} {
	s := make([]interface{}, 0, len(set))
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the set. It does not clone the underlying elements.
func (set Uint64Set) Clone() Uint64Set {
	clonedSet := NewUint64Set()
	for v := range set {
		clonedSet.doAdd(v)
	}
	return clonedSet
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set Uint64Set) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set Uint64Set) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set Uint64Set) Size() int {
	return len(set)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set Uint64Set) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add adds items to the current set, returning the modified set.
func (set Uint64Set) Add(more ...uint64) Uint64Set {
	for _, v := range more {
		set.doAdd(v)
	}
	return set
}

func (set Uint64Set) doAdd(i uint64) {
	set[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set Uint64Set) Contains(i uint64) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines whether a given item is already in the set, returning true if so.
func (set Uint64Set) ContainsAll(i ...uint64) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------------------------------------

// IsSubset determines whether every item in the other set is in this set, returning true if so.
func (set Uint64Set) IsSubset(other Uint64Set) bool {
	for v := range set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset determines whether every item of this set is in the other set, returning true if so.
func (set Uint64Set) IsSuperset(other Uint64Set) bool {
	return other.IsSubset(set)
}

// Append inserts more items into a clone of the set. It returns the augmented set.
// The original set is unmodified.
func (set Uint64Set) Append(more ...uint64) Uint64Set {
	unionedSet := set.Clone()
	for _, v := range more {
		unionedSet.doAdd(v)
	}
	return unionedSet
}

// Union returns a new set with all items in both sets.
func (set Uint64Set) Union(other Uint64Set) Uint64Set {
	unionedSet := set.Clone()
	for v := range other {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set Uint64Set) Intersect(other Uint64Set) Uint64Set {
	intersection := NewUint64Set()
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
func (set Uint64Set) Difference(other Uint64Set) Uint64Set {
	differencedSet := NewUint64Set()
	for v := range set {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set Uint64Set) SymmetricDifference(other Uint64Set) Uint64Set {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear the entire set. Aterwards, it will be an empty set.
func (set *Uint64Set) Clear() {
	*set = NewUint64Set()
}

// Remove a single item from the set.
func (set Uint64Set) Remove(i uint64) {
	delete(set, i)
}

//-------------------------------------------------------------------------------------------------

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
func (set Uint64Set) Send() <-chan uint64 {
	ch := make(chan uint64)
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
func (set Uint64Set) Forall(p func(uint64) bool) bool {
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
func (set Uint64Set) Exists(p func(uint64) bool) bool {
	for v := range set {
		if p(v) {
			return true
		}
	}
	return false
}

// Foreach iterates over the set and executes the function f against each element.
func (set Uint64Set) Foreach(f func(uint64)) {
	for v := range set {
		f(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first uint64 that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set Uint64Set) Find(p func(uint64) bool) (uint64, bool) {

	for v := range set {
		if p(v) {
			return v, true
		}
	}

	var empty uint64
	return empty, false
}

// Filter returns a new Uint64Set whose elements return true for the predicate p.
//
// The original set is not modified
func (set Uint64Set) Filter(p func(uint64) bool) Uint64Set {
	result := NewUint64Set()
	for v := range set {
		if p(v) {
			result[v] = struct{}{}
		}
	}
	return result
}

// Partition returns two new Uint64Sets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't.
//
// The original set is not modified
func (set Uint64Set) Partition(p func(uint64) bool) (Uint64Set, Uint64Set) {
	matching := NewUint64Set()
	others := NewUint64Set()
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return matching, others
}

// Map returns a new Uint64Set by transforming every element with a function f.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set Uint64Set) Map(f func(uint64) uint64) Uint64Set {
	result := NewUint64Set()

	for v := range set {
		k := f(v)
		result[k] = struct{}{}
	}

	return result
}

// MapToString returns a new []string by transforming every element with function f.
// The resulting slice is the same size as the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set Uint64Set) MapToString(f func(uint64) string) []string {
	if set == nil {
		return nil
	}

	result := make([]string, 0, len(set))
	for v := range set {
		result = append(result, f(v))
	}

	return result
}

// FlatMap returns a new Uint64Set by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
// The original set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set Uint64Set) FlatMap(f func(uint64) []uint64) Uint64Set {
	result := NewUint64Set()

	for v := range set {
		for _, x := range f(v) {
			result[x] = struct{}{}
		}
	}

	return result
}

// FlatMapToString returns a new []string by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the set.
// The set is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set Uint64Set) FlatMapToString(f func(uint64) []string) []string {
	if set == nil {
		return nil
	}

	result := make([]string, 0, len(set))
	for v := range set {
		result = append(result, f(v)...)
	}

	return result
}

// CountBy gives the number elements of Uint64Set that return true for the predicate p.
func (set Uint64Set) CountBy(p func(uint64) bool) (result int) {
	for v := range set {
		if p(v) {
			result++
		}
	}
	return
}

//-------------------------------------------------------------------------------------------------
// These methods are included when uint64 is ordered.

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (set Uint64Set) Min() uint64 {
	v := set.MinBy(func(a uint64, b uint64) bool {
		return a < b
	})
	return v
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (set Uint64Set) Max() uint64 {
	v := set.MaxBy(func(a uint64, b uint64) bool {
		return a < b
	})
	return v
}

// MinBy returns an element of Uint64Set containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set Uint64Set) MinBy(less func(uint64, uint64) bool) uint64 {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}
	var m uint64
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

// MaxBy returns an element of Uint64Set containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set Uint64Set) MaxBy(less func(uint64, uint64) bool) uint64 {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}
	var m uint64
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
// These methods are included when uint64 is numeric.

// Sum returns the sum of all the elements in the set.
func (set Uint64Set) Sum() uint64 {
	sum := uint64(0)
	for v := range set {
		sum = sum + v
	}
	return sum
}

//-------------------------------------------------------------------------------------------------

// Equals determines whether two sets are equal to each other, returning true if so.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set Uint64Set) Equals(other Uint64Set) bool {
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
func (set Uint64Set) StringList() []string {
	strings := make([]string, len(set))
	i := 0
	for v := range set {
		strings[i] = fmt.Sprintf("%v", v)
		i++
	}
	return strings
}

// String implements the Stringer interface to render the set as a comma-separated string enclosed in square brackets.
func (set Uint64Set) String() string {
	return set.mkString3Bytes("[", ", ", "]").String()
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (set Uint64Set) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (set Uint64Set) MkString3(before, between, after string) string {
	return set.mkString3Bytes(before, between, after).String()
}

func (set Uint64Set) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
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
func (set Uint64Set) UnmarshalJSON(b []byte) error {
	values := make([]uint64, 0)
	buf := bytes.NewBuffer(b)
	err := json.NewDecoder(buf).Decode(&values)
	if err != nil {
		return err
	}
	set.Add(values...)
	return nil
}

// MarshalJSON implements JSON encoding for this set type.
func (set Uint64Set) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(set.ToSlice())
	return buf.Bytes(), err
}

// StringMap renders the set as a map of strings. The value of each item in the set becomes stringified as a key in the
// resulting map.
func (set Uint64Set) StringMap() map[string]bool {
	strings := make(map[string]bool)
	for v := range set {
		strings[fmt.Sprintf("%v", v)] = true
	}
	return strings
}
