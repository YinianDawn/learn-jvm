package rtda

import "learn/ch07/rtda/heap"

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPC       int
}

func (f Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f Frame) OperandStack() *OperandStack {
	return f.operandStack
}

//func NewFrame(maxLocals, maxStack uint) *Frame {
//	return &Frame{
//		localVars:    newLocalVars(maxLocals),
//		operandStack: newOperandStack(maxStack),
//	}
//}
func (f *Frame) Thread() *Thread {
	return f.thread
}
func (f *Frame) SetNextPC(nextPC int) {
	f.nextPC = nextPC
}

func (f *Frame) NextPC() int {
	return f.nextPC
}

func (f Frame) Method() *heap.Method {
	return f.method
}
func Branch(frame *Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}

//func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
//	return &Frame{
//		thread:       thread,
//		localVars:    newLocalVars(maxLocals),
//		operandStack: newOperandStack(maxStack),
//	}
//}
func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}
func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}
