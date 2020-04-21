package control

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (this *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	this.defaultOffset = reader.ReadInt32()
	this.npairs = reader.ReadInt32()
	this.matchOffsets = reader.ReadInt32s(this.npairs * 2)
}

func (this *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack.PopInt()
	for i := int32(0); i < this.npairs*2; i += 2 {
		if this.matchOffsets[i] == key {
			offset := this.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(this.defaultOffset))
}
