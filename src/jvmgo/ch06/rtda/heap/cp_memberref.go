// cp_memberref.go
package heap

import (
    "jvmgo/ch06/classfile"
)

// 符号引用
type MemberRef struct {
    SymRef
    name        string
    descriptor  string
}

// java 里面一个类不可以出现两个同名的字段,但是JVM是支持的
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
    self.className = refInfo.ClassName()
    self.name, self.descriptor = refInfo.NameAndDescriptor()
}