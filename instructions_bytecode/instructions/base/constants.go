package base

import "JVM-GO/instructions_bytecode/rtda"

type NOP struct {
}

func (this NOP) FetchOperands(reader *BytecodeReader) {
	// do nothing
}
func (this NOP) Execute(frame *rtda.Frame) {
	// do nothing
}
