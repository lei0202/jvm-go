package heap

import "JVM-GO/class_and_object/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.CopyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
