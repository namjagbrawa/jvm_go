package lang

import "github.com/namjagbrawa/jvm_go/native"
import "github.com/namjagbrawa/jvm_go/rtda"
import "github.com/namjagbrawa/jvm_go/rtda/heap"

const jlString = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
}

// public native String intern();
// ()Ljava/lang/String;
func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()  // 获得的是java的string
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
