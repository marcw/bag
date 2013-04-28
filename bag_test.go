package bag

import (
	"fmt"
	"testing"
)

func ExampleNewBag() {
	bag := NewBag()
	bag.Set("foo", "hey, this is a string and here is a number:")
	bag.Set("bar", 42)

	foo, _ := bag.GetString("foo")
	bar, _ := bag.GetInt("bar")

	fmt.Println(foo, bar)
	// Output:
	// hey, this is a string and here is a number: 42
}

func TestSetGet(t *testing.T) {
	bag := NewBag()
	bag.Set("foo", "bar")

	v, ok := bag.GetString("foo")
	if !ok {
		t.Fail()
	}
	if v != "bar" {
		t.Fail()
	}
}

func TestSetGetBool(t *testing.T) {
	bag := NewBag()
	bag.Set("foo", true)
	v, ok := bag.GetBool("foo")
	if !ok {
		t.Fail()
	}
	if v != true {
		t.Fail()
	}
}

func TestSetGetInt(t *testing.T) {
	bag := NewBag()
	bag.Set("foo", 65535)
	v, ok := bag.GetInt("foo")
	if !ok {
		t.Fail()
	}
	if v != 65535 {
		t.Fail()
	}
}

func TestFrom(t *testing.T) {
	data := make(map[string]interface{})
	data["foo"] = "bar"
	bag := From(data)

	v, ok := bag.GetString("foo")
	if !ok {
		t.Fail()
	}
	if v != "bar" {
		t.Fail()
	}
}
