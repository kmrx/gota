package df

import (
	"reflect"
	"testing"
)

func TestDataFrame_New(t *testing.T) {
	a := New(Strings("b"), Ints(1, 2))
	if a.Err() == nil {
		t.Error("Expected error, got success")
	}
	a = New(Strings("b", "a"), NamedInts("Y", 1, 2), Floats(3.0, 4.0))
	if a.Err() != nil {
		t.Error("Expected success, got error")
	}
	expectedNames := []string{"X0", "Y", "X1"}
	receivedNames := a.colnames
	if !reflect.DeepEqual(expectedNames, receivedNames) {
		t.Error(
			"Expected Names:",
			expectedNames,
			"Received Names:",
			receivedNames,
		)
	}
	expectedTypes := []string{"string", "int", "float"}
	receivedTypes := a.coltypes
	if !reflect.DeepEqual(expectedTypes, receivedTypes) {
		t.Error(
			"Expected Types:",
			expectedTypes,
			"Received Types:",
			receivedTypes,
		)
	}
	// TODO: Check that df.colnames == columns.colnames
	// TODO: Check that the address of the columns are different that of the original series
	// TODO: Check that dimensions match
	a = New()
	if a.Err() == nil {
		t.Error("Expected error, got success")
	}
}

func TestDataFrame_Copy(t *testing.T) {
	a := New(NamedStrings("COL.1", "b", "a"), NamedInts("COL.2", 1, 2), NamedFloats("COL.3", 3.0, 4.0))
	b := a.Copy()
	if a.columns[0].Elements.(StringElements)[0] == b.columns[0].Elements.(StringElements)[0] {
		t.Error("Copy error: The memory address should be different even if the content is the same")
	}
}
