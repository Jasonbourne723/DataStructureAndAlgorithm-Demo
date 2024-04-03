package reflecttest

import (
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	name string
	age  int
	tag  []string
}

func (p *Person) Speak(content string) {
	fmt.Printf("%s speak: %v\n", p.name, content)
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func TestReflectValue(t *testing.T) {

	person := Person{}

	person.name = "lilei"

	person.tag = []string{"a", "b"}

	value := reflect.ValueOf(person)

	fmt.Printf("nums of field : %v \n", value.NumField())

	fmt.Printf("name of person : %s \n", value.Field(0))

	fmt.Printf("age of person : %v \n", value.Field(1))

	fmt.Printf("kind : %v \n", value.Kind())

	fmt.Printf("tag of person : %v \n", value.Field(2).Index(1))

}

func TestReflectMethod(t *testing.T) {

	person := Person{}

	person.name = "lilei"

	person.tag = []string{"a", "b"}

	t1 := reflect.TypeOf(&person)

	if method, ok := t1.MethodByName("Speak"); ok {

		fmt.Printf("method.Type.NumIn(): %v\n", method.Type.NumIn())

		fmt.Printf("method.Index: %v\n", method.Index)

	}

}

func TestDeepEqual(t *testing.T) {

	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 2, "a": 1}

	fmt.Printf("m1 == m2 : %v \n", reflect.DeepEqual(m1, m2))

}
