package med

import (
	"testing"

	"github.com/SergeyCherepiuk/share/client/internal"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/ot"
)

func Test_Diff_EmptyLines1(t *testing.T) {
	actual := Diff([]byte{}, []byte{})
	expected := []ot.Operation{}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_EmptyLines2(t *testing.T) {
	actual := Diff([]byte(""), []byte(""))
	expected := []ot.Operation{}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SameLines(t *testing.T) {
	actual := Diff([]byte("test"), []byte("test"))
	expected := []ot.Operation{}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_DifferentLines(t *testing.T) {
	actual := Diff([]byte("asdf"), []byte("qwer"))
	expected := []ot.Operation{
		{Type: ot.SUBSTITUTION, Position: 0, Character: 'q'},
		{Type: ot.SUBSTITUTION, Position: 1, Character: 'w'},
		{Type: ot.SUBSTITUTION, Position: 2, Character: 'e'},
		{Type: ot.SUBSTITUTION, Position: 3, Character: 'r'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLetters_PreviousLonger1(t *testing.T) {
	actual := Diff([]byte("those"), []byte("two"))
	expected := []ot.Operation{
		{Type: ot.SUBSTITUTION, Position: 1, Character: 'w'},
		{Type: ot.DELETION, Position: 4, Character: 'e'},
		{Type: ot.DELETION, Position: 3, Character: 's'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLetters_PreviousLonger2(t *testing.T) {
	actual := Diff([]byte("after"), []byte("tor"))
	expected := []ot.Operation{
		{Type: ot.SUBSTITUTION, Position: 3, Character: 'o'},
		{Type: ot.DELETION, Position: 1, Character: 'f'},
		{Type: ot.DELETION, Position: 0, Character: 'a'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLetters_CurrentLonger1(t *testing.T) {
	actual := Diff([]byte("one"), []byte("some"))
	expected := []ot.Operation{
		{Type: ot.SUBSTITUTION, Position: 1, Character: 'm'},
		{Type: ot.INSERTION, Position: 0, Character: 's'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Diff_SomeMatchingLetters_CurrentLonger2(t *testing.T) {
	actual := Diff([]byte("cat"), []byte("cars"))
	expected := []ot.Operation{
		{Type: ot.SUBSTITUTION, Position: 2, Character: 's'},
		{Type: ot.INSERTION, Position: 2, Character: 'r'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_distance_EmptyLines1(t *testing.T) {
	actual := distance([]byte{}, []byte{})
	expected := [][]int{{0}}

	internal.ShouldBe(t, expected, actual)
}

func Test_distance_EmptyLines2(t *testing.T) {
	actual := distance([]byte(""), []byte(""))
	expected := [][]int{{0}}

	internal.ShouldBe(t, expected, actual)
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

	internal.ShouldBe(t, expected, actual)
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

	internal.ShouldBe(t, expected, actual)
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

	internal.ShouldBe(t, expected, actual)
}

func Test_distance_SomeMatchingLetters_CurrentLonger(t *testing.T) {
	actual := distance([]byte("one"), []byte("some"))
	expected := [][]int{
		{0, 1, 2, 3, 4},
		{1, 1, 1, 2, 3},
		{2, 2, 2, 2, 3},
		{3, 3, 3, 3, 2},
	}

	internal.ShouldBe(t, expected, actual)
}
