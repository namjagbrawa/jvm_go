package misc

import "github.com/namjagbrawa/jvm_go/instructions/base"
import "github.com/namjagbrawa/jvm_go/native"
import "github.com/namjagbrawa/jvm_go/rtda"

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

// private static native void initialize();
// ()V
func initialize(frame *rtda.Frame) {
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}
