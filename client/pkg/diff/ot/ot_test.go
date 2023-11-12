package ot

import (
	"testing"

	"github.com/SergeyCherepiuk/share/client/internal"
	"github.com/SergeyCherepiuk/share/client/pkg/diff"
)

func Test_Diff_EmptyLines1(t *testing.T) {
	actual := Diff([]byte{}, []byte{})
	expected := []diff.Operation{}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_EmptyLines2(t *testing.T) {
	actual := Diff([]byte(""), []byte(""))
	expected := []diff.Operation{}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SameLines(t *testing.T) {
	actual := Diff(
		[]byte("first line\nsecond line"),
		[]byte("first line\nsecond line"),
	)
	expected := []diff.Operation{}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_DifferentLines(t *testing.T) {
	actual := Diff([]byte("first\nsecond"), []byte("123\n456"))
	expected := []diff.Operation{
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
		diff.Substitution{Line: 0, Position: 0, Character: '1'},
		diff.Substitution{Line: 0, Position: 1, Character: '2'},
		diff.Substitution{Line: 0, Position: 2, Character: '3'},
		diff.Deletion{Line: 1, Position: 0},
		diff.Deletion{Line: 1, Position: 0},
		diff.Deletion{Line: 1, Position: 0},
		diff.Substitution{Line: 1, Position: 0, Character: '4'},
		diff.Substitution{Line: 1, Position: 1, Character: '5'},
		diff.Substitution{Line: 1, Position: 2, Character: '6'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_PreviousLonger1(t *testing.T) {
	actual := Diff(
		[]byte("first\nsecond\nthird"),
		[]byte("first\nsecond"),
	)
	expected := []diff.Operation{
		diff.Deletion{Line: 1, Position: 6},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_PreviousLonger2(t *testing.T) {
	actual := Diff(
		[]byte("first\nsecond\nthird"),
		[]byte("first\nthird"),
	)
	expected := []diff.Operation{
		diff.Deletion{Line: 1, Position: 0},
		diff.Deletion{Line: 1, Position: 0},
		diff.Deletion{Line: 1, Position: 0},
		diff.Deletion{Line: 1, Position: 0},
		diff.Deletion{Line: 1, Position: 0},
		diff.Deletion{Line: 1, Position: 0},
		diff.Deletion{Line: 1, Position: 0},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_PreviousLonger3(t *testing.T) {
	actual := Diff(
		[]byte("first\nsecond\nthird"),
		[]byte("second\nthird"),
	)
	expected := []diff.Operation{
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_PreviousLonger4(t *testing.T) {
	actual := Diff(
		[]byte("first1\nsecond2\nthird3"),
		[]byte("first\nsecond"),
	)
	expected := []diff.Operation{
		diff.Deletion{Line: 0, Position: 5},
		diff.Deletion{Line: 1, Position: 6},
		diff.Deletion{Line: 1, Position: 6},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_PreviousLonger5(t *testing.T) {
	actual := Diff(
		[]byte("first1\nsecond2\nthird3"),
		[]byte("first\nthird"),
	)
	expected := []diff.Operation{
		diff.Deletion{Line: 0, Position: 5},
		diff.Deletion{Line: 1, Position: 0},
		diff.Substitution{Line: 1, Position: 0, Character: 't'},
		diff.Substitution{Line: 1, Position: 1, Character: 'h'},
		diff.Substitution{Line: 1, Position: 2, Character: 'i'},
		diff.Substitution{Line: 1, Position: 3, Character: 'r'},
		diff.Deletion{Line: 1, Position: 5},
		diff.Deletion{Line: 1, Position: 5},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_PreviousLonger6(t *testing.T) {
	actual := Diff(
		[]byte("first1\nsecond2\nthird3"),
		[]byte("second\nthird"),
	)
	expected := []diff.Operation{
		diff.Substitution{Line: 0, Position: 0, Character: 's'},
		diff.Substitution{Line: 0, Position: 1, Character: 'e'},
		diff.Substitution{Line: 0, Position: 2, Character: 'c'},
		diff.Substitution{Line: 0, Position: 3, Character: 'o'},
		diff.Substitution{Line: 0, Position: 4, Character: 'n'},
		diff.Substitution{Line: 0, Position: 5, Character: 'd'},
		diff.Deletion{Line: 1, Position: 0},
		diff.Substitution{Line: 1, Position: 0, Character: 't'},
		diff.Substitution{Line: 1, Position: 1, Character: 'h'},
		diff.Substitution{Line: 1, Position: 2, Character: 'i'},
		diff.Substitution{Line: 1, Position: 3, Character: 'r'},
		diff.Deletion{Line: 1, Position: 5},
		diff.Deletion{Line: 1, Position: 5},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
		diff.Deletion{Line: 2, Position: 0},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_CurrentLonger1(t *testing.T) {
	actual := Diff(
		[]byte("first\nsecond"),
		[]byte("first\nsecond\nthird"),
	)
	expected := []diff.Operation{
		diff.Insertion{Line: 1, Position: 6, Character: '\n'},
		diff.Insertion{Line: 2, Position: 0, Character: 't'},
		diff.Insertion{Line: 2, Position: 1, Character: 'h'},
		diff.Insertion{Line: 2, Position: 2, Character: 'i'},
		diff.Insertion{Line: 2, Position: 3, Character: 'r'},
		diff.Insertion{Line: 2, Position: 4, Character: 'd'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_CurrentLonger2(t *testing.T) {
	actual := Diff(
		[]byte("first\nthird"),
		[]byte("first\nsecond\nthird"),
	)
	expected := []diff.Operation{
		diff.Insertion{Line: 1, Position: 0, Character: 's'},
		diff.Insertion{Line: 1, Position: 1, Character: 'e'},
		diff.Insertion{Line: 1, Position: 2, Character: 'c'},
		diff.Insertion{Line: 1, Position: 3, Character: 'o'},
		diff.Insertion{Line: 1, Position: 4, Character: 'n'},
		diff.Insertion{Line: 1, Position: 5, Character: 'd'},
		diff.Insertion{Line: 1, Position: 6, Character: '\n'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_CurrentLonger3(t *testing.T) {
	actual := Diff(
		[]byte("second\nthird"),
		[]byte("first\nsecond\nthird"),
	)
	expected := []diff.Operation{
		diff.Insertion{Line: 0, Position: 0, Character: 'f'},
		diff.Insertion{Line: 0, Position: 1, Character: 'i'},
		diff.Insertion{Line: 0, Position: 2, Character: 'r'},
		diff.Insertion{Line: 0, Position: 3, Character: 's'},
		diff.Insertion{Line: 0, Position: 4, Character: 't'},
		diff.Insertion{Line: 0, Position: 5, Character: '\n'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_CurrentLonger4(t *testing.T) {
	actual := Diff(
		[]byte("first\nsecond"),
		[]byte("first1\nsecond2\nthird3"),
	)
	expected := []diff.Operation{
		diff.Insertion{Line: 0, Position: 5, Character: '1'},
		diff.Insertion{Line: 1, Position: 6, Character: '2'},
		diff.Insertion{Line: 1, Position: 7, Character: '\n'},
		diff.Insertion{Line: 2, Position: 0, Character: 't'},
		diff.Insertion{Line: 2, Position: 1, Character: 'h'},
		diff.Insertion{Line: 2, Position: 2, Character: 'i'},
		diff.Insertion{Line: 2, Position: 3, Character: 'r'},
		diff.Insertion{Line: 2, Position: 4, Character: 'd'},
		diff.Insertion{Line: 2, Position: 5, Character: '3'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_CurrentLonger5(t *testing.T) {
	actual := Diff(
		[]byte("first\nthird"),
		[]byte("first1\nsecond2\nthird3"),
	)
	expected := []diff.Operation{
		diff.Insertion{Line: 0, Position: 5, Character: '1'},
		diff.Insertion{Line: 1, Position: 0, Character: 's'},
		diff.Substitution{Line: 1, Position: 1, Character: 'e'},
		diff.Substitution{Line: 1, Position: 2, Character: 'c'},
		diff.Substitution{Line: 1, Position: 3, Character: 'o'},
		diff.Substitution{Line: 1, Position: 4, Character: 'n'},
		diff.Insertion{Line: 1, Position: 6, Character: '2'},
		diff.Insertion{Line: 1, Position: 7, Character: '\n'},
		diff.Insertion{Line: 2, Position: 0, Character: 't'},
		diff.Insertion{Line: 2, Position: 1, Character: 'h'},
		diff.Insertion{Line: 2, Position: 2, Character: 'i'},
		diff.Insertion{Line: 2, Position: 3, Character: 'r'},
		diff.Insertion{Line: 2, Position: 4, Character: 'd'},
		diff.Insertion{Line: 2, Position: 5, Character: '3'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomematchingLines_CurrentLonger6(t *testing.T) {
	actual := Diff(
		[]byte("second\nthird"),
		[]byte("first1\nsecond2\nthird3"),
	)
	expected := []diff.Operation{
		diff.Substitution{Line: 0, Position: 0, Character: 'f'},
		diff.Substitution{Line: 0, Position: 1, Character: 'i'},
		diff.Substitution{Line: 0, Position: 2, Character: 'r'},
		diff.Substitution{Line: 0, Position: 3, Character: 's'},
		diff.Substitution{Line: 0, Position: 4, Character: 't'},
		diff.Substitution{Line: 0, Position: 5, Character: '1'},
		diff.Insertion{Line: 1, Position: 0, Character: 's'},
		diff.Substitution{Line: 1, Position: 1, Character: 'e'},
		diff.Substitution{Line: 1, Position: 2, Character: 'c'},
		diff.Substitution{Line: 1, Position: 3, Character: 'o'},
		diff.Substitution{Line: 1, Position: 4, Character: 'n'},
		diff.Insertion{Line: 1, Position: 6, Character: '2'},
		diff.Insertion{Line: 1, Position: 7, Character: '\n'},
		diff.Insertion{Line: 2, Position: 0, Character: 't'},
		diff.Insertion{Line: 2, Position: 1, Character: 'h'},
		diff.Insertion{Line: 2, Position: 2, Character: 'i'},
		diff.Insertion{Line: 2, Position: 3, Character: 'r'},
		diff.Insertion{Line: 2, Position: 4, Character: 'd'},
		diff.Insertion{Line: 2, Position: 5, Character: '3'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_LinesSwap1(t *testing.T) {
	actual := Diff(
		[]byte("first\nsecond\nthird"),
		[]byte("second\nfirst\nthird"),
	)
	expected := []diff.Operation{
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
		diff.Insertion{Line: 1, Position: 0, Character: 'f'},
		diff.Insertion{Line: 1, Position: 1, Character: 'i'},
		diff.Insertion{Line: 1, Position: 2, Character: 'r'},
		diff.Insertion{Line: 1, Position: 3, Character: 's'},
		diff.Insertion{Line: 1, Position: 4, Character: 't'},
		diff.Insertion{Line: 1, Position: 5, Character: '\n'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_LinesSwap2(t *testing.T) {
	actual := Diff(
		[]byte("first\nsecond\nthird"),
		[]byte("second\nfirst\nthird\n"),
	)
	expected := []diff.Operation{
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
		diff.Deletion{Line: 0, Position: 0},
		diff.Insertion{Line: 1, Position: 0, Character: 'f'},
		diff.Insertion{Line: 1, Position: 1, Character: 'i'},
		diff.Insertion{Line: 1, Position: 2, Character: 'r'},
		diff.Insertion{Line: 1, Position: 3, Character: 's'},
		diff.Insertion{Line: 1, Position: 4, Character: 't'},
		diff.Insertion{Line: 1, Position: 5, Character: '\n'},
		diff.Insertion{Line: 2, Position: 5, Character: '\n'},
	}

	internal.ShouldBe(t, expected, actual)
}
