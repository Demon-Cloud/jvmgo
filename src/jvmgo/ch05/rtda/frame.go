// frame.go
package rtda

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
    // 下一条指令
    nextPC          int
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
    return &Frame {
        thread:         thread,
        localVars:      newLocalVars(maxLocals),
        operandStack:   newOperandStack(maxStack),
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