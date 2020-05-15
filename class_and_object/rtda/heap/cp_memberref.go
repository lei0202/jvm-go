package heap

import "JVM-GO/class_and_object/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (this *MemberRef) CopyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	this.className = refInfo.ClassName()
	this.name, this.descriptor = refInfo.NameAndDescriptor()
}
