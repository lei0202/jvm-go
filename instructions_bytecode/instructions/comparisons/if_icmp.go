package comparisons

import "JVM-GO/instructions_bytecode/instructions/base"
import "JVM-GO/instructions_bytecode/rtda"

type IF_ICMPEQ struct{ base.BranchInstruction }
type IF_ICMPNE struct{ base.BranchInstruction }
type IF_ICMPLT struct{ base.BranchInstruction }
type IF_ICMPLE struct{ base.BranchInstruction }
type IF_ICMPGT struct{ base.BranchInstruction }
type IF_ICMPGE struct{ base.BranchInstruction }

func (this *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 == v2 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 != v2 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IF_ICMPLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 < v2 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IF_ICMPLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 <= v2 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IF_ICMPGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 > v2 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 >= v2 {
		base.Branch(frame, this.Offset)
	}
}
