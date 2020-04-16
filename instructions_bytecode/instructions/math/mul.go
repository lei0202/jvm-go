package math

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type DMUL struct {
	base.NoOperandsInstruction
}
type IMUL struct {
	base.NoOperandsInstruction
}
type LMUL struct {
	base.NoOperandsInstruction
}
type FMUL struct {
	base.NoOperandsInstruction
}

func (this *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	val := val1 * val2
	stack.PushDouble(val)
}
func (this *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	val := val1 * val2
	stack.PushInt(val)
}
func (this *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	val := val1 * val2
	stack.PushLong(val)
}
func (this *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	val := val1 * val2
	stack.PushFloat(val)
}
