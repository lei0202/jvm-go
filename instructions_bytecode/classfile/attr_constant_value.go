package classfile

type AttributeConstantValue struct {
	constantValueIndex uint16
}

func (self *AttributeConstantValue) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *AttributeConstantValue) ConstantValueIndex() uint16 {
	return self.ConstantValueIndex()
}
