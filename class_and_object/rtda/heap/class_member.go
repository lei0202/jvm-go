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

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

// getters
func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
}

func (this *ClassMember) isAccessibleTo(d *Class) bool {
	if this.IsPublic() {
		return true
	}
	c := this.class
	if this.IsProtected() {
		return d == c || d.isSubClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}
	if !this.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}
