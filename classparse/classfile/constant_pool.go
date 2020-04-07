package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	constantInfos := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		constantInfos[i] = readConstantInfo(reader, constantInfos)
		switch constantInfos[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return constantInfos
}

func (Self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := Self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index")
}

func (Self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := Self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := Self.getUtf8(ntInfo.nameIndex)
	_type := Self.getUtf8(ntInfo.descriptorIndex)
	return name, _type

}

func (Self ConstantPool) getClassName(index uint16) string {
	classInfo := Self.getConstantInfo(index).(*ConstantClassInfo)
	return Self.getUtf8(classInfo.ClassNameIndex)

}

func (Self ConstantPool) getUtf8(index uint16) string {
	utf8Info := Self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
