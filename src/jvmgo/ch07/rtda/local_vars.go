// local_vars.go
package rtda

import (
    "math"
    "jvmgo/ch07/rtda/heap"
)

// 局部变量
type LocalVars []Slot

// 创建局部变量数组
// maxLocals: 最大的局部变量数,这个值是计算出来的
func newLocalVars(maxLocals uint) LocalVars {
    if maxLocals > 0 {
        return make([]Slot, maxLocals)
    }
    return nil
}

// 读写 int 变量
func (self LocalVars) SetInt(index uint, val int32) {
    self[index].num = val
}
func (self LocalVars) GetInt(index uint) int32 {
    return self[index].num
}

// 读写 float 变量,先转成int
func (self LocalVars) SetFloat(index uint, val float32) {
    bits := math.Float32bits(val)
    self[index].num = int32(bits)
}
func (self LocalVars) GetFloat(index uint) float32 {
    bits := uint32(self[index].num)
    return math.Float32frombits(bits)
}

// 读写 long 变量, 需要拆成两个int变量
func (self LocalVars) SetLong(index uint, val int64) {
    self[index].num = int32(val)
    self[index + 1].num = int32(val >> 32)
}
func (self LocalVars) GetLong(index uint) int64 {
    low := uint32(self[index].num)
    high := uint32(self[index+1].num)
    return int64(high) << 32 | int64(low)
}

// 读写 double 变量,需要通过转成 long 来处理
func (self LocalVars) SetDouble(index uint, val float64) {
    bits := math.Float64bits(val)
    self.SetLong(index, int64(bits))
}
func (self LocalVars) GetDouble(index uint) float64 {
    bits := uint64(self.GetLong(index))
    return math.Float64frombits(bits)
}

// 读写引用类型
func (self LocalVars) SetRef(index uint, ref *heap.Object) {
    self[index].ref = ref
}
func (self LocalVars) GetRef(index uint) *heap.Object {
    return self[index].ref
}