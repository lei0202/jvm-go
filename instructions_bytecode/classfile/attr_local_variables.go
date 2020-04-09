package classfile

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTable
}

type LocalVariableTable struct {
	startPc     uint16
	localNumber uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.localVariableTable = make([]*LocalVariableTable, lineNumberTableLength)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVariableTable{
			startPc:     reader.readUint16(),
			localNumber: reader.readUint16(),
		}
	}
}
