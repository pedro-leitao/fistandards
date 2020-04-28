package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// Clean takes a string with a banking identifier and returns a normalized (extraneous characters removed) copy.
func Clean(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ToUpper(s)

	return s
}

// GetParams parses a given string with a grouped regex and returns the grouping as a map.
func GetParams(regEx, s string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(s)
	if match == nil {
		fmt.Printf("Did not match:<%v>", s)
	}

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}
