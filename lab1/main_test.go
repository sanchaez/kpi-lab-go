package main

import (
	"reflect"
	"testing"
)

type testpair struct {
	source   string
	pattern  string
	expected []int
}

type testCases []testpair

var deadCases = testCases{
	{"", "", nil},
	{"no test", "", nil},
	{"", "no origin", nil},
	{"small", "bigger than original", nil},
}

var singleLetter = testCases{
	{"test", "e", []int{1}},
	{"test", "t", []int{0, 3}},
	{"test", "x", nil},
	{"abracadabra", "a", []int{0, 3, 5, 7, 10}},
	{"abracadabra", "x", nil},
	{"She called a storm upon this town.", "s", []int{13, 27}},
}

var multipleLetters = testCases{
	{"test", "es", []int{1}},
	{"test", "te", []int{0}},
	{"test", "test", []int{0}},
	{"abracadabra", "acad", []int{3}},
	{"She called a storm upon this town.", " t", []int{24, 29}},
	{"She called a storm upon this town.", "called", []int{4}},
}

func auxTestLoop(t *testing.T, tests *testCases) {
	for _, x := range *tests {
		result := wildcardMatch(x.source, x.pattern)
		if reflect.DeepEqual(result, x.expected) {
			t.Error(
				"	\ninput:", x.source,
				"	\npattern:", x.pattern,
				"	\nexpected:", x.expected,
				"	\nresult:", result,
			)
		}
	}
}

func TestSingleLetter(t *testing.T) {
	auxTestLoop(t, &singleLetter)
}

func TestMultipleLetters(t *testing.T) {
	auxTestLoop(t, &multipleLetters)
}

func TestDeadCases(t *testing.T) {
	auxTestLoop(t, &deadCases)
}
