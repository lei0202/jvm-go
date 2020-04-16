package math

import (
	"JVM-GO/instructions_bytecode/instructions/base"
	"math"
)
import "JVM-GO/instructions_bytecode/rtda"

type DREM struct {
	base.NoOperandsInstruction
}
type FREM struct {
	base.NoOperandsInstruction
}
type IREM struct {
	base.NoOperandsInstruction
}
type LREM struct {
	base.NoOperandsInstruction
}

func (this *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	val := math.Mod(val1, val2)
	stack.PushDouble(val)
}
func (this *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	if val2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	val := val1 % val2
	stack.PushInt(val)
}
func (this *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	if val2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	val := val1 % val2
	stack.PushLong(val)
}
func (this *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	val := math.Mod(float64(val1), float64(val2))
	stack.PushFloat(float32(val))
}
