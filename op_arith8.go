package z80

func oopADDAn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.AF.Hi = cpu.addU8(a, n)
	cpu.LastOpCycles = 7
}

func oopADDAHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.addU8(a, x)
	cpu.LastOpCycles = 7
}

func oopADDAIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.addU8(a, x)
	cpu.LastOpCycles = 19
}

func oopADDAIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.addU8(a, x)
	cpu.LastOpCycles = 19
}

func oopADCAn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.AF.Hi = cpu.adcU8(a, n)
	cpu.LastOpCycles = 7
}

func oopADCAHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.adcU8(a, x)
	cpu.LastOpCycles = 7
}

func oopADCAIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.adcU8(a, x)
	cpu.LastOpCycles = 19
}

func oopADCAIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.adcU8(a, x)
	cpu.LastOpCycles = 19
}

func oopSUBAn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.AF.Hi = cpu.subU8(a, n)
	cpu.LastOpCycles = 7
}

func oopSUBAHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.subU8(a, x)
	cpu.LastOpCycles = 7
}

func oopSUBAIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.subU8(a, x)
	cpu.LastOpCycles = 19
}

func oopSUBAIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.subU8(a, x)
	cpu.LastOpCycles = 19
}

func oopSBCAn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.AF.Hi = cpu.sbcU8(a, n)
	cpu.LastOpCycles = 7
}

func oopSBCAHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.sbcU8(a, x)
	cpu.LastOpCycles = 7
}

func oopSBCAIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.sbcU8(a, x)
	cpu.LastOpCycles = 19
}

func oopSBCAIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.sbcU8(a, x)
	cpu.LastOpCycles = 19
}

func oopANDn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.AF.Hi = cpu.andU8(a, n)
	cpu.LastOpCycles = 7
}

func oopANDHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.andU8(a, x)
	cpu.LastOpCycles = 7
}

func oopANDIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.andU8(a, x)
	cpu.LastOpCycles = 19
}

func oopANDIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.andU8(a, x)
	cpu.LastOpCycles = 19
}

func oopORn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.AF.Hi = cpu.orU8(a, n)
	cpu.LastOpCycles = 7
}

func oopORHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.orU8(a, x)
	cpu.LastOpCycles = 7
}

func oopORIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.orU8(a, x)
	cpu.LastOpCycles = 19
}

func oopORIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.orU8(a, x)
	cpu.LastOpCycles = 19
}

func oopXORn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.AF.Hi = cpu.xorU8(a, n)
	cpu.LastOpCycles = 7
}

func oopXORHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.xorU8(a, x)
	cpu.LastOpCycles = 7
}

func oopXORIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.xorU8(a, x)
	cpu.LastOpCycles = 19
}

func oopXORIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.xorU8(a, x)
	cpu.LastOpCycles = 19
}

func oopCPn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.cpU8(a, n)
	cpu.LastOpCycles = 7
}

func oopCPHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.cpU8(a, x)
	cpu.LastOpCycles = 7
}

func oopCPIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.cpU8(a, x)
	cpu.LastOpCycles = 19
}

func oopCPIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.cpU8(a, x)
	cpu.LastOpCycles = 19
}

func oopINCHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.incU8(x))
	cpu.LastOpCycles = 11
}

func oopINCIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.incU8(x))
	cpu.LastOpCycles = 23
}

func oopINCIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.incU8(x))
	cpu.LastOpCycles = 23
}

func oopDECHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.decU8(x))
	cpu.LastOpCycles = 11
}

func oopDECIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.decU8(x))
	cpu.LastOpCycles = 23
}

func oopDECIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.decU8(x))
	cpu.LastOpCycles = 23
}

func oopINCIXH(cpu *CPU) {
	v := cpu.incU8(uint8(cpu.IX >> 8))
	cpu.IX = uint16(v)<<8 | cpu.IX&0xff
	cpu.LastOpCycles = 10 // guess
}

func oopDECIXH(cpu *CPU) {
	v := cpu.decU8(uint8(cpu.IX >> 8))
	cpu.IX = uint16(v)<<8 | cpu.IX&0xff
	cpu.LastOpCycles = 10 // guess
}

func oopINCIXL(cpu *CPU) {
	v := cpu.incU8(uint8(cpu.IX))
	cpu.IX = uint16(v) | cpu.IX&0xff00
	cpu.LastOpCycles = 10 // guess
}

func oopDECIXL(cpu *CPU) {
	v := cpu.decU8(uint8(cpu.IX))
	cpu.IX = uint16(v) | cpu.IX&0xff00
	cpu.LastOpCycles = 10 // guess
}

func oopINCIYH(cpu *CPU) {
	v := cpu.incU8(uint8(cpu.IY >> 8))
	cpu.IY = uint16(v)<<8 | cpu.IY&0xff
	cpu.LastOpCycles = 10 // guess
}

func oopDECIYH(cpu *CPU) {
	v := cpu.decU8(uint8(cpu.IY >> 8))
	cpu.IY = uint16(v)<<8 | cpu.IY&0xff
	cpu.LastOpCycles = 10 // guess
}

func oopINCIYL(cpu *CPU) {
	v := cpu.incU8(uint8(cpu.IY))
	cpu.IY = uint16(v) | cpu.IY&0xff00
	cpu.LastOpCycles = 10 // guess
}

func oopDECIYL(cpu *CPU) {
	v := cpu.decU8(uint8(cpu.IY))
	cpu.IY = uint16(v) | cpu.IY&0xff00
	cpu.LastOpCycles = 10 // guess
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func xopADDAb(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Hi
	cpu.AF.Hi = cpu.addU8(a, x)
	cpu.LastOpCycles = 4
}

func xopADDAc(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Lo
	cpu.AF.Hi = cpu.addU8(a, x)
	cpu.LastOpCycles = 4
}

func xopADDAd(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Hi
	cpu.AF.Hi = cpu.addU8(a, x)
	cpu.LastOpCycles = 4
}

func xopADDAe(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Lo
	cpu.AF.Hi = cpu.addU8(a, x)
	cpu.LastOpCycles = 4
}

func xopADDAh(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Hi
	cpu.AF.Hi = cpu.addU8(a, x)
	cpu.LastOpCycles = 4
}

func xopADDAl(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Lo
	cpu.AF.Hi = cpu.addU8(a, x)
	cpu.LastOpCycles = 4
}

func xopADDAa(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.AF.Hi
	cpu.AF.Hi = cpu.addU8(a, x)
	cpu.LastOpCycles = 4
}

func xopADDAixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.AF.Hi = cpu.addU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopADDAixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.AF.Hi = cpu.addU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopADDAiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.AF.Hi = cpu.addU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopADDAiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
	cpu.AF.Hi = cpu.addU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopADCAb(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Hi
	cpu.AF.Hi = cpu.adcU8(a, x)
	cpu.LastOpCycles = 4
}

func xopADCAc(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Lo
	cpu.AF.Hi = cpu.adcU8(a, x)
	cpu.LastOpCycles = 4
}

func xopADCAd(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Hi
	cpu.AF.Hi = cpu.adcU8(a, x)
	cpu.LastOpCycles = 4
}

func xopADCAe(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Lo
	cpu.AF.Hi = cpu.adcU8(a, x)
	cpu.LastOpCycles = 4
}

func xopADCAh(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Hi
	cpu.AF.Hi = cpu.adcU8(a, x)
	cpu.LastOpCycles = 4
}

func xopADCAl(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Lo
	cpu.AF.Hi = cpu.adcU8(a, x)
	cpu.LastOpCycles = 4
}

func xopADCAa(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.AF.Hi
	cpu.AF.Hi = cpu.adcU8(a, x)
	cpu.LastOpCycles = 4
}

func xopADCAixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.AF.Hi = cpu.adcU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopADCAixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.AF.Hi = cpu.adcU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopADCAiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.AF.Hi = cpu.adcU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopADCAiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
	cpu.AF.Hi = cpu.adcU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopSUBAb(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Hi
	cpu.AF.Hi = cpu.subU8(a, x)
	cpu.LastOpCycles = 4
}

func xopSUBAc(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Lo
	cpu.AF.Hi = cpu.subU8(a, x)
	cpu.LastOpCycles = 4
}

func xopSUBAd(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Hi
	cpu.AF.Hi = cpu.subU8(a, x)
	cpu.LastOpCycles = 4
}

func xopSUBAe(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Lo
	cpu.AF.Hi = cpu.subU8(a, x)
	cpu.LastOpCycles = 4
}

func xopSUBAh(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Hi
	cpu.AF.Hi = cpu.subU8(a, x)
	cpu.LastOpCycles = 4
}

func xopSUBAl(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Lo
	cpu.AF.Hi = cpu.subU8(a, x)
	cpu.LastOpCycles = 4
}

func xopSUBAa(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.AF.Hi
	cpu.AF.Hi = cpu.subU8(a, x)
	cpu.LastOpCycles = 4
}

func xopSUBAixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.AF.Hi = cpu.subU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopSUBAixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.AF.Hi = cpu.subU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopSUBAiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.AF.Hi = cpu.subU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopSUBAiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
	cpu.AF.Hi = cpu.subU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopSBCAb(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Hi
	cpu.AF.Hi = cpu.sbcU8(a, x)
	cpu.LastOpCycles = 4
}

func xopSBCAc(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Lo
	cpu.AF.Hi = cpu.sbcU8(a, x)
	cpu.LastOpCycles = 4
}

func xopSBCAd(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Hi
	cpu.AF.Hi = cpu.sbcU8(a, x)
	cpu.LastOpCycles = 4
}

func xopSBCAe(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Lo
	cpu.AF.Hi = cpu.sbcU8(a, x)
	cpu.LastOpCycles = 4
}

func xopSBCAh(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Hi
	cpu.AF.Hi = cpu.sbcU8(a, x)
	cpu.LastOpCycles = 4
}

func xopSBCAl(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Lo
	cpu.AF.Hi = cpu.sbcU8(a, x)
	cpu.LastOpCycles = 4
}

func xopSBCAa(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.AF.Hi
	cpu.AF.Hi = cpu.sbcU8(a, x)
	cpu.LastOpCycles = 4
}

func xopSBCAixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.AF.Hi = cpu.sbcU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopSBCAixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.AF.Hi = cpu.sbcU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopSBCAiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.AF.Hi = cpu.sbcU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopSBCAiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
	cpu.AF.Hi = cpu.sbcU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopANDAb(cpu *CPU) {
	cpu.AF.Hi &= cpu.BC.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, true)
	cpu.LastOpCycles = 4
}

func xopANDAc(cpu *CPU) {
	cpu.AF.Hi &= cpu.BC.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, true)
	cpu.LastOpCycles = 4
}

func xopANDAd(cpu *CPU) {
	cpu.AF.Hi &= cpu.DE.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, true)
	cpu.LastOpCycles = 4
}

func xopANDAe(cpu *CPU) {
	cpu.AF.Hi &= cpu.DE.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, true)
	cpu.LastOpCycles = 4
}

func xopANDAh(cpu *CPU) {
	cpu.AF.Hi &= cpu.HL.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, true)
	cpu.LastOpCycles = 4
}

func xopANDAl(cpu *CPU) {
	cpu.AF.Hi &= cpu.HL.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, true)
	cpu.LastOpCycles = 4
}

func xopANDAa(cpu *CPU) {
	cpu.AF.Hi &= cpu.AF.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, true)
	cpu.LastOpCycles = 4
}

func xopANDixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.AF.Hi = cpu.andU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopANDixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.AF.Hi = cpu.andU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopANDiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.AF.Hi = cpu.andU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopANDiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
	cpu.AF.Hi = cpu.andU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopXORb(cpu *CPU) {
	cpu.AF.Hi ^= cpu.BC.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
	cpu.LastOpCycles = 4
}

func xopXORc(cpu *CPU) {
	cpu.AF.Hi ^= cpu.BC.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
	cpu.LastOpCycles = 4
}

func xopXORd(cpu *CPU) {
	cpu.AF.Hi ^= cpu.DE.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
	cpu.LastOpCycles = 4
}

func xopXORe(cpu *CPU) {
	cpu.AF.Hi ^= cpu.DE.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
	cpu.LastOpCycles = 4
}

func xopXORh(cpu *CPU) {
	cpu.AF.Hi ^= cpu.HL.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
	cpu.LastOpCycles = 4
}

func xopXORl(cpu *CPU) {
	cpu.AF.Hi ^= cpu.HL.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
	cpu.LastOpCycles = 4
}

func xopXORa(cpu *CPU) {
	cpu.AF.Hi ^= cpu.AF.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
	cpu.LastOpCycles = 4
}

func xopXORixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.AF.Hi = cpu.xorU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopXORixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.AF.Hi = cpu.xorU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopXORiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.AF.Hi = cpu.xorU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopXORiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
	cpu.AF.Hi = cpu.xorU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopORb(cpu *CPU) {
	cpu.AF.Hi |= cpu.BC.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
	cpu.LastOpCycles = 4
}

func xopORc(cpu *CPU) {
	cpu.AF.Hi |= cpu.BC.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
	cpu.LastOpCycles = 4
}

func xopORd(cpu *CPU) {
	cpu.AF.Hi |= cpu.DE.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
	cpu.LastOpCycles = 4
}

func xopORe(cpu *CPU) {
	cpu.AF.Hi |= cpu.DE.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
	cpu.LastOpCycles = 4
}

func xopORh(cpu *CPU) {
	cpu.AF.Hi |= cpu.HL.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
	cpu.LastOpCycles = 4
}

func xopORl(cpu *CPU) {
	cpu.AF.Hi |= cpu.HL.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
	cpu.LastOpCycles = 4
}

func xopORa(cpu *CPU) {
	cpu.AF.Hi |= cpu.AF.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
	cpu.LastOpCycles = 4
}

func xopORixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.AF.Hi = cpu.orU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopORixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.AF.Hi = cpu.orU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopORiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.AF.Hi = cpu.orU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopORiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
	cpu.AF.Hi = cpu.orU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopCPb(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Hi
	cpu.cpU8(a, x)
	cpu.LastOpCycles = 4
}

func xopCPc(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Lo
	cpu.cpU8(a, x)
	cpu.LastOpCycles = 4
}

func xopCPd(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Hi
	cpu.cpU8(a, x)
	cpu.LastOpCycles = 4
}

func xopCPe(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Lo
	cpu.cpU8(a, x)
	cpu.LastOpCycles = 4
}

func xopCPh(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Hi
	cpu.cpU8(a, x)
	cpu.LastOpCycles = 4
}

func xopCPl(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Lo
	cpu.cpU8(a, x)
	cpu.LastOpCycles = 4
}

func xopCPa(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.AF.Hi
	cpu.cpU8(a, x)
	cpu.LastOpCycles = 4
}

func xopCPixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.cpU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopCPixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.cpU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopCPiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.cpU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopCPiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
	cpu.cpU8(a, x)
	cpu.LastOpCycles = 19 // guess
}

func xopINCb(cpu *CPU) {
	cpu.BC.Hi = cpu.incU8(cpu.BC.Hi)
	cpu.LastOpCycles = 4
}

func xopINCc(cpu *CPU) {
	cpu.BC.Lo = cpu.incU8(cpu.BC.Lo)
	cpu.LastOpCycles = 4
}

func xopINCd(cpu *CPU) {
	cpu.DE.Hi = cpu.incU8(cpu.DE.Hi)
	cpu.LastOpCycles = 4
}

func xopINCe(cpu *CPU) {
	cpu.DE.Lo = cpu.incU8(cpu.DE.Lo)
	cpu.LastOpCycles = 4
}

func xopINCh(cpu *CPU) {
	cpu.HL.Hi = cpu.incU8(cpu.HL.Hi)
	cpu.LastOpCycles = 4
}

func xopINCl(cpu *CPU) {
	cpu.HL.Lo = cpu.incU8(cpu.HL.Lo)
	cpu.LastOpCycles = 4
}

func xopINCa(cpu *CPU) {
	cpu.AF.Hi = cpu.incU8(cpu.AF.Hi)
	cpu.LastOpCycles = 4
}

func xopDECb(cpu *CPU) {
	cpu.decP8(&cpu.BC.Hi)
	cpu.LastOpCycles = 4
}

func xopDECc(cpu *CPU) {
	cpu.decP8(&cpu.BC.Lo)
	cpu.LastOpCycles = 4
}

func xopDECd(cpu *CPU) {
	cpu.decP8(&cpu.DE.Hi)
	cpu.LastOpCycles = 4
}

func xopDECe(cpu *CPU) {
	cpu.decP8(&cpu.DE.Lo)
	cpu.LastOpCycles = 4
}

func xopDECh(cpu *CPU) {
	cpu.decP8(&cpu.HL.Hi)
	cpu.LastOpCycles = 4
}

func xopDECl(cpu *CPU) {
	cpu.decP8(&cpu.HL.Lo)
	cpu.LastOpCycles = 4
}

func xopDECa(cpu *CPU) {
	cpu.decP8(&cpu.AF.Hi)
	cpu.LastOpCycles = 4
}
