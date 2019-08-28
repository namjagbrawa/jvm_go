package control

import "github.com/namjagbrawa/jvm_go/instructions/base"
import "github.com/namjagbrawa/jvm_go/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
