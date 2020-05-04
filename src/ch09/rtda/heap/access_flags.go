package heap

const (
	ACC_PUBLIC       = 0x0001 // class field method		0000 0000 0000 0001
	ACC_PRIVATE      = 0x0002 // 	   field method		0000 0000 0000 0010
	ACC_PROTECTED    = 0x0004 // 	   field method		0000 0000 0000 0100
	ACC_STATIC       = 0x0008 // 	   field method		0000 0000 0000 1000
	ACC_FINAL        = 0x0010 // class field method		0000 0000 0001 0000
	ACC_SUPER        = 0x0020 // class					0000 0000 0010 0000
	ACC_SYNCHRONIZED = 0x0020 // 			 method
	ACC_VOLATILE     = 0x0040 // 	   field
	ACC_BRIDGE       = 0x0040 // 			 method
	ACC_TRANSIENT    = 0x0080 // 	   field
	ACC_VARARGS      = 0x0080 // 			 method
	ACC_NATIVE       = 0x0100 // 			 method
	ACC_INTERFACE    = 0x0200 // class
	ACC_ABSTRACT     = 0x0400 // class 	     method
	ACC_STRICT       = 0x0800 // 			 method
	ACC_SYNTHETIC    = 0x1000 // class field method
	ACC_ANNOTATION   = 0x2000 // class
	ACC_ENUM         = 0x4000 // class field
)
