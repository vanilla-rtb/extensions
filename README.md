# extensions
Code Generators and Extensions for vanilla-rtb stack 

[![Join the chat at https://gitter.im/vanilla-rtb/Lobby](https://badges.gitter.im/vanilla-rtb/Lobby.svg)](https://gitter.im/vanilla-rtb/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) 

1. To install vanilla-rtb-extensions golang libraries run  ```install.sh``` or execute below commands

```bash
go get github.com/jessevdk/go-flags
go get -d github.com/vanilla-rtb/extensions
mv $GOPATH/src/github.com/vanilla-rtb/extensions/stubs $GOPATH/src/stubs
go install github.com/vanilla-rtb/extensions
go install stubs
```

Make sure you pass ```-d``` flag to  ```go get``` command when installing vanilla-rtb/extensions it will clone our project
without actually installing in the ```$GOPATH\pkg``` folder .

The main reason behind such installation is  ```import "stubs"``` directive in our bidder_generator.go uses relative path allowing you to import your own stubs and  generate your own targeting cache not just what we provide with our examples

2. Download vanilla-rtb extensions to working folder 
```
git clone --recursive clone https://github.com/vanilla-rtb/extensions 
```

3. To execute  generator execute following commands 
```
cd extensions
go run  bidder_generator.go --output-dir . --input-template templates/matcher.tmpl
```

**Our generator is referencing** ``` import "stubs" ``` **without full path to github repo , it's treated as your local package
where you will store all of your future stubs.**

The code you place in the stubs package needs to register your stub classes it's done with
```TypeRegistry``` exported package variable.
Just add more  Objects to the registry  and ```codegen``` will automatically regenerate all your stubs.  

```
var TypeRegistry = map[string]reflect.Type {
	reflect.TypeOf(Domain{}).Name() : reflect.ValueOf(Domain{}).Type(),
}
```

The type passed to registry must be annotated with golang tags for ability to wire in-proc and shared memory correctly 
For types that allocate on the heap there needs to be a conversion ```cpp:"std::string" ipc:"char_string"```

The Domain struct tags implicetely tell generator that the lookup is done by single key ```domain name``` , the sorage structure generated for shared memory acess can be expressed as Map<string, Domain> where string type is a key and Domain type is a value. VanillaRTB relies on those structures when matching campaigns and it uses these set of ```implicitly chained structures``` where output from first becomes input for the next step in matching rule, if any matcher in the chain fails before it reaches the last matcher aka campaign-collector no bid is returned by vanilla-rtb stack.

```
type Domain struct {
    name      string `cpp:"std::string"  ipc:"char_string" is_key:"yes"`
    domain_id uint32 `cpp:"uint32_t" ipc:"uint32_t"`
}
```

Currently we only generate C++ code for vanilla-rtb library but expect our generator produce and API to bridge golang 
types like Domain with C++ IPC layer and later do the same for Java and PHP.

[![Join the chat at https://gitter.im/vanilla-rtb/Lobby](https://badges.gitter.im/vanilla-rtb/Lobby.svg)](https://gitter.im/vanilla-rtb/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) 
