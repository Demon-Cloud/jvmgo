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
}

func NewFrame(maxLocals, maxStack uint) *Frame {
    return &Frame {
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