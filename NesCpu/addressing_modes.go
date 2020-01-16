package nescpu

/* Read operand (if any), advance PC past the current instruction, return value of operand
   int ind_y();
   int zp();
   int immediate();
   int absa();
   signed char relative();
   int absa_y();
   int absa_x();
   int ind();
   void imp();
   void accum();
*/

func (cpu *CPU6502) zpX() uint16 {
	cpu.pc += 2
	baseAddr := cpu.mem.Read(cpu.pc-1, 0)
	return uint16((baseAddr + cpu.xreg) & 0xFF)
}

func (cpu *CPU6502) zpY() uint16 {
	cpu.pc += 2
	baseAddr := cpu.mem.Read(cpu.pc-1, 0)
	return uint16((baseAddr + cpu.yreg) & 0xFF)
}

func (cpu *CPU6502) indX() uint16 {
	cpu.pc += 2
	addr := (cpu.mem.Read(cpu.pc-1, 0) + cpu.xreg) & 0xff
	return uint16(cpu.mem.Read(uint16(addr), 0)) | (uint16(cpu.mem.Read(uint16(addr+1), 0)) * uint16(256))
}

/*
inline int cpu::ind_y() {
    int addr=memory->read(pc+1);
    pc+=2;
    return (((memory->read(addr))|((memory->read((addr+1)&0xFF))<<(8)))+y)&0xFFFF;
}

inline int cpu::zp() {
    int base_addr=memory->read(pc+1);
    pc+=2;
    return base_addr;
}

inline int cpu::immediate() {
    pc+=2;
    return pc-1;
}

inline int cpu::absa() {
    pc+=3;
    return memory->read2(pc-2);
}

inline int cpu::absa_y() {
    pc+=3;
    return (memory->read2(pc-2)+y)&0xFFFF;
}

inline int cpu::absa_x() {
    pc+=3;
    return (memory->read2(pc-2)+x)&0xFFFF;
}

inline signed char cpu::relative() {
    pc+=2;
    return memory->read(pc-1);
}
*/
