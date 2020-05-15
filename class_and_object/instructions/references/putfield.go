package references

import "JVM-GO/class_and_object/instructions/base"
import "JVM-GO/class_and_object/rtda"
import "JVM-GO/class_and_object/rtda/heap"

type PUT_FIELD struct {
	base.Index16Instruction
}

func (this *PUT_FIELD) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(this.Index).(*heap.FieldRef)
	field := fieldRef.ResolveField()

}
