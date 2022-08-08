package utils

import (
	"fem-intro-to-go/05_toolkit/code/exercise_5a_solution/utils"
	"testing"
)

func TestAdd(t *testing.T) {
	expectedValue := 4
	actualValue := utils.Add(2, 2)
	if expectedValue != actualValue {
		t.Errorf("Expected Value did not match! Expected %d Actual %d", expectedValue, actualValue)
	}
}
