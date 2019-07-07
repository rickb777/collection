// Generated from threadsafe/collection.tpl with Type=string
// options: Comparable:true Numeric:<no value> Ordered:<no value> Stringer:true Mutable:always
// by runtemplate v3.4.2
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package collection

// SharedStringSizer defines an interface for sizing methods on string collections.
type SharedStringSizer interface {
	// IsEmpty tests whether SharedStringCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether SharedStringCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// SharedStringMkStringer defines an interface for stringer methods on string collections.
type SharedStringMkStringer interface {
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

// SharedStringCollection defines an interface for common collection methods on string.
type SharedStringCollection interface {
	SharedStringSizer
	SharedStringMkStringer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToSet returns a shallow copy as a set.
	ToSet() *SharedStringSet

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []string

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of SharedStringCollection return true for the predicate p.
	Exists(p func(string) bool) bool

	// Forall verifies that all elements of SharedStringCollection return true for the predicate p.
	Forall(p func(string) bool) bool

	// Foreach iterates over SharedStringCollection and executes the function f against each element.
	Foreach(f func(string))

	// Find returns the first string that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(string) bool) (string, bool)

	// MapToInt returns a new []int by transforming every element with function f.
	// The resulting slice is the same size as the collection. The collection is not modified.
	MapToInt(f func(string) int) []int

	// FlatMapInt returns a new []int by transforming every element with function f
	// that returns zero or more items in a slice. The resulting list may have a different size to the
	// collection. The collection is not modified.
	FlatMapToInt(f func(string) []int) []int

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan string

	// CountBy gives the number elements of SharedStringCollection that return true for the predicate p.
	CountBy(p func(string) bool) int

	// Contains determines whether a given item is already in the collection, returning true if so.
	Contains(v string) bool

	// ContainsAll determines whether the given items are all in the collection, returning true if so.
	ContainsAll(v ...string) bool

	// Clear the entire collection.
	Clear()

	// Add adds items to the current collection.
	Add(more ...string)

	// MinBy returns an element of SharedStringCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(string, string) bool) string

	// MaxBy returns an element of SharedStringCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(string, string) bool) string
}