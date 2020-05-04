package rtda

import "learn/ch07/rtda/heap"

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}
func (self *Thread) PC() int      { return self.pc } // getter
func (self *Thread) SetPC(pc int) { self.pc = pc }   // setter
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

//func (self *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
//	return newFrame(self, maxLocals, maxStack)
//}
func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}

func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}
func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}
func (self *Thread) StackSize() uint {
	return self.stack.Size()
}
