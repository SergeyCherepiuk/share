package lcs

import (
	"testing"

	"github.com/SergeyCherepiuk/share/client/internal"
)

func Test_Diff_NoLines(t *testing.T) {
	actualDel, actualIns := Diff([][]byte{}, [][]byte{})
	expectedDel, expectedIns := []int{}, []int{}

	internal.ShouldBe(t, expectedDel, actualDel)
	internal.ShouldBe(t, expectedIns, actualIns)
}

func Test_Diff_AllLinesMatch(t *testing.T) {
	actualDel, actualIns := Diff(
		[][]byte{[]byte("first"), []byte("second"), []byte("third")},
		[][]byte{[]byte("first"), []byte("second"), []byte("third")},
	)
	expectedDel, expectedIns := []int{}, []int{}

	internal.ShouldBe(t, expectedDel, actualDel)
	internal.ShouldBe(t, expectedIns, actualIns)
}

func Test_Diff_NoMatchingLines(t *testing.T) {
	actualDel, actualIns := Diff(
		[][]byte{[]byte("first"), []byte("second"), []byte("third")},
		[][]byte{[]byte("forth"), []byte("fifth"), []byte("sixth")},
	)
	expectedDel, expectedIns := []int{0, 1, 2}, []int{0, 1, 2}

	internal.ShouldBe(t, actualDel, expectedDel)
	internal.ShouldBe(t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_CurrentLarger1(t *testing.T) {
	actualDel, actualIns := Diff(
		[][]byte{[]byte("first"), []byte("second")},
		[][]byte{[]byte("first"), []byte("second"), []byte("third")},
	)
	expectedDel, expectedIns := []int{}, []int{2}

	internal.ShouldBe(t, actualDel, expectedDel)
	internal.ShouldBe(t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_CurrentLarger2(t *testing.T) {
	actualDel, actualIns := Diff(
		[][]byte{[]byte("first"), []byte("third")},
		[][]byte{[]byte("first"), []byte("second"), []byte("third")},
	)
	expectedDel, expectedIns := []int{}, []int{1}

	internal.ShouldBe(t, actualDel, expectedDel)
	internal.ShouldBe(t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_CurrentLarger3(t *testing.T) {
	actualDel, actualIns := Diff(
		[][]byte{[]byte("second"), []byte("third")},
		[][]byte{[]byte("first"), []byte("second"), []byte("third")},
	)
	expectedDel, expectedIns := []int{}, []int{0}

	internal.ShouldBe(t, actualDel, expectedDel)
	internal.ShouldBe(t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_CurrentLarger4(t *testing.T) {
	actualDel, actualIns := Diff(
		[][]byte{[]byte("second"), []byte("forth")},
		[][]byte{[]byte("first"), []byte("second"), []byte("third"), []byte("forth"), []byte("fifth")},
	)
	expectedDel, expectedIns := []int{}, []int{0, 2, 4}

	internal.ShouldBe(t, actualDel, expectedDel)
	internal.ShouldBe(t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_PreviousLarger1(t *testing.T) {
	actualDel, actualIns := Diff(
		[][]byte{[]byte("first"), []byte("second"), []byte("third")},
		[][]byte{[]byte("first"), []byte("second")},
	)
	expectedDel, expectedIns := []int{2}, []int{}

	internal.ShouldBe(t, actualDel, expectedDel)
	internal.ShouldBe(t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_PreviousLarger2(t *testing.T) {
	actualDel, actualIns := Diff(
		[][]byte{[]byte("first"), []byte("second"), []byte("third")},
		[][]byte{[]byte("first"), []byte("third")},
	)
	expectedDel, expectedIns := []int{1}, []int{}

	internal.ShouldBe(t, actualDel, expectedDel)
	internal.ShouldBe(t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_PreviousLarger3(t *testing.T) {
	actualDel, actualIns := Diff(
		[][]byte{[]byte("first"), []byte("second"), []byte("third")},
		[][]byte{[]byte("second"), []byte("third")},
	)
	expectedDel, expectedIns := []int{0}, []int{}

	internal.ShouldBe(t, actualDel, expectedDel)
	internal.ShouldBe(t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_PreviousLarger4(t *testing.T) {
	actualDel, actualIns := Diff(
		[][]byte{[]byte("first"), []byte("second"), []byte("third"), []byte("forth"), []byte("fifth")},
		[][]byte{[]byte("second"), []byte("forth")},
	)
	expectedDel, expectedIns := []int{0, 2, 4}, []int{}

	internal.ShouldBe(t, actualDel, expectedDel)
	internal.ShouldBe(t, actualIns, expectedIns)
}

func Test_Diff_LinesSwap(t *testing.T) {
	actualDel, actualIns := Diff(
		[][]byte{[]byte("first"), []byte("second")},
		[][]byte{[]byte("second"), []byte("first")},
	)
	expectedDel, expectedIns := []int{0}, []int{1}

	internal.ShouldBe(t, expectedDel, actualDel)
	internal.ShouldBe(t, expectedIns, actualIns)
}

func Test_length_NoLines(t *testing.T) {
	actual := length([][]byte{}, [][]byte{})
	expected := [][]int{{0}}

	internal.ShouldBe(t, expected, actual)
}

func Test_length_AllMatchingLines(t *testing.T) {
	actual := length(
		[][]byte{[]byte("first"), []byte("second"), []byte("third")},
		[][]byte{[]byte("first"), []byte("second"), []byte("third")},
	)
	expected := [][]int{
		{0, 0, 0, 0},
		{0, 1, 1, 1},
		{0, 1, 2, 2},
		{0, 1, 2, 3},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_length_NoMatchingLines(t *testing.T) {
	actual := length(
		[][]byte{[]byte("first"), []byte("second"), []byte("third")},
		[][]byte{[]byte("forth"), []byte("fifth"), []byte("sixth")},
	)
	expected := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_length_SomeMatchingLines_PreviousLarger(t *testing.T) {
	actual := length(
		[][]byte{[]byte("first"), []byte("second"), []byte("third"), []byte("forth"), []byte("fifth")},
		[][]byte{[]byte("second"), []byte("first"), []byte("third")},
	)
	expected := [][]int{
		{0, 0, 0, 0},
		{0, 0, 1, 1},
		{0, 1, 1, 1},
		{0, 1, 1, 2},
		{0, 1, 1, 2},
		{0, 1, 1, 2},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_length_SomeMatchingLines_CurrentLarger(t *testing.T) {
	actual := length(
		[][]byte{[]byte("second"), []byte("first"), []byte("third")},
		[][]byte{[]byte("first"), []byte("second"), []byte("third"), []byte("forth"), []byte("fifth")},
	)
	expected := [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1},
		{0, 1, 1, 2, 2, 2},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_length_LinesSwap(t *testing.T) {
	actual := length(
		[][]byte{[]byte("first"), []byte("second")},
		[][]byte{[]byte("second"), []byte("first")},
	)
	expected := [][]int{
		{0, 0, 0},
		{0, 0, 1},
		{0, 1, 1},
	}

	internal.ShouldBe(t, expected, actual)
}
