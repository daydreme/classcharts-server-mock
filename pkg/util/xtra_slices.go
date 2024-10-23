package util

func Filter[T any](slice []T, test func(T) bool) (ret []T) {
	for _, item := range slice {
		if test(item) {
			ret = append(ret, item)
		}
	}
	return
}

func Map[T any](slice []T, run func(T) T) (ret []T) {
	for _, item := range slice {
		item = run(item)
		ret = append(ret, item)
	}
	return
}
