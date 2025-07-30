package reflecttest

import (
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	name string
	age  int `sd:"Age"`
	En   []string
}

func (p *Person) Speak(content string) {
	fmt.Printf("%s speak: %v\n", p.name, content)
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func TestReflectType(t *testing.T) {
	person := Person{
		name: "wang",
		age:  11,
	}
	tt := reflect.TypeOf(person)
	fmt.Printf("tt.NumField(): %v\n", tt.NumField())

	field := tt.Field(1)
	if field2, exist := tt.FieldByName("name"); exist {
		fmt.Printf("field2.Name: %v\n", field2.Name)
	}

	tag := field.Tag.Get("sd")
	fmt.Printf("tag: %v\n", tag)

	fmt.Printf("tt.Field(1): %v\n", field.Name)

	fmt.Printf("tt.Kind(): %v\n", tt.Kind().String())

}

func TestReflectValue(t *testing.T) {

	person := Person{}

	person.name = "lilei"

	person.En = []string{"a", "b"}

	value := reflect.ValueOf(person)

	fmt.Printf("value.Elem(): %v\n", reflect.ValueOf(&person).Elem().Type())

	fmt.Printf("nums of field : %v \n", value.NumField())
	value.Field(1).SetInt(10)
	//fmt.Printf("name of person : %s \n", value.Field(0).SetInt(10))

	fmt.Printf("age of person : %v \n", value.Field(1))

	fmt.Printf("kind : %v \n", value.Kind())

	fmt.Printf("tag of person : %v \n", value.Field(2).Index(1))

	fmt.Printf("type of person: %v \n", reflect.TypeOf(person))

}

func TestReflectMethod(t *testing.T) {

	person := Person{}

	person.name = "lilei"

	person.En = []string{"a", "b"}

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
