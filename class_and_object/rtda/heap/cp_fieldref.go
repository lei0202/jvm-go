package heap

import "JVM-GO/class_and_object/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.CopyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (this *FieldRef) ResolveField() *Field {
	if this.field == nil {
		this.resolveFieldRef()
	}
	return this.field
}

func (this *FieldRef) resolveFieldRef() {
	d := this.cp.class
	c := this.ResolveClass()
	field := lookupField(c, this.name, this.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	this.field = field
}

func lookupField(c *Class, name string, descriptor string) *Field {

}
