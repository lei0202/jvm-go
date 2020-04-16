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

// 表示跳转指零
type BranchInstruction struct {
	offset int
}

func (This BranchInstruction) FetchOperands(reader *BytecodeReader) {
	This.offset = int(reader.ReadInt16())
}

// 存储和加载类指令需要根据索引存取局部变量表，索引由单字节操作数给出
type Index8Instruction struct {
	Index int
}

func (This Index8Instruction) FetchOperands(reader *BytecodeReader) {
	This.Index = int(reader.ReadInt8())
}

// 这些指令需要访问常量池，常量池索引需要两字节操作数给出
type Index16Instruction struct {
	Index int
}

func (This Index16Instruction) FetchOperands(reader *BytecodeReader) {
	This.Index = int(reader.ReadInt16())
}
