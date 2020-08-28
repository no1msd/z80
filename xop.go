package z80

func xopPOPbc(cpu *CPU) {
	cpu.BC.Lo = cpu.Memory.Get(cpu.SP)
	cpu.BC.Hi = cpu.Memory.Get(cpu.SP + 1)
	cpu.SP += 2
}

func xopPOPde(cpu *CPU) {
	cpu.DE.Lo = cpu.Memory.Get(cpu.SP)
	cpu.DE.Hi = cpu.Memory.Get(cpu.SP + 1)
	cpu.SP += 2
}

func xopPOPhl(cpu *CPU) {
	cpu.HL.Lo = cpu.Memory.Get(cpu.SP)
	cpu.HL.Hi = cpu.Memory.Get(cpu.SP + 1)
	cpu.SP += 2
}

func xopPOPaf(cpu *CPU) {
	cpu.AF.Lo = cpu.Memory.Get(cpu.SP)
	cpu.AF.Hi = cpu.Memory.Get(cpu.SP + 1)
	cpu.SP += 2
}

func xopPUSHbc(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.BC.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.BC.Hi)
}

func xopPUSHde(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.DE.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.DE.Hi)
}

func xopPUSHhl(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.HL.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.HL.Hi)
}

func xopPUSHaf(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.AF.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.AF.Hi)
}

func xopINCbc(cpu *CPU) {
	cpu.BC.Lo++
	if cpu.BC.Lo == 0 {
		cpu.BC.Hi++
	}
}

func xopINCde(cpu *CPU) {
	cpu.DE.Lo++
	if cpu.DE.Lo == 0 {
		cpu.DE.Hi++
	}
}

func xopINChl(cpu *CPU) {
	cpu.HL.Lo++
	if cpu.HL.Lo == 0 {
		cpu.HL.Hi++
	}
}

func xopINCsp(cpu *CPU) {
	cpu.SP++
}

func xopLDbchHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.BC.Hi = cpu.Memory.Get(p)
}

func xopLDbclHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.BC.Lo = cpu.Memory.Get(p)
}

func xopLDdehHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.DE.Hi = cpu.Memory.Get(p)
}

func xopLDdelHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.DE.Lo = cpu.Memory.Get(p)
}

func xopLDhlhHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.HL.Hi = cpu.Memory.Get(p)
}

func xopLDhllHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.HL.Lo = cpu.Memory.Get(p)
}

func xopLDafhHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.AF.Hi = cpu.Memory.Get(p)
}

func xopDECb(cpu *CPU) {
	cpu.decP8(&cpu.BC.Hi)
}

func xopDECc(cpu *CPU) {
	cpu.decP8(&cpu.BC.Lo)
}

func xopDECd(cpu *CPU) {
	cpu.decP8(&cpu.DE.Hi)
}

func xopDECe(cpu *CPU) {
	cpu.decP8(&cpu.DE.Lo)
}

func xopDECh(cpu *CPU) {
	cpu.decP8(&cpu.HL.Hi)
}

func xopDECl(cpu *CPU) {
	cpu.decP8(&cpu.HL.Lo)
}

func xopDECa(cpu *CPU) {
	cpu.decP8(&cpu.AF.Hi)
}

func xopLDbcnn(cpu *CPU, l, h uint8) {
	cpu.BC.Lo = l
	cpu.BC.Hi = h
}

func xopLDdenn(cpu *CPU, l, h uint8) {
	cpu.DE.Lo = l
	cpu.DE.Hi = h
}

func xopLDhlnn(cpu *CPU, l, h uint8) {
	cpu.HL.Lo = l
	cpu.HL.Hi = h
}

func xopLDspnn(cpu *CPU, l, h uint8) {
	cpu.SP = toU16(l, h)
}

func xopXORb(cpu *CPU) {
	cpu.AF.Hi ^= cpu.BC.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopXORc(cpu *CPU) {
	cpu.AF.Hi ^= cpu.BC.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopXORd(cpu *CPU) {
	cpu.AF.Hi ^= cpu.DE.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopXORe(cpu *CPU) {
	cpu.AF.Hi ^= cpu.DE.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopXORh(cpu *CPU) {
	cpu.AF.Hi ^= cpu.HL.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopXORl(cpu *CPU) {
	cpu.AF.Hi ^= cpu.HL.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopXORa(cpu *CPU) {
	cpu.AF.Hi ^= cpu.AF.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopJPnZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ == 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPfZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ != 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPnCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC == 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPfCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC != 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPnPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV == 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPfPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV != 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPnSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS == 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPfSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS != 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopLDHLPb(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.BC.Hi
	cpu.Memory.Set(p, r)
}

func xopLDHLPc(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.BC.Lo
	cpu.Memory.Set(p, r)
}

func xopLDHLPd(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.DE.Hi
	cpu.Memory.Set(p, r)
}

func xopLDHLPe(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.DE.Lo
	cpu.Memory.Set(p, r)
}

func xopLDHLPh(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.HL.Hi
	cpu.Memory.Set(p, r)
}

func xopLDHLPl(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.HL.Lo
	cpu.Memory.Set(p, r)
}

func xopLDHLPa(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.AF.Hi
	cpu.Memory.Set(p, r)
}

func xopCALLnn(cpu *CPU, l, h uint8) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, uint8(cpu.PC))
	cpu.Memory.Set(cpu.SP+1, uint8(cpu.PC>>8))
	cpu.PC = toU16(l, h)
}

func xopCALLnZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ == 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLfZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ != 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLnCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC == 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLfCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC != 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLnPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV == 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLfPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV != 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLnSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS == 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLfSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS != 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopADDHLbc(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.BC.U16()
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
}

func xopADDHLde(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.DE.U16()
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
}

func xopADDHLhl(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.HL.U16()
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
}

func xopADDHLsp(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.SP
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
}

func xopRETnZ(cpu *CPU) {
	if cpu.AF.Lo&maskZ == 0 {
		oopRET(cpu)
	}
}

func xopRETfZ(cpu *CPU) {
	if cpu.AF.Lo&maskZ != 0 {
		oopRET(cpu)
	}
}

func xopRETnC(cpu *CPU) {
	if cpu.AF.Lo&maskC == 0 {
		oopRET(cpu)
	}
}

func xopRETfC(cpu *CPU) {
	if cpu.AF.Lo&maskC != 0 {
		oopRET(cpu)
	}
}

func xopRETnPV(cpu *CPU) {
	if cpu.AF.Lo&maskPV == 0 {
		oopRET(cpu)
	}
}

func xopRETfPV(cpu *CPU) {
	if cpu.AF.Lo&maskPV != 0 {
		oopRET(cpu)
	}
}

func xopRETnS(cpu *CPU) {
	if cpu.AF.Lo&maskS == 0 {
		oopRET(cpu)
	}
}

func xopRETfS(cpu *CPU) {
	if cpu.AF.Lo&maskS != 0 {
		oopRET(cpu)
	}
}
