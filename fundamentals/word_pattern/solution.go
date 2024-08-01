package word_pattern

// Start 18:15
// End 18:21
// URL https://leetcode.com/problems/word-pattern

import (
	"strings"
)

func WordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	valid := true
	p2wDict := map[byte]string{}
	w2pDict := map[string]byte{}

	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		w := words[i]
		pMap, pOk := p2wDict[p] // string
		wMap, wOk := w2pDict[w] // byte
		if !pOk && !wOk {
			p2wDict[p] = w
			w2pDict[w] = p
		} else if pOk && wOk {
			if p != wMap && w != pMap {
				valid = false
				break
			}
		} else {
			valid = false
		}
	}
	return valid
}
