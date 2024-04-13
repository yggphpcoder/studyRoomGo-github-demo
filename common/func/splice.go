package _func

func InStrSplice(splice []string, str string) bool {
	for _, s := range splice {
		if s == str {
			return true
		}
	}
	return false
}

func InIntSplice(splice []int, i int) bool {
	for _, s := range splice {
		if s == i {
			return true
		}
	}
	return false
}
func InInt32Splice(splice []int32, i int32) bool {
	for _, s := range splice {
		if s == i {
			return true
		}
	}
	return false
}
func InInt8Splice(splice []int8, i int8) bool {
	for _, s := range splice {
		if s == i {
			return true
		}
	}
	return false
}
