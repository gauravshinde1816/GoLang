package main

import (
	"fmt"
	"math"
)

// func getName(name string) string {
// 	return "hello" + name
// }

type Shape interface {
	area() float32
}

type Circle struct {
	radius float32
}

type Rectangle struct {
	length  float32
	breadth float32
}

func (c Circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

func getArea(s Shape) float32 {
	return s.area()
}
func main() {
	// fmt.Println("Welcome to our conference booking application")
	// helper.PrintHelperMessage()

	// email, username := "gauravshinde@gmail.com", "gauravs"

	// fmt.Print(getName("Gaurav"))

	// var fruitSlices = []string{"Apple", "Orange", "Grape", "Cherry"}
	// fmt.Print(fruitSlices)
	// fmt.Print(len(fruitSlices))
	// fmt.Print(fruitSlices[1:3])

	// x := 5
	// y := 2

	// if x <= y {
	// 	fmt.Printf("X < = Y")
	// } else if x == y {
	// 	fmt.Println("X == Y")
	// } else {
	// 	fmt.Println(" X > Y")
	// }

	// i := 1

	// for i <= 10 {
	// 	fmt.Print("i :", i)
	// 	i += 1
	// }

	// maps

	// emails := map[string]string{"name": "gauravs", "email": "gauravshinde@gmail.com"}

	// fmt.Println(emails["name"])

	// var ids = []int{14, 15, 25, 36}

	// for _, id := range ids {
	// 	fmt.Print("id : ", id, " ")
	// }

	// for k, v := range emails {
	// 	fmt.Println("key: ", k, "==", "value: ", v)
	// }

	// // pointers

	// a := 3

	// b := &a

	// fmt.Printf("%d", *b)

	circle := Circle{radius: 5}
	rectangle := Rectangle{length: 5, breadth: 4}
	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))

}
