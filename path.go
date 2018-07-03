package gconf

import "strings"

const (
	KeyDelimiter string = "/"
	RootPath     string = "/"
)

func PathCombine(path ...string) string {
	if path == nil || len(path) == 0 {
		return ""
	}

	if path[0] == RootPath {
		path = path[1:]
		return RootPath + strings.Join(path, KeyDelimiter)
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

func GetPaths(key string) []string {
	return strings.Split(key, KeyDelimiter)
}

func FindChildKeys(basePath string, keys []string) []string {
	if keys == nil || len(keys) == 0 {
		return nil
	}

	if basePath == RootPath {
		return keys
	}

	var subKeys []string

	for _, k := range keys {
		if HasPathInKey(basePath, k) {
			subKeys = append(subKeys, k)
		}
	}

	return subKeys
}