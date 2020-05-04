package control

import (
	"learn/ch08/instructions/base"
	"learn/ch08/rtda"
)

type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
