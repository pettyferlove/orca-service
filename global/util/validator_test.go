package util

import (
	"testing"
)

type TestStruct struct {
	Name string `validate:"required"`
}

func TestValidator(t *testing.T) {
	err := InitTranslator("en")
	if err != nil {
		t.Errorf("Failed to initialize translator: %v", err)
	}

	testStruct := TestStruct{Name: ""}
	ok, errMsg := Validate(testStruct)
	if ok {
		t.Errorf("Expected validation to fail, but it passed. Error message: %s", errMsg)
	}

	testStruct.Name = "Test"
	ok, errMsg = Validate(testStruct)
	if !ok {
		t.Errorf("Expected validation to pass, but it failed. Error message: %s", errMsg)
	}
}
