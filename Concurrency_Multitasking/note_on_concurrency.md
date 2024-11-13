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