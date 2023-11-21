package med

import (
	"testing"

	"github.com/SergeyCherepiuk/share/client/internal"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/ot"
)

func Test_distance_Nil(t *testing.T) {
	actual := distance(nil, nil)
	expected := [][]int{{0}}

	internal.ShouldBe(t, expected, actual)
}

func Test_distance_Empty(t *testing.T) {
	actual := distance([]byte{}, []byte{})
	expected := [][]int{{0}}

	internal.ShouldBe(t, expected, actual)
}

func Test_distance_Identical(t *testing.T) {
	actual := distance([]byte("word"), []byte("word"))
	expected := [][]int{
		{0, 1, 2, 3, 4},
		{1, 0, 1, 2, 3},
		{2, 1, 0, 1, 2},
		{3, 2, 1, 0, 1},
		{4, 3, 2, 1, 0},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_distance_SameLength(t *testing.T) {
	actual := distance([]byte("one"), []byte("two"))
	expected := [][]int{
		{0, 1, 2, 3},
		{1, 1, 2, 2},
		{2, 2, 2, 3},
		{3, 3, 3, 3},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_distance_PreviousLonger(t *testing.T) {
	actual := distance([]byte("one1"), []byte("on"))
	expected := [][]int{
		{0, 1, 2},
		{1, 0, 1},
		{2, 1, 0},
		{3, 2, 1},
		{4, 3, 2},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_distance_CurrentLonger(t *testing.T) {
	actual := distance([]byte("tw"), []byte("two2"))
	expected := [][]int{
		{0, 1, 2, 3, 4},
		{1, 0, 1, 2, 3},
		{2, 1, 0, 1, 2},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_MinimumEditDistance_Nil(t *testing.T) {
	actual := MinimumEditDistance(nil, nil)
	expected := []ot.Operation{}

	internal.ShouldBe(t, expected, actual)
}

func Test_MinimumEditDistance_Empty(t *testing.T) {
	actual := MinimumEditDistance([]byte(""), []byte(""))
	expected := []ot.Operation{}

	internal.ShouldBe(t, expected, actual)
}

func Test_MinimumEditDistance_Identical(t *testing.T) {
	actual := MinimumEditDistance([]byte("word"), []byte("word"))
	expected := []ot.Operation{}

	internal.ShouldBe(t, expected, actual)
}

func Test_MinimumEditDistance_SameLength(t *testing.T) {
	actual := MinimumEditDistance([]byte("one"), []byte("two"))
	expected := []ot.Operation{
		{Type: ot.SUBSTITUTION, Position: 0, Character: 't'},
		{Type: ot.SUBSTITUTION, Position: 1, Character: 'w'},
		{Type: ot.SUBSTITUTION, Position: 2, Character: 'o'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_MinimumEditDistance_PreviousLonger1(t *testing.T) {
	actual := MinimumEditDistance([]byte("one1"), []byte("on"))
	expected := []ot.Operation{
		{Type: ot.DELETION, Position: 2, Character: 'e'},
		{Type: ot.DELETION, Position: 3, Character: '1'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_MinimumEditDistance_PreviousLonger2(t *testing.T) {
	actual := MinimumEditDistance([]byte("one1"), []byte("e1"))
	expected := []ot.Operation{
		{Type: ot.DELETION, Position: 0, Character: 'o'},
		{Type: ot.DELETION, Position: 1, Character: 'n'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_MinimumEditDistance_CurrentLonger1(t *testing.T) {
	actual := MinimumEditDistance([]byte("tw"), []byte("two2"))
	expected := []ot.Operation{
		{Type: ot.INSERTION, Position: 2, Character: 'o'},
		{Type: ot.INSERTION, Position: 2, Character: '2'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_MinimumEditDistance_CurrentLonger2(t *testing.T) {
	actual := MinimumEditDistance([]byte("o2"), []byte("two2"))
	expected := []ot.Operation{
		{Type: ot.INSERTION, Position: 0, Character: 't'},
		{Type: ot.INSERTION, Position: 0, Character: 'w'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_MinimumEditDistance_MixedOperations(t *testing.T) {
	actual := MinimumEditDistance([]byte("world_1"), []byte("word_2!"))
	expected := []ot.Operation{
		{Type: ot.DELETION, Position: 3, Character: 'l'},
		{Type: ot.INSERTION, Position: 6, Character: '2'},
		{Type: ot.SUBSTITUTION, Position: 6, Character: '!'},
	}

	internal.ShouldBe(t, expected, actual)
}
