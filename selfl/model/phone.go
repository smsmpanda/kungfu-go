package model

type Phone struct {
	name  string
	color string
}

func (f Phone) GetName() string {
	return f.name
}

func (f Phone) GetColor() string {
	return f.color
}

type PhoneFactory struct {
}

func (f *PhoneFactory) GreatePhone(name string, color string) Phone {
	return Phone{name, color}
}

func Gorun() {
	factory := &PhoneFactory{}
	phone := factory.GreatePhone("iphone", "black")
	//phone.name = "iphone"
	println(phone.GetName())
	println(phone.GetColor())
}
