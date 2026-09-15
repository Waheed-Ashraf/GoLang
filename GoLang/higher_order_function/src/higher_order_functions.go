package src

func MapIntNums(num []int, f func(int) int) []int {
	result := make([]int, len(num))
	for index, number := range num {
		result[index] = f(number)
	}
	return result
}

func FilterSlice(num []int, predicate func(int) bool) []int {

	var result []int
	for _, number := range num {
		if predicate(number) {
			result = append(result, number)
		}
	}
	return result
}

func ReduceSlice(num []int, initial int, reducer func(int, int) int) int {
	var acc = initial
	for _, number := range num {
		acc = reducer(acc, number)
	}
	return acc

}
