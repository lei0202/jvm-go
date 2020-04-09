package classfile

import "math"

// Integer 类型
type ConstantIntegerInfo struct {
	val float32
}

func (Self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	Self.val = math.Float32frombits(bytes)
}

// Long 类型
type ConstantLongInfo struct {
	val int64
}

func (Self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	Self.val = int64(bytes)
}

type ConstantFloatInfo struct {
	val float32
}

func (Self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	Self.val = math.Float32frombits(bytes)
}

// Double 类型
type ConstantDoubleInfo struct {
	val float64
}

func (Self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	Self.val = math.Float64frombits(bytes)
}
