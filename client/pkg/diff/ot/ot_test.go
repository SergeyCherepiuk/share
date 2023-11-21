package ot

import (
	"testing"

	"github.com/SergeyCherepiuk/share/client/internal"
)

func Test_offset_DeletionBeforeOperation(t *testing.T) {
	actualDeletion := offset(
		Operation{Type: DELETION, Position: 0},
		Operation{Type: DELETION, Position: 1},
	)
	actualInsertion := offset(
		Operation{Type: DELETION, Position: 0},
		Operation{Type: INSERTION, Position: 1},
	)
	actualSubstitution := offset(
		Operation{Type: DELETION, Position: 0},
		Operation{Type: SUBSTITUTION, Position: 1},
	)

	internal.ShouldBe(t, -1, actualDeletion)
	internal.ShouldBe(t, -1, actualInsertion)
	internal.ShouldBe(t, -1, actualSubstitution)
}

func Test_offset_InsertionBeforeOperation(t *testing.T) {
	actualDeletion := offset(
		Operation{Type: INSERTION, Position: 0},
		Operation{Type: DELETION, Position: 1},
	)
	actualInsertion := offset(
		Operation{Type: INSERTION, Position: 0},
		Operation{Type: INSERTION, Position: 1},
	)
	actualSubstitution := offset(
		Operation{Type: INSERTION, Position: 0},
		Operation{Type: SUBSTITUTION, Position: 1},
	)

	internal.ShouldBe(t, 1, actualDeletion)
	internal.ShouldBe(t, 1, actualInsertion)
	internal.ShouldBe(t, 1, actualSubstitution)
}

func Test_offset_SubstitutionBeforeOperation(t *testing.T) {
	actualDeletion := offset(
		Operation{Type: SUBSTITUTION, Position: 0},
		Operation{Type: DELETION, Position: 1},
	)
	actualInsertion := offset(
		Operation{Type: SUBSTITUTION, Position: 0},
		Operation{Type: INSERTION, Position: 1},
	)
	actualSubstitution := offset(
		Operation{Type: SUBSTITUTION, Position: 0},
		Operation{Type: SUBSTITUTION, Position: 1},
	)

	internal.ShouldBe(t, 0, actualDeletion)
	internal.ShouldBe(t, 0, actualInsertion)
	internal.ShouldBe(t, 0, actualSubstitution)
}

func Test_Adjust_Nil(t *testing.T) {
	actual := Adjust(nil)

	internal.ShouldBe(t, []Operation{}, actual)
}

func Test_Adjust_Empty(t *testing.T) {
	actual := Adjust([]Operation{})

	internal.ShouldBe(t, []Operation{}, actual)
}

func Test_Adjust_OneOperation(t *testing.T) {
	actual := Adjust([]Operation{
		{Type: INSERTION, Position: 10, Character: 'a'},
	})
	expected := []Operation{
		{Type: INSERTION, Position: 10, Character: 'a'},
	}

	internal.ShouldBe(t, expected, actual)
}

func Test_Adjust_MultipleOperations(t *testing.T) {
	actual := Adjust([]Operation{
		{Type: INSERTION, Position: 10, Character: 'a'},
		{Type: DELETION, Position: 12},
		{Type: DELETION, Position: 13},
		{Type: SUBSTITUTION, Position: 20, Character: 'b'},
		{Type: DELETION, Position: 5},
	})
	expected := []Operation{
		{Type: INSERTION, Position: 10, Character: 'a'},
		{Type: DELETION, Position: 13},
		{Type: DELETION, Position: 13},
		{Type: SUBSTITUTION, Position: 19, Character: 'b'},
		{Type: DELETION, Position: 5},
	}

	internal.ShouldBe(t, expected, actual)
}
