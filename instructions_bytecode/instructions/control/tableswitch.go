package control

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (this *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	this.defaultOffset = reader.ReadInt32()
	this.low = reader.ReadInt32()
	this.high = reader.ReadInt32()
	jumpOffsetsCount := this.high - this.low + 1
	this.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (this *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack.PopInt()
	var offset int
	if index >= this.low && index <= this.high {
		offset = int(this.jumpOffsets[index-this.low])
	} else {
		offset = int(this.defaultOffset)
	}
	base.Branch(frame, offset)
}
