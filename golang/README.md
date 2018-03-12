# golang extensions for vanilla-rtb

1. Generated vanilla-rtb bidder targeting caches from ```templates/matcher.tmpl``` by utilizing golang structures in ```stubs```  is used 
by vanilla-rtb bidders - the generator replaces a need for manual programming as all caches have the same pattern.

2. Generators should be capable of producing  other glue code for interfaceing with our bidder, campaign manager and  other parts of vanilla-rtb infrastructure written in C++.

3. Users should be able to either just use generators and run C++ code directly or generate both C++ and CPPGO layer to integrate with their existing DSP written in Go.

To run integration examples :

Running vanilla-rtb bidder as applcation and linked in library written in Go as a bid handler 
```
npm install
./build/vanilla-rtb-go --config config.cfg 
```

This process builds bid_handler.a from bid_handler.go and links it with bidder.cpp and other vanilla-rtb sources 

The process utlizes following commands 

```
go run  bidder_generator.go --output-dir golang/ --input-template templates/bidder.tmpl -g app -T ico -B APP 
go run  bidder_generator.go --output-dir golang/ --input-template templates/matcher.tmpl -g matchers
cd golang
go build -buildmode=c-archive bid_handler.go

```

Another way ( still work in progress ) is to use bid_handler.go not only as a handler but also as a main entry point 
and have bidder.cpp generted and compiled as a library ( e.g. __main__() instead of main() generated with -B LIB )

```
go run  bidder_generator.go --output-dir golang/ --input-template templates/bidder.tmpl -g app -T ico -B LIB
go run  bidder_generator.go --output-dir golang/ --input-template templates/matcher.tmpl -g matchers
cd golang
go build -buildmode=c-archive bid_handler.go
export CGO_LDFLAGS="-L build bid_handler.a vanilla-rtb-go.a  -L/path/to/boost/lib -lboost_program_options -lboost_system -lboost_regex" 
go build -buildmode=exe bid_handler.go
```

Running second type of build is slightly different as the last command will link everything into bid_handler executable 

```
./bid_handler --config config.cfg 
```

