package callstack

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// TraceStackOneLine 获取调用堆栈信息, 并合并成一行
func TraceStackOneLine() string {
	pcs := make([]uintptr, 16, 16)
	npc := runtime.Callers(2, pcs)

	frameContents := make([]string, 0, npc)
	foundRuntime := false
	endRuntime := false

	for i := 0; i < npc; i++ {
		frame := NewFrame(pcs[i])
		/**
		 * 跳过 recover 调用堆栈
		 *
		 * Note: 因 recover 过程总是从 runtime包开始的,所以只需显示 runtime 出现后的堆栈
		 */
		if !foundRuntime && !endRuntime {
			if strings.HasPrefix(frame.Func(), "runtime.") {
				foundRuntime = true
			}
			continue
		} else if foundRuntime != endRuntime {
			if !strings.HasPrefix(frame.Func(), "runtime.") {
				endRuntime = true
			} else {
				continue
			}
		}

		frameContents = append(frameContents,
			frame.Func()+
				"("+
				frame.File()+
				":"+
				strconv.FormatInt(frame.Line(), 10)+
				")",
		)
	}
	return strings.Join(frameContents, " -> ")
}

// CallFrame get specified number of stack frame
func CallFrame(n int64) (*Frame, error) {
	pcs := make([]uintptr, 16, 16)
	npc := runtime.Callers(2, pcs)

	if n >= int64(npc) {
		return nil, errors.New("out of range")
	}

	return NewFrame(pcs[n]), nil
}
