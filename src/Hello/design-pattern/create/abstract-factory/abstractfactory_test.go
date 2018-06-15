package abstract_factory

import "testing"

func getMainAndDetailOrder(factory OrderFactory) {
	factory.CreateMainOrder().SaveMainOrder()
	factory.CreateDetailOrder().SaveDetailOrder()
}
func TestFactory(t *testing.T) {
	var factory OrderFactory
	factory = &RDBFactory{}
	getMainAndDetailOrder(factory)

	factory = &XMLFactory{}
	getMainAndDetailOrder(factory)
}
