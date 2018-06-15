package prototype
//原型模式使对象能复制自身，并且暴露到接口中，使客户端面向接口编程时，
// 不知道接口实际对象的情况下生成新的对象。

//原型模式配合原型管理器使用，使得客户端在不知道具体类的情况下，
// 通过接口管理器得到新的实例，并且包含部分预设定配置。

type Cloneable interface {
	Clone() Cloneable
}

type PtyManager struct {
	ptyTypes map[string]Cloneable
}
func NewPtyMange() *PtyManager{
	return &PtyManager{
		ptyTypes:make(map[string]Cloneable),
	}
}
func (p *PtyManager)Get(name string)Cloneable  {
	return p.ptyTypes[name]
}
func (p *PtyManager)Set(name string,prototype Cloneable)  {
	p.ptyTypes[name] = prototype
}

type Type1 struct {
	name string
}

func (t *Type1)Clone() Cloneable  {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2)Clone() Cloneable  {
	tc := *t
	return &tc
}
