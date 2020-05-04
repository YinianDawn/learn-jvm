package instructions

import (
	"fmt"
	"learn/ch09/instructions/base"
	"learn/ch09/instructions/comparisons"
	"learn/ch09/instructions/constants"
	"learn/ch09/instructions/control"
	"learn/ch09/instructions/loads"
	"learn/ch09/instructions/math"
	"learn/ch09/instructions/references"
	"learn/ch09/instructions/stack"
	"learn/ch09/instructions/stores"
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
	case 0x05:
		return &constants.ICONST_2{}
	case 0x06:
		return &constants.ICONST_3{}
	case 0x07:
		return &constants.ICONST_4{}
	case 0x08:
		return &constants.ICONST_5{}
	case 0x0A:
		return &constants.LCONST_1{}
	case 0x10:
		return &constants.BIPUSH{}
	case 0x12:
		return &constants.LDC{}
	case 0x14:
		return &constants.LDC2_W{}
	case 0x15:
		return &loads.ILOAD{}
	case 0x19:
		return &loads.ALOAD{}
	case 0x1B:
		return &loads.ILOAD_1{}
	case 0x1C:
		return &loads.ILOAD_2{}
	case 0x1D:
		return &loads.ILOAD_3{}
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
	case 0x2E:
		return &loads.IALOAD{}
	case 0x3C:
		return &stores.ISTORE_1{}
	case 0x32:
		return &loads.AALOAD{}
	case 0x36:
		return &stores.ISTORE{}
	case 0x3A:
		return &stores.ASTORE{}
	case 0x3D:
		return &stores.ISTORE_2{}
	case 0x3E:
		return &stores.ISTORE_3{}
	case 0x40:
		return &stores.LSTORE_1{}
	case 0x4C:
		return &stores.ASTORE_1{}
	case 0x4D:
		return &stores.ASTORE_2{}
	case 0x4E:
		return &stores.ASTORE_3{}
	case 0x4F:
		return &stores.IASTORE{}
	case 0x57:
		return &stack.POP{}
	case 0x59:
		return &stack.DUP{}
	case 0x60:
		return &math.IADD{}
	case 0x61:
		return &math.LADD{}
	case 0x64:
		return &math.ISUB{}
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
	case 0xA2:
		return &comparisons.IF_ICMPGE{}
	case 0xA3:
		return &comparisons.IF_ICMPGT{}
	case 0xA4:
		return &comparisons.IF_ICMPLE{}
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
	case 0xBC:
		return &references.NEW_ARRAY{}
	case 0xBE:
		return &references.ARRAY_LENGTH{}
	case 0xC0:
		return &references.CHECK_CAST{}
	case 0xC1:
		return &references.INSTANCE_OF{}
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
