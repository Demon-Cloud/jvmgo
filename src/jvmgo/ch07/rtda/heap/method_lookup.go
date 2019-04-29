// method_lookup.go
package heap

func LookupMethodInClass(class *Class, anme, descriptor string) *Method {
    for c:= class; c != nil; c = c.superClass {
        for _, method := range c.methods {
            if method.name == name && method.descriptor == descriptor {
                return method
            }
        }
    }

    return nil
}

func lookupMethodInInterfaces(ifaces []*Class, name, descriptor string) *Metho
    for _, iface := range ifaces {
        for _, method := range iface.methods {
            if method.name == name && method.descriptor == descriptor {
                return method
            }
        }
        method := lookupMethodInInterfaces(iface.interfaces, name, descriptor)
        if method != nil {
            return method
        }
    }
    return nil
}

func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
    for _, method := range iface.methods {
        if method.name == name && method.descriptor == descriptor {
            return method
        }
    }
    return lookupMethodInInterfaces(iface.interfaces, name , descriptor)
}