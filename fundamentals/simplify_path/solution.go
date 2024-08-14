package simplify_path

// Start 14:05
// End 14:18
// https://leetcode.com/problems/simplify-path

import (
	"strings"
)

func SimplifyPath(path string) string {
	originalPath := strings.Split(path, "/")
	var canonicalPath []string

	for _, segment := range originalPath {
		if segment == "." || segment == "" {
			continue
		} else if segment == ".." {
			if len(canonicalPath) > 0 {
				canonicalPath = canonicalPath[:len(canonicalPath)-1]
			} else {
				continue
			}
		} else {
			canonicalPath = append(canonicalPath, segment)
		}
	}
	return "/" + strings.Join(canonicalPath, "/")
}
