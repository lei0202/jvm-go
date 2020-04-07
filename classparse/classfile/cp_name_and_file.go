package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (Self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	Self.nameIndex = reader.readUint16()
	Self.descriptorIndex = reader.readUint16()
}
