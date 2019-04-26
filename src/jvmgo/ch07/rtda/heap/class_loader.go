// class_loader.go
package heap

import (
    "fmt"
    "jvmgo/ch07/classfile"
    "jvmgo/ch07/classpath"
)

// 类加载器
type ClassLoader struct {
    cp          *classpath.Classpath
    // 被加载的class, key 是类的完全限定名
    classMap    map[string]*Class
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
    return &ClassLoader {
        cp:         cp,
        classMap:   make(map[string]*Class),
    }
}

// 把类数据加载到方法区
func (self *ClassLoader) LoadClass(name string) *Class {
    if class, ok := self.classMap[name]; ok {
        // 类已经加载
        return class
    }
    return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
    // 将class文件读取到内存,解析class文件,生成类数据
    data, entry := self.readClass(name)
    class := self.defineClass(data)

    // 链接
    link(class)
    fmt.Printf("[Loaded %s from %s]\n", name, entry)
    return class
}

// 将class文件读取到内存
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
    data, entry, err := self.cp.ReadClass(name)
    if err != nil {
        panic("java.lang.ClassNotFoundException: " + name)
    }
    return data, entry
}

func (self *ClassLoader) defineClass(data []byte) *Class {
    class := parseClass(data)
    class.loader = self
    
    // 解析超类和接口的符号引用
    resolveSuperClass(class)
    resolveInterfaces(class)
    
    self.classMap[class.name] = class
    return class
}

func parseClass(data []byte) *Class {
    cf, err := classfile.Parse(data)
    if err != nil {
        panic("java.lang.ClassFormatError")
    }
    return newClass(cf)
}

func resolveSuperClass(class *Class) {
    if class.name != "java/lang/Object" {
        class.superClass = class.loader.LoadClass(class.superClassName)
    }
}

func resolveInterfaces(class *Class) {
    interfaceCount := len(class.interfaceNames)
    if interfaceCount > 0 {
        class.interfaces = make([]*Class, interfaceCount)
        for i, interfaceName := range class.interfaceNames {
            class.interfaces[i] = class.loader.LoadClass(interfaceName)
        }
    }
}

func link(class *Class) {
    // 验证
    verify(class)
    prepare(class)
}

func verify(class *Class) {
    // todo
}

func prepare(class *Class) {
    calcInstanceFieldSlotIds(class)
    calcStaticFieldSlotIds(class)
    allocAndInitStaticVars(class)
}

// 计算实例字段的个数,给它们编号
func calcInstanceFieldSlotIds(class *Class) {
    slotId := uint(0)
    if class.superClass != nil {
        slotId = class.superClass.instanceSlotCount
    }
    for _, field := range class.fields {
        if !field.IsStatic() {
            field.slotId = slotId
            slotId++
            if field.IsLongOrDouble() {
                slotId++
            }
        }
    }
    class.instanceSlotCount = slotId
}

// 计算静态字段的个数
func calcStaticFieldSlotIds(class *Class) {
    slotId := uint(0)
    for _, field := range class.fields {
        if field.IsStatic() {
            field.slotId = slotId
            slotId++
            if field.IsLongOrDouble() {
                slotId++
            }
        }
    }
    class.staticSlotCount = slotId
}

// 给类变量分配空间,并赋初始值
func allocAndInitStaticVars(class *Class) {
    class.staticVars = newSlots(class.staticSlotCount)
    for _, field := range class.fields {
        if field.IsStatic() && field.IsFinal() {
            initStaticFinalVar(class, field)
        }
    }
}

// 从常量池中加载常量值
func initStaticFinalVar(class *Class, field *Field) {
    vars := class.staticVars
    cp := class.constantPool
    cpIndex := field.ConstValueIndex()
    slotId := field.SlotId()

    if cpIndex > 0 {
        switch field.Descriptor() {
        case "Z", "B", "C", "S", "I":
            val := cp.GetConstant(cpIndex).(int32)
            vars.SetInt(slotId, val)
        case "J":
            val := cp.GetConstant(cpIndex).(int64)
            vars.SetLong(slotId, val)
        case "F":
            val := cp.GetConstant(cpIndex).(float32)
            vars.SetFloat(slotId, val)
        case "D":
            val := cp.GetConstant(cpIndex).(float64)
            vars.SetDouble(slotId, val)
        case "Ljava/lang/String;":
            panic("todo")
        }
    }
}