package util

// Filter takes a list (slice) and a predicate (function that takes one element
// of the list, and returns true/false), and returns a filtered list.
// All items for which pred returns true are returned in the new list.
// Returned list has the same type as input list
func Filter[T any](list []T, pred func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range list {
		if pred(item) {
			result = append(result, item)
		}
	}
	return result
}
