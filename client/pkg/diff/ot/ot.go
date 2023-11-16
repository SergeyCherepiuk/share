package ot

import (
	"bytes"

	"github.com/SergeyCherepiuk/share/client/pkg/diff"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/lcs"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/med"
)

func Diff(prev []byte, curr []byte) []diff.Operation {
	var (
		linesPrev       = bytes.SplitAfter(prev, []byte{'\n'})
		linesCurr       = bytes.SplitAfter(curr, []byte{'\n'})
		linesStartsCurr = make([]int, len(linesCurr))
	)

	totalLength := 0
	for i, line := range linesCurr {
		linesStartsCurr[i] = totalLength
		totalLength += len(line)
	}

	operations := make([]diff.Operation, 0)

	deletions, insertions := lcs.Diff(linesPrev, linesCurr)
	deletionsPtr, insertionsPtr := 0, 0
	for deletionsPtr < len(deletions) && insertionsPtr < len(insertions) {
		var (
			iterOps       []diff.Operation
			deletionLine  = deletions[deletionsPtr]
			insertionLine = insertions[insertionsPtr]
		)

		if deletionLine == insertionLine {
			iterOps = med.Diff(linesPrev[deletionLine], linesCurr[insertionLine])
			deletionsPtr++
			insertionsPtr++
		} else if deletionLine < insertionLine {
			iterOps = med.Diff(linesPrev[deletionLine], []byte(""))
			deletionsPtr++
		} else {
			iterOps = med.Diff([]byte(""), linesCurr[insertionLine])
			insertionsPtr++
		}

		for _, operation := range iterOps {
			operation.Position += linesStartsCurr[insertionLine]
			operations = append(operations, operation)
		}
	}

	for ; deletionsPtr < len(deletions); deletionsPtr++ {
		deletionLine := deletions[deletionsPtr]
		for _, operation := range med.Diff(linesPrev[deletionLine], []byte("")) {
			operation.Position += linesStartsCurr[insertions[insertionsPtr-1]] // ???
			operations = append(operations, operation)
		}
	}

	for ; insertionsPtr < len(insertions); insertionsPtr++ {
		insertionLine := insertions[insertionsPtr]
		for _, operation := range med.Diff([]byte(""), linesCurr[insertionLine]) {
			operation.Position += linesStartsCurr[insertionLine]
			operations = append(operations, operation)
		}
	}

	return operations
}
