package method_factory
type Operator interface {
	SetA(int)
	SetB(int)
	GetResult() int
}
type OperatorBase struct {
	a int
	b int
}

func (o *OperatorBase)SetA(a int){
	o.a = a
}
func (o *OperatorBase)SetB(b int)  {
	o.b = b
}
type PlusOperator struct {
	*OperatorBase
}

func (operator *PlusOperator)GetResult()int  {
	return operator.a+operator.b
}

type MinusOperator struct {
	*OperatorBase
}

func (operator *MinusOperator)GetResult()int  {
	return operator.a-operator.b
}

type OperatorFactory interface {
	Create() Operator
}
type PlusOperatorFactory struct {}

func (PlusOperatorFactory)Create() Operator{
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}
type MinusOperatorFactory struct {}

func (MinusOperatorFactory)Create() Operator  {
	return &MinusOperator{
		OperatorBase:&OperatorBase{},
	}
}