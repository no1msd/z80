package z80

func oopLDHLPn(cpu *CPU, n uint8) {
	p := cpu.HL.U16()
	cpu.Memory.Set(p, n)
	cpu.LastOpCycles = 10
}

func oopLDIXdPn(cpu *CPU, d, n uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, n)
	cpu.LastOpCycles = 19
}

func oopLDIYdPn(cpu *CPU, d, n uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, n)
	cpu.LastOpCycles = 19
}

func oopLDABCP(cpu *CPU) {
	p := cpu.BC.U16()
	cpu.AF.Hi = cpu.Memory.Get(p)
	cpu.LastOpCycles = 7
}

func oopLDADEP(cpu *CPU) {
	p := cpu.DE.U16()
	cpu.AF.Hi = cpu.Memory.Get(p)
	cpu.LastOpCycles = 7
}

func oopLDAnnP(cpu *CPU, l, h uint8) {
	p := toU16(l, h)
	cpu.AF.Hi = cpu.Memory.Get(p)
	cpu.LastOpCycles = 13
}

func oopLDBCPA(cpu *CPU) {
	p := cpu.BC.U16()
	cpu.Memory.Set(p, cpu.AF.Hi)
	cpu.LastOpCycles = 7
}

func oopLDDEPA(cpu *CPU) {
	p := cpu.DE.U16()
	cpu.Memory.Set(p, cpu.AF.Hi)
	cpu.LastOpCycles = 7
}

func oopLDnnPA(cpu *CPU, l, h uint8) {
	p := toU16(l, h)
	cpu.Memory.Set(p, cpu.AF.Hi)
	cpu.LastOpCycles = 13
}

func (cpu *CPU) updateFlagIR(d uint8) {
	var nand uint8 = maskS53 | maskZ | maskH | maskPV | maskN
	var or uint8
	or |= d & maskS53
	if d == 0 {
		or |= maskZ
	}
	if cpu.IFF2 {
		or |= maskPV
	}
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func oopLDAI(cpu *CPU) {
	d := cpu.IR.Hi
	cpu.AF.Hi = d
	// update F by d
	// - S is set if the I Register is negative; otherwise, it is
	//   reset.
	// - Z is set if the I Register is 0; otherwise, it is reset.
	// - H is reset.
	// - P/V contains contents of IFF2.
	// - N is reset.
	// - C is not affected.
	// - If an interrupt occurs during execution of this instruction,
	//   the Parity flag contains a 0.
	cpu.updateFlagIR(d)
	cpu.LastOpCycles = 9
}

func oopLDAR(cpu *CPU) {
	d := cpu.IR.Lo
	cpu.AF.Hi = d
	// update F by d
	// - S is set if, R-Register is negative; otherwise, it is reset.
	// - Z is set if the R Register is 0; otherwise, it is reset.
	// - H is reset.
	// - P/V contains contents of IFF2.
	// - N is reset.
	// - C is not affected.
	// - If an interrupt occurs during execution of this instruction,
	//	 the parity flag contains a 0.
	cpu.updateFlagIR(d)
	cpu.LastOpCycles = 9
}

func oopLDIA(cpu *CPU) {
	cpu.IR.Hi = cpu.AF.Hi
	cpu.LastOpCycles = 9
}

func oopLDRA(cpu *CPU) {
	cpu.IR.Lo = cpu.AF.Hi
	cpu.LastOpCycles = 9
}

func oopLDIXHn(cpu *CPU, n uint8) {
	cpu.IX = uint16(n)<<8 | cpu.IX&0xff
	cpu.LastOpCycles = 14 // guess
}

func oopLDIXLn(cpu *CPU, n uint8) {
	cpu.IX = uint16(n) | cpu.IX&0xff00
	cpu.LastOpCycles = 14 // guess
}

func oopLDIYHn(cpu *CPU, n uint8) {
	cpu.IY = uint16(n)<<8 | cpu.IY&0xff
	cpu.LastOpCycles = 14 // guess
}

func oopLDIYLn(cpu *CPU, n uint8) {
	cpu.IY = uint16(n) | cpu.IY&0xff00
	cpu.LastOpCycles = 14 // guess
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func xopLDbHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.BC.Hi = cpu.Memory.Get(p)
	cpu.LastOpCycles = 7
}

func xopLDcHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.BC.Lo = cpu.Memory.Get(p)
	cpu.LastOpCycles = 7
}

func xopLDdHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.DE.Hi = cpu.Memory.Get(p)
	cpu.LastOpCycles = 7
}

func xopLDeHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.DE.Lo = cpu.Memory.Get(p)
	cpu.LastOpCycles = 7
}

func xopLDhHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.HL.Hi = cpu.Memory.Get(p)
	cpu.LastOpCycles = 7
}

func xopLDlHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.HL.Lo = cpu.Memory.Get(p)
	cpu.LastOpCycles = 7
}

func xopLDaHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.AF.Hi = cpu.Memory.Get(p)
	cpu.LastOpCycles = 7
}

func xopLDHLPb(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.BC.Hi
	cpu.Memory.Set(p, r)
	cpu.LastOpCycles = 7
}

func xopLDHLPc(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.BC.Lo
	cpu.Memory.Set(p, r)
	cpu.LastOpCycles = 7
}

func xopLDHLPd(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.DE.Hi
	cpu.Memory.Set(p, r)
	cpu.LastOpCycles = 7
}

func xopLDHLPe(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.DE.Lo
	cpu.Memory.Set(p, r)
	cpu.LastOpCycles = 7
}

func xopLDHLPh(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.HL.Hi
	cpu.Memory.Set(p, r)
	cpu.LastOpCycles = 7
}

func xopLDHLPl(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.HL.Lo
	cpu.Memory.Set(p, r)
	cpu.LastOpCycles = 7
}

func xopLDHLPa(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.AF.Hi
	cpu.Memory.Set(p, r)
	cpu.LastOpCycles = 7
}

func xopLDbn(cpu *CPU, n uint8) {
	cpu.BC.Hi = n
	cpu.LastOpCycles = 7
}

func xopLDcn(cpu *CPU, n uint8) {
	cpu.BC.Lo = n
	cpu.LastOpCycles = 7
}

func xopLDdn(cpu *CPU, n uint8) {
	cpu.DE.Hi = n
	cpu.LastOpCycles = 7
}

func xopLDen(cpu *CPU, n uint8) {
	cpu.DE.Lo = n
	cpu.LastOpCycles = 7
}

func xopLDhn(cpu *CPU, n uint8) {
	cpu.HL.Hi = n
	cpu.LastOpCycles = 7
}

func xopLDln(cpu *CPU, n uint8) {
	cpu.HL.Lo = n
	cpu.LastOpCycles = 7
}

func xopLDan(cpu *CPU, n uint8) {
	cpu.AF.Hi = n
	cpu.LastOpCycles = 7
}

func xopLDbIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.BC.Hi = cpu.Memory.Get(p)
	cpu.LastOpCycles = 19
}

func xopLDcIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.BC.Lo = cpu.Memory.Get(p)
	cpu.LastOpCycles = 19
}

func xopLDdIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.DE.Hi = cpu.Memory.Get(p)
	cpu.LastOpCycles = 19
}

func xopLDeIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.DE.Lo = cpu.Memory.Get(p)
	cpu.LastOpCycles = 19
}

func xopLDhIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.HL.Hi = cpu.Memory.Get(p)
	cpu.LastOpCycles = 19
}

func xopLDlIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.HL.Lo = cpu.Memory.Get(p)
	cpu.LastOpCycles = 19
}

func xopLDaIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.AF.Hi = cpu.Memory.Get(p)
	cpu.LastOpCycles = 19
}

func xopLDbIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.BC.Hi = cpu.Memory.Get(p)
	cpu.LastOpCycles = 19
}

func xopLDcIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.BC.Lo = cpu.Memory.Get(p)
	cpu.LastOpCycles = 19
}

func xopLDdIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.DE.Hi = cpu.Memory.Get(p)
	cpu.LastOpCycles = 19
}

func xopLDeIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.DE.Lo = cpu.Memory.Get(p)
	cpu.LastOpCycles = 19
}

func xopLDhIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.HL.Hi = cpu.Memory.Get(p)
	cpu.LastOpCycles = 19
}

func xopLDlIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.HL.Lo = cpu.Memory.Get(p)
	cpu.LastOpCycles = 19
}

func xopLDaIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.AF.Hi = cpu.Memory.Get(p)
	cpu.LastOpCycles = 19
}

func xopLDIXdPb(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.BC.Hi)
	cpu.LastOpCycles = 19
}

func xopLDIXdPc(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.BC.Lo)
	cpu.LastOpCycles = 19
}

func xopLDIXdPd(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.DE.Hi)
	cpu.LastOpCycles = 19
}

func xopLDIXdPe(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.DE.Lo)
	cpu.LastOpCycles = 19
}

func xopLDIXdPh(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.HL.Hi)
	cpu.LastOpCycles = 19
}

func xopLDIXdPl(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.HL.Lo)
	cpu.LastOpCycles = 19
}

func xopLDIXdPa(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.AF.Hi)
	cpu.LastOpCycles = 19
}

func xopLDIYdPb(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.BC.Hi)
	cpu.LastOpCycles = 19
}

func xopLDIYdPc(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.BC.Lo)
	cpu.LastOpCycles = 19
}

func xopLDIYdPd(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.DE.Hi)
	cpu.LastOpCycles = 19
}

func xopLDIYdPe(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.DE.Lo)
	cpu.LastOpCycles = 19
}

func xopLDIYdPh(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.HL.Hi)
	cpu.LastOpCycles = 19
}

func xopLDIYdPl(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.HL.Lo)
	cpu.LastOpCycles = 19
}

func xopLDIYdPa(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.AF.Hi)
	cpu.LastOpCycles = 19
}
