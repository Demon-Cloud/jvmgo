// object.go
package heap

// Object 对象
type Object struct {
    // class 指针
    class   *Class
    // 实例变量
    fields  Slots
}

