package stack

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type SWAP struct {
	base.NoOperandsInstruction
}

func (this *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
