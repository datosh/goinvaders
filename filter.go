package engine

import "reflect"

// Filter takes a list (slice) and a predicate (function that takes one element
// of the list, and returns true/false), and returns a filtered list.
// All items for which pred returns true are returned in the new list.
// Returned list has the same type as input list
func Filter(list interface{}, pred func(interface{}) bool) interface{} {
	contentType := reflect.TypeOf(list)
	contentValue := reflect.ValueOf(list)

	newContent := reflect.MakeSlice(contentType, 0, 0)
	for i := 0; i < contentValue.Len(); i++ {
		if content := contentValue.Index(i); pred(content.Interface()) {
			newContent = reflect.Append(newContent, content)
		}
	}
	return newContent.Interface()
}
