// operand_stack.go
package rtda

import (
    "math"
    "jvmgo/ch06/rtda/heap"
)

// 操作数栈
type OperandStack struct {
    // 操作数帧大小,编译器可以计算出
    size        uint
    // 用来存放变量的值或引用
    slots       []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
    if maxStack > 0 {
        return &OperandStack {
            slots: make([]Slot, maxStack),
        }
    }
    return nil
}

// 出入栈
func (self *OperandStack) PushInt(val int32) {
    self.slots[self.size].num = val
    self.size++
}

func (self *OperandStack) PopInt() int32 {
    self.size--
    return self.slots[self.size].num
}

func (self *OperandStack) PushFloat(val float32) {
    bits := math.Float32bits(val)
    self.slots[self.size].num = int32(bits)
    self.size++
}

func (self *OperandStack) PopFloat() float32 {
    self.size--
    bits := uint32(self.slots[self.size].num)
    return math.Float32frombits(bits)
}

func (self *OperandStack) PushLong(val int64) {
    self.slots[self.size].num = int32(val)
    self.slots[self.size+1].num = int32(val >> 32)
    self.size += 2
}

func (self *OperandStack) PopLong() int64 {
    self.size -= 2
    low := uint32(self.slots[self.size].num)
    high := uint32(self.slots[self.size+1].num)
    return int64(high)<<32 | int64(low)
}

func (self *OperandStack) PushDouble(val float64) {
    bits := math.Float64bits(val)
    self.PushLong(int64(bits))
}

func (self *OperandStack) PopDouble() float64 {
    bits := uint64(self.PopLong())
    return math.Float64frombits(bits)
}

func (self *OperandStack) PushRef(ref *heap.Object) {
    self.slots[self.size].ref = ref
    self.size++
}

func (self *OperandStack) PopRef() *heap.Object {
    self.size--
    ref := self.slots[self.size].ref
    // 为了让 Go 垃圾收集器回收 Object 结构体实例
    self.slots[self.size].ref = nil
    return ref
}

func (self *OperandStack) PushSlot(slot Slot) {
    self.slots[self.size] = slot
    self.size++
}

func (self *OperandStack) PopSlot() Slot {
    self.size--
    return self.slots[self.size]
}