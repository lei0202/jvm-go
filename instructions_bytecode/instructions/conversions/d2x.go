package conversions

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type D2F struct {
	base.NoOperandsInstruction
}

type D2L struct {
	base.NoOperandsInstruction
}

type D2I struct {
	base.NoOperandsInstruction
}

func (this *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	v := stack.PopDouble()
	result := float32(v)
	stack.PushFloat(result)
}

func (this *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	v := stack.PopDouble()
	result := int64(v)
	stack.PushLong(result)
}

func (this *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	v := stack.PopDouble()
	result := int32(v)
	stack.PushInt(result)
}
