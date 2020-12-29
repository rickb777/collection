package shared

var _ AnyCollection = &AnyList{}
var _ AnyCollection = &AnySet{}
var _ AnyCollection = &AnyQueue{}

var _ IntCollection = &IntList{}
var _ IntCollection = &IntSet{}
var _ IntCollection = &IntQueue{}

var _ Int64Collection = &Int64List{}
var _ Int64Collection = &Int64Set{}
var _ Int64Collection = &Int64Queue{}

var _ StringCollection = &StringList{}
var _ StringCollection = &StringSet{}
var _ StringCollection = &StringQueue{}

var _ UintCollection = &UintList{}
var _ UintCollection = &UintSet{}
var _ UintCollection = &UintQueue{}

var _ Uint64Collection = &Uint64List{}
var _ Uint64Collection = &Uint64Set{}
var _ Uint64Collection = &Uint64Queue{}
