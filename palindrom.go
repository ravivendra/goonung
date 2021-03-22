package goonung

import (
	"errors"
	"strings"
)

// Palindrom : check string is palindrom or not
func Palindrom(texts []string) (map[string]bool, error) {
	var idx, jdx int
	var flag bool
	flags := make(map[string]bool)

	if len(texts) == 0 {
		return flags, errors.New("empty parameter")
	}

	for _, text := range texts {
		text = strings.ToLower(text)

		idx = 0
		jdx = len(text) - 1

		for idx < jdx {
			if text[idx] != text[jdx] {
				flag = false
			} else {
				flag = true
			}

			idx++
			jdx--
		}

		flags[text] = flag
	}

	return flags, nil
}
