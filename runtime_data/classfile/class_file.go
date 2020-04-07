package classfile

import "fmt"

type ClassFile struct { //magicuint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (Self *ClassFile) read(reader *ClassReader) {
	Self.readAndCheckMagic(reader)
	Self.readAndCheckVersion(reader)
	Self.constantPool = readConstantPool(reader)
	Self.accessFlags = reader.readUint16()
	Self.thisClass = reader.readUint16()
	Self.superClass = reader.readUint16()
	Self.interfaces = reader.readUint16s()
	Self.fields = readMembers(reader, Self.constantPool)
	Self.methods = readMembers(reader, Self.constantPool)
	Self.attributes = readAttributes(reader, Self.constantPool)
}

// Getter of majorVersion
func (Self *ClassFile) MajorVersion() uint16 {
	return Self.majorVersion
}
func (Self *ClassFile) AccessFlags() uint16 {
	return Self.accessFlags
}

func (Self *ClassFile) ConstantPool() ConstantPool {
	return Self.constantPool
}

func (Self *ClassFile) Fields() []*MemberInfo {
	return Self.fields
}

func (Self *ClassFile) MinorVersion() uint16 {
	return Self.minorVersion
}

func (Self *ClassFile) Methods() []*MemberInfo {
	return Self.methods
}

func (Self *ClassFile) ThisClassName() string {
	return Self.constantPool.getClassName(Self.thisClass)
}

func (Self *ClassFile) SuperClassName() string {
	if Self.superClass > 0 {
		return Self.constantPool.getClassName(Self.superClass)
	}
	return ""
}

func (Self *ClassFile) InterfaceNames() []string {

	interfaceNames := make([]string, len(Self.interfaces))
	for i, cpIndex := range Self.interfaces {
		interfaceNames[i] = Self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

func (Self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError:magic!")
	}
}

func (Self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	Self.minorVersion = reader.readUint16()
	Self.majorVersion = reader.readUint16()
	switch Self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if Self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}
