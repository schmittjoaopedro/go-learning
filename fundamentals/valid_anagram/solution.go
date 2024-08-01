package valid_anagram

// Start 18:24
// End 18:34
// URL https://leetcode.com/problems/valid-anagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := map[byte]int{}

	for i := 0; i < len(s); i++ {
		count, ok := sMap[s[i]]
		if ok {
			sMap[s[i]] = count + 1
		} else {
			sMap[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		count, ok := sMap[t[i]]
		if ok {
			if count == 0 {
				return false
			} else {
				sMap[t[i]] = count - 1
			}
		} else {
			return false
		}
	}

	return true
}
