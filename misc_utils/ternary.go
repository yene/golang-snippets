package misc_utils

// ternary operator with Go
// cond ? a : b
func Ter[T any](cond bool, a T, b T) T {
	if cond {
		return a
	}
	return b
}
