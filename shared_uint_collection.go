// Generated from threadsafe/collection.tpl with Type=uint
// options: Comparable:true Numeric:true Ordered:true Stringer:true Mutable:always
// by runtemplate v3.5.2
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

// SharedUintSizer defines an interface for sizing methods on uint collections.
type SharedUintSizer interface {
	// IsEmpty tests whether SharedUintCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether SharedUintCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// SharedUintMkStringer defines an interface for stringer methods on uint collections.
type SharedUintMkStringer interface {
	// String implements the Stringer interface to render the list as a comma-separated string enclosed
	// in square brackets.
	String() string

	// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
	MkString(sep string) string

	// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
	MkString3(before, between, after string) string

	// implements json.Marshaler interface {
	MarshalJSON() ([]byte, error)

	// implements json.Unmarshaler interface {
	UnmarshalJSON(b []byte) error

	// StringList gets a list of strings that depicts all the elements.
	StringList() []string
}

// SharedUintCollection defines an interface for common collection methods on uint.
type SharedUintCollection interface {
	SharedUintSizer
	SharedUintMkStringer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToSet returns a shallow copy as a set.
	ToSet() *SharedUintSet

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []uint

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of SharedUintCollection return true for the predicate p.
	Exists(p func(uint) bool) bool

	// Forall verifies that all elements of SharedUintCollection return true for the predicate p.
	Forall(p func(uint) bool) bool

	// Foreach iterates over SharedUintCollection and executes the function f against each element.
	Foreach(f func(uint))

	// Find returns the first uint that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(uint) bool) (uint, bool)

	// MapToString returns a new []string by transforming every element with function f.
	// The resulting slice is the same size as the collection. The collection is not modified.
	MapToString(f func(uint) string) []string

	// FlatMapString returns a new []string by transforming every element with function f
	// that returns zero or more items in a slice. The resulting list may have a different size to the
	// collection. The collection is not modified.
	FlatMapToString(f func(uint) []string) []string

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan uint

	// CountBy gives the number elements of SharedUintCollection that return true for the predicate p.
	CountBy(p func(uint) bool) int

	// Contains determines whether a given item is already in the collection, returning true if so.
	Contains(v uint) bool

	// ContainsAll determines whether the given items are all in the collection, returning true if so.
	ContainsAll(v ...uint) bool

	// Clear the entire collection.
	Clear()

	// Add adds items to the current collection.
	Add(more ...uint)

	// Min returns the minimum value of all the items in the collection. Panics if there are no elements.
	Min() uint

	// Max returns the minimum value of all the items in the collection. Panics if there are no elements.
	Max() uint

	// MinBy returns an element of SharedUintCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(uint, uint) bool) uint

	// MaxBy returns an element of SharedUintCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(uint, uint) bool) uint

	// Sum returns the sum of all the elements in the collection.
	Sum() uint
}
