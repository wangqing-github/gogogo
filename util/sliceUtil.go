package util

func SliceContains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if s == v {
			return true
		}
	}
	return false
}
