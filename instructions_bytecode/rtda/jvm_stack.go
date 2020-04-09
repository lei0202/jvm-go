package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (Self *Stack) push(frame *Frame) {
	if Self.size >= Self.maxSize {
		panic("java.lang.StackOverFlowError")
	}
	if Self._top != nil {
		frame.lower = Self._top
	}
	Self._top = frame
	Self.size++
}

func (Self *Stack) pop() *Frame {
	if Self._top == nil {
		panic("jvm stack is empty!")
	}
	top := Self._top
	Self._top = top.lower
	top.lower = nil
	Self.size--
	return top
}

func (Self *Stack) top() *Frame {
	if Self._top == nil {
		panic("jvm stack is empty!")
	}
	return Self._top
}
