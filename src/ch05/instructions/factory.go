package instructions

import (
	"fmt"
	"learn/ch05/instructions/base"
	"learn/ch05/instructions/comparisons"
	"learn/ch05/instructions/constants"
	"learn/ch05/instructions/control"
	"learn/ch05/instructions/math"
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return &constants.NOP{}
	case 0x01:
		return &constants.ACONST_NULL{}
	case 0x03:
		return &constants.ICONST_0{}
	case 0x04:
		return &constants.ICONST_1{}
	case 0x10:
		return &constants.BIPUSH{}
	case 0x1B:
		return &constants.ILOAD_1{}
	case 0x1C:
		return &constants.ILOAD_2{}
	case 0x3C:
		return &constants.ISTORE_1{}
	case 0x3D:
		return &constants.ISTORE_2{}
	case 0x60:
		return &math.IADD{}
	case 0x84:
		return &math.IINC{}
	case 0xA3:
		return &comparisons.IF_ICMPGT{}
	case 0xA7:
		return &control.GOTO{}
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
