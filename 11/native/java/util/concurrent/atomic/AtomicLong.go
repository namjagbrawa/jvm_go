package atomic

import "github.com/namjagbrawa/jvm_go/native"
import "github.com/namjagbrawa/jvm_go/rtda"

func init() {
	native.Register("java/util/concurrent/atomic/AtomicLong", "VMSupportsCS8", "()Z", vmSupportsCS8)
}

func vmSupportsCS8(frame *rtda.Frame) {
	frame.OperandStack().PushBoolean(false)
}
