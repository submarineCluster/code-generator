package callstack_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"git.code.oa.com/leoshli/code-generator/utils/callstack"
)

func AA() {
	frame, err := callstack.CallFrame(0)
	if err != nil {
		return
	}

	fmt.Printf("%v:%v", filepath.Dir(frame.File()), frame.Line())

	frame.File()
}

func TestCallFrame(t *testing.T) {
	AA()
}
