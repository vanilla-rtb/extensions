package main 

import "C"
import "fmt"

var c  = make(chan string)

//export HandleBid
func HandleBid(s string) {
    go func () {
        msg,ok := <- c
        if ( ok ) {
            fmt.Printf("HandleBid:%s\n" , msg)
        }
    }()
    c <- s
}

func main() {}

