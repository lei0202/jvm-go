package main

import "fmt"
import "strings"
import "JVM-GO/classparse/classpath"
import "JVM-GO/classparse/classfile"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.xJreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	//fmt.Println(cmd.class)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf

}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version:%v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constantscount:%v\n", len(cf.ConstantPool()))
	fmt.Printf("accessflags:0x%x\n", cf.AccessFlags())
	fmt.Printf("thisclass:%v\n", cf.ThisClassName())
	fmt.Printf("superclass:%v\n", cf.SuperClassName())
	fmt.Printf("interfaces:%v\n", cf.InterfaceNames())
	fmt.Printf("fieldscount:%v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("%s\n", f.Name())
	}
	fmt.Printf("methodscount:%v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("%s\n", m.Name())
	}
}
