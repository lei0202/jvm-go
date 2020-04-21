package control

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type GOTO struct {
	base.BranchInstruction
}

func (this *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, this.Offset)
}
