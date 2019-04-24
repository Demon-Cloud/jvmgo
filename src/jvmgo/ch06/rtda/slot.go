// slot.go
package rtda

import (
    "jvmgo/ch06/rtda/heap"
)

// 用来存放变量的值或引用
type Slot struct {
    // 存放整数
    num     int32
    // 存放引用
    ref     *heap.Object
}