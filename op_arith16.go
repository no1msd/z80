package z80

func oopINCIX(cpu *CPU) {
	cpu.IX = cpu.incU16(cpu.IX)
	cpu.LastOpCycles = 10
}

func oopINCIY(cpu *CPU) {
	cpu.IY = cpu.incU16(cpu.IY)
	cpu.LastOpCycles = 10
}

func oopDECIX(cpu *CPU) {
	cpu.IX = cpu.decU16(cpu.IX)
	cpu.LastOpCycles = 10
}

func oopDECIY(cpu *CPU) {
	cpu.IY = cpu.decU16(cpu.IY)
	cpu.LastOpCycles = 10
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func xopINCbc(cpu *CPU) {
	cpu.BC.Lo++
	if cpu.BC.Lo == 0 {
		cpu.BC.Hi++
	}
	cpu.LastOpCycles = 6
}

func xopINCde(cpu *CPU) {
	cpu.DE.Lo++
	if cpu.DE.Lo == 0 {
		cpu.DE.Hi++
	}
	cpu.LastOpCycles = 6
}

func xopINChl(cpu *CPU) {
	cpu.HL.Lo++
	if cpu.HL.Lo == 0 {
		cpu.HL.Hi++
	}
	cpu.LastOpCycles = 6
}

func xopINCsp(cpu *CPU) {
	cpu.SP++
	cpu.LastOpCycles = 6
}

func xopDECbc(cpu *CPU) {
	cpu.BC.Lo--
	if cpu.BC.Lo == 0xff {
		cpu.BC.Hi--
	}
	cpu.LastOpCycles = 6
}

func xopDECde(cpu *CPU) {
	cpu.DE.Lo--
	if cpu.DE.Lo == 0xff {
		cpu.DE.Hi--
	}
	cpu.LastOpCycles = 6
}

func xopDEChl(cpu *CPU) {
	cpu.HL.Lo--
	if cpu.HL.Lo == 0xff {
		cpu.HL.Hi--
	}
	cpu.LastOpCycles = 6
}

func xopDECsp(cpu *CPU) {
	cpu.SP--
	cpu.LastOpCycles = 6
}

func xopADDHLbc(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.BC.U16()
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
	cpu.LastOpCycles = 11
}

func xopADDHLde(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.DE.U16()
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
	cpu.LastOpCycles = 11
}

func xopADDHLhl(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.HL.U16()
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
	cpu.LastOpCycles = 11
}

func xopADDHLsp(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.SP
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
	cpu.LastOpCycles = 11
}

func xopADDIXbc(cpu *CPU) {
	a := cpu.IX
	x := cpu.BC.U16()
	cpu.IX = cpu.addU16(a, x)
	cpu.LastOpCycles = 15
}

func xopADDIXde(cpu *CPU) {
	a := cpu.IX
	x := cpu.DE.U16()
	cpu.IX = cpu.addU16(a, x)
	cpu.LastOpCycles = 15
}

func xopADDIXix(cpu *CPU) {
	a := cpu.IX
	x := cpu.IX
	cpu.IX = cpu.addU16(a, x)
	cpu.LastOpCycles = 15
}

func xopADDIXsp(cpu *CPU) {
	a := cpu.IX
	x := cpu.SP
	cpu.IX = cpu.addU16(a, x)
	cpu.LastOpCycles = 15
}

func xopADDIYbc(cpu *CPU) {
	a := cpu.IY
	x := cpu.BC.U16()
	cpu.IY = cpu.addU16(a, x)
	cpu.LastOpCycles = 15
}

func xopADDIYde(cpu *CPU) {
	a := cpu.IY
	x := cpu.DE.U16()
	cpu.IY = cpu.addU16(a, x)
	cpu.LastOpCycles = 15
}

func xopADDIYiy(cpu *CPU) {
	a := cpu.IY
	x := cpu.IY
	cpu.IY = cpu.addU16(a, x)
	cpu.LastOpCycles = 15
}

func xopADDIYsp(cpu *CPU) {
	a := cpu.IY
	x := cpu.SP
	cpu.IY = cpu.addU16(a, x)
	cpu.LastOpCycles = 15
}

func xopSBCHLbc(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.BC.U16()
	cpu.HL.SetU16(cpu.sbcU16(a, x))
	cpu.LastOpCycles = 15
}

func xopSBCHLde(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.DE.U16()
	cpu.HL.SetU16(cpu.sbcU16(a, x))
	cpu.LastOpCycles = 15
}

func xopSBCHLhl(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.HL.U16()
	cpu.HL.SetU16(cpu.sbcU16(a, x))
	cpu.LastOpCycles = 15
}

func xopSBCHLsp(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.SP
	cpu.HL.SetU16(cpu.sbcU16(a, x))
	cpu.LastOpCycles = 15
}

func xopADCHLbc(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.BC.U16()
	cpu.HL.SetU16(cpu.adcU16(a, x))
	cpu.LastOpCycles = 15
}

func xopADCHLde(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.DE.U16()
	cpu.HL.SetU16(cpu.adcU16(a, x))
	cpu.LastOpCycles = 15
}

func xopADCHLhl(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.HL.U16()
	cpu.HL.SetU16(cpu.adcU16(a, x))
	cpu.LastOpCycles = 15
}

func xopADCHLsp(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.SP
	cpu.HL.SetU16(cpu.adcU16(a, x))
	cpu.LastOpCycles = 15
}
