package shared

import (
	"sort"
)

// Sorted alters the list so that the elements are sorted by their natural ordering.
// Sorting happens in-place; the modified list is returned.
func (list *StringList) Sorted() *StringList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

	sort.Strings(list.m)
	return list
}
