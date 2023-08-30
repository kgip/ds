package ds

import "testing"

type Person interface {
	Sleep()
	Eat()
}

type Student struct {

}

func (Student) Sleep() {

}

func (Student) Eat() {

}

type Teacher struct {

}

func (Teacher) Sleep() {

}

func (*Teacher) Eat() {

}

func TestInter(t *testing.T) {
	var p Person
	p = &Teacher{}
	p.Sleep()
	student, ok := p.(*Teacher)
	t.Log(student, ok)
}


