package main

import "fmt"

// func main() {
// 	goku := Saiyan{}
// 	goku.Name = "Goku"
// 	goku.Power = 9000
// 	fmt.Println(goku)
// }

// func Super(s *Saiyan) {
// }

// func main() {
// 	goku := &Saiyan{"Goku", 9001}
// 	goku.Super()
// 	fmt.Println(goku.Power)
// }

// func main() {
// 	goku := NewSaiyan("David", 18000)
// 	goku.Super()
// 	fmt.Println(goku)
// }

// type Saiyan struct {
// 	Name   string
// 	Power  int
// 	Father *Saiyan
// }

// func (s *Saiyan) Super() {
// 	s.Power += 10000
// }

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
	*Person
	Power int
}

//OverWrite
func (s *Saiyan) Introduce() {
	fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}

func main() {
	goku := &Saiyan{
		Person: &Person{"Goku"},
		Power:  9001,
	}
	goku.Introduce()
}

// func main() {
// 	goku := &Saiyan{
// 		Name:   "Goku",
// 		Power:  9001,
// 		Father: nil,
// 	}
// 	gohan := &Saiyan{
// 		Name:   "Gohan",
// 		Power:  1000,
// 		Father: goku,
// 	}
// 	fmt.Println(gohan)
// }

// func main() {
// 	goku := new(Saiyan)
// 	goku.Name = "Joey"
// 	goku.Power = 9000
// 	goku.Super()
// 	fmt.Println(goku)
// }

// func NewSaiyan(name string, power int) *Saiyan {
// 	return &Saiyan{
// 		Name:  name,
// 		Power: power,
// 	}
// }
// func NewSaiyan(name string, power int) Saiyan {
// 	return Saiyan{
// 		Name:  name,
// 		Power: power,
// 	}
// }
