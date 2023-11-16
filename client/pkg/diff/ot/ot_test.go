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
		{Type: diff.DELETION, Position: 0, Character: 'f'},
		{Type: diff.DELETION, Position: 0, Character: 'i'},
		{Type: diff.SUBSTITUTION, Position: 0, Character: '1'},
		{Type: diff.SUBSTITUTION, Position: 1, Character: '2'},
		{Type: diff.SUBSTITUTION, Position: 2, Character: '3'},
		{Type: diff.DELETION, Position: 4, Character: 's'},
		{Type: diff.DELETION, Position: 4, Character: 'e'},
		{Type: diff.DELETION, Position: 4, Character: 'c'},
		{Type: diff.SUBSTITUTION, Position: 4, Character: '4'},
		{Type: diff.SUBSTITUTION, Position: 5, Character: '5'},
		{Type: diff.SUBSTITUTION, Position: 6, Character: '6'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_PreviousLonger1(t *testing.T) {
	actual := Diff(
		[]byte("first\nsecond\nthird"),
		[]byte("first\nsecond\n"),
	)
	expected := []diff.Operation{
		{Type: diff.DELETION, Position: 12, Character: '\n'},
		{Type: diff.DELETION, Position: 12, Character: 't'},
		{Type: diff.DELETION, Position: 12, Character: 'h'},
		{Type: diff.DELETION, Position: 12, Character: 'i'},
		{Type: diff.DELETION, Position: 12, Character: 'r'},
		{Type: diff.DELETION, Position: 12, Character: 'd'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_PreviousLonger2(t *testing.T) {
	actual := Diff(
		[]byte("first\nsecond\nthird"),
		[]byte("first\nthird"),
	)
	expected := []diff.Operation{
		{Type: diff.DELETION, Position: 0, Character: 's'},
		{Type: diff.DELETION, Position: 0, Character: 'e'},
		{Type: diff.DELETION, Position: 0, Character: 'c'},
		{Type: diff.DELETION, Position: 0, Character: 'o'},
		{Type: diff.DELETION, Position: 0, Character: 'n'},
		{Type: diff.DELETION, Position: 0, Character: 'd'},
		{Type: diff.DELETION, Position: 0, Character: '\n'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_PreviousLonger3(t *testing.T) {
	actual := Diff(
		[]byte("first\nsecond\nthird"),
		[]byte("second\nthird"),
	)
	expected := []diff.Operation{
		{Type: diff.DELETION, Position: 0, Character: 'f'},
		{Type: diff.DELETION, Position: 0, Character: 'i'},
		{Type: diff.DELETION, Position: 0, Character: 'r'},
		{Type: diff.DELETION, Position: 0, Character: 's'},
		{Type: diff.DELETION, Position: 0, Character: 't'},
		{Type: diff.DELETION, Position: 0, Character: '\n'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_PreviousLonger4(t *testing.T) {
	actual := Diff(
		[]byte("first1\nsecond2\nthird3"),
		[]byte("first\nsecond"),
	)
	expected := []diff.Operation{
		{Type: diff.DELETION, Position: 5, Character: '1'},
		{Type: diff.DELETION, Position: 6, Character: '2'},
		{Type: diff.DELETION, Position: 6, Character: '\n'},
		{Type: diff.DELETION, Position: 0, Character: 't'},
		{Type: diff.DELETION, Position: 0, Character: 'h'},
		{Type: diff.DELETION, Position: 0, Character: 'i'},
		{Type: diff.DELETION, Position: 0, Character: 'r'},
		{Type: diff.DELETION, Position: 0, Character: 'd'},
		{Type: diff.DELETION, Position: 0, Character: '3'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_PreviousLonger5(t *testing.T) {
	actual := Diff(
		[]byte("first1\nsecond2\nthird3"),
		[]byte("first\nthird"),
	)
	expected := []diff.Operation{
		{Type: diff.DELETION, Position: 5, Character: '1'},
		{Type: diff.DELETION, Position: 0, Character: 's'},
		{Type: diff.SUBSTITUTION, Position: 0, Character: 't'},
		{Type: diff.SUBSTITUTION, Position: 1, Character: 'h'},
		{Type: diff.SUBSTITUTION, Position: 2, Character: 'i'},
		{Type: diff.SUBSTITUTION, Position: 3, Character: 'r'},
		{Type: diff.DELETION, Position: 5, Character: '2'},
		{Type: diff.DELETION, Position: 5, Character: '\n'},
		{Type: diff.DELETION, Position: 0, Character: 't'},
		{Type: diff.DELETION, Position: 0, Character: 'h'},
		{Type: diff.DELETION, Position: 0, Character: 'i'},
		{Type: diff.DELETION, Position: 0, Character: 'r'},
		{Type: diff.DELETION, Position: 0, Character: 'd'},
		{Type: diff.DELETION, Position: 0, Character: '3'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_PreviousLonger6(t *testing.T) {
	actual := Diff(
		[]byte("first1\nsecond2\nthird3"),
		[]byte("second\nthird"),
	)
	expected := []diff.Operation{
		{Type: diff.SUBSTITUTION, Position: 0, Character: 's'},
		{Type: diff.SUBSTITUTION, Position: 1, Character: 'e'},
		{Type: diff.SUBSTITUTION, Position: 2, Character: 'c'},
		{Type: diff.SUBSTITUTION, Position: 3, Character: 'o'},
		{Type: diff.SUBSTITUTION, Position: 4, Character: 'n'},
		{Type: diff.SUBSTITUTION, Position: 5, Character: 'd'},
		{Type: diff.DELETION, Position: 0, Character: 's'},
		{Type: diff.SUBSTITUTION, Position: 0, Character: 't'},
		{Type: diff.SUBSTITUTION, Position: 1, Character: 'h'},
		{Type: diff.SUBSTITUTION, Position: 2, Character: 'i'},
		{Type: diff.SUBSTITUTION, Position: 3, Character: 'r'},
		{Type: diff.DELETION, Position: 5, Character: '2'},
		{Type: diff.DELETION, Position: 5, Character: '\n'},
		{Type: diff.DELETION, Position: 0, Character: 't'},
		{Type: diff.DELETION, Position: 0, Character: 'h'},
		{Type: diff.DELETION, Position: 0, Character: 'i'},
		{Type: diff.DELETION, Position: 0, Character: 'r'},
		{Type: diff.DELETION, Position: 0, Character: 'd'},
		{Type: diff.DELETION, Position: 0, Character: '3'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_CurrentLonger1(t *testing.T) {
	actual := Diff(
		[]byte("first\nsecond"),
		[]byte("first\nsecond\nthird"),
	)
	expected := []diff.Operation{
		{Type: diff.INSERTION, Position: 12, Character: '\n'},
		{Type: diff.INSERTION, Position: 13, Character: 't'},
		{Type: diff.INSERTION, Position: 14, Character: 'h'},
		{Type: diff.INSERTION, Position: 15, Character: 'i'},
		{Type: diff.INSERTION, Position: 16, Character: 'r'},
		{Type: diff.INSERTION, Position: 17, Character: 'd'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_CurrentLonger2(t *testing.T) {
	actual := Diff(
		[]byte("first\nthird"),
		[]byte("first\nsecond\nthird"),
	)
	expected := []diff.Operation{
		{Type: diff.INSERTION, Position: 6, Character: 's'},
		{Type: diff.INSERTION, Position: 7, Character: 'e'},
		{Type: diff.INSERTION, Position: 8, Character: 'c'},
		{Type: diff.INSERTION, Position: 9, Character: 'o'},
		{Type: diff.INSERTION, Position: 10, Character: 'n'},
		{Type: diff.INSERTION, Position: 11, Character: 'd'},
		{Type: diff.INSERTION, Position: 12, Character: '\n'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_CurrentLonger3(t *testing.T) {
	actual := Diff(
		[]byte("second\nthird"),
		[]byte("first\nsecond\nthird"),
	)
	expected := []diff.Operation{
		{Type: diff.INSERTION, Position: 0, Character: 'f'},
		{Type: diff.INSERTION, Position: 1, Character: 'i'},
		{Type: diff.INSERTION, Position: 2, Character: 'r'},
		{Type: diff.INSERTION, Position: 3, Character: 's'},
		{Type: diff.INSERTION, Position: 4, Character: 't'},
		{Type: diff.INSERTION, Position: 5, Character: '\n'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_CurrentLonger4(t *testing.T) {
	actual := Diff(
		[]byte("first\nsecond"),
		[]byte("first1\nsecond2\nthird3"),
	)
	expected := []diff.Operation{
		{Type: diff.INSERTION, Position: 5, Character: '1'},
		{Type: diff.INSERTION, Position: 13, Character: '2'},
		{Type: diff.INSERTION, Position: 14, Character: '\n'},
		{Type: diff.INSERTION, Position: 15, Character: 't'},
		{Type: diff.INSERTION, Position: 16, Character: 'h'},
		{Type: diff.INSERTION, Position: 17, Character: 'i'},
		{Type: diff.INSERTION, Position: 18, Character: 'r'},
		{Type: diff.INSERTION, Position: 19, Character: 'd'},
		{Type: diff.INSERTION, Position: 20, Character: '3'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLines_CurrentLonger5(t *testing.T) {
	actual := Diff(
		[]byte("first\nthird"),
		[]byte("first1\nsecond2\nthird3"),
	)
	expected := []diff.Operation{
		{Type: diff.INSERTION, Position: 5, Character: '1'},
		{Type: diff.INSERTION, Position: 7, Character: 's'},
		{Type: diff.SUBSTITUTION, Position: 8, Character: 'e'},
		{Type: diff.SUBSTITUTION, Position: 9, Character: 'c'},
		{Type: diff.SUBSTITUTION, Position: 10, Character: 'o'},
		{Type: diff.SUBSTITUTION, Position: 11, Character: 'n'},
		{Type: diff.INSERTION, Position: 13, Character: '2'},
		{Type: diff.INSERTION, Position: 14, Character: '\n'},
		{Type: diff.INSERTION, Position: 15, Character: 't'},
		{Type: diff.INSERTION, Position: 16, Character: 'h'},
		{Type: diff.INSERTION, Position: 17, Character: 'i'},
		{Type: diff.INSERTION, Position: 18, Character: 'r'},
		{Type: diff.INSERTION, Position: 19, Character: 'd'},
		{Type: diff.INSERTION, Position: 20, Character: '3'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomematchingLines_CurrentLonger6(t *testing.T) {
	actual := Diff(
		[]byte("second\nthird"),
		[]byte("first1\nsecond2\nthird3"),
	)
	expected := []diff.Operation{
		{Type: diff.SUBSTITUTION, Position: 0, Character: 'f'},
		{Type: diff.SUBSTITUTION, Position: 1, Character: 'i'},
		{Type: diff.SUBSTITUTION, Position: 2, Character: 'r'},
		{Type: diff.SUBSTITUTION, Position: 3, Character: 's'},
		{Type: diff.SUBSTITUTION, Position: 4, Character: 't'},
		{Type: diff.SUBSTITUTION, Position: 5, Character: '1'},
		{Type: diff.INSERTION, Position: 7, Character: 's'},
		{Type: diff.SUBSTITUTION, Position: 8, Character: 'e'},
		{Type: diff.SUBSTITUTION, Position: 9, Character: 'c'},
		{Type: diff.SUBSTITUTION, Position: 10, Character: 'o'},
		{Type: diff.SUBSTITUTION, Position: 11, Character: 'n'},
		{Type: diff.INSERTION, Position: 13, Character: '2'},
		{Type: diff.INSERTION, Position: 14, Character: '\n'},
		{Type: diff.INSERTION, Position: 15, Character: 't'},
		{Type: diff.INSERTION, Position: 16, Character: 'h'},
		{Type: diff.INSERTION, Position: 17, Character: 'i'},
		{Type: diff.INSERTION, Position: 18, Character: 'r'},
		{Type: diff.INSERTION, Position: 19, Character: 'd'},
		{Type: diff.INSERTION, Position: 20, Character: '3'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_LinesSwap1(t *testing.T) {
	actual := Diff(
		[]byte("first\nsecond\nthird"),
		[]byte("second\nfirst\nthird"),
	)
	expected := []diff.Operation{
		{Type: diff.DELETION, Position: 0, Character: 'f'},
		{Type: diff.DELETION, Position: 0, Character: 'i'},
		{Type: diff.DELETION, Position: 0, Character: 'r'},
		{Type: diff.DELETION, Position: 0, Character: 's'},
		{Type: diff.DELETION, Position: 0, Character: 't'},
		{Type: diff.DELETION, Position: 0, Character: '\n'},
		{Type: diff.INSERTION, Position: 0, Character: 'f'},
		{Type: diff.INSERTION, Position: 1, Character: 'i'},
		{Type: diff.INSERTION, Position: 2, Character: 'r'},
		{Type: diff.INSERTION, Position: 3, Character: 's'},
		{Type: diff.INSERTION, Position: 4, Character: 't'},
		{Type: diff.INSERTION, Position: 5, Character: '\n'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_LinesSwap2(t *testing.T) {
	actual := Diff(
		[]byte("first\nsecond\nthird"),
		[]byte("second\nfirst\nthird\n"),
	)
	expected := []diff.Operation{
		{Type: diff.DELETION, Position: 0, Character: 'f'},
		{Type: diff.DELETION, Position: 0, Character: 'i'},
		{Type: diff.DELETION, Position: 0, Character: 'r'},
		{Type: diff.DELETION, Position: 0, Character: 's'},
		{Type: diff.DELETION, Position: 0, Character: 't'},
		{Type: diff.DELETION, Position: 0, Character: '\n'},
		{Type: diff.INSERTION, Position: 0, Character: 'f'},
		{Type: diff.INSERTION, Position: 1, Character: 'i'},
		{Type: diff.INSERTION, Position: 2, Character: 'r'},
		{Type: diff.INSERTION, Position: 3, Character: 's'},
		{Type: diff.INSERTION, Position: 4, Character: 't'},
		{Type: diff.INSERTION, Position: 5, Character: '\n'},
		{Type: diff.INSERTION, Position: 5, Character: '\n'},
	}

	internal.ShouldBe(t, expected, actual)
}
