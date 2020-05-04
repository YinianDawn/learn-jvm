package heap

import "learn/ch07/classfile"

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
	}
	return methods
}
func (self *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++
	}
}
func (self *Method) MaxLocals() uint {
	return self.maxLocals
}
func (self *Method) MaxStack() uint {
	return self.maxStack
}
func (self *Method) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Method) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}
func (self *Method) Class() *Class {
	return self.class
}
func (self *Method) Code() []byte {
	return self.code
}
func (self *Method) Name() string {
	return self.name
}
func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}
func (self *Method) Descriptor() string {
	return self.descriptor
}
