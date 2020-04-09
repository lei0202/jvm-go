package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (Self *ConstantStringInfo) readInfo(reader *ClassReader) {
	Self.stringIndex = reader.readUint16()
}

func (Self *ConstantStringInfo) String() string {
	return Self.cp.getUtf8(Self.stringIndex)
}
