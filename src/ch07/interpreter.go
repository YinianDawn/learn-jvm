package main

import (
	"fmt"
	"learn/ch07/instructions"
	"learn/ch07/instructions/base"
	"learn/ch07/rtda"
	"learn/ch07/rtda/heap"
)

func interpret(method *heap.Method, logInst bool) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(thread)
	loop(thread, logInst)
}
func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}
func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()

		fmt.Printf("frame: %d -> LocalVars:%v OperandStack:%v\n", thread.StackSize(), frame.LocalVars(), frame.OperandStack())

		pc := frame.NextPC()
		thread.SetPC(pc)
		// decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		if logInst {
			logInstruction(frame, inst)
		}
		// execute
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}
func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}
