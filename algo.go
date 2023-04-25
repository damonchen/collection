package collection

import "context"

// Max max
func Max[T Number | String](iter Iterator[T]) T {
	var r T
	if iter.Next() {
		r = iter.Value()
	}

	for iter.Next() {
		v := iter.Value()
		if r < v {
			r = v
		}
	}
	return r
}

// Min min
func Min[T Number | String](iter Iterator[T]) T {
	var r T
	if iter.Next() {
		r = iter.Value()
	}

	for iter.Next() {
		v := iter.Value()
		if r > v {
			r = v
		}
	}
	return r
}

// Sum sum
func Sum[T Number | String](iter Iterator[T]) T {
	init := Zero[T]()
	return Reduce[T, T](iter, func(acc T, value T) T {
		return acc + value
	}, init)
}

// Close close
type Close func()

// Count count
func Count[T Number](start T, step T) (chan T, Close) {
	ch := make(chan T, 1)

	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		for i := start; i < start; i += step {
			select {
			case <-ctx.Done():
				return
			case ch <- i:

			}
		}
	}()

	return ch, Close(cancelFunc)
}

// Cycle cycle
func Cycle[T any](iter Iterator[T]) (chan T, Close) {
	ch := make(chan T, 1)

	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		for {
			for iter.Next() {
				select {
				case <-ctx.Done():
					return
				case ch <- iter.Value():
				}
			}
		}
	}()

	return ch, Close(cancelFunc)
}

// Repeat while times is -1, repeat times
func Repeat[T any](ele T, times int) (chan T, Close) {
	ch := make(chan T, 1)

	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		if times == -1 {
			for {
				select {
				case <-ctx.Done():
					return
				case ch <- ele:
				}
			}
		} else {
			for i := 0; i < times; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- ele:
				}
			}
		}

	}()

	return ch, Close(cancelFunc)
}

// Key key
type Key[T any, V comparable] func(T any) V

// GroupBy group by
func GroupBy[T any, V comparable](iter Iterator[T], keyFn Key[T, V]) map[V][]T {
	group := map[V][]T{}
	for iter.Next() {
		v := iter.Value()
		key := keyFn(v)

		group[key] = append(group[key], v)
	}
	return group
}

// Keys keys
func Keys[T any](object map[string]T) []string {
	var r []string
	for key := range object {
		r = append(r, key)
	}
	return r
}

// Values values
func Values[T any](object map[string]T) []T {
	var r []T
	for _, v := range object {
		r = append(r, v)
	}
	return r
}
