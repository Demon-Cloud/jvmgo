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