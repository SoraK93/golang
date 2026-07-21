package main

import "fmt"

type Movie struct {
	name   string
	rating float32
}

type Employee struct {
	eid int
	id  int
}

type Circle struct {
	x      int
	y      int
	radius float64
	area   float64
}

type Undergrad struct {
	name   string
	grades []int
}

type Postgrad struct {
	name   string
	grades []int
}

type Student interface {
	getPercentage() int
	getName() string
}

func main() {
	// creating and iterating list of struct
	exMovie1()

	// Comparing struct example
	exMovie2()

	// method example
	exEmployee()

	// Standard struct and method example
	exCircle()

	exGrades()
}

func getMovie(s string, r float32) (m Movie) {
	m = Movie{
		name:   s,
		rating: r,
	}
	return
}

func exMovie1() {
	mov := getMovie("xyz", 2.1)
	mov1 := getMovie("abc", 3.3)
	movies := make([]Movie, 5)
	movies = append(movies, mov)
	movies = append(movies, mov1)
	for _, value := range movies {
		fmt.Println(value)
	}
}

func exMovie2() {
	mov := Movie{"xyz", 2.1}
	mov1 := Movie{"abc", 2.1}
	if mov.rating == mov1.rating || mov != mov1 {
		fmt.Println("condition met")
	} else if mov.rating == mov1.rating {
		fmt.Println("condition_2 met")
	}
}

func (e Employee) get_id() int {
	return e.eid + 10
}

func exEmployee() {
	employees := make([]Employee, 5)
	for i := range employees {
		employees[i] = Employee{eid: i}
		employees[i].id = employees[i].get_id()
		fmt.Printf("%+v\n", employees[i])
	}
}

func (c *Circle) calcArea() {
	const PI float64 = 3.14
	c.area = PI * c.radius * c.radius
}

func exCircle() {
	c := Circle{
		x:      5,
		y:      5,
		radius: 5,
		area:   0,
	}
	fmt.Printf("%+v\n", c)

	c.calcArea()
	fmt.Printf("%+v\n", c)
}

func (p Postgrad) getPercentage() int {
	sum := 0
	for _, v := range p.grades {
		sum += v
	}
	return ((sum * 100) / (len(p.grades) * 200))
}
func (p Postgrad) getName() string {
	return p.name
}

func (u Undergrad) getPercentage() int {
	sum := 0
	for _, v := range u.grades {
		sum += v
	}
	return sum / len(u.grades)
}
func (u Undergrad) getName() string {
	return u.name
}

func printData(s Student) {
	fmt.Println(s.getName())
	fmt.Println(s.getPercentage())
}

func exGrades() {
	u := Undergrad{"Ross", []int{90, 75, 80}}
	p := Postgrad{"Joe", []int{150, 190, 185}}
	printData(u)
	printData(p)
}
