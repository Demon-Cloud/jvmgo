// frame.go
package rtda

import (
    "jvmgo/ch06/rtda/heap"
)

// 栈帧
type Frame struct {
    // 上一帧的地址,用来实现链表数据
    lower           *Frame
    // 局部变量表指针
    localVars       LocalVars
    // 操作数栈指针
    operandStack    *OperandStack
    // 当前线程
    thread          *Thread
    method          *heap.Method
    // 下一条指令
    nextPC          int
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
    return &Frame {
        thread:         thread,
        method:         method,
        localVars:      newLocalVars(method.MaxLocals()),
        operandStack:   newOperandStack(method.MaxStack()),
    }
}

// getters
func (self *Frame) LocalVars() LocalVars {
    return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
    return self.operandStack
}
func (self *Frame) Thread() *Thread {
    return self.thread
}
func (self *Frame) NextPC() int {
    return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
    self.nextPC = nextPC
}

func (self *Frame) Method() *heap.Method {
    return self.method
}