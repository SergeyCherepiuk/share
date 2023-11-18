package ot

const (
	INSERTION = iota
	DELETION
	SUBSTITUTION
)

type Operation struct {
	Type      int  `json:"type"`
	Position  int  `json:"position"`
	Character byte `json:"character"`
}

func Adjust(operations []Operation) []Operation {
	adjusted := make([]Operation, len(operations))
	for i, curr := range operations {
		for _, prev := range operations[:i] {
			curr.Position += offset(prev, curr)
		}
		adjusted[i] = curr
	}
	return adjusted
}

func offset(prev, curr Operation) int {
	if prev.Position <= curr.Position {
		if prev.Type == INSERTION {
			return 1
		} else if prev.Type == DELETION {
			return -1
		}
	}
	return 0
}
