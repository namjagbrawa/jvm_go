package io

import "github.com/namjagbrawa/jvm_go/native"
import "github.com/namjagbrawa/jvm_go/rtda"

func init() {
	native.Register("sun/io/Win32ErrorMode", "setErrorMode", "(J)J", setErrorMode)
}

func setErrorMode(frame *rtda.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}
