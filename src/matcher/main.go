package matcher

import (
	"regexp"
	"strings"
)

func MatchPaths(expected []string, actual []string) []string {
	var foundedFiles []string

	for _, item := range actual {
		if !contains(expected, item) {
			continue
		}

		foundedFiles = append(foundedFiles, item)
	}

	return calculateDifference(expected, foundedFiles)
}

func MatchStructure(expected []string, actual []string) []string {
	var notMatching []string

	for _, item := range actual {
		if findAnyMatching(expected, item) {
			continue
		}

		notMatching = append(notMatching, item)
	}

	return notMatching
}

func findAnyMatching(allowed []string, path string) bool {
	for _, item := range mapWildcards(allowed) {
		r, _ := regexp.Compile(item)

		if r.MatchString(path) {
			return true
		}
	}

	return false
}

func mapWildcards(input []string) []string {
	var result []string

	for _, item := range input {
		// Adds support for ** as wildcard
		var pattern = strings.Replace(item, "**", "\\S*", -1)

		// Adds support for * as wildcard for single file *.extension
		pattern = strings.Replace(item, "*", "\\S*", -1)

		result = append(result, pattern)
	}

	return result
}

/*
Finds items that are present in slice A, but missing in slice B
*/
func calculateDifference(a []string, b []string) []string {
	var missing []string

	for _, item := range a {
		if contains(b, item) {
			continue
		}

		missing = append(missing, item)
	}

	return missing
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
