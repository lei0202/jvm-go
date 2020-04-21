package comparisons

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type IF_ACMPEQ struct{ base.BranchInstruction }
type IF_ACMPNE struct{ base.BranchInstruction }

func (this *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 != ref2 {
		base.Branch(frame, this.Offset)
	}
}
