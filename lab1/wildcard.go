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
		hasHeadingWildcard = true
	}

	startIndex, sourceIndex, wildcardIndex, wasWildcard, isSearchReset := 0, 0, 0, false, false
	wcardLength, sourceLength := len(wildcardTrimmed), len(sourceStr)
	// naїve algorithm
	// Example pattern: ex**p*e (example)
	for startIndex < sourceLength {

		// skip wildcards until the last one
		// guaranteed to have at least one non-wildcard at the end
		if wildcardTrimmed[wildcardIndex] == '*' {
			for wildcardTrimmed[wildcardIndex] == '*' {
				wildcardIndex++
			}
			wasWildcard = true
		}

		// update source index if it mathces the non-wildcard symbol
		// advance wildcard counter
		if sourceStr[sourceIndex] == wildcardTrimmed[wildcardIndex] {
			sourceIndex++
			wildcardIndex++
			wasWildcard = false
		} else if wasWildcard { // wildcard always matches
			sourceIndex++
		} else { // search complete, start anew
			isSearchReset = true
		}

		// if value found add to result
		if wildcardIndex >= wcardLength && !isSearchReset {
			if hasHeadingWildcard {
				foundValues = append(foundValues, 0)
			} else {
				foundValues = append(foundValues, startIndex)
			}
			isSearchReset = true
		}

		// reset (advance) search
		if isSearchReset || sourceIndex == sourceLength {
			wildcardIndex = 0
			startIndex++
			sourceIndex = startIndex
			isSearchReset = false
		}
	}

	return
}
