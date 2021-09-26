package goenv

import "testing"

// TestGetModulePath ...
func TestGetModulePath(t *testing.T) {
	path, err := GetModulePath()
	if err != nil {
		t.Fatalf("GetModulePath fail")
	}
	t.Logf("path=%v", path)
}
