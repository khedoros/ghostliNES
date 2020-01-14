package NesCpu

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

func (this *NesCpu.NesCpu) zp_x() uint16 {
	this.pc += 2
	base_addr := mem.Read(this.pc - 1, 0)
	return (base_addr + this.x) & 0xFF
}

func (this *NesCpu.NesCpu) zp_y() uint16 {
	this.pc += 2
	base_addr := this.mem.Read(this.pc - 1, 0)
	return (base_addr + this.y) & 0xFF
}

func (this *NesCpu.NesCpu) ind_x() uint16 {
	this.pc+=2
	addr := (this.mem.Read(this.pc - 1, 0) + x) & 0xff
	return this.mem.Read(addr,0)|(this.mem.Read(addr+1) * 256)
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
