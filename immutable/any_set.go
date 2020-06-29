// An encapsulated immutable map[interface{}]struct{} used as a set.
// Thread-safe.
//
//
// Generated from immutable/set.tpl with Type=interface{}
// options: Comparable:always Numeric:<no value> Ordered:<no value> Stringer:<no value> Mutable:disabled
// by runtemplate v3.5.4
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

// AnySet is the primary type that represents a set.
type AnySet struct {
	m map[interface{}]struct{}
}

// NewAnySet creates and returns a reference to an empty set.
func NewAnySet(values ...interface{}) *AnySet {
	set := &AnySet{
		m: make(map[interface{}]struct{}),
	}
	for _, i := range values {
		set.m[i] = struct{}{}
	}
	return set
}

// ConvertAnySet constructs a new set containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned set will contain all the values that were correctly converted.
func ConvertAnySet(values ...interface{}) (*AnySet, bool) {
	set := NewAnySet()

	for _, i := range values {
		switch j := i.(type) {
		case interface{}:
			set.m[j] = struct{}{}
		case *interface{}:
			set.m[*j] = struct{}{}
		}
	}

	return set, len(set.m) == len(values)
}

// BuildAnySetFromChan constructs a new AnySet from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildAnySetFromChan(source <-chan interface{}) *AnySet {
	set := NewAnySet()
	for v := range source {
		set.m[v] = struct{}{}
	}
	return set
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (set *AnySet) IsSequence() bool {
	return false
}

// IsSet returns false for lists or queues.
func (set *AnySet) IsSet() bool {
	return true
}

// ToList returns the elements of the set as a list. The returned list is a shallow
// copy; the set is not altered.
func (set *AnySet) ToList() *AnyList {
	if set == nil {
		return nil
	}

	return &AnyList{
		m: set.ToSlice(),
	}
}

// ToSet returns the set; this is an identity operation in this case.
func (set *AnySet) ToSet() *AnySet {
	return set
}

// ToSlice returns the elements of the current set as a slice.
func (set *AnySet) ToSlice() []interface{} {
	if set == nil {
		return nil
	}

	s := make([]interface{}, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// ToInterfaceSlice returns the elements of the current set as a slice of arbitrary type.
func (set *AnySet) ToInterfaceSlice() []interface{} {

	s := make([]interface{}, 0, len(set.m))
	for v := range set.m {
		s = append(s, v)
	}
	return s
}

// Clone returns the same set, which is immutable.
func (set *AnySet) Clone() *AnySet {
	return set
}

//-------------------------------------------------------------------------------------------------

// IsEmpty returns true if the set is empty.
func (set *AnySet) IsEmpty() bool {
	return set.Size() == 0
}

// NonEmpty returns true if the set is not empty.
func (set *AnySet) NonEmpty() bool {
	return set.Size() > 0
}

// Size returns how many items are currently in the set. This is a synonym for Cardinality.
func (set *AnySet) Size() int {
	if set == nil {
		return 0
	}

	return len(set.m)
}

// Cardinality returns how many items are currently in the set. This is a synonym for Size.
func (set *AnySet) Cardinality() int {
	return set.Size()
}

//-------------------------------------------------------------------------------------------------

// Add returns a new set with all original items and all in `more`.
// The original set is not altered.
func (set *AnySet) Add(more ...interface{}) *AnySet {
	newSet := NewAnySet()

	for v := range set.m {
		newSet.doAdd(v)
	}

	for _, v := range more {
		newSet.doAdd(v)
	}

	return newSet
}

func (set *AnySet) doAdd(i interface{}) {
	set.m[i] = struct{}{}
}

// Contains determines whether a given item is already in the set, returning true if so.
func (set *AnySet) Contains(i interface{}) bool {
	if set == nil {
		return false
	}

	_, found := set.m[i]
	return found
}

// ContainsAll determines whether a given item is already in the set, returning true if so.
func (set *AnySet) ContainsAll(i ...interface{}) bool {
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
func (set *AnySet) IsSubset(other *AnySet) bool {
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
func (set *AnySet) IsSuperset(other *AnySet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set *AnySet) Union(other *AnySet) *AnySet {
	if set == nil {
		return other
	}

	if other == nil {
		return set
	}

	unionedSet := NewAnySet()

	for v := range set.m {
		unionedSet.doAdd(v)
	}

	for v := range other.m {
		unionedSet.doAdd(v)
	}

	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set *AnySet) Intersect(other *AnySet) *AnySet {
	if set == nil || other == nil {
		return nil
	}

	intersection := NewAnySet()

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
func (set *AnySet) Difference(other *AnySet) *AnySet {
	if set == nil {
		return nil
	}

	if other == nil {
		return set
	}

	differencedSet := NewAnySet()

	for v := range set.m {
		if !other.Contains(v) {
			differencedSet.doAdd(v)
		}
	}

	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set *AnySet) SymmetricDifference(other *AnySet) *AnySet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Remove removes a single item from the set. A new set is returned that has all the elements except the removed one.
func (set *AnySet) Remove(i interface{}) *AnySet {
	if set == nil {
		return nil
	}

	clonedSet := NewAnySet()

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
func (set *AnySet) Send() <-chan interface{} {
	ch := make(chan interface{})
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
func (set *AnySet) Forall(p func(interface{}) bool) bool {
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
func (set *AnySet) Exists(p func(interface{}) bool) bool {
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

// Foreach iterates over interface{}Set and executes the function f against each element.
func (set *AnySet) Foreach(f func(interface{})) {
	if set == nil {
		return
	}

	for v := range set.m {
		f(v)
	}
}

//-------------------------------------------------------------------------------------------------

// Find returns the first interface{} that returns true for the predicate p. If there are many matches
// one is arbtrarily chosen. False is returned if none match.
func (set *AnySet) Find(p func(interface{}) bool) (interface{}, bool) {

	for v := range set.m {
		if p(v) {
			return v, true
		}
	}

	var empty interface{}
	return empty, false
}

// Filter returns a new AnySet whose elements return true for the predicate p.
func (set *AnySet) Filter(p func(interface{}) bool) *AnySet {
	if set == nil {
		return nil
	}

	result := NewAnySet()

	for v := range set.m {
		if p(v) {
			result.doAdd(v)
		}
	}
	return result
}

// Partition returns two new interface{}Sets whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (set *AnySet) Partition(p func(interface{}) bool) (*AnySet, *AnySet) {
	if set == nil {
		return nil, nil
	}

	matching := NewAnySet()
	others := NewAnySet()

	for v := range set.m {
		if p(v) {
			matching.doAdd(v)
		} else {
			others.doAdd(v)
		}
	}
	return matching, others
}

// Map returns a new AnySet by transforming every element with a function f.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *AnySet) Map(f func(interface{}) interface{}) *AnySet {
	if set == nil {
		return nil
	}

	result := NewAnySet()

	for v := range set.m {
		result.m[f(v)] = struct{}{}
	}

	return result
}

// FlatMap returns a new AnySet by transforming every element with a function f that
// returns zero or more items in a slice. The resulting set may have a different size to the original set.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (set *AnySet) FlatMap(f func(interface{}) []interface{}) *AnySet {
	if set == nil {
		return nil
	}

	result := NewAnySet()

	for v := range set.m {
		for _, x := range f(v) {
			result.m[x] = struct{}{}
		}
	}

	return result
}

// CountBy gives the number elements of AnySet that return true for the predicate p.
func (set *AnySet) CountBy(p func(interface{}) bool) (result int) {

	for v := range set.m {
		if p(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of AnySet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set *AnySet) MinBy(less func(interface{}, interface{}) bool) interface{} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m interface{}
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

// MaxBy returns an element of AnySet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (set *AnySet) MaxBy(less func(interface{}, interface{}) bool) interface{} {
	if set.IsEmpty() {
		panic("Cannot determine the minimum of an empty set.")
	}

	var m interface{}
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
func (set *AnySet) Equals(other *AnySet) bool {
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
