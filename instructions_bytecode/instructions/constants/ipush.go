package constants

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type BIPUSH struct {
	val int8 //push byte
}
type SIPUSH struct {
	val int16 //push short
}

func (this BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	this.val = reader.ReadInt8()
}

func (this BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(this.val)
	frame.OperandStack.PushInt(i)
}

func (this SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	this.val = reader.ReadInt16()
}

func (this SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(this.val)
	frame.OperandStack.PushInt(i)
}
