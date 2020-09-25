// An encapsulated []uint64.
// Thread-safe.
//
// Generated from threadsafe/list.tpl with Type=uint64
// options: Comparable:true Numeric:true Ordered:true StringLike:<no value> Stringer:true
// GobEncode:true Mutable:always ToList:always ToSet:true MapTo:string
// by runtemplate v3.6.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package shared

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

// Uint64List contains a slice of type uint64.
// It encapsulates the slice and provides methods to access or mutate it.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type Uint64List struct {
	s *sync.RWMutex
	m []uint64
}

//-------------------------------------------------------------------------------------------------

// MakeUint64List makes an empty list with both length and capacity initialised.
func MakeUint64List(length, capacity int) *Uint64List {
	return &Uint64List{
		s: &sync.RWMutex{},
		m: make([]uint64, length, capacity),
	}
}

// NewUint64List constructs a new list containing the supplied values, if any.
func NewUint64List(values ...uint64) *Uint64List {
	list := MakeUint64List(len(values), len(values))
	copy(list.m, values)
	return list
}

// ConvertUint64List constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertUint64List(values ...interface{}) (*Uint64List, bool) {
	list := MakeUint64List(0, len(values))

	for _, i := range values {
		switch j := i.(type) {
		case int:
			k := uint64(j)
			list.m = append(list.m, k)
		case *int:
			k := uint64(*j)
			list.m = append(list.m, k)
		case int8:
			k := uint64(j)
			list.m = append(list.m, k)
		case *int8:
			k := uint64(*j)
			list.m = append(list.m, k)
		case int16:
			k := uint64(j)
			list.m = append(list.m, k)
		case *int16:
			k := uint64(*j)
			list.m = append(list.m, k)
		case int32:
			k := uint64(j)
			list.m = append(list.m, k)
		case *int32:
			k := uint64(*j)
			list.m = append(list.m, k)
		case int64:
			k := uint64(j)
			list.m = append(list.m, k)
		case *int64:
			k := uint64(*j)
			list.m = append(list.m, k)
		case uint:
			k := uint64(j)
			list.m = append(list.m, k)
		case *uint:
			k := uint64(*j)
			list.m = append(list.m, k)
		case uint8:
			k := uint64(j)
			list.m = append(list.m, k)
		case *uint8:
			k := uint64(*j)
			list.m = append(list.m, k)
		case uint16:
			k := uint64(j)
			list.m = append(list.m, k)
		case *uint16:
			k := uint64(*j)
			list.m = append(list.m, k)
		case uint32:
			k := uint64(j)
			list.m = append(list.m, k)
		case *uint32:
			k := uint64(*j)
			list.m = append(list.m, k)
		case uint64:
			k := uint64(j)
			list.m = append(list.m, k)
		case *uint64:
			k := uint64(*j)
			list.m = append(list.m, k)
		case float32:
			k := uint64(j)
			list.m = append(list.m, k)
		case *float32:
			k := uint64(*j)
			list.m = append(list.m, k)
		case float64:
			k := uint64(j)
			list.m = append(list.m, k)
		case *float64:
			k := uint64(*j)
			list.m = append(list.m, k)
		}
	}

	return list, len(list.m) == len(values)
}

// BuildUint64ListFromChan constructs a new Uint64List from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildUint64ListFromChan(source <-chan uint64) *Uint64List {
	list := MakeUint64List(0, 0)
	for v := range source {
		list.m = append(list.m, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list *Uint64List) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list *Uint64List) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list *Uint64List) slice() []uint64 {
	if list == nil {
		return nil
	}
	return list.m
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list *Uint64List) ToList() *Uint64List {
	return list
}

// ToSet returns the elements of the list as a set. The returned set is a shallow
// copy; the list is not altered.
func (list *Uint64List) ToSet() *Uint64Set {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	return NewUint64Set(list.m...)
}

// ToSlice returns the elements of the current list as a slice.
func (list *Uint64List) ToSlice() []uint64 {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	s := make([]uint64, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *Uint64List) ToInterfaceSlice() []interface{} {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	s := make([]interface{}, 0, len(list.m))
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the list. It does not clone the underlying elements.
func (list *Uint64List) Clone() *Uint64List {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	return NewUint64List(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
func (list *Uint64List) Get(i int) uint64 {
	list.s.RLock()
	defer list.s.RUnlock()

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list *Uint64List) Head() uint64 {
	list.s.RLock()
	defer list.s.RUnlock()

	return list.m[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns the zero value.
func (list *Uint64List) HeadOption() uint64 {
	if list == nil {
		return 0
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if len(list.m) == 0 {
		return 0
	}
	return list.m[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty or nil.
func (list *Uint64List) Last() uint64 {
	list.s.RLock()
	defer list.s.RUnlock()

	return list.m[len(list.m)-1]
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns the zero value.
func (list *Uint64List) LastOption() uint64 {
	if list == nil {
		return 0
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if len(list.m) == 0 {
		return 0
	}
	return list.m[len(list.m)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list *Uint64List) Tail() *Uint64List {
	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeUint64List(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list *Uint64List) Init() *Uint64List {
	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeUint64List(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether Uint64List is empty.
func (list *Uint64List) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether Uint64List is empty.
func (list *Uint64List) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *Uint64List) Size() int {
	if list == nil {
		return 0
	}

	list.s.RLock()
	defer list.s.RUnlock()

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list *Uint64List) Len() int {
	return list.Size()
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list *Uint64List) Swap(i, j int) {
	list.s.Lock()
	defer list.s.Unlock()

	list.m[i], list.m[j] = list.m[j], list.m[i]
}

//-------------------------------------------------------------------------------------------------

// Contains determines whether a given item is already in the list, returning true if so.
func (list *Uint64List) Contains(v uint64) bool {
	return list.Exists(func(x uint64) bool {
		return v == x
	})
}

// ContainsAll determines whether the given items are all in the list, returning true if so.
// This is potentially a slow method and should only be used rarely.
func (list *Uint64List) ContainsAll(i ...uint64) bool {
	if list == nil {
		return len(i) == 0
	}

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range i {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

// Exists verifies that one or more elements of Uint64List return true for the predicate p.
func (list *Uint64List) Exists(p func(uint64) bool) bool {
	if list == nil {
		return false
	}

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of Uint64List return true for the predicate p.
func (list *Uint64List) Forall(p func(uint64) bool) bool {
	if list == nil {
		return true
	}

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over Uint64List and executes function f against each element.
// The function can safely alter the values via side-effects.
func (list *Uint64List) Foreach(f func(uint64)) {
	if list == nil {
		return
	}

	list.s.Lock()
	defer list.s.Unlock()

	for _, v := range list.m {
		f(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements
// have been consumed. The channel will be closed when all the elements have been sent.
func (list *Uint64List) Send() <-chan uint64 {
	ch := make(chan uint64)
	go func() {
		if list != nil {
			list.s.RLock()
			defer list.s.RUnlock()

			for _, v := range list.m {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// Reverse returns a copy of Uint64List with all elements in the reverse order.
//
// The original list is not modified.
func (list *Uint64List) Reverse() *Uint64List {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

	n := len(list.m)
	result := MakeUint64List(n, n)
	last := n - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// DoReverse alters a Uint64List with all elements in the reverse order.
// Unlike Reverse, it does not allocate new memory.
//
// The list is modified and the modified list is returned.
func (list *Uint64List) DoReverse() *Uint64List {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

	mid := (len(list.m) + 1) / 2
	last := len(list.m) - 1
	for i := 0; i < mid; i++ {
		r := last - i
		if i != r {
			list.m[i], list.m[r] = list.m[r], list.m[i]
		}
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of Uint64List, using a version of the Fisher-Yates shuffle.
//
// The original list is not modified.
func (list *Uint64List) Shuffle() *Uint64List {
	if list == nil {
		return nil
	}

	return list.Clone().doShuffle()
}

// DoShuffle returns a shuffled Uint64List, using a version of the Fisher-Yates shuffle.
//
// The list is modified and the modified list is returned.
func (list *Uint64List) DoShuffle() *Uint64List {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doShuffle()
}

func (list *Uint64List) doShuffle() *Uint64List {
	n := len(list.m)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		list.m[i], list.m[r] = list.m[r], list.m[i]
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Clear the entire collection.
func (list *Uint64List) Clear() {
	if list != nil {
		list.s.Lock()
		defer list.s.Unlock()
		list.m = list.m[:]
	}
}

// Add adds items to the current list. This is a synonym for Append.
func (list *Uint64List) Add(more ...uint64) {
	list.Append(more...)
}

// Append adds items to the current list.
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
func (list *Uint64List) Append(more ...uint64) *Uint64List {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = MakeUint64List(0, len(more))
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doAppend(more...)
}

func (list *Uint64List) doAppend(more ...uint64) *Uint64List {
	list.m = append(list.m, more...)
	return list
}

// DoInsertAt modifies a Uint64List by inserting elements at a given index.
// This is a generalised version of Append.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if the index is out of range.
func (list *Uint64List) DoInsertAt(index int, more ...uint64) *Uint64List {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = MakeUint64List(0, len(more))
		return list.doInsertAt(index, more...)
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doInsertAt(index, more...)
}

func (list *Uint64List) doInsertAt(index int, more ...uint64) *Uint64List {
	if len(more) == 0 {
		return list
	}

	if index == len(list.m) {
		// appending is an easy special case
		return list.doAppend(more...)
	}

	newlist := make([]uint64, 0, len(list.m)+len(more))

	if index != 0 {
		newlist = append(newlist, list.m[:index]...)
	}

	newlist = append(newlist, more...)

	newlist = append(newlist, list.m[index:]...)

	list.m = newlist
	return list
}

//-------------------------------------------------------------------------------------------------

// DoDeleteFirst modifies a Uint64List by deleting n elements from the start of
// the list.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *Uint64List) DoDeleteFirst(n int) *Uint64List {
	list.s.Lock()
	defer list.s.Unlock()
	return list.doDeleteAt(0, n)
}

// DoDeleteLast modifies a Uint64List by deleting n elements from the end of
// the list.
//
// The list is modified and the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *Uint64List) DoDeleteLast(n int) *Uint64List {
	list.s.Lock()
	defer list.s.Unlock()
	return list.doDeleteAt(len(list.m)-n, n)
}

// DoDeleteAt modifies a Uint64List by deleting n elements from a given index.
//
// The list is modified and the modified list is returned.
// Panics if the index is out of range or n is large enough to take the index out of range.
func (list *Uint64List) DoDeleteAt(index, n int) *Uint64List {
	list.s.Lock()
	defer list.s.Unlock()
	return list.doDeleteAt(index, n)
}

func (list *Uint64List) doDeleteAt(index, n int) *Uint64List {
	if n == 0 {
		return list
	}

	newlist := make([]uint64, 0, len(list.m)-n)

	if index != 0 {
		newlist = append(newlist, list.m[:index]...)
	}

	index += n

	if index != len(list.m) {
		newlist = append(newlist, list.m[index:]...)
	}

	list.m = newlist
	return list
}

//-------------------------------------------------------------------------------------------------

// DoKeepWhere modifies a Uint64List by retaining only those elements that match
// the predicate p. This is very similar to Filter but alters the list in place.
//
// The list is modified and the modified list is returned.
func (list *Uint64List) DoKeepWhere(p func(uint64) bool) *Uint64List {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doKeepWhere(p)
}

func (list *Uint64List) doKeepWhere(p func(uint64) bool) *Uint64List {
	result := make([]uint64, 0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result = append(result, v)
		}
	}

	list.m = result
	return list
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of Uint64List containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *Uint64List) Take(n int) *Uint64List {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if n >= len(list.m) {
		return list
	}

	result := MakeUint64List(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of Uint64List without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *Uint64List) Drop(n int) *Uint64List {
	if list == nil || n == 0 {
		return list
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if n >= len(list.m) {
		return nil
	}

	result := MakeUint64List(0, 0)
	result.m = list.m[n:]
	return result
}

// TakeLast returns a slice of Uint64List containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
//
// The original list is not modified.
func (list *Uint64List) TakeLast(n int) *Uint64List {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	l := len(list.m)
	if n >= l {
		return list
	}

	result := MakeUint64List(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of Uint64List without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *Uint64List) DropLast(n int) *Uint64List {
	if list == nil || n == 0 {
		return list
	}

	list.s.RLock()
	defer list.s.RUnlock()

	l := len(list.m)
	if n >= l {
		return nil
	}

	result := MakeUint64List(0, 0)
	result.m = list.m[:l-n]
	return result
}

// TakeWhile returns a new Uint64List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
//
// The original list is not modified.
func (list *Uint64List) TakeWhile(p func(uint64) bool) *Uint64List {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeUint64List(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new Uint64List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
//
// The original list is not modified.
func (list *Uint64List) DropWhile(p func(uint64) bool) *Uint64List {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeUint64List(0, 0)
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

// Find returns the first uint64 that returns true for predicate p.
// False is returned if none match.
func (list *Uint64List) Find(p func(uint64) bool) (uint64, bool) {
	if list == nil {
		return 0, false
	}

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if p(v) {
			return v, true
		}
	}

	var empty uint64
	return empty, false
}

// Filter returns a new Uint64List whose elements return true for predicate p.
//
// The original list is not modified. See also DoKeepWhere (which does modify the original list).
func (list *Uint64List) Filter(p func(uint64) bool) *Uint64List {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeUint64List(0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new Uint64Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified
func (list *Uint64List) Partition(p func(uint64) bool) (*Uint64List, *Uint64List) {
	if list == nil {
		return nil, nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	matching := MakeUint64List(0, len(list.m))
	others := MakeUint64List(0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new Uint64List by transforming every element with function f.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *Uint64List) Map(f func(uint64) uint64) *Uint64List {
	if list == nil {
		return nil
	}

	result := MakeUint64List(len(list.m), len(list.m))
	list.s.RLock()
	defer list.s.RUnlock()

	for i, v := range list.m {
		result.m[i] = f(v)
	}

	return result
}

// MapToString returns a new []string by transforming every element with function f.
// The resulting slice is the same size as the list.
// The list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *Uint64List) MapToString(f func(uint64) string) []string {
	if list == nil {
		return nil
	}

	result := make([]string, len(list.m))
	list.s.RLock()
	defer list.s.RUnlock()

	for i, v := range list.m {
		result[i] = f(v)
	}

	return result
}

// FlatMap returns a new Uint64List by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *Uint64List) FlatMap(f func(uint64) []uint64) *Uint64List {
	if list == nil {
		return nil
	}

	result := MakeUint64List(0, len(list.m))
	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		result.m = append(result.m, f(v)...)
	}

	return result
}

// FlatMapToString returns a new []string by transforming every element with function f that
// returns zero or more items in a slice. The resulting slice may have a different size to the list.
// The list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *Uint64List) FlatMapToString(f func(uint64) []string) []string {
	if list == nil {
		return nil
	}

	result := make([]string, 0, len(list.m))
	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		result = append(result, f(v)...)
	}

	return result
}

// CountBy gives the number elements of Uint64List that return true for the predicate p.
func (list *Uint64List) CountBy(p func(uint64) bool) (result int) {
	if list == nil {
		return 0
	}

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if p(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of Uint64List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *Uint64List) MinBy(less func(uint64, uint64) bool) uint64 {
	list.s.RLock()
	defer list.s.RUnlock()

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

// MaxBy returns an element of Uint64List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *Uint64List) MaxBy(less func(uint64, uint64) bool) uint64 {
	list.s.RLock()
	defer list.s.RUnlock()

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

// DistinctBy returns a new Uint64List whose elements are unique, where equality is defined by the equal function.
func (list *Uint64List) DistinctBy(equal func(uint64, uint64) bool) *Uint64List {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeUint64List(0, len(list.m))
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
func (list *Uint64List) IndexWhere(p func(uint64) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list *Uint64List) IndexWhere2(p func(uint64) bool, from int) int {
	list.s.RLock()
	defer list.s.RUnlock()

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list *Uint64List) LastIndexWhere(p func(uint64) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list *Uint64List) LastIndexWhere2(p func(uint64) bool, before int) int {
	list.s.RLock()
	defer list.s.RUnlock()

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
// These methods are included when uint64 is comparable.

// Equals determines if two lists are equal to each other.
// If they both are the same size and have the same items in the same order, they are considered equal.
// Order of items is not relevent for sets to be equal.
// Nil lists are considered to be empty.
func (list *Uint64List) Equals(other *Uint64List) bool {
	if list == nil {
		if other == nil {
			return true
		}
		other.s.RLock()
		defer other.s.RUnlock()
		return len(other.m) == 0
	}

	if other == nil {
		list.s.RLock()
		defer list.s.RUnlock()
		return len(list.m) == 0
	}

	list.s.RLock()
	other.s.RLock()
	defer list.s.RUnlock()
	defer other.s.RUnlock()

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

type sortableUint64List struct {
	less func(i, j uint64) bool
	m    []uint64
}

func (sl sortableUint64List) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableUint64List) Len() int {
	return len(sl.m)
}

func (sl sortableUint64List) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list *Uint64List) SortBy(less func(i, j uint64) bool) *Uint64List {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

	sort.Sort(sortableUint64List{less, list.m})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list *Uint64List) StableSortBy(less func(i, j uint64) bool) *Uint64List {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

	sort.Stable(sortableUint64List{less, list.m})
	return list
}

//-------------------------------------------------------------------------------------------------
// These methods are included when uint64 is ordered.

// Sorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list *Uint64List) Sorted() *Uint64List {
	return list.SortBy(func(a, b uint64) bool {
		return a < b
	})
}

// StableSorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list *Uint64List) StableSorted() *Uint64List {
	return list.StableSortBy(func(a, b uint64) bool {
		return a < b
	})
}

// Min returns the first element containing the minimum value, when compared to other elements.
// Panics if the collection is empty.
func (list *Uint64List) Min() uint64 {
	list.s.RLock()
	defer list.s.RUnlock()

	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
	}

	v := list.m[0]
	m := v
	for i := 1; i < l; i++ {
		v := list.m[i]
		if v < m {
			m = v
		}
	}
	return m
}

// Max returns the first element containing the maximum value, when compared to other elements.
// Panics if the collection is empty.
func (list *Uint64List) Max() (result uint64) {
	list.s.RLock()
	defer list.s.RUnlock()

	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}

	v := list.m[0]
	m := v
	for i := 1; i < l; i++ {
		v := list.m[i]
		if v > m {
			m = v
		}
	}
	return m
}

//-------------------------------------------------------------------------------------------------
// These methods are included when uint64 is numeric.

// Sum returns the sum of all the elements in the list.
func (list *Uint64List) Sum() uint64 {
	list.s.RLock()
	defer list.s.RUnlock()

	sum := uint64(0)
	for _, v := range list.m {
		sum = sum + v
	}
	return sum
}

//-------------------------------------------------------------------------------------------------

// StringList gets a list of strings that depicts all the elements.
func (list *Uint64List) StringList() []string {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	strings := make([]string, len(list.m))
	for i, v := range list.m {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list *Uint64List) String() string {
	return list.MkString3("[", ", ", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list *Uint64List) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list *Uint64List) MkString3(before, between, after string) string {
	if list == nil {
		return ""
	}

	return list.mkString3Bytes(before, between, after).String()
}

func (list Uint64List) mkString3Bytes(before, between, after string) *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(before)
	sep := ""

	list.s.RLock()
	defer list.s.RUnlock()

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
func (list *Uint64List) UnmarshalJSON(b []byte) error {
	list.s.Lock()
	defer list.s.Unlock()

	return json.Unmarshal(b, &list.m)
}

// MarshalJSON implements JSON encoding for this list type.
func (list Uint64List) MarshalJSON() ([]byte, error) {
	list.s.RLock()
	defer list.s.RUnlock()

	buf, err := json.Marshal(list.m)
	return buf, err
}

//-------------------------------------------------------------------------------------------------

// GobDecode implements 'gob' decoding for this list type.
// You must register uint64 with the 'gob' package before this method is used.
func (list *Uint64List) GobDecode(b []byte) error {
	list.s.Lock()
	defer list.s.Unlock()

	buf := bytes.NewBuffer(b)
	return gob.NewDecoder(buf).Decode(&list.m)
}

// GobEncode implements 'gob' encoding for this list type.
// You must register uint64 with the 'gob' package before this method is used.
func (list Uint64List) GobEncode() ([]byte, error) {
	list.s.RLock()
	defer list.s.RUnlock()

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(list.m)
	return buf.Bytes(), err
}
