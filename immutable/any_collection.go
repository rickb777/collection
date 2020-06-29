// Generated from immutable/collection.tpl with Type=interface{}
// options: Comparable:<no value> Numeric:<no value> Ordered:<no value> Stringer:<no value> Mutable:disabled
// by runtemplate v3.5.4
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package immutable

// AnySizer defines an interface for sizing methods on interface{} collections.
type AnySizer interface {
	// IsEmpty tests whether AnyCollection is empty.
	IsEmpty() bool

	// NonEmpty tests whether AnyCollection is empty.
	NonEmpty() bool

	// Size returns the number of items in the list - an alias of Len().
	Size() int
}

// AnyCollection defines an interface for common collection methods on interface{}.
type AnyCollection interface {
	AnySizer

	// IsSequence returns true for lists and queues.
	IsSequence() bool

	// IsSet returns false for lists and queues.
	IsSet() bool

	// ToSet returns a shallow copy as a set.
	ToSet() *AnySet

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

	// MinBy returns an element of AnyCollection containing the minimum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
	// element is returned. Panics if there are no elements.
	MinBy(less func(interface{}, interface{}) bool) interface{}

	// MaxBy returns an element of AnyCollection containing the maximum value, when compared to other elements
	// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
	// element is returned. Panics if there are no elements.
	MaxBy(less func(interface{}, interface{}) bool) interface{}
}
