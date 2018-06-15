package abstract_factory

import "fmt"

//抽象工厂 即生成工厂的工厂
type MainOrderDao interface {
	SaveMainOrder()
}
type DetailOrderDao interface {
	SaveDetailOrder()
}

type OrderFactory interface {
	CreateMainOrder() MainOrderDao
	CreateDetailOrder() DetailOrderDao
}

type RDBMainOrder struct{}
type RDBDetailOrder struct{}
type RDBFactory struct{}

func (dao *RDBMainOrder) SaveMainOrder() {
	fmt.Print("rdb main order save\n")
}

func (dao *RDBDetailOrder) SaveDetailOrder() {
	fmt.Print("rdb detail order save\n")
}
func (factory *RDBFactory) CreateMainOrder() MainOrderDao {
	return &RDBMainOrder{}
}
func (factory *RDBFactory) CreateDetailOrder() DetailOrderDao {
	return &RDBDetailOrder{}
}

type XMLMainOrder struct{}
type XMLDetailOrder struct{}
type XMLFactory struct{}

func (dao *XMLMainOrder) SaveMainOrder() {
	fmt.Print("xml main order save\n")
}

func (dao *XMLDetailOrder) SaveDetailOrder() {
	fmt.Print("xml detail order save\n")
}

func (factory *XMLFactory) CreateMainOrder() MainOrderDao {
	return &XMLMainOrder{}
}
func (factory *XMLFactory) CreateDetailOrder() DetailOrderDao {
	return &XMLDetailOrder{}
}
