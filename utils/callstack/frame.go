package callstack

import (
	"runtime"
)

// Frame in trace stack
type Frame struct {
	pc uintptr

	fn *runtime.Func
}

// NewFrame constructor
func NewFrame(pc uintptr) *Frame {

	return &Frame{
		pc: pc - 1,
		fn: runtime.FuncForPC(pc - 1),
	}
}

// Func 获取函数名称
func (f *Frame) Func() string {
	if f.fn == nil {
		return "unknown"
	}
	return f.fn.Name()
}

// File 获取文件名
func (f *Frame) File() string {
	if f.fn == nil {
		return "unknown"
	}
	file, _ := f.fn.FileLine(f.pc)
	return file
}

// Line 获取行数
func (f *Frame) Line() int64 {
	if f.fn == nil {
		return 0
	}
	_, line := f.fn.FileLine(f.pc)
	return int64(line)
}
