package z80

import "math/bits"

func oopINAnP(cpu *CPU, n uint8) {
	cpu.AF.Hi = cpu.ioIn(n)
	cpu.LastOpCycles = 11
}

func (cpu *CPU) updateFlagIObZ() {
	var nand uint8 = maskZ
	var or uint8
	if cpu.BC.Hi == 0 {
		or |= maskZ
	}
	or |= maskN
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func oopINI(cpu *CPU) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.updateFlagIObZ()
	cpu.LastOpCycles = 16
}

func oopINIR(cpu *CPU) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.updateFlagIObZ()
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
		cpu.LastOpCycles = 21
	} else {
		cpu.LastOpCycles = 16
	}
}

func oopIND(cpu *CPU) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.updateFlagIObZ()
	cpu.LastOpCycles = 16
}

func oopINDR(cpu *CPU) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.updateFlagIObZ()
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
		cpu.LastOpCycles = 21
	} else {
		cpu.LastOpCycles = 16
	}
}

func oopOUTnPA(cpu *CPU, n uint8) {
	cpu.ioOut(n, cpu.AF.Hi)
	cpu.LastOpCycles = 11
}

func oopOUTI(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.updateFlagIObZ()
	cpu.LastOpCycles = 16
}

func oopOTIR(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.updateFlagIObZ()
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
		cpu.LastOpCycles = 21
	} else {
		cpu.LastOpCycles = 16
	}
}

func oopOUTD(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.updateFlagIObZ()
	cpu.LastOpCycles = 16
}

func oopOTDR(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.updateFlagIObZ()
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
		cpu.LastOpCycles = 21
	} else {
		cpu.LastOpCycles = 16
	}
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func (cpu *CPU) updateIOIn(r uint8) {
	var nand uint8 = maskS53 | maskZ | maskH | maskPV | maskN
	var or uint8
	or |= r & maskS53
	if r == 0 {
		or |= maskZ
	}
	or |= (uint8(bits.OnesCount8(r)%2) - 1) & maskPV
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func xopINbCP(cpu *CPU) {
	r := cpu.ioIn(cpu.BC.Lo)
	cpu.BC.Hi = r
	cpu.updateIOIn(r)
	cpu.LastOpCycles = 12
}

func xopINcCP(cpu *CPU) {
	r := cpu.ioIn(cpu.BC.Lo)
	cpu.BC.Lo = r
	cpu.updateIOIn(r)
	cpu.LastOpCycles = 12
}

func xopINdCP(cpu *CPU) {
	r := cpu.ioIn(cpu.BC.Lo)
	cpu.DE.Hi = r
	cpu.updateIOIn(r)
	cpu.LastOpCycles = 12
}

func xopINeCP(cpu *CPU) {
	r := cpu.ioIn(cpu.BC.Lo)
	cpu.DE.Lo = r
	cpu.updateIOIn(r)
	cpu.LastOpCycles = 12
}

func xopINhCP(cpu *CPU) {
	r := cpu.ioIn(cpu.BC.Lo)
	cpu.HL.Hi = r
	cpu.updateIOIn(r)
	cpu.LastOpCycles = 12
}

func xopINlCP(cpu *CPU) {
	r := cpu.ioIn(cpu.BC.Lo)
	cpu.HL.Lo = r
	cpu.updateIOIn(r)
	cpu.LastOpCycles = 12
}

func xopINaCP(cpu *CPU) {
	r := cpu.ioIn(cpu.BC.Lo)
	cpu.AF.Hi = r
	cpu.updateIOIn(r)
	cpu.LastOpCycles = 12
}

func xopOUTCPb(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.BC.Hi)
	cpu.LastOpCycles = 12
}

func xopOUTCPc(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.BC.Lo)
	cpu.LastOpCycles = 12
}

func xopOUTCPd(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.DE.Hi)
	cpu.LastOpCycles = 12
}

func xopOUTCPe(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.DE.Lo)
	cpu.LastOpCycles = 12
}

func xopOUTCPh(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.HL.Hi)
	cpu.LastOpCycles = 12
}

func xopOUTCPl(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.HL.Lo)
	cpu.LastOpCycles = 12
}

func xopOUTCPa(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.AF.Hi)
	cpu.LastOpCycles = 12
}
