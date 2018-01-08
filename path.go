package gconf

import "strings"

const (
	KeyDelimiter string = ":"
	RootPath     string = ""
)

func PathCombine(path ...string) string {
	if path == nil || len(path) == 0 {
		return ""
	}

	return strings.Join(path, KeyDelimiter)
}

func GetSectionKey(path string) string {
	if path == "" {
		return ""
	}

	idx := strings.LastIndex(path, KeyDelimiter)

	if idx == -1 {
		return path
	}

	return path[idx+1:]
}

func GetParentPath(path string) string {
	if path == "" {
		return ""
	}

	idx := strings.LastIndex(path, KeyDelimiter)

	if idx == -1 {
		return ""
	}

	return path[0:idx]
}

func HasPathInKey(path, key string) bool {
	return strings.HasPrefix(strings.ToLower(key), strings.ToLower(path))
}