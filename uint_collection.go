// Generated from simple/collection.tpl with Type=uint
// options: Comparable:true Numeric:true Ordered:true Stringer:true Mutable:always
// by runtemplate v3.7.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

// UintSizer defines an interface for sizing methods on uint collections.
type UintSizer interface {
	// IsEmpty tests whether UintCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether UintCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// UintMkStringer defines an interface for stringer methods on uint collections.
type UintMkStringer interface {
	// String implements the Stringer interface to render the list as a comma-separated string enclosed
	// in square brackets.
	String() string

	// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
	MkString(sep string) string

	// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
	MkString3(before, between, after string) string

	// implements json.Marshaler interface {
	//MarshalJSON() ([]byte, error)

	// implements json.Unmarshaler interface {
	//UnmarshalJSON(b []byte) error

	// StringList gets a list of strings that depicts all the elements.
	StringList() []string
}

// UintCollection defines an interface for common collection methods on uint.
type UintCollection interface {
	UintSizer
	UintMkStringer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToSet returns a shallow copy as a set.
	ToSet() UintSet

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []uint

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of UintCollection return true for the predicate p.
	Exists(p func(uint) bool) bool

	// Forall verifies that all elements of UintCollection return true for the predicate p.
	Forall(p func(uint) bool) bool

	// Foreach iterates over UintCollection and executes the function f against each element.
	Foreach(f func(uint))

	// Find returns the first uint that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(uint) bool) (uint, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan uint

	// CountBy gives the number elements of UintCollection that return true for the predicate p.
	CountBy(p func(uint) bool) int

	// Contains determines whether a given item is already in the collection, returning true if so.
	Contains(v uint) bool

	// ContainsAll determines whether the given items are all in the collection, returning true if so.
	ContainsAll(v ...uint) bool

	// Add adds items to the current collection.
	//Add(more ...uint)

	// Min returns the minimum value of all the items in the collection. Panics if there are no elements.
	Min() uint

	// Max returns the minimum value of all the items in the collection. Panics if there are no elements.
	Max() uint

	// MinBy returns an element of UintCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(uint, uint) bool) uint

	// MaxBy returns an element of UintCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(uint, uint) bool) uint

	// Sum returns the sum of all the elements in the collection.
	Sum() uint
}
