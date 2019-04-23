package classfile

import (
	"encoding/binary"
)

// Class 文件读取类
// Java虚拟机规范定义了u1,u2,u4是那种类型表示1,2,4个字节无符号整数,分别对应go的uint8,uint16,uint32
// 相同类型的多条数据一般按表(table)的形式存储在class文件中,表由表头和表项(item)构成,表头是u2或u4整数,
// 假设表头是n,后面就紧跟着n个表项数据
/*
	Java 虚拟机中的class文件结构如下:
ClassFile {
	u4 				magic;
	u2 				minor_version;
	u2 				major_version;
	u2 				constant_pool_count;
	cp_info 		constant_pool[constant_pool_count-1];
	u2 				access_flags;
	u2 				this_class;
	u2 				super_class;
	u2 				interfaces_count;
	u2 			    interfaces[interfaces_count];
	u2 				fields_count;
	field_info 		fields[fields_count];
	u2 				methods_count;
	method_info 	methods[methods_count];
	u2 				attributes_count;
	attribute_info 	attributes[attributes_count];
}
*/
type ClassReader struct {
	data []byte
}

// 读取 u1 类型数据
func (self *ClassReader) readUint8() uint8 {
    val := self.data[0]
    // go使用reslice语法跳过已经读取的数据,Java虚拟机默认是使用索引的方式记录读取位置
    self.data = self.data[1:]
    return val;
}

// 读取 u2 类型数据
func (self *ClassReader) readUint16() uint16 {
    // BigEndian 可以从[]byte中解码多字节数据
    val := binary.BigEndian.Uint16(self.data)
    self.data = self.data[2:]
    return val
}

// 读取 u4 类型数据
func (self *ClassReader) readUint32() uint32 {
    val := binary.BigEndian.Uint32(self.data)
    self.data = self.data[4:]
    return val
}

// 读取 u8 类型数据,Java 虚拟机并没有定义u8
func (self *ClassReader) readUint64() uint64 {
    val := binary .BigEndian.Uint64(self.data)
    self.data = self.data[8:]
    return val
}

// 读取 uint16表,表的大小由开头的uint16数据指出
func (self *ClassReader) readUint16s() []uint16 {
    n := self.readUint16()
    s := make([]uint16, n)
    for i := range s {
        s[i] = self.readUint16()
    }
    return s
}

// 读取指定数量的字节
func (self *ClassReader) readBytes(n uint32) []byte {
    bytes := self.data[:n]
    self.data = self.data[n:]
    return bytes
}