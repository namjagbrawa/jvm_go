package reserved

import "github.com/namjagbrawa/jvm_go/instructions/base"
import "github.com/namjagbrawa/jvm_go/rtda"
import "github.com/namjagbrawa/jvm_go/native"
import _ "github.com/namjagbrawa/jvm_go/native/java/io"
import _ "github.com/namjagbrawa/jvm_go/native/java/lang"
import _ "github.com/namjagbrawa/jvm_go/native/java/security"
import _ "github.com/namjagbrawa/jvm_go/native/java/util/concurrent/atomic"
import _ "github.com/namjagbrawa/jvm_go/native/sun/io"
import _ "github.com/namjagbrawa/jvm_go/native/sun/misc"
import _ "github.com/namjagbrawa/jvm_go/native/sun/reflect"

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
