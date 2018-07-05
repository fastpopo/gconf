package gconf

import (
	"strings"
	"fmt"
	"strconv"
)

const (
	PathDelimiter  = "/"
	RootPath       = "/"
	ArrayDelimiter = "$"
)

func PathCombine(path ...string) string {
	if path == nil || len(path) == 0 {
		return ""
	}

	p := strings.Join(path, PathDelimiter)

	if p == "" {
		return ""
	}

	pathEntity := NewStringSplitter(p).Split(PathDelimiter, true)

	return RootPath + strings.Join(pathEntity, PathDelimiter)
}

func GetSectionKey(path string) string {
	if path == "" {
		return ""
	}

	if path == RootPath {
		return RootPath
	}

	idx := strings.LastIndex(path, PathDelimiter)

	if idx == -1 {
		return path
	}

	return path[idx+1:]
}

func GetParentPath(path string) string {
	if path == "" || path == RootPath {
		return ""
	}

	idx := strings.LastIndex(path, PathDelimiter)

	if idx == -1 || idx == 0 {
		return RootPath
	}

	return path[0:idx]
}

func HasPathInKey(path, key string) bool {
	return strings.HasPrefix(strings.ToLower(key), strings.ToLower(path))
}

func GetPaths(key string) []string {
	paths := make([]string, 0)
	paths = append(paths, RootPath)
	paths = append(NewStringSplitter(key).Split(PathDelimiter, true))
	return paths
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

func FindChildPairs(basePath string, pairs []KeyValuePair) []KeyValuePair {
	if pairs == nil || len(pairs) == 0 {
		return nil
	}

	if basePath == RootPath {
		return pairs
	}

	var subPairs []KeyValuePair

	for _, p := range pairs {
		if HasPathInKey(basePath, p.Key) {
			subPairs = append(subPairs, p)
		}
	}

	return subPairs
}

func IsArrayIndex(path string) bool {
	sectionKey := GetSectionKey(path)

	if sectionKey == "" {
		return false
	}

	if len(sectionKey) >= 2 {
		return false
	}

	idxStr := sectionKey[1:]

	_, err := strconv.Atoi(idxStr)
	if err != nil {
		return false
	}

	return true
}

func IsArrayPath(path string, keys []string) bool {
	if path == "" {
		return false
	}

	if keys == nil || len(keys) == 0 {
		return false
	}

	checkIdxString := GetArrayIndexPath(path, 0)

	if checkIdxString == "" {
		return false
	}

	for _, k := range keys {
		if HasPathInKey(checkIdxString, k) {
			return true
		}
	}

	return false
}

func GetArrayIndex(idx int) string {
	if idx < 0 {
		return ""
	}

	return ArrayDelimiter + fmt.Sprint(idx)
}

func GetArrayIndexPath(path string, idx int) string {
	if path == "" {
		return ""
	}

	if idx < 0 {
		return ""
	}

	return PathCombine(path, GetArrayIndex(idx))
}

func LengthOfArrayPath(path string, keys []string) int {
	maxIndex := -1

	if path == "" {
		return maxIndex
	}

	if keys == nil || len(keys) == 0 {
		return maxIndex
	}

	checkIdxString := PathCombine(path, ArrayDelimiter)

	for _, k := range keys {
		if HasPathInKey(checkIdxString, k) {
			index := parseIndexOfArray(checkIdxString, k)

			if index > maxIndex {
				maxIndex = index
			}
		}
	}

	if maxIndex == -1 {
		return -1
	}

	// increase size for it's length not index
	return maxIndex + 1
}

func parseIndexOfArray(checkIdxString, key string) int {
	newKey := strings.Replace(key, checkIdxString, "", -1)

	if newKey == "" {
		return -1
	}

	idx := strings.Index(newKey, PathDelimiter)

	if idx != -1 {
		newKey = newKey[:idx]
	}

	findIdx, err := strconv.Atoi(newKey)
	if err != nil {
		return -1
	}

	return findIdx
}