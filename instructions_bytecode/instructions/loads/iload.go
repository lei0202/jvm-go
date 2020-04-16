package loads

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type ILOAD struct {
	base.Index8Instruction
}

type ILOAD_0 struct {
	base.NoOperandsInstruction
}
type ILOAD_1 struct {
	base.NoOperandsInstruction
}
type ILOAD_2 struct {
	base.NoOperandsInstruction
}
type ILOAD_3 struct {
	base.NoOperandsInstruction
}
type ILOAD_4 struct {
	base.NoOperandsInstruction
}

func _load(frame *rtda.Frame, index uint) {
	val := frame.LocalVars.GetInt(index)
	frame.OperandStack.PushInt(val)
}

func (this *ILOAD) Execute(frame *rtda.Frame) {
	_load(frame, uint(this.Index))
}

func (this *ILOAD_0) Execute(frame *rtda.Frame) {
	_load(frame, uint(0))
}

func (this *ILOAD_1) Execute(frame *rtda.Frame) {
	_load(frame, uint(1))
}

func (this *ILOAD_2) Execute(frame *rtda.Frame) {
	_load(frame, uint(2))
}

func (this *ILOAD_3) Execute(frame *rtda.Frame) {
	_load(frame, uint(3))
}

func (this *ILOAD_4) Execute(frame *rtda.Frame) {
	_load(frame, uint(4))
}
