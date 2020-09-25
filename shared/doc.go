// package shared provides shareable collection types for core Go built-in types.
// These are safe for sharing across multiple goroutines because they lock access
// using mutexes. However, take care because any sequence of multiple operations
// will not be atomic.
//
package shared
