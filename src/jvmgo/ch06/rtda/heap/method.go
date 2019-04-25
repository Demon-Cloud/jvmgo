// method.go
package heap

import (
    "jvmgo/ch06/classfile"
)

// 方法信息
type Method struct {
    ClassMember
    // 操作数栈
    maxStack        uint
    // 局部变量
    maxLocals       uint
    code            []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
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
        self.maxStack = codeAttr.MaxStack()
        self.maxLocals = codeAttr.MaxLocals()
        self.code = codeAttr.Code()
    }
}

func (self *Method) IsSynchronized() bool {
    return 0 != self.accessFlags&ACC_SYNCHRONIZED
}
func (self *Method) IsBridge() bool {
    return 0 != self.accessFlags&ACC_BRIDGE
}
func (self *Method) IsVarargs() bool {
    return 0 != self.accessFlags&ACC_VARARGS
}
func (self *Method) IsNative() bool {
    return 0 != self.accessFlags&ACC_NATIVE
}
func (self *Method) IsAbstract() bool {
    return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Method) IsStrict() bool {
    return 0 != self.accessFlags&ACC_STRICT
}

// getters
func (self *Method) MaxStack() uint {
    return self.maxStack
}
func (self *Method) MaxLocals() uint {
    return self.maxLocals
}
func (self *Method) Code() []byte {
    return self.code
}