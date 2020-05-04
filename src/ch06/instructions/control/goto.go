package control

import (
	"learn/ch06/instructions/base"
	"learn/ch06/rtda"
)

type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
