package valid_palindrome

// https://leetcode.com/problems/valid-palindrome

func IsPalindrome(s string) bool {

	head := 0
	tail := len(s) - 1
	palindrome := true
	for head < tail {
		if !isAlphanumeric(s[head]) {
			head++
		} else if !isAlphanumeric(s[tail]) {
			tail--
		} else {
			if !isEqual(s[head], s[tail]) {
				palindrome = false
			}
			head++
			tail--
		}
	}

	return palindrome
}

func isAlphanumeric(c byte) bool {
	return ("A"[0] <= c && c <= "Z"[0]) ||
		("a"[0] <= c && c <= "z"[0]) ||
		("0"[0] <= c && c <= "9"[0])
}

func isEqual(c1 byte, c2 byte) bool {
	// Convert upper case to lower case by offseting the
	// upper case character to the lower case space
	if "A"[0] <= c1 && c1 <= "Z"[0] {
		c1 = c1 + ("a"[0] - "A"[0])
	}
	if "A"[0] <= c2 && c2 <= "Z"[0] {
		c2 = c2 + ("a"[0] - "A"[0])
	}
	return c1 == c2
}
