package math

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type DDIV struct {
	base.NoOperandsInstruction
}
type IDIV struct {
	base.NoOperandsInstruction
}
type LDIV struct {
	base.NoOperandsInstruction
}
type FDIV struct {
	base.NoOperandsInstruction
}

func (this *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	val := val1 / val2
	stack.PushDouble(val)
}
func (this *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	val := val1 / val2
	stack.PushInt(val)
}
func (this *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	val := val1 / val2
	stack.PushLong(val)
}
func (this *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	val := val1 / val2
	stack.PushFloat(val)
}
