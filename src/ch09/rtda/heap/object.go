package heap

type Object struct {
	class *Class
	data  interface{}
	extra interface{}
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}
func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (self *Object) Fields() *Slots {
	return self.data.(*Slots)
}

func (self *Object) Class() *Class {
	return self.class
}

func (self *Object) Extra() interface{} {
	return self.extra
}
