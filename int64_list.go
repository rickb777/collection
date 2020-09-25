// A simple type derived from []int64
// Not thread-safe.
//
// Generated from simple/list.tpl with Type=int64
// options: Comparable:true Numeric:true Ordered:true StringLike:<no value> Stringer:true
// GobEncode:true Mutable:always ToList:always ToSet:true MapTo:string
// by runtemplate v3.6.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// Int64List is a slice of type int64. Use it where you would use []int64.
// To add items to the list, simply use the normal built-in append function.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type Int64List []int64

//-------------------------------------------------------------------------------------------------

// MakeInt64List makes an empty list with both length and capacity initialised.
func MakeInt64List(length, capacity int) Int64List {
	return make(Int64List, length, capacity)
}

// NewInt64List constructs a new list containing the supplied values, if any.
func NewInt64List(values ...int64) Int64List {
	list := MakeInt64List(len(values), len(values))
	copy(list, values)
	return list
}

// ConvertInt64List constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertInt64List(values ...interface{}) (Int64List, bool) {
	list := MakeInt64List(0, len(values))

	for _, i := range values {
		switch j := i.(type) {
		case int:
			k := int64(j)
			list = append(list, k)
		case *int:
			k := int64(*j)
			list = append(list, k)
		case int8:
			k := int64(j)
			list = append(list, k)
		case *int8:
			k := int64(*j)
			list = append(list, k)
		case int16:
			k := int64(j)
			list = append(list, k)
		case *int16:
			k := int64(*j)
			list = append(list, k)
		case int32:
			k := int64(j)
			list = append(list, k)
		case *int32:
			k := int64(*j)
			list = append(list, k)
		case int64:
			k := int64(j)
			list = append(list, k)
		case *int64:
			k := int64(*j)
			list = append(list, k)
		case uint:
			k := int64(j)
			list = append(list, k)
		case *uint:
			k := int64(*j)
			list = append(list, k)
		case uint8:
			k := int64(j)
			list = append(list, k)
		case *uint8:
			k := int64(*j)
			list = append(list, k)
		case uint16:
			k := int64(j)
			list = append(list, k)
		case *uint16:
			k := int64(*j)
			list = append(list, k)
		case uint32:
			k := int64(j)
			list = append(list, k)
		case *uint32:
			k := int64(*j)
			list = append(list, k)
		case uint64:
			k := int64(j)
			list = append(list, k)
		case *uint64:
			k := int64(*j)
			list = append(list, k)
		case float32:
			k := int64(j)
			list = append(list, k)
		case *float32:
			k := int64(*j)
			list = append(list, k)
		case float64:
			k := int64(j)
			list = append(list, k)
		case *float64:
			k := int64(*j)
			list = append(list, k)
		}
	}

	return list, len(list) == len(values)
}

// BuildInt64ListFromChan constructs a new Int64List from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildInt64ListFromChan(source <-chan int64) Int64List {
	list := MakeInt64List(0, 0)
	for v := range source {
		list = append(list, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list Int64List) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list Int64List) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list Int64List) slice() []int64 {
	return list
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list Int64List) ToList() Int64List {
	return list
}

// ToSet returns the elements of the list as a set. The returned set is a shallow
// copy; the list is not altered.
func (list Int64List) ToSet() Int64Set {
	if list == nil {
		return nil
	}

	return NewInt64Set(list...)
}

// ToSlice returns the elements of the list as a slice, which is an identity operation in this case,
// because the simple list is merely a dressed-up slice.
func (list Int64List) ToSlice() []int64 {
	return list
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list Int64List) ToInterfaceSlice() []interface{} {
	s := make([]interface{}, 0, len(list))
	for _, v := range list {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the list. It does not clone the underlying elements.
func (list Int64List) Clone() Int64List {
	return NewInt64List(list...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
// The simple list is a dressed-up slice and normal slice operations will also work.
func (list Int64List) Get(i int) int64 {
	return list[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list Int64List) Head() int64 {
	return list[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns the zero value.
func (list Int64List) HeadOption() int64 {
	if list.IsEmpty() {
		return 0
	}
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list Int64List) Last() int64 {
	return list[len(list)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns the zero value.
func (list Int64List) LastOption() int64 {
	if list.IsEmpty() {
		return 0
	}
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list Int64List) Tail() Int64List {
	return list[1:]
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list Int64List) Init() Int64List {
	return list[:len(list)-1]
}

// IsEmpty tests whether Int64List is empty.
func (list Int64List) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether Int64List is empty.
func (list Int64List) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list Int64List) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list Int64List) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list Int64List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

//-------------------------------------------------------------------------------------------------

// Contains determines whether a given item is already in the list, returning true if so.
func (list Int64List) Contains(v int64) bool {
	return list.Exists(func(x int64) bool {
		return x == v
	})
}

// ContainsAll determines whether the given items are all in the list, returning true if so.
// This is potentially a slow method and should only be used rarely.
func (list Int64List) ContainsAll(i ...int64) bool {
	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of Int64List return true for the predicate p.
func (list Int64List) Exists(p func(int64) bool) bool {
	for _, v := range list {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of Int64List return true for the predicate p.
func (list Int64List) Forall(p func(int64) bool) bool {
	for _, v := range list {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over Int64List and executes function f against each element.
func (list Int64List) Foreach(f func(int64)) {
	for _, v := range list {
		f(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements
// have been consumed. The channel will be closed when all the elements have been sent.
func (list Int64List) Send() <-chan int64 {
	ch := make(chan int64)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// Reverse returns a copy of Int64List with all elements in the reverse order.
//
// The original list is not modified.
func (list Int64List) Reverse() Int64List {
	n := len(list)
	result := MakeInt64List(n, n)
	last := n - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// DoReverse alters a Int64List with all elements in the reverse order.
// Unlike Reverse, it does not allocate new memory.
//
// The list is modified and the modified list is returned.
func (list Int64List) DoReverse() Int64List {
	mid := (len(list) + 1) / 2
	last := len(list) - 1
	for i := 0; i < mid; i++ {
		r := last - i
		if i != r {
			list[i], list[r] = list[r], list[i]
		}
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of Int64List, using a version of the Fisher-Yates shuffle.
//
// The original list is not modified.
func (list Int64List) Shuffle() Int64List {
	if list == nil {
		return nil
	}

	return list.Clone().DoShuffle()
}

// DoShuffle returns a shuffled Int64List, using a version of the Fisher-Yates shuffle.
//
// The list is modified and the modified list is returned.
func (list Int64List) DoShuffle() Int64List {
	if list == nil {
		return nil
	}

	n := len(list)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		list[i], list[r] = list[r], list[i]
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of Int64List containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list Int64List) Take(n int) Int64List {
	if n >= len(list) {
		return list
	}
	return list[0:n]
}

// Drop returns a slice of Int64List without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list Int64List) Drop(n int) Int64List {
	if n == 0 {
		return list
	}

	l := len(list)
	if n < l {
		return list[n:]
	}
	return list[l:]
}

// TakeLast returns a slice of Int64List containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
//
// The original list is not modified.
func (list Int64List) TakeLast(n int) Int64List {
	l := len(list)
	if n >= l {
		return list
	}
	return list[l-n:]
}

// DropLast returns a slice of Int64List without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list Int64List) DropLast(n int) Int64List {
	if n == 0 {
		return list
	}

	l := len(list)
	if n > l {
		return list[l:]
	}
	return list[0 : l-n]
}

// TakeWhile returns a new Int64List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
//
// The original list is not modified.
func (list Int64List) TakeWhile(p func(int64) bool) Int64List {
	result := MakeInt64List(0, 0)
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new Int64List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
//
// The original list is not modified.
func (list Int64List) DropWhile(p func(int64) bool) Int64List {
	result := MakeInt64List(0, 0)
	adding := false

	for _, v := range list {
		if adding || !p(v) {
			adding = true
			result = append(result, v)
		}
	}

	return result
}

//-------------------------------------------------------------------------------------------------

// Find returns the first int64 that returns true for predicate p.
// False is returned if none match.
func (list Int64List) Find(p func(int64) bool) (int64, bool) {

	for _, v := range list {
		if p(v) {
			return v, true
		}
	}

	var empty int64
	return empty, false
}

// Filter returns a new Int64List whose elements return true for predicate p.
//
// The original list is not modified.
func (list Int64List) Filter(p func(int64) bool) Int64List {
	result := MakeInt64List(0, len(list))

	for _, v := range list {
		if p(v) {
			result = append(result, v)
		}
	}

	return result
}

// Partition returns two new Int64Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified.
func (list Int64List) Partition(p func(int64) bool) (Int64List, Int64List) {
	matching := MakeInt64List(0, len(list))
	others := MakeInt64List(0, len(list))

	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}

	return matching, others
}

// Map returns a new Int64List by transforming every element with function f.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list Int64List) Map(f func(int64) int64) Int64List {
	result := MakeInt64List(0, len(list))

	for _, v := range list {
		result = append(result, f(v))
	}

	return result
}

// MapToString returns a new []string by transforming every element with function f.
// The resulting slice is the same size as the list.
// The list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list Int64List) MapToString(f func(int64) string) []string {
	if list == nil {
		return nil
	}

	result := make([]string, len(list))
	for i, v := range list {
		result[i] = f(v)
	}

	return result
}

// FlatMap returns a new Int64List by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list Int64List) FlatMap(f func(int64) []int64) Int64List {
	result := MakeInt64List(0, len(list))

	for _, v := range list {
		result = append(result, f(v)...)
	}

	return result
}

// FlatMapToString returns a new []string by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the list.
// The list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list Int64List) FlatMapToString(f func(int64) []string) []string {
	if list == nil {
		return nil
	}

	result := make([]string, 0, len(list))
	for _, v := range list {
		result = append(result, f(v)...)
	}

	return result
}

// CountBy gives the number elements of Int64List that return true for the predicate p.
func (list Int64List) CountBy(p func(int64) bool) (result int) {
	for _, v := range list {
		if p(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of Int64List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list Int64List) MinBy(less func(int64, int64) bool) int64 {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
	}

	m := 0
	for i := 1; i < l; i++ {
		if less(list[i], list[m]) {
			m = i
		}
	}

	return list[m]
}

// MaxBy returns an element of Int64List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list Int64List) MaxBy(less func(int64, int64) bool) int64 {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}

	m := 0
	for i := 1; i < l; i++ {
		if less(list[m], list[i]) {
			m = i
		}
	}

	return list[m]
}

// DistinctBy returns a new Int64List whose elements are unique, where equality is defined by the equal function.
func (list Int64List) DistinctBy(equal func(int64, int64) bool) Int64List {
	result := MakeInt64List(0, len(list))
Outer:
	for _, v := range list {
		for _, r := range result {
			if equal(v, r) {
				continue Outer
			}
		}
		result = append(result, v)
	}
	return result
}

// IndexWhere finds the index of the first element satisfying predicate p. If none exists, -1 is returned.
func (list Int64List) IndexWhere(p func(int64) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list Int64List) IndexWhere2(p func(int64) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list Int64List) LastIndexWhere(p func(int64) bool) int {
	return list.LastIndexWhere2(p, len(list))
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list Int64List) LastIndexWhere2(p func(int64) bool, before int) int {
	if before < 0 {
		before = len(list)
	}
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int64 is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items in the same order, they are considered equal.
// Order of items is not relevent for sets to be equal.
func (list Int64List) Equals(other Int64List) bool {
	if list == nil {
		return len(other) == 0
	}

	if other == nil {
		return len(list) == 0
	}

	if len(list) != len(other) {
		return false
	}

	for i, v := range list {
		if v != other[i] {
			return false
		}
	}

	return true
}

//-------------------------------------------------------------------------------------------------

type sortableInt64List struct {
	less func(i, j int64) bool
	m    []int64
}

func (sl sortableInt64List) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableInt64List) Len() int {
	return len(sl.m)
}

func (sl sortableInt64List) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list Int64List) SortBy(less func(i, j int64) bool) Int64List {
	sort.Sort(sortableInt64List{less, list})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list Int64List) StableSortBy(less func(i, j int64) bool) Int64List {
	sort.Stable(sortableInt64List{less, list})
	return list
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int64 is ordered.

// Sorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list Int64List) Sorted() Int64List {
	return list.SortBy(func(a, b int64) bool {
		return a < b
	})
}

// StableSorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list Int64List) StableSorted() Int64List {
	return list.StableSortBy(func(a, b int64) bool {
		return a < b
	})
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list Int64List) Min() int64 {
	m := list.MinBy(func(a int64, b int64) bool {
		return a < b
	})
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list Int64List) Max() (result int64) {
	m := list.MaxBy(func(a int64, b int64) bool {
		return a < b
	})
	return m
}

//-------------------------------------------------------------------------------------------------
// These methods are included when int64 is numeric.

// Sum returns the sum of all the elements in the list.
func (list Int64List) Sum() int64 {
	sum := int64(0)
	for _, v := range list {
		sum = sum + v
	}
	return sum
}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (list Int64List) StringList() []string {
	strings := make([]string, len(list))
	for i, v := range list {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list Int64List) String() string {
	return list.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list Int64List) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list Int64List) MkString3(before, between, after string) string {
	if list == nil {
		return ""
	}

	return list.mkString3Bytes(before, between, after).String()
}

func (list Int64List) mkString3Bytes(before, between, after string) *strings.Builder {
	b := &strings.Builder{}
	b.WriteString(before)
	sep := ""
	for _, v := range list {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		sep = between
	}
	b.WriteString(after)
	return b
}

//-------------------------------------------------------------------------------------------------

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this list type.
// You must register int64 with the 'gob' package before this method is used.
func (list Int64List) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&list)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register int64 with the 'gob' package before this method is used.
func (list Int64List) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(list)
	return buf.Bytes(), err
}
