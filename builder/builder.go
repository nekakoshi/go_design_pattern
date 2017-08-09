package builder

import (
	"strconv"
)

type Color string
type Make string
type Model string

const (
	BLUE Color = "blue"
	RED  Color = "red"
)

type ICar interface {
	Drive() string
	Stop() string
}

type ICarBuilder interface {
	SetSpeed(int) ICarBuilder
	SetPaint(Color) ICarBuilder
	Build() ICar
}

type carBuilder struct {
	speedOption int
	color       Color
}

func (cb *carBuilder) SetSpeed(speed int) ICarBuilder {
	cb.speedOption = speed
	return cb
}

func (cb *carBuilder) SetPaint(color Color) ICarBuilder {
	cb.color = color
	return cb
}

func (cb *carBuilder) Build() ICar {
	return &car{
		speed: cb.speedOption,
		color: cb.color,
	}
}

func NewCarBuilder() ICarBuilder {
	return &carBuilder{}
}

type car struct {
	speed int
	color Color
}

func (c *car) Drive() string {
	return "Driving at speed: " + strconv.Itoa(c.speed)
}

func (c *car) Stop() string {
	return "Stopping a " + string(c.color) + " car"
}

//func main() {
//	builder := NewCarBuilder()
//	car := builder.SetSpeed(50).SetPaint(BLUE).Build()
//
//	fmt.Println(car.Drive())
//	fmt.Println(car.Stop())
//}
