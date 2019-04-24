// cp_interface_methodref.go
package heap

import (
    "jvmgo/ch06/classfile"
)

// 接口方法符号引用
type InterfaceMethodRef struct {
    MemberRef
    method      *Method
}

func newInterfaceMethodRef(co *ConstantPool, 
    refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
    ref := &InterfaceMethodRef{}
    ref.cp = cp
    ref.copyMemberRefInfo(&refInfo.ConstantMemberInfo)
    return ref
}