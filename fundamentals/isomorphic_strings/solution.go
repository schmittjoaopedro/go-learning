package isomorphic_strings

// Start 17:53
// End 18:10
// URL https://leetcode.com/problems/isomorphic-strings

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s2tDict := map[byte]byte{}
	t2sDict := map[byte]byte{}
	valid := true

	for i := 0; i < len(s); i++ {
		sChar := s[i]
		tChar := t[i]
		sMap, sOk := s2tDict[sChar]
		tMap, tOk := t2sDict[tChar]

		if !sOk && !tOk {
			s2tDict[sChar] = tChar
			t2sDict[tChar] = sChar
		} else if sOk && tOk {
			if sMap != tChar || tMap != sChar {
				valid = false
				break
			}
		} else {
			valid = false
			break
		}
	}
	return valid
}
