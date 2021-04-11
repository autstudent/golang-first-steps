# Channels and Go Routines

Channels and Go Routines are struct in Golang to manage concurring programing, basically async calls.

## Go Routines

Golang triggers a main Go routine when you execute a golang program by default. Then, if you can execute parallel routines, it is possible to perform this new launched by the following command:

```$bash
// Previous function call
checkWeb(url)

// Create a new Go routine
go checkWeb(url)
```

## Channels

Channels are used to communicate different go routines. The main idea of channels is to have communication between go routines in order to be sure that the main routine can finish after all previous go routines are finished.

It is possible to create channels to communicate some Go Routines, and when these channels are created, it is required to define the type of the information which go through every channel. For example:

```$bash
c := make(chan string)
```

Once the channel has been created, it is possible to use it passing this value as an argument in a specific function:

```$bash
c := make(chan string)
go checkWeb(url, c)

func checkWeb(url string, c chan string) {
    ...
}
```

Lastly, it is required to send the messages through the channel. For this actions, it is required to make use of "<-" in order to send values to the channel, assign values to a new variable through the channel or make use of this value directly:

```$bash
c := make(chan string)
for _, url := range webs {
    go checkWeb(url, c)
}
fmt.Println(<- c)

func checkWeb(url string, c chan string) {
    resp, err := http.Get(url)
    ...
    c <- url + " - " + strconv.Itoa(resp.StatusCode)
}
```

NOTE: This implementation only prints the first call message, why??

### blocking channels

If the previous example is reviewed, this program only print the first call message because **blocking channel** is not implement or partially. Basically, this important feature allows programs to wait for a specific data in the channel in order to finish the main Go Routine.

When a set of calls are performed, it is required wait for the data emitted by **every function call** applying a specific structure that waits for this specific number of events or data sent. To achieve this goal, it is possible to use a _for_ loop:

```$bash
for i := 0; i < len(webs); i++ {
    fmt.Println(<-c)
}
```

Lastly, if it is required to wait **infinitely** for a specific channel messages. For this implementation, it is required:

```$bash
for l := range c {
    go checkWeb(l, c)
}
```

### Function Literals

It is possible to simplify the previous scenario including a function literals. A _function literal_ represents an anonymous function which refers to variables defined in a surrounding function.

It is useful to implement specific functions when it is required to implement specific actions with channels and Go routines:

```$bash
for l := range c {

    // Define a function which is called automatically by itself passing l argument every call
    // otherwise, we had a constant value and it always call the same url
    go func(url string) {
        time.Sleep(1 * time.Second)
        checkWeb(url, c)
    }(l)

}
```

- Bad example

```$bash
for l := range c {
    go func() {
        time.Sleep(1 * time.Second)
        checkWeb(l, c)
    }()
}
```
