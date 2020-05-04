package control

import (
	"learn/ch05/instructions/base"
	"learn/ch05/rtda"
)

type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
