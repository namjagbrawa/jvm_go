package rtda

import "github.com/namjagbrawa/jvm_go/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
