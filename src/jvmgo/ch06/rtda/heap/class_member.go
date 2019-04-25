// class_member.go
package heap

import (
    "jvmgo/ch06/calssfile"
)

// 成员信息,保存属性和方法
type ClassMember struct {
    accessFlags     uint16
    name            string
    descriptor      string
    // 字段所属的类
    class           *Class
}

// 从class中复制数据
func (self *ClassMember) copyMemberInfo(memberInfo *class.MemberInfo) {
    self.accessFlags = memberInfo.AccessFlas()
    self.name = memberInfo.Name()
    self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) isAccessibleTo(d *Class) bool {
    if self.IsPublic() {
        return true
    }

    c := self.class
    if self.IsProtected() {
        return d == c || d.isSubClassOf(c) ||
            c.getPackageName() == d.getPackageName()
    }
    if !self.IsPrivate() {
        return c.getPackageName() == d.getPackageName()
    }

    return d == c
}