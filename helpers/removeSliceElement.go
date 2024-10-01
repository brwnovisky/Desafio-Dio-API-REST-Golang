package helpers

import "reflect"

func findIndex[T any](slice []T, target T) int {
	for i, v := range slice {
		if reflect.DeepEqual(v, target) {
			return i
		}
	}
	return -1
}

func RemoveSliceElement[T any](slice []T, target T) []T {
	index := findIndex(slice, target)
	if index != -1 {
		return append(slice[:index], slice[index+1:]...)
	}
	return slice
}
