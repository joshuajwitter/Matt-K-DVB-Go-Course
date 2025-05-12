package lesson_12

import (
	"encoding/json"
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func DoLesson() {

	firstPart()
	secondPart()
	thirdPart()
	fourthPart()
}

func firstPart() {
	employeeMap := map[string]*Employee{}

	var employeeJosh Employee

	// setting the fields manually
	employeeJosh.Name = "Josh"
	employeeJosh.Number = 1
	employeeJosh.Hired = time.Now()

	// adding the plus to placeholder gives us field names
	fmt.Printf("%T %+[1]v\n", employeeJosh)

	// struct literal
	var employeeCharlotte = Employee{
		Name:   "Charlotte",
		Number: 1, // need the comma
		// no need to set all fields
	}

	fmt.Printf("%T %#[1]v\n", employeeCharlotte)

	employeeCharlotte.Boss = &employeeJosh

	employeeMap[employeeJosh.Name] = &employeeJosh
	employeeMap[employeeCharlotte.Name] = &employeeCharlotte

	printEmployeeBossName(employeeJosh.Name, employeeMap)
	printEmployeeBossName(employeeCharlotte.Name, employeeMap)
	printEmployeeBossName("Gina", employeeMap)
}

func printEmployeeBossName(name string, employeeMap map[string]*Employee) {
	if employee, ok := employeeMap[name]; ok {
		if employee.Boss == nil {
			fmt.Printf("Employee %s has no boss\n", name)
		} else {
			fmt.Printf("Employee %s's boss is %s\n", name, employee.Boss.Name)
		}
	} else {
		fmt.Printf("No employee with the name %s\n", name)
	}
}

func secondPart() {
	album1 := struct {
		title string
	}{
		"The White Album",
	}

	album2 := struct {
		title string
	}{
		"The Black Album",
	}

	fmt.Println(album1, album2)

	// we can assign one to the other even through they are anonymous types because they are
	// structurally the same, this does not work for named types though
	album1 = album2
	fmt.Println(album1, album2)
}

func thirdPart() {
	// you can use struct tags like this
	v1 := struct {
		X int `json:"foo"` // this struct tag shows us how to encode the
	}{1}
	v2 := struct {
		X int `json:"foo"`
	}{2}

	// you could not do this if these were named types
	v1 = v2

	// even if they did have different names, you could still convert them like this:
	// v1 = T1(v2)

	fmt.Println(v1)
}

// example of json encoding of a struct
type Response struct {
	Page  int      `json:"page"`
	Words []string `json:"words,omitempty"` // nice way to hide empty fields from the json
}

func fourthPart() {
	r := &Response{Page: 1, Words: []string{"up", "down", "in", "out"}}

	// encode some json
	j, _ := json.Marshal(r) // ignore the error for this example, normally handle it
	fmt.Printf("%#v\n", r)
	fmt.Printf("%#v\n", j) // the json Marshal gives you a byte slice which we need to make a string
	fmt.Println(string(j))

	// unmarshall the json into a response object
	var r2 Response
	_ = json.Unmarshal(j, &r2)
	fmt.Printf("%#v\n", r2)

	// note that the field is not exported to json due to `omitempty`
	r3 := &Response{Page: 1, Words: []string{}}
	j3, _ := json.Marshal(r3) // ignore the error for this example, normally handle it
	fmt.Println(string(j3))
}
