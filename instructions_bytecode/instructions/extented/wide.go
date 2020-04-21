package extented

import (
	"JVM-GO/instructions_bytecode/instructions/base"
	"JVM-GO/instructions_bytecode/instructions/loads"
	"JVM-GO/instructions_bytecode/instructions/math"
	"JVM-GO/instructions_bytecode/instructions/stores"
	"JVM-GO/instructions_bytecode/rtda"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

// todo
func (this *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUInt8()
	switch opcode {
	case 0x15:
		inst := &loads.ILOAD{} //iload
		inst.Index = int(reader.ReadUInt16())
		this.modifiedInstruction = inst
	case 0x16: //lload
		inst := &loads.ILOAD{}
		inst.Index = int(reader.ReadUInt16())
		this.modifiedInstruction = inst
	case 0x17: //fload
		inst := &loads.ILOAD{}
		inst.Index = int(reader.ReadUInt16())
		this.modifiedInstruction = inst
	case 0x18: //dload
		inst := &loads.ILOAD{}
		inst.Index = int(reader.ReadUInt16())
		this.modifiedInstruction = inst
	case 0x19: //aload
		inst := &loads.ILOAD{}
		inst.Index = int(reader.ReadUInt16())
		this.modifiedInstruction = inst
	case 0x36: //istore
		inst := &stores.LSTORE{}
		inst.Index = int(reader.ReadUInt16())
		this.modifiedInstruction = inst
	case 0x37: //lstore
		inst := &stores.LSTORE{}
		inst.Index = int(reader.ReadUInt16())
		this.modifiedInstruction = inst
	case 0x38: //fstore
		inst := &stores.LSTORE{}
		inst.Index = int(reader.ReadUInt16())
		this.modifiedInstruction = inst
	case 0x39: //dstore
		inst := &stores.LSTORE{}
		inst.Index = int(reader.ReadUInt16())
		this.modifiedInstruction = inst
	case 0x3a: //astore
		inst := &stores.LSTORE{}
		inst.Index = int(reader.ReadUInt16())
		this.modifiedInstruction = inst
	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUInt16())
		inst.Const = int32(reader.ReadUInt16())
		this.modifiedInstruction = inst //iinc
	case 0xa9:
		panic("Unsupported opcode: 0xa9!") //ret
	}
}

func (this *WIDE) Execute(frame *rtda.Frame) {
	this.modifiedInstruction.Execute(frame)
}
