package rtda

import "JVM-GO/instructions_bytecode/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
