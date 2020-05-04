package main

import (
	"fmt"
	"learn/ch06/instructions"
	"learn/ch06/instructions/base"
	"learn/ch06/rtda"
	"learn/ch06/rtda/heap"
)

//func interpret(methodInfo *classfile.MemberInfo) {
//	codeAttr := methodInfo.CodeAttribute()
//	maxLocals := codeAttr.MaxLocals()
//	maxStack := codeAttr.MaxStack()
//	bytecode := codeAttr.Code()
//	thread := rtda.NewThread()
//	frame := thread.NewFrame(maxLocals, maxStack)
//	thread.PushFrame(frame)
//	defer catchErr(frame)
//	loop(thread, bytecode)
//}
func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, method.Code())
}
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}
func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		fmt.Printf("LocalVars:%v OperandStack:%v\n", frame.LocalVars(), frame.OperandStack())
		pc := frame.NextPC()
		thread.SetPC(pc)
		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
