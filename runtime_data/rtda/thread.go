package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (Self *Thread) PC() int {
	return Self.pc
}

func (Self *Thread) SetPc(pc int) {
	Self.pc = pc
}

func (Self *Thread) PopFrame() *Frame {
}

func (Self *Thread) PushFrame(frame *Frame) {
}

func (Self *Thread) CurrentFrame() *Frame {
}
