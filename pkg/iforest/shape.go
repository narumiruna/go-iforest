package iforest

type Shape []int

func (s Shape) Dim() int {
	return len(s)
}

func (s Shape) Size(i int) int {
	if i < 0 || i >= len(s) {
		panic("index out of range")
	}
	return s[i]
}
