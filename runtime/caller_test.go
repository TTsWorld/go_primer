package runtime_test

import (
	"fmt"
	"runtime"
	"testing"
)

// runtime.Caller 是 Go 语言 runtime 包中的关键函数，用于在运行时获取调用栈信息（如文件名、行号、函数名），常用于日志增强、错误追踪等场景。

/*
func Caller(skip int) (pc uintptr, file string, line int, ok bool)

	​skip 参数​：控制跳过栈帧的层级
	skip=0：返回 Caller 自身调用位置（文件、行号）。
	skip=1：返回调用 Caller 的直接上层函数的位置（最常用）。
	skip=N：返回向上跳过 N 层的调用者位置。

	​返回值​：
	pc：程序计数器，可通过 runtime.FuncForPC(pc) 获取函数名。
	file：源文件完整路径（需用 path/filepath.Base() 提取文件名）。
	line：调用点在文件中的行号。
	ok：是否成功获取信息（必须检查）
*/
func TestCaller(t *testing.T) {
	pcs := make([]uintptr, 10)
	n := runtime.Callers(0, pcs) // 从当前栈帧开始
	frames := runtime.CallersFrames(pcs[:n])
	for {
		frame, more := frames.Next()
		fmt.Printf("File: %s, Line: %d, Func: %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
		fmt.Println("--------------------------------")
	}

}
