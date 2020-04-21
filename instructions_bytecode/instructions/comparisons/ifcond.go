package comparisons

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type IFEQ struct{ base.BranchInstruction }
type IFNE struct{ base.BranchInstruction }
type IFLT struct{ base.BranchInstruction }
type IFLE struct{ base.BranchInstruction }
type IFGT struct{ base.BranchInstruction }
type IFGE struct{ base.BranchInstruction }

func (this *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	if val == 0 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	if val != 0 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	if val < 0 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	if val <= 0 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	if val > 0 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	if val >= 0 {
		base.Branch(frame, this.Offset)
	}
}
