package main 
//#cgo darwin  LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
//#cgo !darwin LDFLAGS: -lstdc++ -lrt -Wl,-unresolved-symbols=ignore-all
//#include <stdlib.h>
//#ifdef __cplusplus
//extern "C" {
//#endif
// static void* allocArgv(int argc) {
//    return malloc(sizeof(char *) * argc);
//}
//  void RunBidder( int argc , char **argv);
//#ifdef __cplusplus
//}
//#endif
import "C"
import "fmt"
import "os"
import "unsafe"

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

func main() {
    argv := os.Args
    argc := C.int(len(argv))
    c_argv := (*[0xfff]*C.char)(C.allocArgv(argc))
    defer C.free(unsafe.Pointer(c_argv))

    for i, arg := range argv {
        c_argv[i] = C.CString(arg)
        defer C.free(unsafe.Pointer(c_argv[i]))
    }
    C.RunBidder(argc, (**C.char)(unsafe.Pointer(c_argv)))
}

