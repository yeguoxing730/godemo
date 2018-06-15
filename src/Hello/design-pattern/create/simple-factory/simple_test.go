package simple_factory

import "testing"

func TestTomato_GetName(t *testing.T) {
	tomato := NewProduct(2)
	tomato.GetName()
}
func TestMice_GetName(t *testing.T) {
	mice := NewProduct(1)
	mice.GetName()
}
