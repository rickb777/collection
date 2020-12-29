// Generated from simple/collection.tpl with Type=interface{}
// options: Comparable:true Numeric:<no value> Ordered:<no value> Stringer:true Mutable:always
// by runtemplate v3.7.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

// AnySizer defines an interface for sizing methods on interface{} collections.
type AnySizer interface {
	// IsEmpty tests whether AnyCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether AnyCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the collection - an alias of Len().
	Size() int
}

// AnyMkStringer defines an interface for stringer methods on interface{} collections.
type AnyMkStringer interface {
	// String implements the Stringer interface to render the collection as a comma-separated string enclosed
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

	// StringList gets a collection of strings that depicts all the elements.
	StringList() []string
}

// AnyCollection defines an interface for common collection methods on interface{}.
type AnyCollection interface {
	AnySizer
	AnyMkStringer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []interface{}

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of AnyCollection return true for the predicate p.
	Exists(p func(interface{}) bool) bool

	// Forall verifies that all elements of AnyCollection return true for the predicate p.
	Forall(p func(interface{}) bool) bool

	// Foreach iterates over AnyCollection and executes the function f against each element.
	Foreach(f func(interface{}))

	// Find returns the first interface{} that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(interface{}) bool) (interface{}, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan interface{}

	// CountBy gives the number elements of AnyCollection that return true for the predicate p.
	CountBy(p func(interface{}) bool) int

	// Contains determines whether a given item is already in the collection, returning true if so.
	Contains(v interface{}) bool

	// ContainsAll determines whether the given items are all in the collection, returning true if so.
	ContainsAll(v ...interface{}) bool

	// Add adds items to the current collection.
	//Add(more ...interface{})

	// MinBy returns an element of AnyCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(interface{}, interface{}) bool) interface{}

	// MaxBy returns an element of AnyCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(interface{}, interface{}) bool) interface{}

	// Fold aggregates all the values in the collection using a supplied function, starting from some initial value.
	Fold(initial interface{}, fn func(interface{}, interface{}) interface{}) interface{}
}

// AnySequence defines an interface for sequence methods on interface{}.
type AnySequence interface {
	AnyCollection

	// Head gets the first element in the sequence. Head plus Tail include the whole sequence. Head is the opposite of Last.
	Head() interface{}

	// HeadOption gets the first element in the sequence, if possible.
	HeadOption() (interface{}, bool)

	// Last gets the last element in the sequence. Init plus Last include the whole sequence. Last is the opposite of Head.
	Last() interface{}

	// LastOption gets the last element in the sequence, if possible.
	LastOption() (interface{}, bool)
}
