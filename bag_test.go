package bag

import (
	"fmt"
	"testing"
)

func ExampleNewBag() {
	bag := NewBag()
	bag.Set("foo", "hey, this is a string and here is a number:")
	bag.Set("bar", 42)

	fmt.Println(bag.GetString("foo"), bag.GetInt("bar"))
	// Output:
	// hey, this is a string and here is a number: 42
}

func TestSetGet(t *testing.T) {
	bag := NewBag()
	bag.Set("foo", "bar")

	if bag.GetString("foo") != "bar" {
		t.Fail()
	}
}

func TestSetGetBool(t *testing.T) {
	bag := NewBag()
	bag.Set("foo", true)
	if bag.GetBool("foo") != true {
		t.Fail()
	}
}

func TestSetGetInt(t *testing.T) {
	bag := NewBag()
	bag.Set("foo", 65535)
	if bag.GetInt("foo") != 65535 {
		t.Fail()
	}
}

func TestFrom(t *testing.T) {
	data := make(map[string]interface{})
	data["foo"] = "bar"
	bag := From(data)

	if bag.GetString("foo") != "bar" {
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	bag := NewBag()
	bag.Set("foo", "bar")
	b2 := From(bag.Map())

	if b2.GetString("foo") != "bar" {
		t.Fail()
	}
}

func TestHas(t *testing.T) {
	bag := NewBag()
	bag.Set("foo", "bar")

	if !bag.Has("foo") {
		t.Fail()
	}
	if bag.Has("bar") {
		t.Fail()
	}
}

