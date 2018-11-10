package function

import "strings"

// Title returns a copy of the string s with all Unicode letters that begin words
// mapped to their title case.
//go:generate basil generate Title
func Title(s string) string {
	return strings.Title(s)
}