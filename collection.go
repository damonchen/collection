package collection

type Integer interface {
	~int | ~uint | ~uint8 | ~int8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64
}

type Float interface {
	~float64 | ~float32
}

type Number interface {
	Integer | Float
}

type String interface {
	~string
}

// Array array for any element
type Array[V any] struct {
	elements []V
}

// Iterator iterator for array
type Iterator[T any] interface {
	Next() bool
	Value() T
}

// SliceIterator Iterator impl
type SliceIterator[T any] struct {
	Elements []T
	value    T
	index    int
}

// NewSliceIterator Create an iterator over the slice xs
func NewSliceIterator[T any](xs []T) Iterator[T] {
	return &SliceIterator[T]{
		Elements: xs,
	}
}

func (i *SliceIterator[T]) Next() bool {
	if i.index < len(i.Elements) {
		i.value = i.Elements[i.index]
		i.index += 1
		return true
	}
	return false
}

func (i *SliceIterator[T]) Value() T {
	return i.value
}

// Index slice index
func Index[T comparable](iter Iterator[T], val T) int {
  for i:=0; iter.Next(); i++ {
    v := iter.Value()
    if v == val {
      return i
    }
  }
  return -1
}

func Contain[T comparable](iter Iterator[T], val T) bool {
  return Index(iter, val) != -1
}


type mapIterator[T any, V any] struct {
	source Iterator[T]
	mapper func(T) V
}

// Next has next iterator
func (iter *mapIterator[T, V]) Next() bool {
	return iter.source.Next()
}

func (iter *mapIterator[T, V]) Value() V {
	value := iter.source.Value()
	return iter.mapper(value)
}

// Map from one iterator to another iterator
func Map[T any, V any](iter Iterator[T], f func(T) V) Iterator[V] {
	return &mapIterator[T, V]{
		iter, f,
	}
}

type filterIterator[T any] struct {
	source    Iterator[T]
	predicate func(T) bool
}

func (i *filterIterator[T]) Next() bool {
	for i.source.Next() {
		if i.predicate(i.source.Value()) {
			return true
		}
	}
	return false
}

func (i *filterIterator[T]) Value() T {
	return i.source.Value()
}

// Filter filter iterator
func Filter[T any](iter Iterator[T], predicate func(T) bool) Iterator[T] {
	return &filterIterator[T]{source: iter, predicate: predicate}
}

// Collect iterator to collection
func Collect[T any](iter Iterator[T]) []T {
	var xs []T

	for iter.Next() {
		xs = append(xs, iter.Value())
	}

	return xs
}

// Reducer reducer func define
type Reducer[T, V any] func(acc T, value V) T

// Reduce values iterated over to a single value
func Reduce[T, V any](iter Iterator[V], f Reducer[T, V], initAcc T) T {
	acc := initAcc
	for iter.Next() {
		acc = f(acc, iter.Value())
	}
	return acc
}

// Mapper iterator to map func define
type Mapper[T comparable, V any] func(v V) T

// ToMap to map
func ToMap[T comparable, V any](iter Iterator[V], f Mapper[T, V]) map[T]V {
	var r = map[T]V{}
	for iter.Next() {
		v := iter.Value()
		key := f(v)
		r[key] = v
	}
	return r
}


