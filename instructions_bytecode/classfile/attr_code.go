package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)

}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	length := reader.readUint16()
	table := make([]*ExceptionTableEntry, length)
	for i := range table {
		table[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return table
}

func (this *CodeAttribute) MaxStack() uint {
	return uint(this.maxStack)
}

func (this *CodeAttribute) MaxLocals() uint {
	return uint(this.maxLocals)
}

func (this *CodeAttribute) Code() []byte {
	return this.code
}
