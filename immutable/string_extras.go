package immutable

import "sort"

// Sorted returns a new copy of the list in which the elements are sorted by their natural ordering.
func (list *StringList) Sorted() *StringList {
	if list == nil {
		return nil
	}

	result := NewStringList(list.m...)
	sort.Strings(result.m)
	return result
}
