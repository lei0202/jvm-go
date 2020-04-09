package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (Self *ClassReader) readUint8() uint8 {
	val := Self.data[0]
	Self.data = Self.data[1:]
	return val
}

func (Self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(Self.data)
	Self.data = Self.data[2:]
	return val
}

func (Self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(Self.data)
	Self.data = Self.data[4:]
	return val
}

//u4
func (Self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(Self.data)
	Self.data = Self.data[8:]
	return val
}

func (Self *ClassReader) readUint16s() []uint16 {
	n := Self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = Self.readUint16()
	}
	return s
}

func (Self *ClassReader) readBytes(n uint32) []byte {
	bytes := Self.data[:n]
	Self.data = Self.data[n:]
	return bytes
}
