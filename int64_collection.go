// Generated from simple/collection.tpl with Type=int64
// options: Comparable:true Numeric:true Ordered:true Stringer:true Mutable:always
// by runtemplate v3.4.2
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

// Int64Sizer defines an interface for sizing methods on int64 collections.
type Int64Sizer interface {
	// IsEmpty tests whether Int64Collection is empty.
	IsEmpty() bool

	// NonEmpty tests whether Int64Collection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// Int64MkStringer defines an interface for stringer methods on int64 collections.
type Int64MkStringer interface {
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

// Int64Collection defines an interface for common collection methods on int64.
type Int64Collection interface {
	Int64Sizer
	Int64MkStringer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToSet returns a shallow copy as a set.
	ToSet() Int64Set

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []int64

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of Int64Collection return true for the predicate p.
	Exists(p func(int64) bool) bool

	// Forall verifies that all elements of Int64Collection return true for the predicate p.
	Forall(p func(int64) bool) bool

	// Foreach iterates over Int64Collection and executes the function f against each element.
	Foreach(f func(int64))

	// Find returns the first int64 that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(int64) bool) (int64, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan int64

	// CountBy gives the number elements of Int64Collection that return true for the predicate p.
	CountBy(p func(int64) bool) int

	// Contains determines whether a given item is already in the collection, returning true if so.
	Contains(v int64) bool

	// ContainsAll determines whether the given items are all in the collection, returning true if so.
	ContainsAll(v ...int64) bool

	// Add adds items to the current collection.
	//Add(more ...int64)

	// Min returns the minimum value of all the items in the collection. Panics if there are no elements.
	Min() int64

	// Max returns the minimum value of all the items in the collection. Panics if there are no elements.
	Max() int64

	// MinBy returns an element of Int64Collection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(int64, int64) bool) int64

	// MaxBy returns an element of Int64Collection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(int64, int64) bool) int64

	// Sum returns the sum of all the elements in the collection.
	Sum() int64
}
