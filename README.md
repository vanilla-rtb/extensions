# extensions
Code Generators and Extensions for vanilla-rtb stack 

to execute matcher cache genertor run following commands or execute ```install.sh```

```bash
go get github.com/jessevdk/go-flags
go get -d github.com/vanilla-rtb/extensions
mv $GOPATH/src/github.com/vanilla-rtb/extensions/stubs $GOPATH/src/stubs
go install github.com/vanilla-rtb/extensions
go install stubs
```

Make sure you pass ```-d``` flag to  ```go get``` command when installing vanilla-rtb/extensions it will clone our project
without actually installing in the ```$GOPATH\pkg``` folder .


To execute  generator execute following commands 
```
cd $GOPATH/src/github.com/vanilla-rtb/extensions
go run  bidder_generator.go --output-dir . --input-template golang/matcher.tmpl
```

Our generator is referencing ```import "stubs"``` without full path to github repo , it's treated as your local package
where you will store all of your future stubs.

The code you place in the stubs package needs to register your stub classes it's done with
```TypeRegistry``` exported package variable.
Just add more  Objects to the registry  and ```codegen``` will automatically regenerate all your stubs.  

```
var TypeRegistry = map[string]reflect.Type {
	reflect.TypeOf(Domain{}).Name() : reflect.ValueOf(Domain{}).Type(),
}
```

