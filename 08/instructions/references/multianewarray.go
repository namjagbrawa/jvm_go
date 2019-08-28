package references

import "github.com/namjagbrawa/jvm_go/instructions/base"
import "github.com/namjagbrawa/jvm_go/rtda"
import "github.com/namjagbrawa/jvm_go/rtda/heap"

// Create new multidimensional array
type MULTI_ANEW_ARRAY struct {
	index      uint16
	dimensions uint8
}

func (self *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.index = reader.ReadUint16()
	self.dimensions = reader.ReadUint8()
}
func (self *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(uint(self.index)).(*heap.ClassRef)
	arrClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	counts := popAndCheckCounts(stack, int(self.dimensions))
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)
}

func popAndCheckCounts(stack *rtda.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()  // 右维度往左维度POP和存储
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}

	return counts
}

// 建立多维数组
func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := arrClass.NewArray(count)

	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}

	return arr
}
