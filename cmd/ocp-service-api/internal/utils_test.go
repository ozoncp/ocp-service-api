package utils

import (
	"reflect"
	"testing"
)

// SplitToBatches
func TestSplitToBatches(t *testing.T) {
	originalSlice := []string{"42", "qw", "3", "1", "qr", "rqw", "51"}
	expected := [][]string{{"42", "qw"}, {"3", "1"}, {"qr", "rqw"}, {"51"}}
	got := SplitToBatches(originalSlice, 2)

	isEqual := reflect.DeepEqual(got, expected)
	if isEqual == false {
		t.Errorf("not equal")
	}
}
func TestSplitToBatchesOneBatch(t *testing.T) {
	originalSlice := []string{"42", "qw", "3", "1", "qr", "rqw", "51"}
	expected := [][]string{{"42", "qw", "3", "1", "qr", "rqw", "51"}}
	got := SplitToBatches(originalSlice, 7)

	isEqual := reflect.DeepEqual(got, expected)
	if isEqual == false {
		t.Errorf("not equal")
	}
}
func TestSplitToBatchesEmpty(t *testing.T) {
	originalSlice := []string{}
	var expected [][]string
	got := SplitToBatches(originalSlice, 2)

	isEqual := reflect.DeepEqual(got, expected)

	if isEqual == false {
		t.Errorf("not equal")
	}
}
func TestSplitToBatchesOneNotFullBatch(t *testing.T) {
	originalSlice := []string{"42", "qw", "3", "1", "qr", "rqw", "51"}
	expected := [][]string{{"42", "qw", "3", "1", "qr", "rqw", "51"}}
	got := SplitToBatches(originalSlice, 15)

	isEqual := reflect.DeepEqual(got, expected)

	if isEqual == false {
		t.Errorf("not equal")
	}
}
func TestSplitToBatchesOneElement(t *testing.T) {
	originalSlice := []string{"42"}
	expected := [][]string{{"42"}}
	got := SplitToBatches(originalSlice, 1)

	isEqual := reflect.DeepEqual(got, expected)

	if isEqual == false {
		t.Errorf("not equal")
	}
}

// ReverseMapKeysAndValues

func TestReverseMapKeysAndValues(t *testing.T) {
	original := map[string]string{"a": "1", "b": "2", "c": "3"}
	expected := map[string]string{"1": "a", "2": "b", "3": "c"}
	got := ReverseMapKeysAndValues(original)

	isEqual := reflect.DeepEqual(got, expected)

	if isEqual == false {
		t.Errorf("not equal")
	}
}
func TestReverseMapKeysAndValuesEmpty(t *testing.T) {
	original := map[string]string{}
	expected := map[string]string{}
	got := ReverseMapKeysAndValues(original)

	isEqual := reflect.DeepEqual(got, expected)

	if isEqual == false {
		t.Errorf("not equal")
	}
}

// FilterOut

func TestFilterOut(t *testing.T) {
	original := []string{"Marv", "124q", "kreks"}
	expected := []string{"124q", "kreks"}
	got := FilterOut(original)

	isEqual := reflect.DeepEqual(got, expected)

	if isEqual == false {
		t.Errorf("not equal")
	}
}
