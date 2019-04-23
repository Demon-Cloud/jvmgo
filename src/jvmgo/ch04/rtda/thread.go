// thread.go
package rtda

// 线程
type Thread struct {
    pc int
    // Java虚拟机栈指针
    stack *Stack
}

func NewThread() *Thread {
    return &Thread {
        // 创建的Stack最多可容纳的栈帧数
        stack: newStack(1024),
    }
}

// 入栈
func (self *Thread) PushFrame(frame *Frame) {
    self.stack.push(frame)
}

// 出栈
func (self *Thread) PopFrame() *Frame {
    return self.stack.pop()
}

// 返回当前栈帧
func (self *Thread) CurrentFrame() *Frame {
    return self.stack.top()
}

