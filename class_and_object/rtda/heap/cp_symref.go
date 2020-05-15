package heap

type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (this *SymRef) ResolveClass() *Class {
	if this.class == nil {
		this.resolveClassRef()
	}
	return this.class
}

//通俗地讲，如果类D通过符号引用N引用类C的话，要解析N，先用D的类加载器加载C，然后检查D是否有权限访问C，如果没有，则抛出IllegalAccessError异常。
func (this *SymRef) resolveClassRef() {
	d := this.cp.class
	c := d.loader.LoadClass(this.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	this.class = c
}
