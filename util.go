package collection

// Zero zero
func Zero[T any]() T {
	return *new(T)
}

// IsZero is zero
func IsZero[T comparable](v T) bool {
	return v == *new(T)
}
