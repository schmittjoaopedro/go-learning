package ransom_note

// Start 17:43
// End 17:49
// URL https://leetcode.com/problems/ransom-note

func CanConstruct(ransomNote string, magazine string) bool {
	dict := [26]int{}
	for i := 0; i < len(magazine); i++ {
		pos := int(magazine[i]) - int('a')
		dict[pos]++
	}
	for i := 0; i < len(ransomNote); i++ {
		pos := int(ransomNote[i]) - int('a')
		dict[pos]--
		if dict[pos] < 0 {
			return false
		}
	}

	return true
}
