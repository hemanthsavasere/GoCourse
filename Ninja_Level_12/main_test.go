package main

import (
	"GoCourse/Ninja_Level_12/dog"
	"testing"
)

func myTest(t *testing.T) {
	type test struct {
		data   int
		answer int

	tests := []test{}
	tests = make([]test, 10)

	for i := 0; i < 10; i++ {
		tests[i].data = i + 2
		tests[i].answer = (i + 1) * 7
	}

	for _, v := range tests {
		x := dog.Years(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}
	}

}
