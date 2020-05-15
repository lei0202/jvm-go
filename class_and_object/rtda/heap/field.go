package heap

import "JVM-GO/class_and_object/classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfFields := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfFields)
		fields[i].copyAttributes(cfFields)
	}
	return fields
}
func (self *Field) IsVolatile() bool {
	return 0 != self.accessFlags&ACC_VOLATILE
}
func (self *Field) IsTransient() bool {
	return 0 != self.accessFlags&ACC_TRANSIENT
}
func (self *Field) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

func (this *Field) isLongOrDouble() bool {
	return this.descriptor == "J" || this.descriptor == "D"
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}
func (self *Field) SlotId() uint {
	return self.slotId
}

func (this *Field) copyAttributes(info *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		this.constValueIndex = uint(valAttr.ConstantValueIndex())
	}

}
