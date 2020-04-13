package base

import "JVM-GO/instructions_bytecode/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

// 表示没有操作数的指令
type NoOperandsInstruction struct {
}

func (This NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}
