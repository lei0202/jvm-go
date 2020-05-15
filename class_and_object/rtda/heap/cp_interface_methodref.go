package heap

import "JVM-GO/class_and_object/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.CopyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
