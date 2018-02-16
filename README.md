# extensions
Code Generators and Extensions for vanilla-rtb stack 

I. To install vanilla-rtb-extensions golang libraries run  ```install.sh``` or execute below commands

```bash
go get github.com/jessevdk/go-flags
go get -d github.com/vanilla-rtb/extensions
mv $GOPATH/src/github.com/vanilla-rtb/extensions/stubs $GOPATH/src/stubs
go install github.com/vanilla-rtb/extensions
go install stubs
```

Make sure you pass ```-d``` flag to  ```go get``` command when installing vanilla-rtb/extensions it will clone our project
without actually installing in the ```$GOPATH\pkg``` folder .

II. download vanilla-rtb extensions to working folder 
```
git clone --recursive clone https://github.com/vanilla-rtb/extensions 
```

III. To execute  generator execute following commands 
```
cd $GOPATH/src/github.com/vanilla-rtb/extensions
go run  bidder_generator.go --output-dir . --input-template templates/matcher.tmpl
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

[![Join the chat at https://gitter.im/vanilla-rtb/Lobby](https://badges.gitter.im/vanilla-rtb/Lobby.svg)](https://gitter.im/vanilla-rtb/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) 
