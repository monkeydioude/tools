package tools

import (
	"testing"
)

func TestIFailOnIncorrectPattern(t *testing.T) {
	v, err := MatchAndFind("[[]", "")

	if err == nil || v != nil {
		t.Error("Err must be returned")
	}
}

func TestIFailIfStringDoesNotMatchPattern(t *testing.T) {
	v, err := MatchAndFind("t", "f")

	if err == nil || v != nil {
		t.Error("Err must be returned")
	}
}

func TestIMatchAndRetrieveTargetParts(t *testing.T) {
	trial, err := MatchAndFind("f(.+)", "ffinal fantasy")
	goal := []string{"ffinal fantasy", "final fantasy"}

	if err != nil || len(trial) <= 0 {
		t.Error("Err should not be returned or trial must no be empty")
	}

	if goal[0] != trial[0] || goal[1] != trial[1] {
		t.Error("pattern and target do not match")
	}
}
