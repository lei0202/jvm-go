package references

import "JVM-GO/class_and_object/instructions/base"
import "JVM-GO/class_and_object/rtda"
import "JVM-GO/class_and_object/rtda/heap"

//new指令的操作数是一个uint16索引，来自字节码。通过这个索引，
// 可以从当前类的运行时常量池中找到一个类符号引用。解析这个类符号引用，
// 拿到类数据，然后创建对象，并把对象引用推入栈顶，new指令的工作就完成了。
type NEW struct {
	base.Index16Instruction
}

func (this *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(this.Index).(*heap.ClassRef)
	class := classRef.ResolveClass()
	if class.IsAbstract() || class.IsInterface() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
