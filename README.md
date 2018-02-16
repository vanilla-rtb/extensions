# extensions
Code Generators and Extensions for vanilla-rtb stack 

to execute matcher cache genertor run following command
```bash
go get github.com/jessevdk/go-flags
go get -d github.com/vanilla-rtb/extensions
mv $GOPATH/src/github.com/vanilla-rtb/extensions/stubs $GOPATH/src/stubs
go install github.com/vanilla-rtb/extensions
go install stubs
cd $GOPATH/src/github.com/vanilla-rtb/extensions
go run  bidder_generator.go --output-dir . --input-template golang/matcher.tmpl
```

Make sure you pass ```-d``` flag with ```go get``` when installing vanilla-rtb/extensions it will clone our project
without actually installing in the ```$GOPATH\pkg``` folder .
Our generator is referencing ```import "stubs"``` without full path to github repo , it's treated as your local package
where you will store all of your future stubs.

The code you place in the stubs package neededs to register your stub classes it's done with a trait we provided
```TypeRegistry``` exported package variable.
Just add more  Objects to the registry  and ```codegen``` will automatically regenerate all your stubs.  

```
var TypeRegistry = map[string]reflect.Type {
	reflect.TypeOf(Domain{}).Name() : reflect.ValueOf(Domain{}).Type(),
}
```
