package heap

import "learn/ch08/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func (r MethodRef) Name() string {
	return r.name
}

func (r MethodRef) Descriptor() string {
	return r.descriptor
}

func newMethodRef(cp *ConstantPool,
	refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
