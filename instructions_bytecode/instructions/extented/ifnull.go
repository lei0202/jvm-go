package extented

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type IFNULL struct {
	base.BranchInstruction // branch if null
}

type IFNONULL struct {
	base.BranchInstruction // branch if not null
}

func (this *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack.PopRef()
	if ref == nil {
		base.Branch(frame, this.Offset)
	}
}

func (this *IFNONULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack.PopRef()
	if ref != nil {
		base.Branch(frame, this.Offset)
	}
}
