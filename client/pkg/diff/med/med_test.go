package med

import (
	"testing"

	"github.com/SergeyCherepiuk/share/client/internal"
	"github.com/SergeyCherepiuk/share/client/pkg/diff"
)

func Test_Diff_EmptyLines1(t *testing.T) {
	actual := Diff([]byte{}, []byte{}, 0)
	expected := []diff.Operation{}

	internal.ShouldBe[[]diff.Operation](t, expected, actual)
}

func Test_Diff_EmptyLines2(t *testing.T) {
	actual := Diff([]byte(""), []byte(""), 0)
	expected := []diff.Operation{}

	internal.ShouldBe[[]diff.Operation](t, expected, actual)
}

func Test_Diff_SameLines(t *testing.T) {
	actual := Diff([]byte("test"), []byte("test"), 0)
	expected := []diff.Operation{}

	internal.ShouldBe[[]diff.Operation](t, expected, actual)
}

func Test_Diff_DifferentLines(t *testing.T) {
	actual := Diff([]byte("asdf"), []byte("qwer"), 0)
	expected := []diff.Operation{
		diff.Substitution{Line: 0, Position: 3, Character: 'r'},
		diff.Substitution{Line: 0, Position: 2, Character: 'e'},
		diff.Substitution{Line: 0, Position: 1, Character: 'w'},
		diff.Substitution{Line: 0, Position: 0, Character: 'q'},
	}

	internal.ShouldBe[[]diff.Operation](t, expected, actual)
}

func Test_Diff_SomeMatchingLetters_PreviousLonger1(t *testing.T) {
	actual := Diff([]byte("those"), []byte("two"), 0)
	expected := []diff.Operation{
		diff.Deletion{Line: 0, Position: 4},
		diff.Deletion{Line: 0, Position: 3},
		diff.Substitution{Line: 0, Position: 1, Character: 'w'},
	}

	internal.ShouldBe[[]diff.Operation](t, expected, actual)
}

func Test_Diff_SomeMatchingLetters_PreviousLonger2(t *testing.T) {
	actual := Diff([]byte("after"), []byte("tor"), 0)
	expected := []diff.Operation{
		diff.Substitution{Line: 0, Position: 3, Character: 'o'},
		diff.Deletion{Line: 0, Position: 1},
		diff.Deletion{Line: 0, Position: 0},
	}

	internal.ShouldBe[[]diff.Operation](t, expected, actual)
}

func Test_Diff_SomeMatchingLetters_CurrentLonger1(t *testing.T) {
	actual := Diff([]byte("one"), []byte("some"), 0)
	expected := []diff.Operation{
		diff.Substitution{Line: 0, Position: 1, Character: 'm'},
		diff.Insertion{Line: 0, Position: 0, Character: 's'},
	}

	internal.ShouldBe[[]diff.Operation](t, expected, actual)
}

func Test_Diff_SomeMatchingLetters_CurrentLonger2(t *testing.T) {
	actual := Diff([]byte("cat"), []byte("cars"), 0)
	expected := []diff.Operation{
		diff.Substitution{Line: 0, Position: 2, Character: 's'},
		diff.Insertion{Line: 0, Position: 2, Character: 'r'},
	}

	internal.ShouldBe[[]diff.Operation](t, expected, actual)
}

func Test_distance_EmptyLines1(t *testing.T) {
	actual := distance([]byte{}, []byte{})
	expected := [][]int{{0}}

	internal.ShouldBe[[][]int](t, expected, actual)
}

func Test_distance_EmptyLines2(t *testing.T) {
	actual := distance([]byte(""), []byte(""))
	expected := [][]int{{0}}

	internal.ShouldBe[[][]int](t, expected, actual)
}

func Test_distance_SameLines(t *testing.T) {
	actual := distance([]byte("test"), []byte("test"))
	expected := [][]int{
		{0, 1, 2, 3, 4},
		{1, 0, 1, 2, 3},
		{2, 1, 0, 1, 2},
		{3, 2, 1, 0, 1},
		{4, 3, 2, 1, 0},
	}

	internal.ShouldBe[[][]int](t, expected, actual)
}

func Test_distance_DifferentLines(t *testing.T) {
	actual := distance([]byte("asdf"), []byte("qwer"))
	expected := [][]int{
		{0, 1, 2, 3, 4},
		{1, 1, 2, 3, 4},
		{2, 2, 2, 3, 4},
		{3, 3, 3, 3, 4},
		{4, 4, 4, 4, 4},
	}

	internal.ShouldBe[[][]int](t, expected, actual)
}

func Test_distance_SomeMatchingLetters_PreviousLonger(t *testing.T) {
	actual := distance([]byte("those"), []byte("two"))
	expected := [][]int{
		{0, 1, 2, 3},
		{1, 0, 1, 2},
		{2, 1, 1, 2},
		{3, 2, 2, 1},
		{4, 3, 3, 2},
		{5, 4, 4, 3},
	}

	internal.ShouldBe[[][]int](t, expected, actual)
}

func Test_distance_SomeMatchingLetters_CurrentLonger(t *testing.T) {
	actual := distance([]byte("one"), []byte("some"))
	expected := [][]int{
		{0, 1, 2, 3, 4},
		{1, 1, 1, 2, 3},
		{2, 2, 2, 2, 3},
		{3, 3, 3, 3, 2},
	}

	internal.ShouldBe[[][]int](t, expected, actual)
}
