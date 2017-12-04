package main

import "strings"

//Match does string matching with `*` wildcard support.
//Returns a slice with N positions of all matched wildcardStr substrings in sourceStr.
//Trailing wildcards `str***` are ignored.
//With heading wildcards `***str` position 1 will be returned N times.
func Match(sourceStr, wildcardStr string) (foundValues []int) {
	startIndex, sourceIndex, wildcardIndex, hasHeadingWildcard := 0, 0, 0, false
	foundValues = nil

	//if wildcard is empty pattern match always fails
	if wildcardStr == "" {
		return
	}

	//remove trailing *
	wildcardTrimmed := strings.TrimLeft(wildcardStr, "*")
	// if nothing left pattern always matches all the string
	if wildcardTrimmed == "" {
		foundValues = append(foundValues, 0)
		return
	}

	if strings.HasPrefix(wildcardTrimmed, "*") {
		wildcardTrimmed = strings.TrimRight(wildcardTrimmed, "*")
		hasHeadingWildcard = true //TODO: do 2 versions of the algorithm
	}
	//TODO: add wildcard-in-the-middle case here

	// naїve algorithm
	for sourceIndex < len(sourceStr) {
		if sourceStr[sourceIndex] == wildcardTrimmed[wildcardIndex] {
			// symbol matched
			if wildcardIndex == len(wildcardTrimmed)-1 {
				// pattern matched!
				if hasHeadingWildcard {
					foundValues = append(foundValues, 0)
				}
				foundValues = append(foundValues, startIndex)
				wildcardIndex = 0
				startIndex++
				sourceIndex = startIndex
			} else {
				// pattern not matched yet
				wildcardIndex++
				sourceIndex++
			}
		} else if wildcardIndex != 0 {
			// were matching previously
			// reset wcard AND source indexes
			wildcardIndex = 0
			startIndex++
			sourceIndex = startIndex
		} else {
			// not matching yet, just advance
			startIndex++
			sourceIndex++
		}
	}

	return
}
