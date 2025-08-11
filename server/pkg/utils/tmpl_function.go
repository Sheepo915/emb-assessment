package utils

func Seq(start, end int) []int {
	s := make([]int, end-start+1)
	for i := range s {
		s[i] = start + i
	}
	return s
}

func Sub(left, right int) int {
	return left - right
}

func Add(left, right int) int {
	return left + right
}
