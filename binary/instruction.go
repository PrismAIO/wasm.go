package binary

type Expr = []Instruction

type Instruction struct {
	Opcode byte
	Args   interface{}
}

// block & loop
type BlockArgs struct {
	BT     BlockType
	Instrs []Instruction
}

type IfArgs struct {
	BT      BlockType
	Instrs1 []Instruction
	Instrs2 []Instruction
}

type BrTableArgs struct {
	Labels  []LabelIdx
	Default LabelIdx
}

// MiscArgs are the arguments of a 0xFC instruction from bulk memory or
// tables: Sub is the opcode after the prefix, Idx its data, element or table
// indices in encoding order.
type MiscArgs struct {
	Sub uint8
	Idx [2]uint32
}

type MemArg struct {
	Align  uint32
	Offset uint32
}

func (instr Instruction) GetOpname() string {
	return opnames[instr.Opcode]
}
func (instr Instruction) String() string {
	return opnames[instr.Opcode]
}
