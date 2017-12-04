package main

import "strings"

//Match does string matching with `*` wildcard support.
//Returns a slice with N positions of all matched wildcardStr substrings in sourceStr.
//Trailing wildcards `str***` are ignored.
//With heading wildcards `***str` position 1 will be returned N times.
func Match(sourceStr, wildcardStr string) (foundValues []int) {
	hasHeadingWildcard := false
	foundValues = nil

	//if wildcard is empty pattern match always fails
	if wildcardStr == "" {
		return
	}

	//remove trailing *
	wildcardTrimmed := strings.TrimRight(wildcardStr, "*")
	// if nothing left pattern always matches all the string
	if wildcardTrimmed == "" {
		foundValues = []int{0}
		return
	}

	if strings.HasPrefix(wildcardTrimmed, "*") {
		wildcardTrimmed = strings.TrimLeft(wildcardTrimmed, "*")
		hasHeadingWildcard = true //TODO: do 2 versions of the algorithm
	}

	startIndex, sourceIndex, wildcardIndex, wasWildcard, isSearchReset := 0, 0, 0, false, false
	wcardLength, sourceLength := len(wildcardTrimmed), len(sourceStr)
	// naїve algorithm
	// Example pattern: ex**p*e (example)
	for startIndex < sourceLength {
		if wildcardTrimmed[wildcardIndex] == '*' {
			// skip wildcards until the last one
			// guaranteed to have at least one non-wildcard at the end
			for wildcardTrimmed[wildcardIndex] == '*' {
				wildcardIndex++
			}
			wasWildcard = true
		}

		if sourceStr[sourceIndex] == wildcardTrimmed[wildcardIndex] {
			sourceIndex++
			wildcardIndex++
			wasWildcard = false
		} else if wasWildcard {
			sourceIndex++
		} else {
			isSearchReset = true
		}

		if wildcardIndex >= wcardLength && !isSearchReset {
			if hasHeadingWildcard {
				foundValues = append(foundValues, 0)
			} else {
				foundValues = append(foundValues, startIndex)
			}
			isSearchReset = true
		}

		if isSearchReset || sourceIndex == sourceLength {
			wildcardIndex = 0
			startIndex++
			sourceIndex = startIndex
			isSearchReset = false
		}
	}

	return
}
