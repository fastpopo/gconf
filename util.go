package gconf

import "strings"

type StringSplitter struct {
	str string
}

func NewStringSplitter(str string) StringSplitter {
	return StringSplitter{str: str}
}

func (s StringSplitter) Split(sep string, removeIfEmpty bool) []string {
	var ret []string

	values := strings.Split(s.str, sep)

	for _, v := range values {
		if removeIfEmpty && v == "" {
			continue
		}

		ret = append(ret, v)
	}

	return ret
}
