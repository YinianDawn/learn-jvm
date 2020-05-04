package math

import (
	"learn/ch05/instructions/base"
	"learn/ch05/rtda"
)

type DADD struct{ base.NoOperandsInstruction }
type FADD struct{ base.NoOperandsInstruction }
type IADD struct{ base.NoOperandsInstruction }
type LADD struct{ base.NoOperandsInstruction }

func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	if 2<<16 < result || result < -2<<16 {
		panic("java.lang.ArithmeticException: 2^16 < or < -2^16")
	}
	stack.PushInt(result)
}
