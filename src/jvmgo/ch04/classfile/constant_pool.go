// constant_pool.go
package classfile

// 常量池,表头给出的常量池大小比实际大1
// 有效的常量池索引是 1~n - 1,0是无效索引
// CONSTANT_Long_info 和 CONSTANT_Double_infp 各占两个位置
type ConstantPool []ConstantInfo

// 读取常量池
func readConstantPool(reader *ClassReader) ConstantPool {
    cpCount := int(reader.readUint16())
    cp := make([]ConstantInfo, cpCount)
    // 索引从1开始
    for index := 1; index < cpCount; index++ {
        cp[index] = readConstantInfo(reader, cp)
        switch cp[index].(type) {
            // 这两个分别占两个位置
        case *ConstantLongInfo, *ConstantDoubleInfo:
            index++
        } 
    }

    return cp
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
    if cpInfo := self[index]; cpInfo != nil {
        return cpInfo
    }
    panic("Invalid constant pool index!")
}

func (self  ConstantPool) getNameAndType(index uint16) (string, string) {
    ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
    name := self.getUtf8(ntInfo.nameIndex)
    _type := self.getUtf8(ntInfo.descriptorIndex)
    return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
    classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
    return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {
    utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
    return utf8Info.str
}