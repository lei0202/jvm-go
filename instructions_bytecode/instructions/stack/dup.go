package stack

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type DUP struct {
	base.NoOperandsInstruction
}
type DUP_X1 struct {
	base.NoOperandsInstruction
}
type DUP_X2 struct {
	base.NoOperandsInstruction
}
type DUP2 struct {
	base.NoOperandsInstruction
}
type DUP2_X1 struct {
	base.NoOperandsInstruction
}
type DUP2_x2 struct {
	base.NoOperandsInstruction
}

func (this *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
