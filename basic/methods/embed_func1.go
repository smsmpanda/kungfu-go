package methods

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func Gorunfw32() {
	c := new(Customer)
	c.Name = "jiangrui"
	c.log = new(Log)
	c.log.msg = "output loging..."

	c = &Customer{"Shilijuan", &Log{"1 - Yes we can!"}}
	c.Log().Add("2 - After me the world will be a better place!")

	fmt.Println(c.Log())

	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone exhibits multiple behaviors...")
	fmt.Println("It exhibits behavior of a Camera: ", cp.TakeAPicture())
	fmt.Println("It works like a Phone too: ", cp.Call())
}

func (c *Customer) Log() *Log {
	return c.log
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring Ring"
}

type CameraPhone struct {
	Camera
	Phone
}
