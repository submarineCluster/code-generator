package callstack_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"git.code.oa.com/tencent_abtest/code-generator/utils/callstack"
)

// AA ...
func AA() {
	// CallFrame ...
	frame, err := callstack.CallFrame(0)
	if err != nil {
		return
	}

	fmt.Printf("%v:%v", filepath.Dir(frame.File()), frame.Line())

	frame.File()
}

// TestCallFrame ...
func TestCallFrame(t *testing.T) {
	AA()
}
