package z80

func oopLDIXnn(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.IX = nn
	cpu.LastOpCycles = 14
}

func oopLDIYnn(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.IY = nn
	cpu.LastOpCycles = 14
}

func oopLDHLnnP(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.HL.Lo = cpu.Memory.Get(nn)
	cpu.HL.Hi = cpu.Memory.Get(nn + 1)
	cpu.LastOpCycles = 20
}

func oopLDIXnnP(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.IX = cpu.readU16(nn)
	cpu.LastOpCycles = 20
}

func oopLDIYnnP(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.IY = cpu.readU16(nn)
	cpu.LastOpCycles = 20
}

func oopLDnnPHL(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.HL.U16())
	cpu.LastOpCycles = 16
}

func oopLDnnPIX(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.IX)
	cpu.LastOpCycles = 20
}

func oopLDnnPIY(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.IY)
	cpu.LastOpCycles = 20
}

func oopLDSPHL(cpu *CPU) {
	cpu.SP = cpu.HL.U16()
	cpu.LastOpCycles = 6
}

func oopLDSPIX(cpu *CPU) {
	cpu.SP = cpu.IX
	cpu.LastOpCycles = 10
}

func oopLDSPIY(cpu *CPU) {
	cpu.SP = cpu.IY
	cpu.LastOpCycles = 10
}

func oopPUSHIX(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.IX)
	cpu.LastOpCycles = 15
}

func oopPUSHIY(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.IY)
	cpu.LastOpCycles = 15
}

func oopPOPIX(cpu *CPU) {
	cpu.IX = cpu.readU16(cpu.SP)
	cpu.SP += 2
	cpu.LastOpCycles = 14
}

func oopPOPIY(cpu *CPU) {
	cpu.IY = cpu.readU16(cpu.SP)
	cpu.SP += 2
	cpu.LastOpCycles = 14
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func xopPUSHbc(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.BC.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.BC.Hi)
	cpu.LastOpCycles = 11
}

func xopPUSHde(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.DE.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.DE.Hi)
	cpu.LastOpCycles = 11
}

func xopPUSHhl(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.HL.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.HL.Hi)
	cpu.LastOpCycles = 11
}

func xopPUSHaf(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.AF.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.AF.Hi)
	cpu.LastOpCycles = 11
}

func xopPOPbc(cpu *CPU) {
	cpu.BC.Lo = cpu.Memory.Get(cpu.SP)
	cpu.BC.Hi = cpu.Memory.Get(cpu.SP + 1)
	cpu.SP += 2
	cpu.LastOpCycles = 10
}

func xopPOPde(cpu *CPU) {
	cpu.DE.Lo = cpu.Memory.Get(cpu.SP)
	cpu.DE.Hi = cpu.Memory.Get(cpu.SP + 1)
	cpu.SP += 2
	cpu.LastOpCycles = 10
}

func xopPOPhl(cpu *CPU) {
	cpu.HL.Lo = cpu.Memory.Get(cpu.SP)
	cpu.HL.Hi = cpu.Memory.Get(cpu.SP + 1)
	cpu.SP += 2
	cpu.LastOpCycles = 10
}

func xopPOPaf(cpu *CPU) {
	cpu.AF.Lo = cpu.Memory.Get(cpu.SP)
	cpu.AF.Hi = cpu.Memory.Get(cpu.SP + 1)
	cpu.SP += 2
	cpu.LastOpCycles = 10
}

func xopLDbcnn(cpu *CPU, l, h uint8) {
	cpu.BC.Lo = l
	cpu.BC.Hi = h
	cpu.LastOpCycles = 10
}

func xopLDdenn(cpu *CPU, l, h uint8) {
	cpu.DE.Lo = l
	cpu.DE.Hi = h
	cpu.LastOpCycles = 10
}

func xopLDhlnn(cpu *CPU, l, h uint8) {
	cpu.HL.Lo = l
	cpu.HL.Hi = h
	cpu.LastOpCycles = 10
}

func xopLDspnn(cpu *CPU, l, h uint8) {
	cpu.SP = toU16(l, h)
	cpu.LastOpCycles = 10
}

func xopLDnnPbc(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.BC.U16())
	cpu.LastOpCycles = 20
}

func xopLDnnPde(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.DE.U16())
	cpu.LastOpCycles = 20
}

func xopLDnnPhl(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.HL.U16())
	cpu.LastOpCycles = 20 // guess
}

func xopLDnnPsp(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.SP)
	cpu.LastOpCycles = 20
}

func xopLDbcnnP(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.BC.SetU16(cpu.readU16(nn))
	cpu.LastOpCycles = 20
}

func xopLDdennP(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.DE.SetU16(cpu.readU16(nn))
	cpu.LastOpCycles = 20
}

func xopLDhlnnP(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.HL.SetU16(cpu.readU16(nn))
	cpu.LastOpCycles = 20 // guess
}

func xopLDspnnP(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.SP = cpu.readU16(nn)
	cpu.LastOpCycles = 20
}
