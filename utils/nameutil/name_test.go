package nameutil

import (
	"strings"
	"testing"
)

func TestNameTitle(t *testing.T) {
	name := "resourceName"
	t.Logf("title: %v", strings.Title(name))
}
