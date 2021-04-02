# Deeper into Go

We are working here with an application with different pieces:

- newDeck
- print
- shuffle
- deal
- saveToFile
- newDeckFromFile

## Structs

Structs are data models which allows define a new category of object like dictionaries.

With this in mind, it is possible to create a modify this structures with the following sentences:

´´´$bash
type person struct {
	firstName string
	lastName string
}
...
alex1 := person{"Alex1","Anderson1"}

alex2 := person{firstName: "Alex2", lastName: "Anderson2"}

var alex3 person
alex2.firstName = "Alex3"
alex2.lastName = "Anderson3"
´´´

## Embedded Structs

On the other hand, it is possible to embedded structs in multiple levels. For example:

´´´$bash
type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}
...
jim := person{
    firstName: "Jim",
    lastName: "Party",
    contact: contactInfo{
        email: "jim@gmail.com",
        zipCode: 28021,
    },
}
fmt.Printf("%+v", jim)
´´´

## Structs with Receiver functions

Functions with receivers and complex structs work in the same way that the simple ones. The following example includes a new function _print_ which receives _person_ type:

´´´$bash
type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

jim := person{
    firstName: "Jim",
    lastName: "Party",
    contactInfo: contactInfo{
        email: "jim@gmail.com",
        zipCode: 28021,
    },
}

´´´

## Pointers

It is required to bear in mind that when you create a new object, this new object has a specific memory position. If you pass this object to a function with a receiver, this function will create a copy of this object creating another one. 

This concept is important because you can not modify an object passing it to a specific funcion. For example:

´´´$bash
jim := person{
    firstName: "Jim",
    lastName: "Party",
    contactInfo: contactInfo{
        email: "jim@gmail.com",
        zipCode: 28021,
    },
}
jim.updateName("pepe")
jim.print()

// This function will create a new object and a new object reference in memory (pointer)
func (p person) updateName(newFirstName string)  {
	p.firstName = newFirstName
}

// This function print Jim as firtName because as a new object reference which is copied from the original object and it is not the same as the previous one
func (p person) print(){
	fmt.Printf("%+v", p)
}
´´´

In order to be able to update the original object, you need to update name in the following way:

´´´$bash
jim := person{
    firstName: "Jim",
    lastName: "Party",
    contactInfo: contactInfo{
        email: "jim@gmail.com",
        zipCode: 28021,
    },
}
jimPointer := &jim
jimPointer.updateName("pepe")
jim.print()

func (pointerToPerson *person) updateName(newFirstName string)  {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print(){
	fmt.Printf("%+v", p)
}
´´´