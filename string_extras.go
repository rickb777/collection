package collection

import (
	"sort"
)

// Sorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list StringList) Sorted() StringList {
	sort.Strings(list)
	return list
}
