package heap

import "fmt"
import "JVM-GO/class_and_object/classfile"

type Constant interface {
}
type ConstantPool struct {
	class  *Class
	consts []Constant
}
