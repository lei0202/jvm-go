package classfile

type ConstantClassInfo struct {
	cp             ConstantPool
	ClassNameIndex uint16
}

func (Self *ConstantClassInfo) readInfo(reader *ClassReader) {
	Self.ClassNameIndex = reader.readUint16()
}

func (Self *ConstantClassInfo) Name() string {
	return Self.cp.getUtf8(Self.ClassNameIndex)
}
