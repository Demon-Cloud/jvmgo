// cp_fieldref.go
package heap

import (
    "jvmgo/ch06/classfile"
)

// 字段符号引用
type FieldRef struct {
    MemberRef
    // 缓存解析后的字段指针
    field       *Field
}

func newFieldRef(cp *ConstantPool,
        refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
    ref := &FieldRef{}
    ref.cp = cp
    ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
    return ref
}