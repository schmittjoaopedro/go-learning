package is_subsequence

func IsSubsequence(s string, t string) bool {
	if len(t) < len(s) {
		return false
	}

	s_idx := 0
	t_idx := 0
	for s_idx < len(s) && t_idx < len(t) {
		if s[s_idx] == t[t_idx] {
			s_idx++
		}
		t_idx++
	}

	return s_idx == len(s)
}
