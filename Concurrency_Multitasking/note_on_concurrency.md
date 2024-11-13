# Note of concurrency

## Regular execution

```go
func main(){
    greet("Nice to meet you")
    greet("How are you?")
    slowGreet("How...are..you?")
    greet("Have a good day!")
}
```

## Concurrency

* Statements with the `go` command running are run as `subroutines` to the main program.

```go
funct main(){
    go greet("Nice to meet you!")
    go greet("How are you?")
    go slowGreet("How...are...you?")
    go greet("Have a good day!")
}
```

* So all four "greet" commands will execute at the same time

* Keep in mind though print statements don't occur due to subroutine/goroutine behavior

    * Functions in subroutines run in a `non-blocking` state. Console Out statements get dispatched and do not return values
        * Dispatched like a truck sent from a trucking center. Think!!!

### Channels in Go

* Go offers a concept of `Channel` to help report back information from the "Truck" (Subroutine) back to the Trucking center (Main caller)

```go

//Function
func slowGreet(phrase string, doneChan chan bool){
    time.Sleep(3 *time.Second) //BIG TASK
    fmt.Println("Hello!", phrase)
    doneChan <- true
}

done := make(chan bool)
go slowGreet("How...are..you?", done)
fmt.Println(<- done) // Waiting for channel to emit data, don't HAVE to print it but go will only continue after the channel emits
```

* The "Channel" is a communication device, it doesn't *need* to report data back. Only that a process is finished.
   * In effect, you could use the same *channel* to emit data after subsequent calls are done.
   * THIS however, will result in a "race" condition, make sure to emit as many times as subroutines as you have.

### Channels as Slices

```go
dones := make([]chan bool, 4)

dones[0] = make(chan bool)
go greet("Nice to meet you!", dones[0])

dones[1] = make(chan bool)
go greet("How are you", dones[1])

dones[2] = make(chan bool)
go slowGreet("How...are...you?", dones[2])

dones[3] = make(chan bool)
go greet("Have a good day", dones[3])

// Iterate and emit data of finished tasks
for _, done := range dones {
    <-done
}
```

### Single Done Channel

```go
func slowGreet(phrase string, doneChan chan bool){
    time.Sleep(3*time.Second)
    fmt.Println("Hello!", phrase)
    doneChan <- true
    close(doneChan) //This keyword should be used on the operation you think will
                    // take the longest.
}


done := make(chan bool)
go greet("Nice to meet you!", done)

go greet("How are you", done)

go slowGreet("How...are...you?", done)

go greet("Have a good day", done)

// Iterate and emit data of finished tasks
for _ := range done {
}
```