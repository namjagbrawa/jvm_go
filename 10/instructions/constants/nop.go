package constants

import "github.com/namjagbrawa/jvm_go/instructions/base"
import "github.com/namjagbrawa/jvm_go/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
