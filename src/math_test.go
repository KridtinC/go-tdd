package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testSum struct {
	arg1, arg2, result int
}

func TestSum(t *testing.T) {

	var testScenario = []testSum{
		{1, 2, 3},
		{-1, 2, 1},
		{0, 1, 1},
		{1, -4, -3},
	}

	for _, testcase := range testScenario {
		var actual = Sum(testcase.arg1, testcase.arg2)
		assert.Equal(t, testcase.result, actual)
	}

}
