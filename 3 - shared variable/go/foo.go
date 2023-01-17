// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
    "time"

)


var i = 0


func server(ch1 chan int,ch2 chan int) {
    for {
        select {
        case <-ch1:
            i++
            //Println(i)
        case <-ch2:
            i--
            //Println(i)
        }
    }



}
func incrementing(ch chan int,done chan int) {
    //TODO: increment i 1000000 times
    for j := 0; j < 1000036; j++ {
        ch <- 1
    }

    done <- 1
    
}

func decrementing(ch chan int,done chan int) {
    //TODO: decrement i 1000000 times
    for j := 0; j < 1000000; j++ {
        ch <- 1
    }
    done <- 1

}

func main() {
    // What does GOMAXPROCS do? What happens if you set it to 1?
    ch1 := make(chan int,1)
    ch2 := make(chan int, 1)
    done1 := make(chan int, 1)
    done2 := make(chan int, 1)
    runtime.GOMAXPROCS(3)
    go incrementing(ch1,done1)
    go decrementing(ch2,done2)
    go server(ch1,ch2)
	 
    // TODO: Spawn both functions as goroutines
	
    // We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
    // We will do it properly with channels soon. For now: Sleep.
    <-done1
    <-done2
    ch1 <- 1
    
    ch2 <- 1
    time.Sleep(100000)
    Println((int)time.Millisecond)

    Println("The magic number is:", i)
}
