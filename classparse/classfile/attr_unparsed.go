package classfile

type UnparsedAttribute struct {
	name string
	len  uint32
	info []byte
}

func (Self *UnparsedAttribute) readInfo(reader *ClassReader) {
	Self.info = reader.readBytes(Self.len)
}
