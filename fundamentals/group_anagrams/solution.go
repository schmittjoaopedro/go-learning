package group_anagrams

// Start 18:37
// End 18:47
// URL https://leetcode.com/problems/group-anagrams

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	groupsDict := map[string][]string{}
	for _, str := range strs {
		key := sortWord(str)
		group, ok := groupsDict[key]
		if ok {
			groupsDict[key] = append(group, str)
		} else {
			groupsDict[key] = []string{str}
		}
	}
	groups := [][]string{}
	for _, value := range groupsDict {
		groups = append(groups, value)
	}
	return groups
}

func sortWord(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
