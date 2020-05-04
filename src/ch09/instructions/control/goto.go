package control

import (
	"learn/ch09/instructions/base"
	"learn/ch09/rtda"
)

type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
