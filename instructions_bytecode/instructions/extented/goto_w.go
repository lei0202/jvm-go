package extented

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type GOTO_W struct {
	offset int // 32 byte
}

func (this *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	this.offset = int(reader.ReadUInt32())
}

func (this *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, this.offset)
}
