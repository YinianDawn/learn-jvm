package extended

import (
	"learn/ch07/instructions/base"
	"learn/ch07/rtda"
)

type IFNULL struct{ base.BranchInstruction }    // Branch if reference is null
type IFNONNULL struct{ base.BranchInstruction } // Branch if reference not null

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}
