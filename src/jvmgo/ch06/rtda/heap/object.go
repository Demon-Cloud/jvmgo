// object.go
package heap

// Object 对象
type Object struct {
    // class 指针
    class   *Class
    // 实例变量
    fields  Slots
}

func newObject(class *Class) *Object {
    return &Object {
        class:  class,
        fields: newSlots(class.instanceSlotCount),
    }
}

// getters
func (self *Object) Class() *Class {
    return self.class
}
func (self *Object) Fields() Slots {
    return self.fields
}

func (self *Object) IsInstanceOf(class *Class) bool {
    return class.isAssignableFrom(self.class)
}