package longest_substring_without_repeating_characters

// https://leetcode.com/problems/longest-substring-without-repeating-characters
/*
Start: 18:50
End: 19:40
*/

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	dict := map[byte]int{} // {a:0, b:2, c:0}
	p1, p2 := 0, 0         // 7, 7
	length := 1            // 3
	lastSaw := s[p1]       // b
	dict[lastSaw] = 1
	diff := true // false
	for {
		if diff {
			p2++
			if p2 >= len(s) {
				break
			}
			c := s[p2]        // b
			k := key(dict, c) // 1
			if k == 0 {
				length = max(length, p2-p1+1)
			} else {
				diff = false
				lastSaw = c
			}
			dict[c]++
		} else {
			c := s[p1] // b
			p1++
			dict[c]--
			if lastSaw == c {
				diff = true
			}
		}
	}

	return length
}

func key(dict map[byte]int, key byte) int {
	i, ok := dict[key]
	if ok {
		return i
	} else {
		dict[key] = 0
		return 0
	}
}
