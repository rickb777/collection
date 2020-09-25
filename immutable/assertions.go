package immutable

var _ IntCollection = &IntList{}
var _ IntCollection = &IntSet{}

var _ StringCollection = &StringList{}
var _ StringCollection = &StringSet{}
