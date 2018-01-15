package wildcard

import (
	"sort"
)

func hasMatched(sourceStr, wildcardStr string) bool {
	var sourceIndex, wildcardIndex int
	wcardLength, sourceLength := len(wildcardStr), len(sourceStr)

	// search
	// loop to find an entry starting from startIndex
	for wildcardIndex < wcardLength &&
		sourceIndex < sourceLength &&
		wildcardStr[wildcardIndex] == sourceStr[sourceIndex] {

		wildcardIndex++
		sourceIndex++

		// skip wildcards until the last one
		// guaranteed to have at least one non-wildcard at the end
		if wildcardIndex < wcardLength &&
			wildcardStr[wildcardIndex] == '*' {
			for wildcardIndex < wcardLength &&
				wildcardStr[wildcardIndex] == '*' {
				wildcardIndex++
			}

			// wildcardIndex points to non-wildcard character in wildcardStr
			// loop the string until it is found in sourceStr
			if wildcardIndex < wcardLength {
				for sourceIndex < sourceLength &&
					wildcardStr[wildcardIndex] != sourceStr[sourceIndex] {
					sourceIndex++
				}
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
	length := len(*str)
	begin, end := 0, length
	for i := 0; i < length && (*str)[i] == '*'; i++ {
		begin++
	}
	for i := length - 1; i >= 0 && (*str)[i] == '*'; i-- {
		end--
	}
	if begin <= end {
		*str = (*str)[begin:end]
	} else {
		*str = ""
	}

	if begin > 0 {
		return true
	}

	return false
}

type matchRoutine func(string, string, int, chan int)

func matchGoroutine(sourceStr, wildcardStr string,
	startIndex int,
	channel chan int) {
	if hasMatched(sourceStr[startIndex:], wildcardStr) {
		channel <- startIndex
	}
}

func matchGoroutineWcard(sourceStr, wildcardStr string,
	startIndex int,
	channel chan int) {
	if hasMatched(sourceStr[startIndex:], wildcardStr) {
		channel <- 0
	}
}

//Match does string matching with `*` wildcard support.
//Returns a slice with N positions of all matched wildcardStr substrings in sourceStr.
//Trailing wildcards `str***` are ignored.
//With heading wildcards `***str` position 1 will be returned N times.
func Match(sourceStr, wildcardStr string) []int {
	hasHeadingWildcard := readyWildcardString(&wildcardStr)

	if wildcardStr == "" {
		return []int{0}
	}

	if len(wildcardStr) > len(sourceStr) {
		return nil
	}

	var fn matchRoutine
	if hasHeadingWildcard {
		fn = matchGoroutineWcard
	} else {
		fn = matchGoroutine
	}

	// make channels
	sourceLength := len(sourceStr)
	const parallelWorkers int = 4
	in := make(chan int, parallelWorkers)
	out := make(chan int, sourceLength)
	done := make(chan bool, parallelWorkers)

	// spawn producer
	go func() {
		for i := 0; i < sourceLength; i++ {
			if sourceStr[i] == wildcardStr[0] {
				in <- i
			}
		}

		close(in)
	}()

	// spawn consumers
	for p := 0; p < parallelWorkers; p++ {
		go func() {
			// search the substrings in v
			for v := range in {
				fn(sourceStr, wildcardStr, v, out)
			}
			// denote completion of this goroutine
			done <- true
		}()
	}

	// spawn waiter
	go func() {
		// wait for workers
		for p := 0; p < parallelWorkers; p++ {
			<-done
		}

		// finish workers
		close(out)
	}()

	// capture values as they come
	var foundValues []int
	for value := range out {
		foundValues = append(foundValues, value)
	}

	sort.Ints(foundValues)
	return foundValues
}
