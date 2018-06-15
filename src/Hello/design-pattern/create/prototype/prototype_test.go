package prototype

import "testing"

var manager *PtyManager
func init()  {
	manager = NewPtyMange()
	t1 :=&Type1{name:"type1"}
	manager.Set("t1",t1)
	t2 :=&Type2{name:"type2"}
	manager.Set("t2",t2)
}
func TestPtyManager_Get(t *testing.T) {
	c := manager.Get("t1").Clone()
	t1 := c.(*Type1)
	if t1.name != "type1" {
		t.Fatal("error")
	}

	c2 := manager.Get("t2").Clone()
	t2 := c2.(*Type2)
	if t2.name != "type2" {
		t.Fatal("error")
	}
}
