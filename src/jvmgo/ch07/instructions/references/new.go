// new.go
package references

import (
    "jvmgo/ch07/instructions/base"
    "jvmgo/ch07/rtda"
    "jvmgo/ch07/rtda/heap"
)

// create new object
type NEW struct { base.Index16Instruction }

func (self *NEW) Execute(frame *rtda.Frame) {
    cp := frame.Method().Class().ConstantPool()
    classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
    class := classRef.ResolvedClass()
    // todo: init class

    if class.IsInterface() || class.IsAbstract() {
        panic("java.lang.InstantiationError")
    }

    ref := class.NewObject()
    frame.OperandStack().PushRef(ref)
}