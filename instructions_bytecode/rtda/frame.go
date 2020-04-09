package rtda

type Frame struct {
	lower *Frame
	LocalVars
	OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		LocalVars:    newLocalVars(maxLocals),
		OperandStack: newOperandStack(maxStack),
	}
}
