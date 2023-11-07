package ot

import (
	"slices"
)

type matrix [][]int

func (m matrix) lastIndex() index {
	return index{
		row: len(m) - 1,
		col: len(m[len(m)-1]) - 1,
	}
}

func (m matrix) get(i index) int {
	return m[i.row][i.col]
}

type index struct {
	row int
	col int
}

func MinimumEditDistance(old, curr []byte) []operation {
	return walkDistanceMatrix(
		old, curr, getDistanceMatrix(old, curr),
	)
}

func getDistanceMatrix(old, curr []byte) matrix {
	distance := make([][]int, len(old)+1)
	for i := 0; i < len(distance); i++ {
		distance[i] = make([]int, len(curr)+1)
	}

	for i := 1; i < len(old)+1; i++ {
		distance[i][0] = int(i)
	}
	for i := 1; i < len(curr)+1; i++ {
		distance[0][i] = (i)
	}

	for i := 1; i < len(distance); i++ {
		for j := 1; j < len(distance[i]); j++ {
			if old[i-1] == curr[j-1] {
				distance[i][j] = distance[i-1][j-1]
				continue
			}

			insertion := distance[i][j-1]
			deletion := distance[i-1][j]
			replacement := distance[i-1][j-1]
			distance[i][j] = min(insertion, deletion, replacement) + 1
		}
	}

	return distance
}

func walkDistanceMatrix(old, curr []byte, distance matrix) []operation {
	operations := []operation{}

	idx := distance.lastIndex()
	for idx.row > 0 && idx.col > 0 {
		var (
			substitutionIdx = index{row: idx.row - 1, col: idx.col - 1}
			deletionIdx     = index{row: idx.row - 1, col: idx.col}
			insertionIdx    = index{row: idx.row, col: idx.col - 1}
		)

		if distance.get(substitutionIdx) == min(
			distance.get(substitutionIdx),
			distance.get(deletionIdx),
			distance.get(insertionIdx),
		) {
			currDistance := distance.get(idx)
			idx = substitutionIdx

			if currDistance == distance.get(substitutionIdx) {
				continue
			}

			operations = append(operations, substitution{
				Position:  substitutionIdx.col,
				Character: curr[substitutionIdx.col],
			})
		} else if distance.get(insertionIdx) <= distance.get(deletionIdx) {
			idx = insertionIdx
			operations = append(operations, insertion{
				Position:  insertionIdx.col,
				Character: curr[insertionIdx.col],
			})
		} else {
			idx = deletionIdx
			operations = append(operations, deletion{Position: idx.row})
		}
	}

	for idx.col > 0 {
		operations = append(operations, insertion{
			Position:  idx.col - 1,
			Character: curr[idx.col-1],
		})
		idx = index{row: idx.row, col: idx.col - 1}
	}

	for idx.row > 0 {
		operations = append(operations, deletion{Position: idx.row - 1})
		idx = index{row: idx.row - 1, col: idx.col}
	}

	slices.Reverse(operations)
	return operations
}
