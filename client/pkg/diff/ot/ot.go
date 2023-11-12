package ot

import (
	"strings"

	"github.com/SergeyCherepiuk/share/client/pkg/diff"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/lcs"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/med"
)

func Diff(prev, curr []byte) []diff.Operation {
	var (
		linesPrev = strings.SplitAfter(string(prev), "\n")
		linesCurr = strings.SplitAfter(string(curr), "\n")
	)

	deletions, insertions := lcs.Diff(linesPrev, linesCurr)
	return diffRec(deletions, insertions, linesPrev, linesCurr)
}

// TODO: Refactor
func diffRec(deletions, insertions []int, prev, curr []string) []diff.Operation {
	if len(deletions) == 0 && len(insertions) == 0 {
		return []diff.Operation{}
	} else if len(insertions) == 0 {
		return append(
			deletionsFromLine([]byte(prev[(deletions)[0]]), (deletions)[0]),
			diffRec(deletions[1:], insertions, prev, curr)...,
		)
	} else if len(deletions) == 0 {
		return append(
			insertionsFromLine([]byte(curr[insertions[0]]), insertions[0]),
			diffRec(deletions, insertions[1:], prev, curr)...,
		)
	}

	if deletions[0] == insertions[0] {
		return append(
			med.Diff(
				[]byte(prev[deletions[0]]),
				[]byte(curr[insertions[0]]),
				deletions[0],
			),
			diffRec(deletions[1:], insertions[1:], prev, curr)...,
		)
	} else if deletions[0] < insertions[0] {
		return append(
			deletionsFromLine([]byte(prev[deletions[0]]), deletions[0]),
			diffRec(deletions[1:], insertions, prev, curr)...,
		)
	} else {
		return append(
			insertionsFromLine([]byte(curr[insertions[0]]), insertions[0]),
			diffRec(deletions, insertions[1:], prev, curr)...,
		)
	}
}

func deletionsFromLine(line []byte, lineNumber int) []diff.Operation {
	ops := make([]diff.Operation, len(line))
	for i := range ops {
		ops[i] = diff.Deletion{
			Line:     lineNumber,
			Position: 0,
		}
	}
	return ops
}

func insertionsFromLine(line []byte, lineNumber int) []diff.Operation {
	ops := make([]diff.Operation, len(line))
	for i := range ops {
		ops[i] = diff.Insertion{
			Line:      lineNumber,
			Position:  i,
			Character: line[i],
		}
	}
	return ops
}
