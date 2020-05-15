package heap

import (
	"JVM-GO/class_and_object/classfile"
	"fmt"
)
import "JVM-GO/class_and_object/classpath"

type ClassLoader struct {
	cp       *classpath.Classpath
	classMap map[string]*Class // loaded class
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

func (this *ClassLoader) LoadClass(name string) *Class {
	if class, ok := this.classMap[name]; ok {
		return class
	} else {
		return this.loadNonArrayClass(name)

	}

}

func (this *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := this.readClass(name)
	class := this.defineClass(data)
	link(class)
	fmt.Printf("[Loaded%sfrom%s]\n", name, entry)
	return class

}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsFinal() && field.IsStatic() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.slotId
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

func verify(class *Class) {

}

func (this *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := this.cp.ReadClass(name)
	if err != nil {
		panic("Java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (this *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = this
	resolveSuperClass(class) // 解析类符号应用
	resolveInterfaces(class)
	this.classMap[class.name] = class
	return class
}

func resolveInterfaces(class *Class) {
	count := len(class.interfaceNames)
	if count > 0 {
		class.interfaces = make([]*Class, count)
		for i, interfaceNames := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceNames)
		}
	}
}

// 除Object类以外，所有类有且只有一个超类
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError!")
	}
	return NewClass(cf)
}
