package med

import (
	"slices"

	"github.com/SergeyCherepiuk/share/client/pkg/diff/ot"
)

func Diff(prev []byte, curr []byte) []ot.Operation {
	operations := make([]ot.Operation, 0)

	distance := distance(prev, curr)

	i, j := len(prev), len(curr)
	for i > 0 && j > 0 {
		var (
			substitutionDist = distance[i-1][j-1]
			insertionDist    = distance[i][j-1]
			deletionDist     = distance[i-1][j]
		)

		if substitutionDist == min(
			substitutionDist, insertionDist, deletionDist,
		) {
			i, j = i-1, j-1

			if distance[i][j] == distance[i+1][j+1] {
				continue
			}

			operations = append(operations, ot.Operation{
				Type:      ot.SUBSTITUTION,
				Position:  i,
				Character: curr[j],
			})
		} else if insertionDist <= deletionDist {
			j--
			operations = append(operations, ot.Operation{
				Type:      ot.INSERTION,
				Position:  i,
				Character: curr[j],
			})
		} else {
			i--
			operations = append(operations, ot.Operation{
				Type:      ot.DELETION,
				Position:  i,
				Character: prev[i],
			})
		}
	}

	for i > 0 {
		i--
		operations = append(operations, ot.Operation{
			Type:      ot.DELETION,
			Position:  i,
			Character: prev[i],
		})
	}

	for j > 0 {
		j--
		operations = append(operations, ot.Operation{
			Type:      ot.INSERTION,
			Position:  i,
			Character: curr[j],
		})
	}

	slices.Reverse(operations)
	return operations
}

func distance(prev, curr []byte) [][]int {
	distance := make([][]int, len(prev)+1)
	for i := 0; i < len(distance); i++ {
		distance[i] = make([]int, len(curr)+1)
	}

	for i := 1; i < len(prev)+1; i++ {
		distance[i][0] = i
	}

	for i := 1; i < len(curr)+1; i++ {
		distance[0][i] = i
	}

	for i := 1; i < len(distance); i++ {
		for j := 1; j < len(distance[i]); j++ {
			if prev[i-1] == curr[j-1] {
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
