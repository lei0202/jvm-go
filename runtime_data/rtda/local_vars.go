package rtda

import "math"

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (Self LocalVars) SetInt(index uint, val int32) {
	Self[index].num = val
}

func (Self LocalVars) GetInt(index uint) int32 {
	return Self[index].num
}

func (Self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	Self[index].num = int32(bits)
}

func (Self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(Self[index].num)
	return math.Float32frombits(bits)
}

func (Self LocalVars) SetLong(index uint, val int64) {
	Self[index].num = int32(val)
	Self[index+1].num = int32(val >> 32)
}

func (Self LocalVars) GetLong(index uint) int64 {
	low := Self[index].num
	high := Self[index].num
	return int64(high)<<32 | int64(low)
}

func (Self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	Self.SetLong(index, int64(bits))
}

func (Self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(Self.GetLong(index))
	return math.Float64frombits(bits)
}

func (Self LocalVars) SetRef(index uint, ref *Object) {
	Self[index].ref = ref
}

func (Self LocalVars) GetRef(index uint) *Object {
	return Self[index].ref
}
