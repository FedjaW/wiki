# Channels

## WHAT IS CONCURRENCY?

Concurrency is the ability to perform multiple tasks at the same time. Typically, our code is executed one line at a time, one after the other. This is called _sequential execution_ or _synchronous execution_.

If the computer we're running our code on has multiple cores, we can even execute multiple tasks at exactly the same time. If we're running on a single core, a single core executes code at almost the same time by switching between tasks very quickly. Either way, the code we write looks the same in Go and takes advantage of whatever resources are available.

## HOW DOES CONCURRENCY WORK IN GO?

Go was designed to be concurrent, which is a trait fairly unique to Go. It excels at performing many tasks simultaneously safely using a simple syntax.

There isn't a popular programming language in existence where spawning concurrent execution is quite as elegant, at least in my opinion.

Concurrency is as simple as using the `go` keyword when calling a function:

```go
go doSomething()
```

In the example above, doSomething() will be executed concurrently with the rest of the code in the function. The go keyword is used to spawn a new goroutine.

## WHAT ARE CHANNELS?

Channels are a typed, thread-safe queue. Channels allow different goroutines to communicate with each other.

### CREATE A CHANNEL

Like maps and slices, channels must be created _before_ use. They also use the same `make` keyword:

```go
ch := make(chan int)
```

### SEND DATA TO A CHANNEL

```go
ch <- 69
```

The `<-` operator is called the _channel operator_. Data flows in the direction of the arrow. This operation will _block_ until another goroutine is ready to receive the value

### RECEIVE DATA FROM A CHANNEL

```go
v := <-ch
```

This reads and removes a value from the channel and saves it into the variable `v`. This operation will _block_ until there is a value in the channel to be read.

## BLOCKING AND DEADLOCKS

A deadlock is when a group of goroutines are all blocking so none of them can continue. This is a common bug that you need to watch out for in concurrent programming.

## TOKENS

Empty structs are often used as `tokens` in Go programs. In this context, a token is a unary value. In other words, we don't _care_ what is passed through the channel. We care _when and if_ it is passed.

We can block and wait until _something_ is sent on a channel using the following syntax

```go
<-ch
```

This will block until it pops a single item off the channel, then continue, discarding the item.

## BUFFERED CHANNELS

Channels can _optionally_ be buffered.

### CREATING A CHANNEL WITH A BUFFER

You can provide a buffer length as the second argument to `make()` to create a buffered channel:

```go
ch := make(chan int, 100)
```

Sending on a buffered channel only blocks when the buffer is _full_.

Receiving blocks only when the buffer is _empty_.

## CLOSING CHANNELS IN GO

Channels can be explicitly closed by a _sender_:

```go
ch := make(chan int)

// do some stuff with the channel

close(ch)
```

### CHECKING IF A CHANNEL IS CLOSED

Similar to the `ok` value when accessing data in a `map`, receivers can check the `ok` value when receiving from a channel to test if a channel was closed.

```go
v, ok := <-ch
```

ok is `false` if the channel is empty and closed.

### DON'T SEND ON A CLOSED CHANNEL

Sending on a closed channel will cause a panic. A panic on the main goroutine will cause the entire program to crash, and a panic in any other goroutine will cause that goroutine to crash.

Closing isn't necessary. There's nothing wrong with leaving channels open, they'll still be garbage collected if they're unused. You should close channels to indicate explicitly to a receiver that nothing else is going to come across.

## RANGE

Similar to slices and maps, channels can be ranged over.

```go
for item := range ch {
    // item is the next value received from the channel
}
```

This example will receive values over the channel (blocking at each iteration if nothing new is there) and will exit only when the channel is closed.

## SELECT

Sometimes we have a single goroutine listening to multiple channels and want to process data in the order it comes through each channel.

A `select` statement is used to listen to multiple channels at the same time. It is similar to a `switch` statement but for channels.

```go
select {
case i, ok := <- chInts:
    fmt.Println(i)
case s, ok := <- chStrings:
    fmt.Println(s)
}
```

The first channel with a value ready to be received will fire and its body will execute. If multiple channels are ready at the same time one is chosen randomly. The `ok` variable in the example above refers to whether or not the channel has been closed by the sender yet.

### SELECT DEFAULT CASE

The `default` case in a `select` statement executes _immediately_ if no other channel has a value ready. A `default` case stops the `select` statement from blocking.

```go
select {
case v := <-ch:
    // use v
default:
    // receiving from ch would block
    // so do something else
}
```

## READ-ONLY CHANNELS

A channel can be marked as read-only by casting it from a `chan` to a `<-chan` type. For example:

```go
func main() {
    ch := make(chan int)
    readCh(ch)
}

func readCh(ch <-chan int) {
    // ch can only be read from
    // in this function
}

```

## WRITE-ONLY CHANNELS

The same goes for write-only channels, but the arrow's position moves.

```go
func writeCh(ch chan<- int) {
    // ch can only be written to
    // in this function
}
```

## MORE NOTES ON CHANNELS

### A SEND TO A NIL CHANNEL BLOCKS FOREVER

```go
var c chan string // c is nil
c <- "let's get started" // blocks
```

### A RECEIVE FROM A NIL CHANNEL BLOCKS FOREVER

```go
var c chan string // c is nil
fmt.Println(<-c) // blocks
```

### A SEND TO A CLOSED CHANNEL PANICS

```go
var c = make(chan int, 100)
close(c)
c <- 1 // panic: send on closed channel
```

### A RECEIVE FROM A CLOSED CHANNEL RETURNS THE ZERO VALUE IMMEDIATELY

```go
var c = make(chan int, 100)
close(c)
fmt.Println(<-c) // 0
```
