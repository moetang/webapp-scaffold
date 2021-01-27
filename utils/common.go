package utils

import "testing"

func AssertTrue(t *testing.T, j bool) {
	if !j {
		t.Fatal("assertTrue failed")
	}
}

func AssertFalse(t *testing.T, j bool) {
	if j {
		t.Fatal("assertFalse failed")
	}
}

func AssertEquals(t *testing.T, a, b interface{}) {
	if a != b {
		t.Fatal("assertEqual fail")
	}
}
