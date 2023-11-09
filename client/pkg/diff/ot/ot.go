package ot

import (
	"strings"

	"github.com/SergeyCherepiuk/share/client/pkg/diff"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/lcs"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/med"
)

func Diff(prev, curr []byte) []diff.Operation {
	var (
		linesPrev = strings.Split(string(prev), "\n")
		linesCurr = strings.Split(string(curr), "\n")
	)

	var operations []diff.Operation

	deletions, insertions := lcs.Diff(linesPrev, linesCurr)
	for len(deletions) != 0 && len(insertions) != 0 {
		var (
			ops            []diff.Operation
			prevLineNumber = deletions[0]
			prevLine       = linesPrev[prevLineNumber]
			currLineNumber = insertions[0]
			currLine       = linesCurr[currLineNumber]
		)

		if currLineNumber == prevLineNumber {
			ops = med.Diff([]byte(prevLine), []byte(currLine), prevLineNumber)
			deletions, insertions = deletions[1:], insertions[1:]
		} else if prevLineNumber < currLineNumber {
			ops = med.Diff([]byte(prevLine), []byte(""), prevLineNumber)
			deletions = deletions[1:]
		} else {
			ops = med.Diff([]byte(""), []byte(currLine), currLineNumber)
			insertions = insertions[1:]
		}

		operations = append(operations, ops...)
	}

	for len(deletions) > 0 {
		var (
			prevLineNumber = deletions[0]
			prevLine       = linesPrev[prevLineNumber]
		)

		deletions = deletions[1:]
		operations = append(operations, med.Diff(
			[]byte(prevLine), []byte(""), prevLineNumber)...,
		)
	}

	for len(insertions) > 0 {
		var (
			currLineNumber = insertions[0]
			currLine       = linesCurr[currLineNumber]
		)

		insertions = insertions[1:]
		operations = append(operations, med.Diff(
			[]byte(""), []byte(currLine), currLineNumber)...,
		)
	}

	return operations
}
