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

It is possible to list the arguments and the list of return including errors. For example:

```$bash
type bot interface{
    getGreeting(int, string,int) string
    // getGreeting(int, string,int) (string, error)
}
...
func (spanishBot) getGreeting(a int, b string, c int) string  {
    return "Hola!" + strconv.Itoa(a) + b + strconv.Itoa(c)
}
```

In addition, it is important to bear in mind:

- Interfaces are not generic types
- Interfaces are implicit because you are not necessary to declare a link between a type to an interface (_Not require code as -> type enblishBot implements..._)
- Interfaces are a contract to help us manage types
- Interfaces are not easy to understand and working with them (_Reuse code located on standard libs_)

## Http Package

Http Package has different functions, for example:

- func (c *Client) Get(url string) (resp *Response, err error)

This function has different types, for example:

- Header Header
- Body io.ReadCloser

If it is required to read the Body content, it will be necessary to work with this new type (io.ReadCloser) and searching more information about it.

### Reader Interface (io packages)

Golang has a specific interface implemente to read information from diferrent resources and tranforms them into a byte slice ([]byte). For example: 

- Http Response
- Text file on hard drive
- Image file on hard drive
- User entering text into the command line

In the following list, the type _io.ReadCloser_ which contains the Reader and Closer interfaces, in package io, is described:

- io.ReadCloser
  - Reader -> io.Reader Interface
    - Read ([]byte) (int,error) -> function
  - Closer -> io.Closer Interface
    - Close() (error) -> function

The _Reader_ integration in the code can be performed following the next steps:

- Create an empty byte slice using _make_
- Implement or call the respective function with this new argument
- Print this new value converting it to string

```$bash
...
resp, err := http.Get("http://google.com")
...
bs := make([]byte, 99999)
resp.Body.Read(bs)
fmt.Println(string(bs))
...
```

### Writer Interface (io packages)

On the contrary, Golang has a specific interface to write some information from a byte slice to the following third party sites:

- Outgoing Http Request
- Text file on hard drive
- Image file on hard drive
- Terminal


### io.Copy

When it is required to move some data from a reader to a writer, for example move a content file to the terminal, golang has a specific function implemented. 

The following example implements a data translation between a response's body to stdout terminal:

```$bash
...
// Log information in a terminal nicely
io.Copy(os.Stdout, resp.Body)
...
```

NOTE: _os.Stdout_ -> _type File_ -> _func Write_ which implements "Write" interface
