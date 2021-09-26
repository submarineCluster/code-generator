package nameutil

import (
	"strings"
	"testing"
)

// TestNameTitle ...
func TestNameTitle(t *testing.T) {
	name := "resourceName"
	t.Logf("title: %v", strings.Title(name))
}
