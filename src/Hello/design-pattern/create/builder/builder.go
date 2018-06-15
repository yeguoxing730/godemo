package builder

type Builder interface {
	BuildEngine() string
	BuildPowerDriven() string
	BuildGearBox() string
}
type Car struct {
	Engine string
	PowerDriven string
	GearBox string
	builder Builder
}
func NewCar(builder Builder) Car{
	car := Car{}
	car.Engine = builder.BuildEngine()
	car.PowerDriven = builder.BuildPowerDriven()
	car.GearBox = builder.BuildGearBox()
	return car
}
type DongFengBuilder struct {}

func (car *DongFengBuilder)BuildEngine() string  {
	return "dong feng engine"
}

func (car *DongFengBuilder)BuildPowerDriven() string  {
	return "dong feng power driven"
}

func (car *DongFengBuilder)BuildGearBox() string  {
	return "dong feng gearbox "
}

type ChangChengBuilder struct {}

func (car *ChangChengBuilder)BuildEngine() string  {
	return "chang cheng engine"
}

func (car *ChangChengBuilder)BuildPowerDriven() string  {
	return "chang cheng power driven"
}

func (car *ChangChengBuilder)BuildGearBox() string  {
	return "chang cheng gearbox "
}