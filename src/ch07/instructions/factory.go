package instructions

import (
	"fmt"
	"learn/ch07/instructions/base"
	"learn/ch07/instructions/comparisons"
	"learn/ch07/instructions/constants"
	"learn/ch07/instructions/control"
	"learn/ch07/instructions/loads"
	"learn/ch07/instructions/math"
	"learn/ch07/instructions/references"
	"learn/ch07/instructions/stack"
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
	case 0x0A:
		return &constants.LCONST_1{}
	case 0x10:
		return &constants.BIPUSH{}
	case 0x12:
		return &constants.LDC{}
	case 0x14:
		return &constants.LDC2_W{}
	case 0x1B:
		return &loads.ILOAD_1{}
	case 0x1C:
		return &loads.ILOAD_2{}
	case 0x1E:
		return &loads.LLOAD_0{}
	case 0x1F:
		return &loads.LLOAD_1{}
	case 0x2A:
		return &loads.ALOAD_0{}
	case 0x2B:
		return &loads.ALOAD_1{}
	case 0x2C:
		return &loads.ALOAD_2{}
	case 0x2D:
		return &loads.ALOAD_3{}
	case 0x3C:
		return &constants.ISTORE_1{}
	case 0x3D:
		return &constants.ISTORE_2{}
	case 0x40:
		return &constants.LSTORE_1{}
	case 0x4C:
		return &constants.ASTORE_1{}
	case 0x4D:
		return &constants.ASTORE_2{}
	case 0x4E:
		return &constants.ASTORE_3{}
	case 0x57:
		return &stack.POP{}
	case 0x59:
		return &stack.DUP{}
	case 0x60:
		return &math.IADD{}
	case 0x61:
		return &math.LADD{}
	case 0x65:
		return &math.LSUB{}
	case 0x84:
		return &math.IINC{}
	case 0x94:
		return &comparisons.LCMP{}
	case 0x99:
		return &comparisons.IFEQ{}
	case 0x9D:
		return &comparisons.IFGT{}
	case 0xA3:
		return &comparisons.IF_ICMPGT{}
	case 0xA6:
		return &comparisons.IF_ACMPNE{}
	case 0xA7:
		return &control.GOTO{}
	case 0xAC:
		return &control.IRETURN{}
	case 0xAD:
		return &control.LRETURN{}
	case 0xB1:
		return &control.RETURN{}
	case 0xB2:
		return &references.GET_STATIC{}
	case 0xB3:
		return &references.PUT_STATIC{}
	case 0xB4:
		return &references.GET_FIELD{}
	case 0xB5:
		return &references.PUT_FIELD{}
	case 0xB6:
		return &references.INVOKE_VIRTUAL{}
	case 0xB7:
		return &references.INVOKE_SPECIAL{}
	case 0xB8:
		return &references.INVOKE_STATIC{}
	case 0xB9:
		return &references.INVOKE_INTERFACE{}
	case 0xBB:
		return &references.NEW{}
	case 0xC0:
		return &references.CHECK_CAST{}
	case 0xC1:
		return &references.INSTANCE_OF{}
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
