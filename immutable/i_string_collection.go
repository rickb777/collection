// Generated from immutable/collection.tpl with Type=string
// options: Comparable:true Numeric:<no value> Ordered:<no value> Stringer:true Mutable:disabled
// by runtemplate v3.4.2
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

// IStringSizer defines an interface for sizing methods on string collections.
type IStringSizer interface {
	// IsEmpty tests whether IStringCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether IStringCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// IStringMkStringer defines an interface for stringer methods on string collections.
type IStringMkStringer interface {
	// String implements the Stringer interface to render the list as a comma-separated string enclosed
	// in square brackets.
	String() string

	// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
	MkString(sep string) string

	// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
	MkString3(before, between, after string) string

	// implements json.Marshaler interface {
	MarshalJSON() ([]byte, error)

	// StringList gets a list of strings that depicts all the elements.
	StringList() []string
}

// IStringCollection defines an interface for common collection methods on string.
type IStringCollection interface {
	IStringSizer
	IStringMkStringer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToSet returns a shallow copy as a set.
	ToSet() *IStringSet

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []string

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of IStringCollection return true for the predicate p.
	Exists(p func(string) bool) bool

	// Forall verifies that all elements of IStringCollection return true for the predicate p.
	Forall(p func(string) bool) bool

	// Foreach iterates over IStringCollection and executes the function f against each element.
	Foreach(f func(string))

	// Find returns the first string that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(string) bool) (string, bool)

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan string

	// CountBy gives the number elements of IStringCollection that return true for the predicate p.
	CountBy(p func(string) bool) int

	// Contains determines whether a given item is already in the collection, returning true if so.
	Contains(v string) bool

	// ContainsAll determines whether the given items are all in the collection, returning true if so.
	ContainsAll(v ...string) bool

	// MinBy returns an element of IStringCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(string, string) bool) string

	// MaxBy returns an element of IStringCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(string, string) bool) string
}
