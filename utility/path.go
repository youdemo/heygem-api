package util

import "github.com/gogf/gf/v2/text/gstr"

const Separator = "/"

func JoinPaths(paths ...string) string {
	var s string
	for _, path := range paths {
		if s != "" {
			s += Separator
		}
		s += gstr.TrimRight(path, Separator)
	}
	return s
}
