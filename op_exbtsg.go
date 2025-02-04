package z80

func oopEXDEHL(cpu *CPU) {
	cpu.HL, cpu.DE = cpu.DE, cpu.HL
	cpu.LastOpCycles = 4
}

func oopEXAFAF(cpu *CPU) {
	cpu.AF, cpu.Alternate.AF = cpu.Alternate.AF, cpu.AF
	cpu.LastOpCycles = 4
}

func oopEXX(cpu *CPU) {
	cpu.BC, cpu.Alternate.BC = cpu.Alternate.BC, cpu.BC
	cpu.DE, cpu.Alternate.DE = cpu.Alternate.DE, cpu.DE
	cpu.HL, cpu.Alternate.HL = cpu.Alternate.HL, cpu.HL
	cpu.LastOpCycles = 4
}

func oopEXSPPHL(cpu *CPU) {
	v := cpu.readU16(cpu.SP)
	cpu.writeU16(cpu.SP, cpu.HL.U16())
	cpu.HL.SetU16(v)
	cpu.LastOpCycles = 19
}

func oopEXSPPIX(cpu *CPU) {
	v := cpu.readU16(cpu.SP)
	cpu.writeU16(cpu.SP, cpu.IX)
	cpu.IX = v
	cpu.LastOpCycles = 23
}

func oopEXSPPIY(cpu *CPU) {
	v := cpu.readU16(cpu.SP)
	cpu.writeU16(cpu.SP, cpu.IY)
	cpu.IY = v
	cpu.LastOpCycles = 23
}

func (cpu *CPU) updateFlagLDID(a uint8) {
	var nand uint8 = mask53 | maskH | maskPV | maskN
	var or uint8
	if cpu.BC.Lo != 0 || cpu.BC.Hi != 0 {
		or |= maskPV
	}
	a += cpu.AF.Hi
	or |= a&mask3 | (a<<4)&mask5
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func oopLDI(cpu *CPU) {
	de := cpu.DE.U16()
	hl := cpu.HL.U16()
	bc := cpu.BC.U16()
	a := cpu.Memory.Get(hl)
	cpu.Memory.Set(de, a)
	cpu.DE.SetU16(de + 1)
	cpu.HL.SetU16(hl + 1)
	cpu.BC.SetU16(bc - 1)
	cpu.updateFlagLDID(a)
	cpu.LastOpCycles = 16
}

func oopLDIR(cpu *CPU) {
	oopLDI(cpu)
	if cpu.AF.Lo&maskPV != 0 { // cpu.BC != 0
		cpu.PC -= 2
		cpu.LastOpCycles = 21
	} else {
		cpu.LastOpCycles = 16
	}
}

func oopLDD(cpu *CPU) {
	de := cpu.DE.U16()
	hl := cpu.HL.U16()
	bc := cpu.BC.U16()
	a := cpu.Memory.Get(hl)
	cpu.Memory.Set(de, a)
	cpu.DE.SetU16(de - 1)
	cpu.HL.SetU16(hl - 1)
	cpu.BC.SetU16(bc - 1)
	cpu.updateFlagLDID(a)
	cpu.LastOpCycles = 16
}

func oopLDDR(cpu *CPU) {
	oopLDD(cpu)
	if cpu.AF.Lo&maskPV != 0 { // cpu.BC != 0
		cpu.PC -= 2
		cpu.LastOpCycles = 21
	} else {
		cpu.LastOpCycles = 16
	}
}

func (cpu *CPU) updateFlagCPx(r, a, b uint8) {
	c := r ^ a ^ b
	var nand uint8 = maskS53 | maskZ | maskH | maskPV | maskN
	var or uint8
	or |= r & maskS
	if r == 0 {
		or |= maskZ
	}
	or |= c & maskH
	if cpu.BC.Lo != 0 || cpu.BC.Hi != 0 {
		or |= maskPV
	}
	or |= maskN

	r2 := r - (c & 0x10 >> 4)
	or |= r2 & 0x02 << 4 // mask5
	or |= r2 & mask3
	//if r&0x0f == 8 && c&maskH != 0 {
	//	or &= ^uint8(mask3)
	//}

	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func oopCPI(cpu *CPU) {
	hl := cpu.HL.U16()
	bc := cpu.BC.U16()
	a := cpu.AF.Hi
	x := cpu.Memory.Get(hl)
	r := a - x
	cpu.HL.SetU16(hl + 1)
	cpu.BC.SetU16(bc - 1)
	cpu.updateFlagCPx(r, a, x)
	cpu.LastOpCycles = 16
}

func oopCPIR(cpu *CPU) {
	oopCPI(cpu)
	// cpu.BC != 0 && A - (HL) != 0
	if cpu.AF.Lo&maskPV != 0 && cpu.AF.Lo&maskZ == 0 {
		cpu.PC -= 2
		cpu.LastOpCycles = 21
	} else {
		cpu.LastOpCycles = 16
	}
}

func oopCPD(cpu *CPU) {
	hl := cpu.HL.U16()
	bc := cpu.BC.U16()
	a := cpu.AF.Hi
	x := cpu.Memory.Get(hl)
	r := a - x
	cpu.HL.SetU16(hl - 1)
	cpu.BC.SetU16(bc - 1)
	cpu.updateFlagCPx(r, a, x)
	cpu.LastOpCycles = 16
}

func oopCPDR(cpu *CPU) {
	oopCPD(cpu)
	// cpu.BC != 0 && A - (HL) != 0
	if cpu.AF.Lo&maskPV != 0 && cpu.AF.Lo&maskZ == 0 {
		cpu.PC -= 2
		cpu.LastOpCycles = 21
	} else {
		cpu.LastOpCycles = 16
	}
}
