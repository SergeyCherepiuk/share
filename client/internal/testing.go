package internal

import (
	"reflect"
	"testing"
)

func ShouldBe[T any](t *testing.T, expected, actual T) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Actual should be equal to expected\nexpected: %+v\nactual: %+v\n", expected, actual)
	}
}

func ShouldNotBe[T any](t *testing.T, expected, actual T) {
	if reflect.DeepEqual(expected, actual) {
		t.Errorf("Actual should not be equal to expected\nexpected: %+v\nactual: %+v\n", expected, actual)
	}
}
