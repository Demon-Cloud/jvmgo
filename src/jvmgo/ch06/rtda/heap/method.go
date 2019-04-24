// method.go
package heap

import (
    "jvmgo/ch06/classfile"
)

// 方法信息
type Method struct {
    CLassMember
    // 操作数栈
    maxStack        uint
    // 局部变量
    maxLocals       uint
    code            []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) {
    methods := make([]*Method, len(cfMethods))

    for i, cfMethod := range cfMethods {
        methods[i] = &Method{}
        methods[i].class = class
        methods[i].copyMemberInfo(cfMethod)
        methods[i].copyAttributes(cfMethod)
    }
    return methods
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
    if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
        self.maxStack : codeAttr.MaxStack()
        self.maxLocals = codeAttr.maxLocals()
        self.code = codeAttr.Code()
    }
}