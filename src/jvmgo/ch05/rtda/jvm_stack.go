// jvm_stack.go
package rtda

// 栈信息
type Stack struct {
    // 最多可容纳的栈帧数
    maxSize         uint
    // 栈的大小
    size            uint
    // 保存栈顶指针
    _top            *Frame
}

func newStack(maxSize uint) *Stack {
    return &Stack{
        maxSize: maxSize,
    }
}

// 入栈
func (self *Stack) push(frame *Frame) {
    // 栈溢出
    if self.size >= self.maxSize {
        panic("java.lang.StackOverflowError")
    }
    // lower 是当前帧的上一帧地址
    if self._top != nil {
        frame.lower = self._top
    }
    self._top = frame
    self.size++
}

// 出栈
func (self *Stack) pop() *Frame {
    // 默认栈帧不可能为空
    if self._top == nil {
        panic("jvm stack is empty!")
    }
     top := self._top
     self._top = top.lower
     top.lower = nil
     self.size--

     return top
}

// 获取栈顶帧
func (self *Stack) top() *Frame {
    if self._top == nil {
        panic("jvm stack is empty!")
    }
    return self._top
}