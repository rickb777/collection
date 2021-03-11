// An encapsulated immutable []interface{}.
// Thread-safe.
//
//
// Generated from immutable/list.tpl with Type=interface{}
// options: Comparable:true Numeric:<no value> Integer:<no value> Ordered:<no value>
//          StringLike:<no value> StringParser:<no value> Stringer:true
// by runtemplate v3.10.1
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package immutable

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// AnyList contains a slice of type interface{}. It is designed
// to be immutable - ideal for race-free reference lists etc.
// It encapsulates the slice and provides methods to access it.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type AnyList struct {
	m []interface{}
}

//-------------------------------------------------------------------------------------------------

func newAnyList(length, capacity int) *AnyList {
	return &AnyList{
		m: make([]interface{}, length, capacity),
	}
}

// NewAnyList constructs a new list containing the supplied values, if any.
func NewAnyList(values ...interface{}) *AnyList {
	list := newAnyList(len(values), len(values))
	copy(list.m, values)
	return list
}

// ConvertAnyList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertAnyList(values ...interface{}) (*AnyList, bool) {
	list := newAnyList(0, len(values))

	for _, i := range values {
		switch j := i.(type) {
		case interface{}:
			k := interface{}(j)
			list.m = append(list.m, k)
		case *interface{}:
			k := interface{}(*j)
			list.m = append(list.m, k)
		}
	}

	return list, len(list.m) == len(values)
}

// BuildAnyListFromChan constructs a new AnyList from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildAnyListFromChan(source <-chan interface{}) *AnyList {
	list := newAnyList(0, 0)
	for v := range source {
		list.m = append(list.m, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list *AnyList) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list *AnyList) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list *AnyList) slice() []interface{} {
	if list == nil {
		return nil
	}
	return list.m
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list *AnyList) ToList() *AnyList {
	return list
}

// ToSet returns the elements of the list as a set. The returned set is a shallow
// copy; the list is not altered.
func (list *AnyList) ToSet() *AnySet {
	if list == nil {
		return nil
	}

	return NewAnySet(list.m...)
}

// ToSlice returns the elements of the current list as a slice.
func (list *AnyList) ToSlice() []interface{} {

	s := make([]interface{}, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *AnyList) ToInterfaceSlice() []interface{} {

	s := make([]interface{}, 0, len(list.m))
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns the same list, which is immutable.
func (list *AnyList) Clone() *AnyList {
	return list
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
func (list *AnyList) Get(i int) interface{} {
	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list *AnyList) Head() interface{} {
	return list.m[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns the zero value.
func (list *AnyList) HeadOption() (interface{}, bool) {
	if list == nil || len(list.m) == 0 {
		var v interface{}
		return v, false
	}
	return list.m[0], true
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list *AnyList) Last() interface{} {
	return list.m[len(list.m)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns the zero value.
func (list *AnyList) LastOption() (interface{}, bool) {
	if list == nil || len(list.m) == 0 {
		var v interface{}
		return v, false
	}
	return list.m[len(list.m)-1], true
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list *AnyList) Tail() *AnyList {
	result := newAnyList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list *AnyList) Init() *AnyList {
	result := newAnyList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether AnyList is empty.
func (list *AnyList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether AnyList is empty.
func (list *AnyList) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *AnyList) Size() int {
	if list == nil {
		return 0
	}

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
func (list *AnyList) Len() int {
	return list.Size()
}

//-------------------------------------------------------------------------------------------------

// Contains determines whether a given item is already in the list, returning true if so.
func (list *AnyList) Contains(v interface{}) bool {
	return list.Exists(func(x interface{}) bool {
		return x == v
	})
}

// ContainsAll determines whether the given items are all in the list, returning true if so.
// This is potentially a slow method and should only be used rarely.
func (list *AnyList) ContainsAll(i ...interface{}) bool {
	if list == nil {
		return len(i) == 0
	}

	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of AnyList return true for the predicate p.
func (list *AnyList) Exists(p func(interface{}) bool) bool {
	if list == nil {
		return false
	}

	for _, v := range list.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of AnyList return true for the predicate p.
func (list *AnyList) Forall(p func(interface{}) bool) bool {
	if list == nil {
		return true
	}

	for _, v := range list.m {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over AnyList and executes function f against each element.
// The function receives copies that do not alter the list elements when they are changed.
func (list *AnyList) Foreach(f func(interface{})) {
	if list == nil {
		return
	}

	for _, v := range list.m {
		f(v)
	}
}

// Send returns a channel that will send all the elements in order. A goroutine is created to
// send the elements; this only terminates when all the elements have been consumed. The
// channel will be closed when all the elements have been sent.
func (list *AnyList) Send() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		if list != nil {
			for _, v := range list.m {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// Reverse returns a copy of AnyList with all elements in the reverse order.
func (list *AnyList) Reverse() *AnyList {
	if list == nil {
		return nil
	}

	n := len(list.m)
	result := newAnyList(n, n)
	last := n - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of AnyList, using a version of the Fisher-Yates shuffle.
func (list *AnyList) Shuffle() *AnyList {
	if list == nil {
		return nil
	}

	result := NewAnyList(list.m...)
	n := len(result.m)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		result.m[i], result.m[r] = result.m[r], result.m[i]
	}
	return result
}

// Append returns a new list with all original items and all in `more`; they retain their order.
// The original list is not altered.
func (list *AnyList) Append(more ...interface{}) *AnyList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		return NewAnyList(more...)
	}

	newList := NewAnyList(list.m...)
	newList.doAppend(more...)
	return newList
}

func (list *AnyList) doAppend(more ...interface{}) {
	list.m = append(list.m, more...)
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of AnyList containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *AnyList) Take(n int) *AnyList {
	if list == nil || n >= len(list.m) {
		return list
	}

	result := newAnyList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of AnyList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *AnyList) Drop(n int) *AnyList {
	if list == nil || n == 0 {
		return list
	}

	if n >= len(list.m) {
		return nil
	}

	result := newAnyList(0, 0)
	result.m = list.m[n:]
	return result
}

// TakeLast returns a slice of AnyList containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *AnyList) TakeLast(n int) *AnyList {
	if list == nil {
		return nil
	}

	l := len(list.m)
	if n >= l {
		return list
	}

	result := newAnyList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of AnyList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
func (list *AnyList) DropLast(n int) *AnyList {
	if list == nil || n == 0 {
		return list
	}

	l := len(list.m)
	if n >= l {
		return nil
	}

	result := newAnyList(0, 0)
	result.m = list.m[:l-n]
	return result
}

// TakeWhile returns a new AnyList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
func (list *AnyList) TakeWhile(p func(interface{}) bool) *AnyList {
	if list == nil {
		return nil
	}

	result := newAnyList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new AnyList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
func (list *AnyList) DropWhile(p func(interface{}) bool) *AnyList {
	if list == nil {
		return nil
	}

	result := newAnyList(0, 0)
	adding := false

	for _, v := range list.m {
		if adding || !p(v) {
			adding = true
			result.m = append(result.m, v)
		}
	}

	return result
}

//-------------------------------------------------------------------------------------------------

// Find returns the first interface{} that returns true for predicate p.
// False is returned if none match.
func (list *AnyList) Find(p func(interface{}) bool) (interface{}, bool) {
	if list == nil {
		return nil, false
	}

	for _, v := range list.m {
		if p(v) {
			return v, true
		}
	}

	var empty interface{}
	return empty, false
}

// Filter returns a new AnyList whose elements return true for predicate p.
func (list *AnyList) Filter(p func(interface{}) bool) *AnyList {
	if list == nil {
		return nil
	}

	result := newAnyList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new interface{}Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list *AnyList) Partition(p func(interface{}) bool) (*AnyList, *AnyList) {
	if list == nil {
		return nil, nil
	}

	matching := newAnyList(0, len(list.m)/2)
	others := newAnyList(0, len(list.m)/2)

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new AnyList by transforming every element with function f.
// The resulting list is the same size as the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *AnyList) Map(f func(interface{}) interface{}) *AnyList {
	if list == nil {
		return nil
	}

	result := newAnyList(len(list.m), len(list.m))

	for i, v := range list.m {
		result.m[i] = f(v)
	}

	return result
}

// FlatMap returns a new AnyList by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *AnyList) FlatMap(f func(interface{}) []interface{}) *AnyList {
	if list == nil {
		return nil
	}

	result := newAnyList(0, len(list.m))

	for _, v := range list.m {
		result.m = append(result.m, f(v)...)
	}

	return result
}

// CountBy gives the number elements of AnyList that return true for the predicate p.
func (list *AnyList) CountBy(p func(interface{}) bool) (result int) {
	if list == nil {
		return 0
	}

	for _, v := range list.m {
		if p(v) {
			result++
		}
	}
	return
}

// Fold aggregates all the values in the list using a supplied function, starting from some initial value.
func (list *AnyList) Fold(initial interface{}, fn func(interface{}, interface{}) interface{}) interface{} {
	m := initial
	for _, v := range list.m {
		m = fn(m, v)
	}

	return m
}

// MinBy returns an element of AnyList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *AnyList) MinBy(less func(interface{}, interface{}) bool) interface{} {
	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
	}

	m := 0
	for i := 1; i < l; i++ {
		if less(list.m[i], list.m[m]) {
			m = i
		}
	}
	return list.m[m]
}

// MaxBy returns an element of AnyList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *AnyList) MaxBy(less func(interface{}, interface{}) bool) interface{} {
	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(list.m[m], list.m[i]) {
			m = i
		}
	}

	return list.m[m]
}

// DistinctBy returns a new AnyList whose elements are unique, where equality is defined by the equal function.
func (list *AnyList) DistinctBy(equal func(interface{}, interface{}) bool) *AnyList {
	if list == nil {
		return nil
	}

	result := newAnyList(0, len(list.m))
Outer:
	for _, v := range list.m {
		for _, r := range result.m {
			if equal(v, r) {
				continue Outer
			}
		}
		result.m = append(result.m, v)
	}
	return result
}

// IndexWhere finds the index of the first element satisfying predicate p. If none exists, -1 is returned.
func (list *AnyList) IndexWhere(p func(interface{}) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list *AnyList) IndexWhere2(p func(interface{}) bool, from int) int {

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list *AnyList) LastIndexWhere(p func(interface{}) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list *AnyList) LastIndexWhere2(p func(interface{}) bool, before int) int {

	if before < 0 {
		before = len(list.m)
	}
	for i := len(list.m) - 1; i >= 0; i-- {
		v := list.m[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}

//-------------------------------------------------------------------------------------------------
// These methods are included when interface{} is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items in the same order, they are considered equal.
// Order of items is not relevent for sets to be equal.
// Nil lists are considered to be empty.
func (list *AnyList) Equals(other *AnyList) bool {
	if list == nil {
		return other == nil || len(other.m) == 0
	}

	if other == nil {
		return len(list.m) == 0
	}

	if len(list.m) != len(other.m) {
		return false
	}

	for i, v := range list.m {
		if v != other.m[i] {
			return false
		}
	}

	return true
}

//-------------------------------------------------------------------------------------------------

type sortableAnyList struct {
	less func(i, j interface{}) bool
	m    []interface{}
}

func (sl sortableAnyList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableAnyList) Len() int {
	return len(sl.m)
}

func (sl sortableAnyList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy returns a new list in which the elements are sorted by a specified ordering.
func (list *AnyList) SortBy(less func(i, j interface{}) bool) *AnyList {
	if list == nil {
		return nil
	}

	result := NewAnyList(list.m...)
	sort.Sort(sortableAnyList{less, result.m})
	return result
}

// StableSortBy returns a new list in which the elements are sorted by a specified ordering.
// The algorithm keeps the original order of equal elements.
func (list *AnyList) StableSortBy(less func(i, j interface{}) bool) *AnyList {
	if list == nil {
		return nil
	}

	result := NewAnyList(list.m...)
	sort.Stable(sortableAnyList{less, result.m})
	return result
}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (list *AnyList) StringList() []string {

	strings := make([]string, len(list.m))
	for i, v := range list.m {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *AnyList) String() string {
	return list.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *AnyList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *AnyList) MkString3(before, between, after string) string {
	if list == nil {
		return ""
	}

	return list.mkString3Bytes(before, between, after).String()
}

func (list AnyList) mkString3Bytes(before, between, after string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""

	for _, v := range list.m {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

// UnmarshalJSON implements JSON decoding for this list type.
func (list *AnyList) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &list.m)
}

// MarshalJSON implements JSON encoding for this list type.
func (list AnyList) MarshalJSON() ([]byte, error) {
	buf, err := json.Marshal(list.m)
	return buf, err
}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this list type.
// You must register interface{} with the 'gob' package before this method is used.
func (list *AnyList) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&list.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register interface{} with the 'gob' package before this method is used.
func (list AnyList) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(list.m)
	return buf.Bytes(), err
}
