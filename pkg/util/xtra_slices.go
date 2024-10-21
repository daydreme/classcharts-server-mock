package util

func Filter[T any](slice []T, test func(T) bool) (ret []T) {
	for _, item := range slice {
		if test(item) {
			ret = append(ret, item)
		}
	}
	return
}
