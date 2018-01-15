package wildcard

import (
	"sort"
)

func hasMatched(sourceStr, wildcardStr string) bool {
	sourceIndex, wildcardIndex := 0, 0
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

			// wildcardIndex points to non-wildcard character in wildcardTrimmed
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
	//var sem = make(chan struct{}, 100)

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

	// make in channel
	sourceLength := len(sourceStr)
	in := make(chan int, sourceLength/8)

	go func() {
		for i := 0; i < sourceLength; i++ {
			if sourceStr[i] == wildcardStr[0] {
				in <- i
			}
		}

		close(in)
	}()

	out := make(chan int, sourceLength)
	done := make(chan bool)

	// spawn parallel workers
	const parallelWorkers int = 4
	for p := 0; p < parallelWorkers; p++ {
		go func() {
			// Search the substrings, starting from v
			for v := range in {
				fn(sourceStr, wildcardStr, v, out)
			}
			// Denote completion of this goroutine.
			done <- true
		}()
	}

	go func() {
		// receive P values from done channel
		for p := 0; p < parallelWorkers; p++ {
			<-done
		}
		// At this point we know that all worker goroutines has completed
		// and no more data is coming, so close the output channel.
		close(out)
	}()

	var foundValues []int
	for value := range out {
		foundValues = append(foundValues, value)
	}

	sort.Ints(foundValues)
	return foundValues
}
