package helper

func GetNotDefault[T comparable](val1 T, val2 T) T {
	var defaultT T
	if defaultT == val2 || val2 == val1 {
		return val1
	}
	return val2
}
