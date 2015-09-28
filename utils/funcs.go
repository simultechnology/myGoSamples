package utils

func Sum(nums ...int) (result int) {

	for _, n := range nums {
		result += n
	}
	return
}
