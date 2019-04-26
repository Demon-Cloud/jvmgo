// class.go
package heap

import (
    "strings"
    "jvmgo/ch07/classfile"
)

// 方法区中保存的类信息
type Class struct {
    // 类的访问标志,16bit,具体含义定义在 heap\access_flags.go 中
    accessFlags             uint16
    // 类名,都是完全限定的:java/lang/Object
    name                    string
    // 超类名
    superClassName          string
    // 接口名
    interfaceNames          []string
    // 运行时常量池指针
    constantPool            *ConstantPool
    // 字段表
    fields                  []*Field
    // 方法表
    methods                 []*Method
    // 类加载器指针
    loader                  *ClassLoader
    // 超类指针
    superClass              *Class
    // 接口指针
    interfaces              []*Class
    instanceSlotCount       uint
    // 变量占据的空间大小
    staticSlotCount         uint
    // 实例占据的空间大小
    staticVars              Slots
}

func newClass(cf *classfile.ClassFile) *Class {
    class := &Class{}
    class.accessFlags = cf.AccessFlags()
    class.name = cf.ClassName()
    class.superClassName = cf.SuperClassName()
    class.interfaceNames = cf.InterfaceNames()
    class.constantPool = newConstantPool(class, cf.ConstantPool())
    class.fields = newFields(class, cf.Fields())
    class.methods = newMethods(class, cf.Methods())
    return class
}

func (self *Class) NewObject() *Object {
    return newObject(self)
}

// 判断类标志
func (self *Class) IsPublic() bool {
    return 0 != self.accessFlags & ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
    return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
    return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
    return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
    return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
    return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
    return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
    return 0 != self.accessFlags&ACC_ENUM
}

// getters
func (self *Class) ConstantPool() *ConstantPool {
    return self.constantPool
}
func (self *Class) StaticVars() Slots {
    return self.staticVars
}

func (self *Class) isAccessibleTo(other *Class) bool {
    return self.IsPublic() || self.getPackageName() == other.getPackageName()
}

func (self *Class) getPackageName() string {
    if i := strings.LastIndex(self.name, "/"); i >= 0 {
        return self.name[:i]
    }
    return ""
}

func (self *Class) GetMainMethod() *Method {
    return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
    for _, method := range self.methods {
        if method.IsStatic() &&
            method.name == name && method.descriptor == descriptor {
            return method
        }
    }
    return nil
}
