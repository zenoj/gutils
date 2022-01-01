package stack

type stack struct {
	s []int
}

func mkStack() stack{
	return stack{s:make([]int, 0)}
}
