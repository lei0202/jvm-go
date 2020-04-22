package heap

import "JVM-GO/class_and_object/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
}

func (this *Method) copyAttributes(info *classfile.MemberInfo) {
	if codeAttr := info.CodeAttribute(); codeAttr != nil {
		this.maxLocals = codeAttr.MaxLocals()
		this.maxStack = codeAttr.MaxStack()
		this.code = codeAttr.Code()
	}
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}
