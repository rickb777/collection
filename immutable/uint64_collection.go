// Generated from immutable/collection.tpl with Type=uint64
// options: Comparable:true Numeric:<no value> Integer:true Ordered:true Stringer:true Mutable:disabled
// by runtemplate v3.10.0
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package immutable

// Uint64Sizer defines an interface for sizing methods on uint64 collections.
type Uint64Sizer interface {
	// IsEmpty tests whether Uint64Collection is empty.
	IsEmpty() bool

	// NonEmpty tests whether Uint64Collection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// Uint64MkStringer defines an interface for stringer methods on uint64 collections.
type Uint64MkStringer interface {
	// String implements the Stringer interface to render the list as a comma-separated string enclosed
	// in square brackets.
	String() string

	// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
	MkString(sep string) string

	// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
	MkString3(before, between, after string) string

	// implements json.Marshaler interface {
	MarshalJSON() ([]byte, error)

	// StringList gets a slice of strings that depicts all the elements.
	StringList() []string
}

// Uint64Collection defines an interface for common collection methods on uint64.
type Uint64Collection interface {
	Uint64Sizer
	Uint64MkStringer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToSlice returns a shallow copy as a plain slice.
	ToSlice() []uint64

	// ToInterfaceSlice returns a shallow copy as a slice of arbitrary type.
	ToInterfaceSlice() []interface{}

	// Exists verifies that one or more elements of Uint64Collection return true for the predicate p.
	Exists(p func(uint64) bool) bool

	// Forall verifies that all elements of Uint64Collection return true for the predicate p.
	Forall(p func(uint64) bool) bool

	// Foreach iterates over Uint64Collection and executes the function f against each element.
	Foreach(f func(uint64))

	// Find returns the first uint64 that returns true for the predicate p.
	// False is returned if none match.
	Find(p func(uint64) bool) (uint64, bool)

	// MapToString returns a new []string by transforming every element with function f.
	// The resulting slice is the same size as the collection. The collection is not modified.
	MapToString(f func(uint64) string) []string

	// FlatMapString returns a new []string by transforming every element with function f
	// that returns zero or more items in a slice. The resulting slice may have a different size to the
	// collection. The collection is not modified.
	FlatMapToString(f func(uint64) []string) []string

	// Send returns a channel that will send all the elements in order. Can be used with the plumbing code, for example.
	// A goroutine is created to send the elements; this only terminates when all the elements have been consumed
	Send() <-chan uint64

	// CountBy gives the number elements of Uint64Collection that return true for the predicate p.
	CountBy(p func(uint64) bool) int

	// Contains determines whether a given item is already in the collection, returning true if so.
	Contains(v uint64) bool

	// ContainsAll determines whether the given items are all in the collection, returning true if so.
	ContainsAll(v ...uint64) bool

	// Min returns the minimum value of all the items in the collection. Panics if there are no elements.
	Min() uint64

	// Max returns the minimum value of all the items in the collection. Panics if there are no elements.
	Max() uint64

	// MinBy returns an element of Uint64Collection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(uint64, uint64) bool) uint64

	// MaxBy returns an element of Uint64Collection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(uint64, uint64) bool) uint64

	// Fold aggregates all the values in the collection using a supplied function, starting from some initial value.
	Fold(initial uint64, fn func(uint64, uint64) uint64) uint64

	// Sum returns the sum of all the elements in the collection.
	Sum() uint64
}

// Uint64Sequence defines an interface for sequence methods on uint64.
type Uint64Sequence interface {
	Uint64Collection

	// Head gets the first element in the sequence. Head plus Tail include the whole sequence. Head is the opposite of Last.
	Head() uint64

	// HeadOption gets the first element in the sequence, if possible.
	HeadOption() (uint64, bool)

	// Last gets the last element in the sequence. Init plus Last include the whole sequence. Last is the opposite of Head.
	Last() uint64

	// LastOption gets the last element in the sequence, if possible.
	LastOption() (uint64, bool)
}
