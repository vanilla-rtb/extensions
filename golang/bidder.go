package main 
// #cgo darwin  LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
// #cgo !darwin LDFLAGS: -Wl,-unresolved-symbols=ignore-all
//#ifdef __cplusplus
//extern "C" {
//#endif
//  void RunBidder( char*, char*, char*);
//#ifdef __cplusplus
//}
//#endif
import "C"
import "fmt"
import "os"
import "unsafe"
import "github.com/jessevdk/go-flags"

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


type Options struct {
        Config flags.Filename `long:"config" description:"configuration file relative path" default:"etc/config.cfg"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

//currently building as 'go build -buildmode=c-archive bid_handler.go'
//the command generates bid_handler.a and bid_handler.h
//the bidder.cpp from vanilla-rtb has main() istead of __main__()
//however can be built as executable and callback handler in one file
//go build -buildmode=exe bid_handler.go 
//above needs -ldflags parameter to link with  bid_handler.a and boost libraries 
func main() {
    bidderOptions := []string { "--config" , string(options.Config) }
    C.RunBidder( (*C.char)(unsafe.Pointer(&os.Args[0])) , (*C.char)(unsafe.Pointer(&bidderOptions[0])), (*C.char)(unsafe.Pointer(&bidderOptions[1])) )
}

