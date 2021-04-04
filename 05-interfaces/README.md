# Interfaces

An interface is a collection of method signatures that a Type can implement (using methods). Hence interface defines (not declares) the behavior of the object (of the type Type).

For example, a Dog can walk and bark. If an interface defines method signatures for walk and bark while Dog implements walk and bark methods, then Dog is said to implement that interface.

"Every function we ever write has to be rewritten to accommodate different types even if the logic in it is identical?"

## Define a interface

```$bash
type bot interface{
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func main()  {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string  {
	// BUSINESS LOGIC Here...
	return "Hi There!"
}

func (spanishBot) getGreeting() string  {
	// BUSINESS LOGIC Here...
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
```

In the previous example, it is possible to see a couple of structs named **englishBot** and **spanishBot**. When a structs contains a _getGreeting_ function with a return of type string, it is automatically promoted to an **honorary member** of type "bot" which has been defined above with a set of "rules" in order to limit the structs which can be integrated here.

In _printGreeting_ function, it is possible to use **englishBot** and **spanishBot** type interchangeably because of both of them meet the interface requirements.

## Rules of Interfaces

