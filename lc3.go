package main

import "fmt"

// registers

const (
	R0        int = 0
	R1        int = 1
	R2        int = 2
	R3        int = 3
	R4        int = 4
	R5        int = 5
	R6        int = 6
	R7        int = 7
	PC        int = 8
	COND      int = 9
	REG_COUNT int = 10
)

var reg [REG_COUNT]uint16

// op-codes

const (
	OP_BR   int = 0  // branch
	OP_ADD  int = 1  // add
	OP_LD   int = 2  // load
	OP_ST   int = 3  // store
	OP_JSR  int = 4  // jump register
	OP_AND  int = 5  // bitwise and
	OP_LDR  int = 6  // load register
	OP_STR  int = 7  // store register
	OP_RTI  int = 8  // unused
	OP_NOT  int = 9  // bitwise not
	OP_LDI  int = 10 // load indirect
	OP_STI  int = 11 // store indirect
	OP_JMP  int = 12 // jump
	OP_RES  int = 13 // reserved
	OP_LEA  int = 14 // load effective address
	OP_TRAP int = 15 // execute trap
)

// condition flags

const (
	FL_POS = 1 << 0
	FL_ZRO = 1 << 1
	FL_NEG = 1 << 2
)

// update conditional flags

func update_flags(r uint16) {
	if reg[r] == 0 {
		reg[COND] = FL_ZRO
	} else if reg[r] >> 15 {
		reg[COND] = FL_NEG
	} else {
		reg[COND] = FL_POS
	}
}

// extending any number so its sign is it's MSB (two's complement)

func sign_extend(x uint16, bit_count int)

func main() {
	const PC_START = 0x3000
	reg[PC] = PC_START

	running := true
	for running {
		instr := mem_read(reg[PC])
		reg[PC]++
		op := instr >> 12

		switch op {
		case OP_ADD:
			dr = (instr >> 9) & 0x7  // destination reg
			r1 = (instr >> 6) & 0x7  // first source register
			imm = (instr >> 5) & 0x1 // are we in immediate mode

			if imm {
				imm5 := sign_extend(instr&0x1F, 5)
				reg[dr] = reg[r1] + imm5
			} else {
				r2 := instr & 0x7 // second source register
				reg[dr] = reg[r1] + reg[r2]
			}

			update_flags(dr)
		case OP_AND:
			dr = (instr >> 9) & 0x7  // destination reg
			r1 = (instr >> 6) & 0x7  // first source register
			imm = (instr >> 5) & 0x1 // are we in immediate mode

			if imm {
				imm5 := sign_extend(instr&0x1F, 5)
				reg[dr] = reg[r1] & imm5 // AND operation on register and immediate value
			} else {
				r2 := instr & 0x7
				reg[dr] = reg[r1] & reg[r2] // AND operation on two registers
			}

			update_flags(dr)
		case OP_NOT:
			dr = (instr >> 9) & 0x7 // destination reg
			r1 = (instr >> 6) & 0x7 // source reg
			reg[dr] = ^reg[r1]      // NOT operation on source reg
			update_flags(dr)
		case OP_BR:
			cond_flag = (instr >> 9) & 0x7 // get condition flags for testing

			if cond_flag & reg[COND] {
				pc_offset = sign_extend(instr&0x1ff, 9) // get PC offset for branching
				reg[PC] += pc_offset                    // add to the program counter
			}
		case OP_JMP:
			base_reg = (instr >> 6) & 0x7 // base register number
			reg[PC] = reg[base_reg]       // set program counter to contents of base reg
		case OP_JSR:
			reg[7] = reg[PC]
			long_flag := (instr >> 11) & 0x1

			if long_flag { // JSR
				pc_offset = sign_extend(instr&0x7ff, 9)
				reg[PC] += pc_offset
			} else { // JSRR
				base_reg = (instr >> 6) & 0x7
				reg[PC] = reg[base_reg] // obtain program counter from base reg
			}
		case OP_LD:
			dr = (instr >> 9) & 0x7                 // destination reg
			pc_offset = sign_extend(instr&0x1ff, 9) // get pc_offset for loading a value

			reg[dr] = mem_read(reg[PC] + pc_offset) // loading value from PC + pc_offset into the dest reg
			update_flags(dr)                        // update cond flags
		case OP_LDI:
			dr = (instr >> 9) & 0x7                 // destination reg
			pc_offset = sign_extend(instr&0x1ff, 9) // get PC offset for mem-read location
			reg[dr] = mem_read(mem_read(reg[PC] + pc_offset))
			update_flags(dr) // update cond flags
		case OP_LDR:
		case OP_LEA:
		case OP_ST:
		case OP_STI:
		case OP_STR:
		case OP_TRAP:
		case OP_RES:
		case OP_RTI:
		default:
			break
		}
	}
}
