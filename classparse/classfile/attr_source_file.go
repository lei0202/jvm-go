package classfile

type AttributeSourceFile struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (Self *AttributeSourceFile) readInfo(reader *ClassReader) {
	Self.sourceFileIndex = reader.readUint16()
}

func (Self *AttributeSourceFile) FileName() string {
	return Self.cp.getUtf8(Self.sourceFileIndex)
}
