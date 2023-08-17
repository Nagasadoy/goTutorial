package mathSlice

func SumSlice(s []int) int {
	sum := 0

	for _, el := range s {
		sum += el
	}

	return sum
}

func MapSlice(s []int, op func(int) int) {
	for i, v := range s {
		s[i] = op(v)
	}
}

func FoldSlice(
	s []int,
	op func(int, int) int,
	init int,
) int {
	result := init
	for _, v := range s {
		result = op(result, v)
	}

	return result
}
