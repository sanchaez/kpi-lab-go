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
	{"She called a storm upon this town.", " t", []int{23, 28}},
	{"She called a storm upon this town.", "called", []int{4}},
}

var wildcardSingle = testCases{
	{"test", "*es", []int{0}},
	{"test", "te*", []int{0}},
	{"test", "t*t", []int{0}},
	{"abracadabra", "ac**ad", []int{3}},
	{"abracadabra", "a*a", []int{0, 3, 5, 7}},
	{"She called a storm upon this town.", "st*m", []int{13}},
	{"She called a storm upon this town.", "***st*m", []int{0}},
	{"test", "*ak", nil},
	{"test", "dd*", nil},
	{"test", "t*dt", nil},
}

var wildcardMultiple = testCases{
	{"abracadabra", "ac*o*ad", nil},
	{"abracadabra", "a*b*a", []int{0, 3, 5, 7}},
	{"She called a storm upon this town.", "a*st*m", []int{5, 11}},
}

func auxTestLoop(t *testing.T, tests *testCases) {
	for _, x := range *tests {
		result := Match(x.source, x.pattern)
		if !reflect.DeepEqual(result, x.expected) {
			t.Errorf(
				"\ninput:		%#v\n"+
					"pattern:	%#v\n"+
					"expected:	%#v\n"+
					"result:		%#v\n", x.source, x.pattern, x.expected, result,
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

func TestWildcardSingle(t *testing.T) {
	auxTestLoop(t, &wildcardSingle)
}

func TestWildcardMultiple(t *testing.T) {
	auxTestLoop(t, &wildcardMultiple)
}
