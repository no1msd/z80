package z80

import "math/bits"

func (cpu *CPU) updateFlagRL(r uint8) {
	var nand uint8 = mask53 | maskH | maskN | maskC
	var or = (r >> 7) & maskC
	or |= r << 1 & mask53
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func oopRLCA(cpu *CPU) {
	a := cpu.AF.Hi
	a2 := a<<1 | a>>7
	cpu.AF.Hi = a2
	cpu.updateFlagRL(a)
	cpu.LastOpCycles = 4
}

func oopRLA(cpu *CPU) {
	a := cpu.AF.Hi
	a2 := a<<1 | cpu.AF.Lo&maskC
	cpu.AF.Hi = a2
	cpu.updateFlagRL(a)
	cpu.LastOpCycles = 4
}

func (cpu *CPU) updateFlagRR(r uint8) {
	var nand uint8 = mask53 | maskH | maskN | maskC
	var or = r & maskC
	or |= r >> 1 & mask53
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func oopRRCA(cpu *CPU) {
	a := cpu.AF.Hi
	a2 := a>>1 | a<<7
	cpu.AF.Hi = a2
	cpu.updateFlagRR(a)
	cpu.LastOpCycles = 4
}

func oopRRA(cpu *CPU) {
	a := cpu.AF.Hi
	a2 := a >> 1
	if cpu.flagC() {
		a2 |= 0x80
	}
	cpu.AF.Hi = a2
	cpu.updateFlagRR(a)
	cpu.LastOpCycles = 4
}

func oopRLCIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.rlcU8(cpu.Memory.Get(p)))
	cpu.LastOpCycles = 23
}

func oopRLCIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.rlcU8(cpu.Memory.Get(p)))
	cpu.LastOpCycles = 23
}

func oopRLIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.rlU8(cpu.Memory.Get(p)))
	cpu.LastOpCycles = 23
}

func oopRLIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.rlU8(cpu.Memory.Get(p)))
	cpu.LastOpCycles = 23
}

func oopRRCIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.rrcU8(cpu.Memory.Get(p)))
	cpu.LastOpCycles = 23
}

func oopRRCIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.rrcU8(cpu.Memory.Get(p)))
	cpu.LastOpCycles = 23
}

func oopRRIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.rrU8(cpu.Memory.Get(p)))
	cpu.LastOpCycles = 23
}

func oopRRIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.rrU8(cpu.Memory.Get(p)))
	cpu.LastOpCycles = 23
}

func oopSLAIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.slaU8(cpu.Memory.Get(p)))
	cpu.LastOpCycles = 23
}

func oopSLAIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.slaU8(cpu.Memory.Get(p)))
	cpu.LastOpCycles = 23
}

func oopSRAIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.sraU8(cpu.Memory.Get(p)))
	cpu.LastOpCycles = 23
}

func oopSRAIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.sraU8(cpu.Memory.Get(p)))
	cpu.LastOpCycles = 23
}

func oopSRLIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.srlU8(cpu.Memory.Get(p)))
	cpu.LastOpCycles = 23
}

func oopSRLIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.srlU8(cpu.Memory.Get(p)))
	cpu.LastOpCycles = 23
}

func (cpu *CPU) updateFlagRxD(r uint8) {
	var nand uint8 = maskS53 | maskZ | maskH | maskPV | maskN
	var or uint8
	or |= uint8(r) & maskS53
	if uint8(r) == 0 {
		or |= maskZ
	}
	or |= (uint8(bits.OnesCount8(r)%2) - 1) & maskPV
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func oopRLD(cpu *CPU) {
	p := cpu.HL.U16()
	a := cpu.AF.Hi
	b := cpu.Memory.Get(p)
	a2 := a&0xf0 | b>>4
	b2 := b<<4 | a&0x0f
	cpu.Memory.Set(p, b2)
	cpu.AF.Hi = a2
	cpu.updateFlagRxD(a2)
	cpu.LastOpCycles = 18
}

func oopRRD(cpu *CPU) {
	p := cpu.HL.U16()
	a := cpu.AF.Hi
	b := cpu.Memory.Get(p)
	a2 := a&0xf0 | b&0x0f
	b2 := a<<4 | b>>4
	cpu.Memory.Set(p, b2)
	cpu.AF.Hi = a2
	cpu.updateFlagRxD(a2)
	cpu.LastOpCycles = 18
}

func oopSL1IXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	v := cpu.sl1U8(cpu.Memory.Get(p))
	cpu.Memory.Set(p, v)
	cpu.LastOpCycles = 23 // guess
}

func oopSL1IYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	v := cpu.sl1U8(cpu.Memory.Get(p))
	cpu.Memory.Set(p, v)
	cpu.LastOpCycles = 23 // guess
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func xopRLCb(cpu *CPU) {
	cpu.BC.Hi = cpu.rlcU8(cpu.BC.Hi)
	cpu.LastOpCycles = 8
}

func xopRLCc(cpu *CPU) {
	cpu.BC.Lo = cpu.rlcU8(cpu.BC.Lo)
	cpu.LastOpCycles = 8
}

func xopRLCd(cpu *CPU) {
	cpu.DE.Hi = cpu.rlcU8(cpu.DE.Hi)
	cpu.LastOpCycles = 8
}

func xopRLCe(cpu *CPU) {
	cpu.DE.Lo = cpu.rlcU8(cpu.DE.Lo)
	cpu.LastOpCycles = 8
}

func xopRLCh(cpu *CPU) {
	cpu.HL.Hi = cpu.rlcU8(cpu.HL.Hi)
	cpu.LastOpCycles = 8
}

func xopRLCl(cpu *CPU) {
	cpu.HL.Lo = cpu.rlcU8(cpu.HL.Lo)
	cpu.LastOpCycles = 8
}

func xopRLCa(cpu *CPU) {
	cpu.AF.Hi = cpu.rlcU8(cpu.AF.Hi)
	cpu.LastOpCycles = 8
}

func xopRLCHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.rlcU8(x)
	cpu.Memory.Set(p, x)
	cpu.LastOpCycles = 15
}

func xopRRCb(cpu *CPU) {
	cpu.BC.Hi = cpu.rrcU8(cpu.BC.Hi)
	cpu.LastOpCycles = 8
}

func xopRRCc(cpu *CPU) {
	cpu.BC.Lo = cpu.rrcU8(cpu.BC.Lo)
	cpu.LastOpCycles = 8
}

func xopRRCd(cpu *CPU) {
	cpu.DE.Hi = cpu.rrcU8(cpu.DE.Hi)
	cpu.LastOpCycles = 8
}

func xopRRCe(cpu *CPU) {
	cpu.DE.Lo = cpu.rrcU8(cpu.DE.Lo)
	cpu.LastOpCycles = 8
}

func xopRRCh(cpu *CPU) {
	cpu.HL.Hi = cpu.rrcU8(cpu.HL.Hi)
	cpu.LastOpCycles = 8
}

func xopRRCl(cpu *CPU) {
	cpu.HL.Lo = cpu.rrcU8(cpu.HL.Lo)
	cpu.LastOpCycles = 8
}

func xopRRCa(cpu *CPU) {
	cpu.AF.Hi = cpu.rrcU8(cpu.AF.Hi)
	cpu.LastOpCycles = 8
}

func xopRRCHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.rrcU8(x)
	cpu.Memory.Set(p, x)
	cpu.LastOpCycles = 15
}

func xopRLb(cpu *CPU) {
	cpu.BC.Hi = cpu.rlU8(cpu.BC.Hi)
	cpu.LastOpCycles = 8
}

func xopRLc(cpu *CPU) {
	cpu.BC.Lo = cpu.rlU8(cpu.BC.Lo)
	cpu.LastOpCycles = 8
}

func xopRLd(cpu *CPU) {
	cpu.DE.Hi = cpu.rlU8(cpu.DE.Hi)
	cpu.LastOpCycles = 8
}

func xopRLe(cpu *CPU) {
	cpu.DE.Lo = cpu.rlU8(cpu.DE.Lo)
	cpu.LastOpCycles = 8
}

func xopRLh(cpu *CPU) {
	cpu.HL.Hi = cpu.rlU8(cpu.HL.Hi)
	cpu.LastOpCycles = 8
}

func xopRLl(cpu *CPU) {
	cpu.HL.Lo = cpu.rlU8(cpu.HL.Lo)
	cpu.LastOpCycles = 8
}

func xopRLa(cpu *CPU) {
	cpu.AF.Hi = cpu.rlU8(cpu.AF.Hi)
	cpu.LastOpCycles = 8
}

func xopRLHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.rlU8(x)
	cpu.Memory.Set(p, x)
	cpu.LastOpCycles = 15
}

func xopRRb(cpu *CPU) {
	cpu.BC.Hi = cpu.rrU8(cpu.BC.Hi)
	cpu.LastOpCycles = 8
}

func xopRRc(cpu *CPU) {
	cpu.BC.Lo = cpu.rrU8(cpu.BC.Lo)
	cpu.LastOpCycles = 8
}

func xopRRd(cpu *CPU) {
	cpu.DE.Hi = cpu.rrU8(cpu.DE.Hi)
	cpu.LastOpCycles = 8
}

func xopRRe(cpu *CPU) {
	cpu.DE.Lo = cpu.rrU8(cpu.DE.Lo)
	cpu.LastOpCycles = 8
}

func xopRRh(cpu *CPU) {
	cpu.HL.Hi = cpu.rrU8(cpu.HL.Hi)
	cpu.LastOpCycles = 8
}

func xopRRl(cpu *CPU) {
	cpu.HL.Lo = cpu.rrU8(cpu.HL.Lo)
	cpu.LastOpCycles = 8
}

func xopRRa(cpu *CPU) {
	cpu.AF.Hi = cpu.rrU8(cpu.AF.Hi)
	cpu.LastOpCycles = 8
}

func xopRRHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.rrU8(x)
	cpu.Memory.Set(p, x)
	cpu.LastOpCycles = 15
}

func xopSLAb(cpu *CPU) {
	cpu.BC.Hi = cpu.slaU8(cpu.BC.Hi)
	cpu.LastOpCycles = 8
}

func xopSLAc(cpu *CPU) {
	cpu.BC.Lo = cpu.slaU8(cpu.BC.Lo)
	cpu.LastOpCycles = 8
}

func xopSLAd(cpu *CPU) {
	cpu.DE.Hi = cpu.slaU8(cpu.DE.Hi)
	cpu.LastOpCycles = 8
}

func xopSLAe(cpu *CPU) {
	cpu.DE.Lo = cpu.slaU8(cpu.DE.Lo)
	cpu.LastOpCycles = 8
}

func xopSLAh(cpu *CPU) {
	cpu.HL.Hi = cpu.slaU8(cpu.HL.Hi)
	cpu.LastOpCycles = 8
}

func xopSLAl(cpu *CPU) {
	cpu.HL.Lo = cpu.slaU8(cpu.HL.Lo)
	cpu.LastOpCycles = 8
}

func xopSLAa(cpu *CPU) {
	cpu.AF.Hi = cpu.slaU8(cpu.AF.Hi)
	cpu.LastOpCycles = 8
}

func xopSLAHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.slaU8(x)
	cpu.Memory.Set(p, x)
	cpu.LastOpCycles = 15
}

func xopSRAb(cpu *CPU) {
	cpu.BC.Hi = cpu.sraU8(cpu.BC.Hi)
	cpu.LastOpCycles = 8
}

func xopSRAc(cpu *CPU) {
	cpu.BC.Lo = cpu.sraU8(cpu.BC.Lo)
	cpu.LastOpCycles = 8
}

func xopSRAd(cpu *CPU) {
	cpu.DE.Hi = cpu.sraU8(cpu.DE.Hi)
	cpu.LastOpCycles = 8
}

func xopSRAe(cpu *CPU) {
	cpu.DE.Lo = cpu.sraU8(cpu.DE.Lo)
	cpu.LastOpCycles = 8
}

func xopSRAh(cpu *CPU) {
	cpu.HL.Hi = cpu.sraU8(cpu.HL.Hi)
	cpu.LastOpCycles = 8
}

func xopSRAl(cpu *CPU) {
	cpu.HL.Lo = cpu.sraU8(cpu.HL.Lo)
	cpu.LastOpCycles = 8
}

func xopSRAa(cpu *CPU) {
	cpu.AF.Hi = cpu.sraU8(cpu.AF.Hi)
	cpu.LastOpCycles = 8
}

func xopSRAHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.sraU8(x)
	cpu.Memory.Set(p, x)
	cpu.LastOpCycles = 15
}

func xopSL1b(cpu *CPU) {
	cpu.BC.Hi = cpu.sl1U8(cpu.BC.Hi)
	cpu.LastOpCycles = 8
}

func xopSL1c(cpu *CPU) {
	cpu.BC.Lo = cpu.sl1U8(cpu.BC.Lo)
	cpu.LastOpCycles = 8
}

func xopSL1d(cpu *CPU) {
	cpu.DE.Hi = cpu.sl1U8(cpu.DE.Hi)
	cpu.LastOpCycles = 8
}

func xopSL1e(cpu *CPU) {
	cpu.DE.Lo = cpu.sl1U8(cpu.DE.Lo)
	cpu.LastOpCycles = 8
}

func xopSL1h(cpu *CPU) {
	cpu.HL.Hi = cpu.sl1U8(cpu.HL.Hi)
	cpu.LastOpCycles = 8
}

func xopSL1l(cpu *CPU) {
	cpu.HL.Lo = cpu.sl1U8(cpu.HL.Lo)
	cpu.LastOpCycles = 8
}

func xopSL1a(cpu *CPU) {
	cpu.AF.Hi = cpu.sl1U8(cpu.AF.Hi)
	cpu.LastOpCycles = 8
}

func xopSL1HLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.sl1U8(x)
	cpu.Memory.Set(p, x)
	cpu.LastOpCycles = 8
}

func xopSRLb(cpu *CPU) {
	cpu.BC.Hi = cpu.srlU8(cpu.BC.Hi)
	cpu.LastOpCycles = 8
}

func xopSRLc(cpu *CPU) {
	cpu.BC.Lo = cpu.srlU8(cpu.BC.Lo)
	cpu.LastOpCycles = 8
}

func xopSRLd(cpu *CPU) {
	cpu.DE.Hi = cpu.srlU8(cpu.DE.Hi)
	cpu.LastOpCycles = 8
}

func xopSRLe(cpu *CPU) {
	cpu.DE.Lo = cpu.srlU8(cpu.DE.Lo)
	cpu.LastOpCycles = 8
}

func xopSRLh(cpu *CPU) {
	cpu.HL.Hi = cpu.srlU8(cpu.HL.Hi)
	cpu.LastOpCycles = 8
}

func xopSRLl(cpu *CPU) {
	cpu.HL.Lo = cpu.srlU8(cpu.HL.Lo)
	cpu.LastOpCycles = 8
}

func xopSRLa(cpu *CPU) {
	cpu.AF.Hi = cpu.srlU8(cpu.AF.Hi)
	cpu.LastOpCycles = 8
}

func xopSRLHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.srlU8(x)
	cpu.Memory.Set(p, x)
	cpu.LastOpCycles = 15
}
