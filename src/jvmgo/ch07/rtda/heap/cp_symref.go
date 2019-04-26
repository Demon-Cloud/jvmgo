// cp_symref.go
package heap

// 符号引用公共类
type SymRef struct {
    // 符号引用所在的运行时常量池指针
    cp          *ConstantPool
    // 类的完全限定名
    className   string
    // 缓存解析后的类信息
    class       *Class
}

func (self *SymRef) ResolvedClass() *Class {
    // 判断类符号是否已解析
    if self.class == nil {
        self.resolveClassRef()
    }
    return self.class
}

func (self *SymRef) resolveClassRef() {
    d := self.cp.class
    c := d.loader.LoadClass(self.className)
    if !c.isAccessibleTo(d) {
        panic("java.lang.IllegalAccessError")
    }
    self.class = c
}