package classfile

type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

type ConstantFieldrefInfo struct {
	ConstantMemberRefInfo
}
type ConstantMethodrefInfo struct {
	ConstantMemberRefInfo
}
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberRefInfo
}

func (Self *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	Self.classIndex = reader.readUint16()
	Self.nameAndTypeIndex = reader.readUint16()
}

func (Self *ConstantMemberRefInfo) className() string {
	return Self.cp.getClassName(Self.classIndex)
}

func (Self *ConstantMemberRefInfo) Name() (string, string) {
	return Self.cp.getNameAndType(Self.nameAndTypeIndex)
}
