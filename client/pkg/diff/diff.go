package diff

import (
	"bytes"

	"github.com/SergeyCherepiuk/share/client/pkg/diff/lcs"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/med"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/ot"
)

func Diff(prev, curr []byte) []ot.Operation {
	var (
		operations = make([]ot.Operation, 0)

		prevLines = bytes.SplitAfter(prev, []byte{'\n'})
		currLines = bytes.SplitAfter(curr, []byte{'\n'})

		prevLinesStarts = lineStarts(prevLines)
		currLinesStarts = lineStarts(currLines)
	)

	deletions, insertions := lcs.Diff(prevLines, currLines)
	i, j := 0, 0
	for i < len(deletions) && j < len(insertions) {
		deletionLineNumber, insertionLineNumber := deletions[i], insertions[j]

		if deletionLineNumber < insertionLineNumber {
			i++
			for k, chr := range prevLines[deletionLineNumber] {
				operations = append(operations, ot.Operation{
					Type:      ot.DELETION,
					Position:  k + prevLinesStarts[deletionLineNumber],
					Character: chr,
				})
			}
		} else if deletionLineNumber > insertionLineNumber {
			j++
			for k, chr := range currLines[insertionLineNumber] {
				operations = append(operations, ot.Operation{
					Type:      ot.INSERTION,
					Position:  k + currLinesStarts[insertionLineNumber],
					Character: chr,
				})
			}
		} else {
			i, j = i+1, j+1
			ops := med.Diff(prevLines[deletionLineNumber], currLines[insertionLineNumber])
			for _, operation := range ops {
				operation.Position += prevLinesStarts[deletionLineNumber]
				operations = append(operations, operation)
			}
		}
	}

	for ; i < len(deletions); i++ {
		for _, chr := range prevLines[deletions[i]] {
			operations = append(operations, ot.Operation{
				Type:      ot.DELETION,
				Position:  prevLinesStarts[deletions[i]] - 1,
				Character: chr,
			})
		}
	}

	for ; j < len(insertions); j++ {
		for _, chr := range currLines[insertions[j]] {
			operations = append(operations, ot.Operation{
				Type:      ot.INSERTION,
				Position:  currLinesStarts[insertions[j]] - 1,
				Character: chr,
			})
		}
	}

	return ot.Adjust(operations)
}

func lineStarts(lines [][]byte) []int {
	starts := make([]int, len(lines))
	for i, total := 0, 0; i < len(starts); i++ {
		starts[i] = total
		total += len(lines[i])
	}
	return starts
}
