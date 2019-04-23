// slot.go
package rtda

// 用来存放变量的值或引用
type Slot struct {
    // 存放整数
    num     int32
    // 存放引用
    ref     *Object
}