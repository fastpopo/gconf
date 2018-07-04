package gconf

import (
	"testing"
)

func TestPathCombine(t *testing.T) {
	returned := PathCombine(RootPath, RootPath, RootPath)
	expected := RootPath
	if returned != expected {
		t.Errorf("path combine with only root path, expected: %s, returned: %s", expected, returned)
	}

	returned = PathCombine("", "", "")
	expected = RootPath
	if returned != expected {
		t.Errorf("path combine with invalid format, expected: %s, returned: %s", expected, returned)
	}

	returned = PathCombine(RootPath, RootPath, "/test/", "/value")
	expected = "/test/value"
	if returned != expected {
		t.Errorf("path combine with various format, expected: %s, returned: %s", expected, returned)
	}

	returned = PathCombine(RootPath, "", "/test/", "/value/")
	expected = "/test/value"
	if returned != expected {
		t.Errorf("path combine with various format, expected: %s, returned: %s", expected, returned)
	}
}