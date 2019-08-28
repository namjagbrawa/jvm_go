package references

import "github.com/namjagbrawa/jvm_go/instructions/base"
import "github.com/namjagbrawa/jvm_go/rtda"
import "github.com/namjagbrawa/jvm_go/rtda/heap"

// Create new object
type NEW struct{ base.Index16Instruction }

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	// todo: init class

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
