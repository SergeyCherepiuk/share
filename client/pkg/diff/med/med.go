package med

import (
	"slices"

	"github.com/SergeyCherepiuk/share/client/pkg/diff"
	"github.com/SergeyCherepiuk/share/client/pkg/internal"
)

func Diff(prev, curr []byte, line int) []diff.Operation {
	return walk(prev, curr, line, distance(prev, curr))
}

func distance(old, curr []byte) internal.Matrix {
	distance := make(internal.Matrix, len(old)+1)
	for i := 0; i < len(distance); i++ {
		distance[i] = make([]int, len(curr)+1)
	}

	for i := 1; i < len(old)+1; i++ {
		distance[i][0] = i
	}

	for i := 1; i < len(curr)+1; i++ {
		distance[0][i] = i
	}

	for i := 1; i < len(distance); i++ {
		for j := 1; j < len(distance[i]); j++ {
			if old[i-1] == curr[j-1] {
				distance[i][j] = distance[i-1][j-1]
				continue
			}

			substitution := distance[i-1][j-1]
			insertion := distance[i][j-1]
			deletion := distance[i-1][j]
			distance[i][j] = min(substitution, insertion, deletion) + 1
		}
	}

	return distance
}

func walk(prev, curr []byte, line int, distance internal.Matrix) []diff.Operation {
	operations := []diff.Operation{}

	idx := distance.LastIndex()
	for idx.Row > 0 && idx.Col > 0 {
		var (
			substitutionIdx = internal.Index{Row: idx.Row - 1, Col: idx.Col - 1}
			deletionIdx     = internal.Index{Row: idx.Row - 1, Col: idx.Col}
			insertionIdx    = internal.Index{Row: idx.Row, Col: idx.Col - 1}
		)

		if distance.Get(substitutionIdx) == min(
			distance.Get(substitutionIdx),
			distance.Get(deletionIdx),
			distance.Get(insertionIdx),
		) {
			currDistance := distance.Get(idx)
			idx = substitutionIdx

			if currDistance == distance.Get(substitutionIdx) {
				continue
			}

			operations = append(operations, diff.Substitution{
				Line:      line,
				Position:  substitutionIdx.Col,
				Character: curr[substitutionIdx.Col],
			})
		} else if distance.Get(insertionIdx) <= distance.Get(deletionIdx) {
			idx = insertionIdx
			operations = append(operations, diff.Insertion{
				Line:      line,
				Position:  insertionIdx.Col,
				Character: curr[insertionIdx.Col],
			})
		} else {
			idx = deletionIdx
			operations = append(operations, diff.Deletion{
				Line:     line,
				Position: idx.Row,
			})
		}
	}

	for idx.Col > 0 {
		operations = append(operations, diff.Insertion{
			Line:      line,
			Position:  idx.Col - 1,
			Character: curr[idx.Col-1],
		})
		idx = internal.Index{Row: idx.Row, Col: idx.Col - 1}
	}

	for idx.Row > 0 {
		operations = append(operations, diff.Deletion{
			Line:     line,
			Position: idx.Row - 1,
		})
		idx = internal.Index{Row: idx.Row - 1, Col: idx.Col}
	}

	slices.Reverse(operations)
	return operations
}
