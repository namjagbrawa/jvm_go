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
	if !class.InitStarted() {
		frame.RevertNextPC()  // 从Reader的下一个指令的位置回到当前指令的位置，下次等初始化完类后，再来实例化对象，重新执行一次，修改frame的下一步指令，还是指向自己的
		base.InitClass(frame.Thread(), class)
		return
	}

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
