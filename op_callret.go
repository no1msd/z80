package z80

func oopRETI(cpu *CPU) {
	cpu.PC = cpu.readU16(cpu.SP)
	cpu.SP += 2

	if cpu.INT != nil {
		cpu.INT.ReturnINT()
	}
	cpu.LastOpCycles = 14
}

func oopRETN(cpu *CPU) {
	cpu.PC = cpu.readU16(cpu.SP)
	cpu.SP += 2
	cpu.IFF1 = cpu.IFF2

	cpu.InNMI = false
	cpu.LastOpCycles = 14
}

func oopRET(cpu *CPU) {
	cpu.PC = cpu.readU16(cpu.SP)
	cpu.SP += 2
	cpu.LastOpCycles = 10
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func xopCALLnn(cpu *CPU, l, h uint8) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, uint8(cpu.PC))
	cpu.Memory.Set(cpu.SP+1, uint8(cpu.PC>>8))
	cpu.PC = toU16(l, h)
	cpu.LastOpCycles = 17
}

func xopCALLnZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ == 0 {
		xopCALLnn(cpu, l, h)
		cpu.LastOpCycles = 17
	} else {
		cpu.LastOpCycles = 10
	}
}

func xopCALLfZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ != 0 {
		xopCALLnn(cpu, l, h)
		cpu.LastOpCycles = 17
	} else {
		cpu.LastOpCycles = 10
	}
}

func xopCALLnCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC == 0 {
		xopCALLnn(cpu, l, h)
		cpu.LastOpCycles = 17
	} else {
		cpu.LastOpCycles = 10
	}
}

func xopCALLfCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC != 0 {
		xopCALLnn(cpu, l, h)
		cpu.LastOpCycles = 17
	} else {
		cpu.LastOpCycles = 10
	}
}

func xopCALLnPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV == 0 {
		xopCALLnn(cpu, l, h)
		cpu.LastOpCycles = 17
	} else {
		cpu.LastOpCycles = 10
	}
}

func xopCALLfPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV != 0 {
		xopCALLnn(cpu, l, h)
		cpu.LastOpCycles = 17
	} else {
		cpu.LastOpCycles = 10
	}
}

func xopCALLnSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS == 0 {
		xopCALLnn(cpu, l, h)
		cpu.LastOpCycles = 17
	} else {
		cpu.LastOpCycles = 10
	}
}

func xopCALLfSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS != 0 {
		xopCALLnn(cpu, l, h)
		cpu.LastOpCycles = 17
	} else {
		cpu.LastOpCycles = 10
	}
}

func xopRETnZ(cpu *CPU) {
	if cpu.AF.Lo&maskZ == 0 {
		oopRET(cpu)
		cpu.LastOpCycles = 11
	} else {
		cpu.LastOpCycles = 5
	}
}

func xopRETfZ(cpu *CPU) {
	if cpu.AF.Lo&maskZ != 0 {
		oopRET(cpu)
		cpu.LastOpCycles = 11
	} else {
		cpu.LastOpCycles = 5
	}
}

func xopRETnC(cpu *CPU) {
	if cpu.AF.Lo&maskC == 0 {
		oopRET(cpu)
		cpu.LastOpCycles = 11
	} else {
		cpu.LastOpCycles = 5
	}
}

func xopRETfC(cpu *CPU) {
	if cpu.AF.Lo&maskC != 0 {
		oopRET(cpu)
		cpu.LastOpCycles = 11
	} else {
		cpu.LastOpCycles = 5
	}
}

func xopRETnPV(cpu *CPU) {
	if cpu.AF.Lo&maskPV == 0 {
		oopRET(cpu)
		cpu.LastOpCycles = 11
	} else {
		cpu.LastOpCycles = 5
	}
}

func xopRETfPV(cpu *CPU) {
	if cpu.AF.Lo&maskPV != 0 {
		oopRET(cpu)
		cpu.LastOpCycles = 11
	} else {
		cpu.LastOpCycles = 5
	}
}

func xopRETnS(cpu *CPU) {
	if cpu.AF.Lo&maskS == 0 {
		oopRET(cpu)
		cpu.LastOpCycles = 11
	} else {
		cpu.LastOpCycles = 5
	}
}

func xopRETfS(cpu *CPU) {
	if cpu.AF.Lo&maskS != 0 {
		oopRET(cpu)
		cpu.LastOpCycles = 11
	} else {
		cpu.LastOpCycles = 5
	}
}

func xopRST00(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0000
	cpu.LastOpCycles = 11
}

func xopRST08(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0008
	cpu.LastOpCycles = 11
}

func xopRST10(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0010
	cpu.LastOpCycles = 11
}

func xopRST18(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0018
	cpu.LastOpCycles = 11
}

func xopRST20(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0020
	cpu.LastOpCycles = 11
}

func xopRST28(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0028
	cpu.LastOpCycles = 11
}

func xopRST30(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0030
	cpu.LastOpCycles = 11
}

func xopRST38(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0038
	cpu.LastOpCycles = 11
}
