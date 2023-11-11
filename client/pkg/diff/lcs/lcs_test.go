package lcs

import (
	"testing"

	"github.com/SergeyCherepiuk/share/client/internal"
)

func Test_Diff_NoLines(t *testing.T) {
	actualDel, actualIns := Diff([]string{}, []string{})
	expectedDel, expectedIns := []int{}, []int{}

	internal.ShouldBe[[]int](t, expectedDel, actualDel)
	internal.ShouldBe[[]int](t, expectedIns, actualIns)
}

func Test_Diff_AllLinesMatch(t *testing.T) {
	actualDel, actualIns := Diff(
		[]string{"first", "second", "third"},
		[]string{"first", "second", "third"},
	)
	expectedDel, expectedIns := []int{}, []int{}

	internal.ShouldBe[[]int](t, expectedDel, actualDel)
	internal.ShouldBe[[]int](t, expectedIns, actualIns)
}

func Test_Diff_NoMatchingLines(t *testing.T) {
	actualDel, actualIns := Diff(
		[]string{"first", "second", "third"},
		[]string{"forth", "fifth", "sixth"},
	)
	expectedDel, expectedIns := []int{0, 1, 2}, []int{0, 1, 2}

	internal.ShouldBe[[]int](t, actualDel, expectedDel)
	internal.ShouldBe[[]int](t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_CurrentLarger1(t *testing.T) {
	actualDel, actualIns := Diff(
		[]string{"first", "second"},
		[]string{"first", "second", "third"},
	)
	expectedDel, expectedIns := []int{}, []int{2}

	internal.ShouldBe[[]int](t, actualDel, expectedDel)
	internal.ShouldBe[[]int](t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_CurrentLarger2(t *testing.T) {
	actualDel, actualIns := Diff(
		[]string{"first", "third"},
		[]string{"first", "second", "third"},
	)
	expectedDel, expectedIns := []int{}, []int{1}

	internal.ShouldBe[[]int](t, actualDel, expectedDel)
	internal.ShouldBe[[]int](t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_CurrentLarger3(t *testing.T) {
	actualDel, actualIns := Diff(
		[]string{"second", "third"},
		[]string{"first", "second", "third"},
	)
	expectedDel, expectedIns := []int{}, []int{0}

	internal.ShouldBe[[]int](t, actualDel, expectedDel)
	internal.ShouldBe[[]int](t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_CurrentLarger4(t *testing.T) {
	actualDel, actualIns := Diff(
		[]string{"second", "forth"},
		[]string{"first", "second", "third", "forth", "fifth"},
	)
	expectedDel, expectedIns := []int{}, []int{0, 2, 4}

	internal.ShouldBe[[]int](t, actualDel, expectedDel)
	internal.ShouldBe[[]int](t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_PreviousLarger1(t *testing.T) {
	actualDel, actualIns := Diff(
		[]string{"first", "second", "third"},
		[]string{"first", "second"},
	)
	expectedDel, expectedIns := []int{2}, []int{}

	internal.ShouldBe[[]int](t, actualDel, expectedDel)
	internal.ShouldBe[[]int](t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_PreviousLarger2(t *testing.T) {
	actualDel, actualIns := Diff(
		[]string{"first", "second", "third"},
		[]string{"first", "third"},
	)
	expectedDel, expectedIns := []int{1}, []int{}

	internal.ShouldBe[[]int](t, actualDel, expectedDel)
	internal.ShouldBe[[]int](t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_PreviousLarger3(t *testing.T) {
	actualDel, actualIns := Diff(
		[]string{"first", "second", "third"},
		[]string{"second", "third"},
	)
	expectedDel, expectedIns := []int{0}, []int{}

	internal.ShouldBe[[]int](t, actualDel, expectedDel)
	internal.ShouldBe[[]int](t, actualIns, expectedIns)
}

func Test_Diff_SomeMatchingLines_PreviousLarger4(t *testing.T) {
	actualDel, actualIns := Diff(
		[]string{"first", "second", "third", "forth", "fifth"},
		[]string{"second", "forth"},
	)
	expectedDel, expectedIns := []int{0, 2, 4}, []int{}

	internal.ShouldBe[[]int](t, actualDel, expectedDel)
	internal.ShouldBe[[]int](t, actualIns, expectedIns)
}

func Test_length_NoLines(t *testing.T) {
	actual := length([]string{}, []string{})
	expected := [][]int{{0}}

	internal.ShouldBe[[][]int](t, expected, actual)
}

func Test_length_AllMatchingLines(t *testing.T) {
	actual := length(
		[]string{"first", "second", "third"},
		[]string{"first", "second", "third"},
	)
	expected := [][]int{
		{0, 0, 0, 0},
		{0, 1, 1, 1},
		{0, 1, 2, 2},
		{0, 1, 2, 3},
	}

	internal.ShouldBe[[][]int](t, expected, actual)
}

func Test_length_NoMatchingLines(t *testing.T) {
	actual := length(
		[]string{"first", "second", "third"},
		[]string{"forth", "fifth", "sixth"},
	)
	expected := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	internal.ShouldBe[[][]int](t, expected, actual)
}

func Test_length_SomeMatchingLines_PreviousLarger(t *testing.T) {
	actual := length(
		[]string{"first", "second", "third", "forth", "fifth"},
		[]string{"second", "first", "third"},
	)
	expected := [][]int{
		{0, 0, 0, 0},
		{0, 0, 1, 1},
		{0, 1, 1, 1},
		{0, 1, 1, 2},
		{0, 1, 1, 2},
		{0, 1, 1, 2},
	}

	internal.ShouldBe[[][]int](t, expected, actual)
}

func Test_length_SomeMatchingLines_CurrentLarger(t *testing.T) {
	actual := length(
		[]string{"second", "first", "third"},
		[]string{"first", "second", "third", "forth", "fifth"},
	)
	expected := [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1},
		{0, 1, 1, 2, 2, 2},
	}

	internal.ShouldBe[[][]int](t, expected, actual)
}
