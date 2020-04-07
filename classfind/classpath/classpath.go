package classpath

import (
	"fmt"
	"os"
)

import "path/filepath"

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	// !!!!!!
	fmt.Println(jreOption, cpOption)
	cp.parseBootAndExtClasspath(jreOption)
	fmt.Println(cp.bootClasspath, cp.extClasspath)
	cp.parseUserClasspath(cpOption)
	fmt.Println(cp.userClasspath)
	return cp
}

func (mySelf *Classpath) parseBootAndExtClasspath(jreOption string) {
	// jre/lib/*
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	mySelf.bootClasspath = newWildcardEntry(jreLibPath)
	//jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	mySelf.extClasspath = newWildcardEntry(jreExtPath)

}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Cannot find jreFolder!")

}

func exists(path string) bool {
	// returns a file info
	if _, err := os.Stat(path); err != nil {

		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (mySelf *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	mySelf.userClasspath = newEntry(cpOption)
}

func (mySelf *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"

	data, entry, err := mySelf.bootClasspath.readClass(className)
	fmt.Println(data, entry, err)
	fmt.Println(mySelf.bootClasspath)
	if data, entry, err := mySelf.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := mySelf.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return mySelf.userClasspath.readClass(className)
}

func (mySelf *Classpath) String() string {
	return mySelf.userClasspath.String()
}
