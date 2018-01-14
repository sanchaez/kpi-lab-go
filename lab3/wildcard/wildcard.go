package wildcard

import (
	"sort"
	"strings"
	"sync"
)

func hasMatched(sourceStr, wildcardStr string, startIndex int) bool {
	sourceIndex, wildcardIndex, wasWildcard := startIndex, 0, false
	wcardLength, sourceLength := len(wildcardStr), len(sourceStr)

	// search
	wildcardIndex = 0

	// loop to find an entry starting from startIndex
	for wildcardIndex < wcardLength &&
		sourceIndex < sourceLength &&
		wildcardStr[wildcardIndex] == sourceStr[sourceIndex] {

		wildcardIndex++
		sourceIndex++

		// skip wildcards until the last one
		// guaranteed to have at least one non-wildcard at the end
		wasWildcard = false
		if wildcardIndex < wcardLength &&
			wildcardStr[wildcardIndex] == '*' {
			wasWildcard = true
			for wildcardIndex < wcardLength &&
				wildcardStr[wildcardIndex] == '*' {
				wildcardIndex++
			}
		}

		// wildcardIndex points to non-wildcard character in wildcardTrimmed
		// loop the string until it is found in sourceStr
		if wasWildcard && wildcardIndex < wcardLength {
			for sourceIndex < sourceLength &&
				wildcardStr[wildcardIndex] != sourceStr[sourceIndex] {
				sourceIndex++
			}
		}
	}

	// if value found add to results
	if wildcardIndex >= wcardLength {
		return true
	}

	return false
}

func readyWildcardString(str *string) bool {
	*str = strings.TrimRight(*str, "*")
	if strings.HasPrefix(*str, "*") {
		*str = strings.TrimLeft(*str, "*")
		return true
	}
	return false
}

type matchRoutine func(string, string, int, chan int, *sync.WaitGroup)

func matchGoroutine(sourceStr, wildcardStr string,
	startIndex int,
	channel chan int,
	wg *sync.WaitGroup) {
	if hasMatched(sourceStr, wildcardStr, startIndex) {
		channel <- startIndex
	}

	wg.Done()
}

func matchGoroutineWcard(sourceStr, wildcardStr string,
	startIndex int,
	channel chan int,
	wg *sync.WaitGroup) {
	if hasMatched(sourceStr, wildcardStr, startIndex) {
		channel <- 0
	}

	wg.Done()
}

//Match does string matching with `*` wildcard support.
//Returns a slice with N positions of all matched wildcardStr substrings in sourceStr.
//Trailing wildcards `str***` are ignored.
//With heading wildcards `***str` position 1 will be returned N times.
func Match(sourceStr, wildcardStr string) (foundValues []int) {
	hasHeadingWildcard := readyWildcardString(&wildcardStr)
	foundValues = nil
	//var sem = make(chan struct{}, 100)

	if wildcardStr == "" {
		foundValues = []int{0}
		return
	}

	if len(wildcardStr) > len(sourceStr) {
		return
	}

	sourceLength := len(sourceStr)

	resultChannel := make(chan int, sourceLength)
	wg := sync.WaitGroup{}
	var fn matchRoutine

	if hasHeadingWildcard {
		fn = matchGoroutineWcard
	} else {
		fn = matchGoroutine
	}

	wg.Add(sourceLength)

	for index := 0; index < sourceLength; index++ {
		go fn(sourceStr, wildcardStr, index, resultChannel, &wg)
	}

	wg.Wait()
	close(resultChannel)

	for value := range resultChannel {
		foundValues = append(foundValues, value)
	}

	sort.Ints(foundValues)
	return
}
