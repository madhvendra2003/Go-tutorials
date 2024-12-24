package main

import "fmt"

func main() {
	// car := car{}
	// car.brand = "Hyundai"
	// fmt.Println(car)

	// myCar := struct {
	// 	Make  string
	// 	Model string
	// }{
	// 	Make:  "Tesla",
	// 	Model: "model3",
	// }

	// truck := truck{}
	// truck.model = "hello"

	// fmt.Println(" %v and %v", myCar, truck)

	r := rect{56, 75}

	fmt.Println(test(r))

	auth := AuthenticationInfo{
		"Madhvendra9",
		"hellp@madhvendra",
	}

	fmt.Println(auth.GetBasicAuth())

}

func test(sh shape) float64 {
	return sh.area()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type messageToSend struct {
	phoneNumber int
	message     string
}

type car struct {
	brand      string
	model      string
	doors      int
	mileage    int
	frontWheel wheel
	backWheel  wheel
}

type wheel struct {
	radius   int
	material string
}

type truck struct {
	car
	bedsize int
}

type rect struct {
	width  int
	height int
}

func (r rect) area() float64 {
	return (float64(r.width)) * float64(r.height)
}

func (r rect) perimeter() float64 {
	return 2 * (float64(r.height + r.width))
}

type AuthenticationInfo struct {
	Username string
	Password string
}

func (auth AuthenticationInfo) GetBasicAuth() string {

	ans := fmt.Sprintf("Autherization : basic %v: %v", auth.Username, auth.Password)
	return ans

}

type shape interface {
	area() float64
	perimeter() float64
}
