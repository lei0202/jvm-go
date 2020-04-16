package math

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type DSUB struct {
	base.NoOperandsInstruction
}
type ISUB struct {
	base.NoOperandsInstruction
}
type LSUB struct {
	base.NoOperandsInstruction
}
type FSUB struct {
	base.NoOperandsInstruction
}

func (this *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	val := val1 - val2
	stack.PushDouble(val)
}
func (this *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	val := val1 - val2
	stack.PushInt(val)
}
func (this *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	val := val1 - val2
	stack.PushLong(val)
}
func (this *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	val := val1 - val2
	stack.PushFloat(val)
}
