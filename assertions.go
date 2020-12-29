package collection

var _ AnyCollection = &AnyList{}
var _ AnyCollection = &AnySet{}

var _ IntCollection = &IntList{}
var _ IntCollection = &IntSet{}

var _ Int64Collection = &Int64List{}
var _ Int64Collection = &Int64Set{}

var _ StringCollection = &StringList{}
var _ StringCollection = &StringSet{}

var _ UintCollection = &UintList{}
var _ UintCollection = &UintSet{}

var _ Uint64Collection = &Uint64List{}
var _ Uint64Collection = &Uint64Set{}
