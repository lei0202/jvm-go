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

// tableswitch 指令操作码的后面有0-3字节的padding，以保证defaultOffset在字节码中的地址是4的倍数
func (this *BytecodeReader) SkipPadding() {
	for this.pc%4 != 0 {
		this.ReadUInt8()
	}
}

func (this *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = this.ReadInt32()
	}
	return ints
}

func (this *BytecodeReader) PC() int {
	return this.pc
}
