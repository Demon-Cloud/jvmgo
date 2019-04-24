// cp_methodref.go
package heap

import (
    "jvmgo/ch06/classfile"
)

// 方法符号引用
type MethodRef struct {
    MemberRef
    method      *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethoderfInfo) *MethodRef {
    ref := &MethodRef{}
    ref.cp = cp
    ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
    return ref
}