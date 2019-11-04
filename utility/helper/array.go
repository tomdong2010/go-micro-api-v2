package helper

func StrArrContains(arr []string, e string) bool {
	for _, str := range arr {
		if str == e {
			return true
		}
	}
	return false
}

func IntArrContains(arr []int, e int) bool {
	for _, n := range arr {
		if n == e {
			return true
		}
	}
	return false
}

// 字符串数组做差集
func StrArrDiff(arr []string, diff []string) []string{
	var result = make([]string, 0, len(arr))

	for _, a := range arr {
		if StrArrContains(diff, a) == false {
			result = append(result, a)
		}
	}

	return result
}

// 字符串数组去重
func StrArrUnique(arr []string) []string {
	var uniqueArr = make([]string, 0, len(arr))

	for _, v := range arr {
		if StrArrContains(uniqueArr, v) == false {
			uniqueArr = append(uniqueArr, v)
		}
	}

	return uniqueArr
}