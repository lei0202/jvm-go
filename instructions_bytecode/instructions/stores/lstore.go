package stores

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type LSTORE struct {
	base.Index8Instruction
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}
type LSTORE_1 struct {
	base.NoOperandsInstruction
}
type LSTORE_2 struct {
	base.NoOperandsInstruction
}
type LSTORE_3 struct {
	base.NoOperandsInstruction
}
type LSTORE_4 struct {
	base.NoOperandsInstruction
}

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack.PopLong()
	frame.LocalVars.SetLong(index, val)
}

func (this *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(this.Index))
}

func (this *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(0))
}

func (this *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(1))
}

func (this *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(2))
}

func (this *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(3))
}

func (this *LSTORE_4) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(4))
}
