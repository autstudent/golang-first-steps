package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	// contactInfo === contactInfo contactInfo
	contact contactInfo
}

func main()  {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var alex2 person 
	alex2.firstName = "Alex2"
	alex2.lastName = "Anderson2"
	fmt.Println(alex2)
	fmt.Printf("%+v", alex2)

	jim := person{
		firstName: "Jim",
		lastName: "Party",
	    // contactInfo: contactInfo{
		contact: contactInfo{
			email: "jim@gmail.com",
			zipCode: 28021,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("Pepe")
	jim.print()

	sam := person{
		firstName: "Sam",
		lastName: "Bar",
		contact: contactInfo{
			email: "sam@gmail.com",
			zipCode: 28043,
		},
	}
	sam.updateName("Pepe")
	sam.print()
}

func (pointerToPerson *person) updateName(newFirstName string)  {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print(){
	fmt.Printf("%+v", p)
}