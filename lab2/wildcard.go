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

	startIndex, sourceIndex, wildcardIndex, wasWildcard := 0, 0, 0, false
	wcardLength, sourceLength := len(wildcardTrimmed), len(sourceStr)
	for startIndex < sourceLength {
		wildcardIndex = 0
		sourceIndex = startIndex
		for wildcardIndex < wcardLength && 
		sourceIndex < sourceLength &&
		wildcardTrimmed[wildcardIndex] == sourceStr[sourceIndex] {
			wildcardIndex++
			sourceIndex++
			wasWildcard = false
			if (wildcardIndex < wcardLength && 
				wildcardTrimmed[wildcardIndex] == '*') {
				wasWildcard = true
				for wildcardIndex < wcardLength && 
				wildcardTrimmed[wildcardIndex] == '*' {
					wildcardIndex++
				}
			}

			if (wasWildcard && wildcardIndex < wcardLength) {
				for sourceIndex < sourceLength &&
				wildcardTrimmed[wildcardIndex] != sourceStr[sourceIndex] {
						sourceIndex++
					}
			}
		}

		if wildcardIndex >= wcardLength {
			if hasHeadingWildcard {
				foundValues = append(foundValues, 0)
			} else {
				foundValues = append(foundValues, startIndex)
			}
		}

		startIndex++
	}

	return
}
