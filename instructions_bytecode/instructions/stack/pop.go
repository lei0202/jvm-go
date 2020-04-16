package stack

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type POP struct {
	base.NoOperandsInstruction
}
type POP2 struct {
	base.NoOperandsInstruction
}

func (this *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	stack.PopSlot()
}

func (this *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	stack.PopSlot()
	stack.PopSlot()
}
