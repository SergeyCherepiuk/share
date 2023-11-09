package internal

type Matrix [][]int

func (m Matrix) LastIndex() Index {
	return Index{
		Row: len(m) - 1,
		Col: len(m[len(m)-1]) - 1,
	}
}

func (m Matrix) Get(i Index) int {
	return m[i.Row][i.Col]
}

type Index struct {
	Row int
	Col int
}
