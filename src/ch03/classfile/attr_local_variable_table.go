package classfile

type LocalVariableTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}
type LocalVariableTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, localVariableTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
