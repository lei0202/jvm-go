package main

import "fmt"
import "JVM-GO/instructions_bytecode/classfile"
import "JVM-GO/instructions_bytecode/instructions"
import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, bytecode)
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPc(pc)
		//decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUInt8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		//execute
		fmt.Printf("pc:%2dinst:%T%v\n", pc, inst, inst)
		inst.Execute(frame)
	}

}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars)
		fmt.Printf("OperandStack:%v\n", frame.OperandStack)
		panic(r)
	}
}
