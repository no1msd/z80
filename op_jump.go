package z80

func oopJPnn(cpu *CPU, l, h uint8) {
	cpu.PC = toU16(l, h)
	cpu.LastOpCycles = 10
}

func oopJRe(cpu *CPU, off uint8) {
	cpu.PC = addrOff(cpu.PC, off)
	cpu.LastOpCycles = 12
}

func oopJRCe(cpu *CPU, off uint8) {
	if cpu.flagC() {
		cpu.PC = addrOff(cpu.PC, off)
		cpu.LastOpCycles = 12
	} else {
		cpu.LastOpCycles = 7
	}
}

func oopJRNCe(cpu *CPU, off uint8) {
	if !cpu.flagC() {
		cpu.PC = addrOff(cpu.PC, off)
		cpu.LastOpCycles = 12
	} else {
		cpu.LastOpCycles = 7
	}
}

func oopJRZe(cpu *CPU, off uint8) {
	if cpu.flagZ() {
		cpu.PC = addrOff(cpu.PC, off)
		cpu.LastOpCycles = 12
	} else {
		cpu.LastOpCycles = 7
	}
}

func oopJRNZe(cpu *CPU, off uint8) {
	if !cpu.flagZ() {
		cpu.PC = addrOff(cpu.PC, off)
		cpu.LastOpCycles = 12
	} else {
		cpu.LastOpCycles = 7
	}
}

func oopJPHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.PC = p
	cpu.LastOpCycles = 4
}

func oopJPIXP(cpu *CPU) {
	p := cpu.IX
	cpu.PC = p
	cpu.LastOpCycles = 8
}

func oopJPIYP(cpu *CPU) {
	p := cpu.IY
	cpu.PC = p
	cpu.LastOpCycles = 8
}

func oopDJNZe(cpu *CPU, off uint8) {
	cpu.BC.Hi--
	if cpu.BC.Hi != 0 {
		cpu.PC = addrOff(cpu.PC, off)
		cpu.LastOpCycles = 13
	} else {
		cpu.LastOpCycles = 8
	}
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func xopJPnZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ == 0 {
		cpu.PC = toU16(l, h)
	}
	cpu.LastOpCycles = 10
}

func xopJPfZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ != 0 {
		cpu.PC = toU16(l, h)
	}
	cpu.LastOpCycles = 10
}

func xopJPnCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC == 0 {
		cpu.PC = toU16(l, h)
	}
	cpu.LastOpCycles = 10
}

func xopJPfCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC != 0 {
		cpu.PC = toU16(l, h)
	}
	cpu.LastOpCycles = 10
}

func xopJPnPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV == 0 {
		cpu.PC = toU16(l, h)
	}
	cpu.LastOpCycles = 10
}

func xopJPfPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV != 0 {
		cpu.PC = toU16(l, h)
	}
	cpu.LastOpCycles = 10
}

func xopJPnSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS == 0 {
		cpu.PC = toU16(l, h)
	}
	cpu.LastOpCycles = 10
}

func xopJPfSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS != 0 {
		cpu.PC = toU16(l, h)
	}
	cpu.LastOpCycles = 10
}
