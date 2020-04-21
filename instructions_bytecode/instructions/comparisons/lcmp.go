package comparisons

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

// 比较LONG类型
type LCMP struct {
	base.NoOperandsInstruction
}

func (this *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
