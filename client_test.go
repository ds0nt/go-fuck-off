package fuck_test

import (
	"fmt"
	"testing"

	"github.com/ds0nt/go-fuck-off"
)

const (
	name      = "Jesus"
	from      = "ds0nt"
	thing     = "religion"
	reference = "Book of Ds0nt 1:5"
	company   = "Zend"
	tool      = "PHP"
	reaction  = "fix your shit"
	do        = "wash"
	something = "dishes"
	subject   = "Pretty much everyone"
	noun      = "PHP"
)

func ExampleFuckOff() {
	msg, err := fuck.Off("America", "ds0nt")
	if err != nil {
		panic(err)
	}
	fmt.Println(msg)
	// Output: Fuck off, Jesus. - ds0nt
}

func TestClient(t *testing.T) {
	msg, err := fuck.Version()
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}

	msg, err = fuck.Operations()
	if err != nil {
		t.Fail()
	} else {
		fmt.Println("Operations: OK")
	}

	msg, err = fuck.Thing(thing, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}

	msg, err = fuck.Off(name, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.You(name, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.This(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.That(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Everything(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Everyone(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Donut(name, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Shakespeare(name, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Linus(name, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.King(name, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Pink(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Life(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Chainsaw(name, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Outside(name, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Thanks(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Flying(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Fascinating(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Madison(name, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Cool(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Field(from, name, reference)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Nugget(name, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Yoda(name, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Ballmer(name, company, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.What(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Because(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.CanIUse(tool, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Bye(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Diabetes(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Bus(name, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Xmas(name, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Bday(name, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Awesome(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Tucker(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Bucket(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Family(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Shutup(name, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Zayn(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.KeepCalm(reaction, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.DoSomething(do, something, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Mornin(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Thumbs(subject, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Retard(from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
	msg, err = fuck.Greed(noun, from)
	if err != nil {
		t.Fail()
	} else {
		fmt.Println(msg)
	}
}
