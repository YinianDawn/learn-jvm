package constants

import (
	"learn/ch08/instructions/base"
	"learn/ch08/rtda"
)

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// 什么也不用做
}
