package builder

import (
	"testing"
	"fmt"
)

func TestCar_Build(t *testing.T) {
	dfBuilder := &DongFengBuilder{}
	dfCar := NewCar(dfBuilder)
	fmt.Println(dfCar)


	chBuilder := &ChangChengBuilder{}
	chCar := NewCar(chBuilder)
	fmt.Println(chCar)
}
