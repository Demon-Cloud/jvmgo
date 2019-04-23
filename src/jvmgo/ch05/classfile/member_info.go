package classfile

// 存储字段和方法信息
type MemberInfo struct {
    // 常量池指针
    cp                  ConstantPool
    accessFlags         uint16
    nameIndex           uint16
    descriptorIndex     uint16
    attributes          []AttributeInfo
}

// 读取字段表或方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
    memberCount := reader.readUint16()
    members := make([]*MemberInfo, memberCount)
    for index := range members {
        members[index] = readMember(reader, cp)
    }
    return members
}

// 读取字段或方法数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
    return &MemberInfo {
        cp :                cp,
        accessFlags:        reader.readUint16(),
        nameIndex:          reader.readUint16(),
        descriptorIndex:    reader.readUint16(),
        attributes:         readAttributes(reader, cp),
    }
}

// 从常量池中查找字段或方法名
func (self *MemberInfo) Name() string {
    return self.cp.getUtf8(self.nameIndex)
}

// 从常量池中查找字段或方法描述
func (self *MemberInfo) Descriptor() string {
    return self.cp.getUtf8(self.descriptorIndex)
}