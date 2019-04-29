// cp_interface_methodref.go
package heap

import (
    "jvmgo/ch07/classfile"
)

// 接口方法符号引用
type InterfaceMethodRef struct {
    MemberRef
    method      *Method
}

func newInterfaceMethodRef(cp *ConstantPool, 
    refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
    ref := &InterfaceMethodRef{}
    ref.cp = cp
    ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
    return ref
}

func (self *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
    if self.method == nil {
        self.resolveInterfaceMethodRef()
    }
    return self.method
}

// jvms8 5.4.3.4
func (self *InterfaceMethodRef) resolveInterfaceMethodRef() {
    d := self.cp.class
    c := self.ResolvedClass()
    if !c.IsInterface() {
        panic("java.lang.IncompatibleClassChangeError")
    }
    method := lookupInterfaceMethod(c, self.name, self.descriptor)
    if method == nil {
        panic("java.lang.NoSuchMethodError")
    }
    if !method.isAccessibleTo(d) {
        panic("java.lang/IllegalAccessError")
    }
    self.method = method
}