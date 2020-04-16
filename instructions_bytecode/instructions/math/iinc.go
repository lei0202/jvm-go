package math

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type IINC struct {
	Index uint
	Const int32
}

func (this *IINC) FetchOperand(reader *base.BytecodeReader) {
	this.Index = uint(reader.ReadUInt8())
	this.Const = reader.ReadInt32()

}

func (this *IINC) Execute(frame *rtda.Frame) {
	localvars := frame.LocalVars
	val := localvars.GetInt(this.Index)
	val += this.Const
	localvars.SetInt(this.Index, val)
}
