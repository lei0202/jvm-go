package heap

import "JVM-GO/class_and_object/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (this *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	this.accessFlags = memberInfo.AccessFlags()
	this.name = memberInfo.Name()
	this.descriptor = memberInfo.Descriptor()
}
