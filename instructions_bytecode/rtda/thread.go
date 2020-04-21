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
	return Self.stack.pop()
}

func (Self *Thread) PushFrame(frame *Frame) {
	Self.stack.push(frame)
}

func (Self *Thread) CurrentFrame() *Frame {
	return Self.stack.top()
}

func (Self *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return NewFrame(Self, maxLocals, maxStack)
}
