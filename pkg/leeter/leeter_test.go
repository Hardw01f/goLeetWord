package leeter

import (
	"testing"
)

func TestDefaultLeeter(t *testing.T) {
	var testStrings []string
	testStrings = append(testStrings, "hey")
	result, err := DefaultLeeter(testStrings)
	if err != nil {
		t.Fatal(err)
	}

	if result != "h3y" {
		t.Fatalf("Test failed : expected->%s, autual->%s", "h3y", result)
	}

}
