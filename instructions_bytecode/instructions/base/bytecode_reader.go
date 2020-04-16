package base

type BytecodeReader struct {
	code []byte // 存放字节码
	pc   int    // 记录当前光标
}

// basically a constructor
func (this BytecodeReader) Reset(code []byte, pc int) {
	this.code = code
	this.pc = pc
}

func (this BytecodeReader) ReadUInt8() uint8 {
	i := this.code[this.pc]
	this.pc++
	return i
}

func (this BytecodeReader) ReadInt8() int8 {
	return int8(this.ReadUInt8())
}

func (this BytecodeReader) ReadUInt16() uint16 {
	i1 := uint16(this.ReadInt8())
	i2 := uint16(this.ReadInt8())
	i := i1<<8 | i2
	return i
}

func (this BytecodeReader) ReadInt16() int16 {
	return int16(this.ReadUInt16())
}

func (this BytecodeReader) ReadUInt32() uint32 {
	i1 := uint32(this.ReadInt8())
	i2 := uint32(this.ReadInt8())
	i3 := uint32(this.ReadInt8())
	i4 := uint32(this.ReadInt8())
	i := i1<<24 | i2<<16 | i3<<8 | i4
	return i
}

func (this BytecodeReader) ReadInt32() int32 {
	return int32(this.ReadUInt32())
}
