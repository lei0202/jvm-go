package rtda

type Frame struct {
	lower        *Frame
	LocalVars    LocalVars
	OperandStack OperandStack
	thread       *Thread
	nextPC       int
}

func (this *Frame) NextPC() int {
	return this.nextPC
}

func (this *Frame) SetNextPC(pc int) {
	this.nextPC = pc
}

func (this *Frame) Thread() *Thread {
	return this.thread
}

func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		LocalVars:    newLocalVars(maxLocals),
		OperandStack: newOperandStack(maxStack),
	}
}
