package math

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type DADD struct {
	base.NoOperandsInstruction
}
type IADD struct {
	base.NoOperandsInstruction
}
type LADD struct {
	base.NoOperandsInstruction
}
type FADD struct {
	base.NoOperandsInstruction
}

func (this *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	val := val1 + val2
	stack.PushDouble(val)
}
func (this *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	val := val1 + val2
	stack.PushInt(val)
}
func (this *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	val := val1 + val2
	stack.PushLong(val)
}
func (this *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	val := val1 + val2
	stack.PushFloat(val)
}
