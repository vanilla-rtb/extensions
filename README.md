[![alt text][1.1]][1]
[![alt text][2.1]][2]
[![alt text][3.1]][3]
[![alt text][4.1]][4]
[![alt text][5.1]][5]

[1.1]: http://i.imgur.com/tXSoThF.png (twitter icon with padding)
[2.1]: http://i.imgur.com/P3YfQoD.png (facebook icon with padding)
[3.1]: http://i.imgur.com/yCsTjba.png (google plus icon with padding)
[4.1]: http://i.imgur.com/YckIOms.png (tumblr icon with padding)
[5.1]: http://i.imgur.com/0o48UoR.png (github icon with padding)

[1]: http://www.twitter.com/vanilla_rtb
[2]: http://www.linkedin.com/company/vanillartb
[3]: https://plus.google.com/+VladimirVenediktov
[4]: http://forkbid.com
[5]: http://www.github.com/vanilla-rtb

# VanillaRTB extensions
Code Generators and Extensions for vanilla-rtb stack 

[![Join the chat at https://gitter.im/vanilla-rtb/Lobby](https://badges.gitter.im/vanilla-rtb/Lobby.svg)](https://gitter.im/vanilla-rtb/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) 

1. To install vanilla-rtb-extensions golang libraries run  ```install.sh``` or execute below commands

```bash
rm -rf $GOPATH/src/github.com/vanilla-rtb/
rm -rf $GOPATH/src/stubs
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

3. To generate mathers and standalone application code execute following commands 
```
cd extensions
go run  bidder_generator.go --output-dir . --input-template templates/matcher.tmpl -g matchers
go run  bidder_generator.go --output-dir . --input-template templates/biddergo.tmpl -g app -T ico -B APP
```
4. To generate bidder library for later binding to other languages like NodeJS/Go/etc run following command sequence  
```
go run  bidder_generator.go --output-dir . --input-template templates/matcher.tmpl -g matchers
go run  ../bidder_generator.go --output-dir . --input-template ../templates/biddergo.tmpl -g app -T ico -B LIB
```
**For more information HELP in generator itself**
```
go run bidder_generator.go --help
```

**Our generator is referencing** ``` import "stubs" ``` **without full path to github repo , it's treated as your local package
where you will store all of your future stubs.**

The code you place in the stubs package needs to register your stub classes it's done with
```TypeRegistry``` exported package variable.
Just add more  Objects to the registry  and ```codegen``` will automatically regenerate all your stubs.  

```
var TypeRegistry = []reflect.Type{
    reflect.ValueOf(Domain{}).Type(),
    reflect.ValueOf(ICOCampaign{}).Type(),
    reflect.ValueOf(Geo{}).Type(),
    reflect.ValueOf(GeoCampaign{}).Type(),
}
```

It also needs to group stubs types by targeting model as shown below , the order in array coresponds to order executed by bidder in real time 

```
//agregate  targetings based on the bidder model the execution in the bidder will preserve as order of declaration
var Targetings = map[string][]reflect.Type{
    "ico": []reflect.Type{
        reflect.ValueOf(Domain{}).Type(),
        reflect.ValueOf(ICOCampaign{}).Type(),
    },
    "geo": []reflect.Type{
        reflect.ValueOf(Geo{}).Type(),
        reflect.ValueOf(GeoCampaign{}).Type(),
    },
}
```

The type passed to registry must be annotated with golang tags for ability to wire in-proc and shared memory correctly 
For types that allocate on the heap there needs to be a conversion ```cpp:"std::string" ipc:"shared_string"```

The Domain struct tags implicetely tells generator that the lookup is done by single key ```domain name``` , the sorage structure generated for shared memory acess can be expressed as Map<string, Domain> where string type is a key and Domain type is a value. VanillaRTB relies on those structures when matching campaigns and it uses these set of ```implicitly chained structures``` where output from first becomes input for the next step in matching rule, if any matcher in the chain fails before it reaches the last matcher aka campaign-collector no bid is returned by vanilla-rtb stack.
The terminal function selectingthe Ads is not included in the stubs it's implicit it comes as library by default accepting collection of Campaigns.
When looking at "ico" or "geo" examples in both cases collection of ICOCampaign or GeoCampaign is automatically fed into terminal AdSelector which matches based on the campaign ids and size of ad. 

```
type Domain struct {
    name      string `cpp:"std::string"  ipc:"shared_string" is_key:"yes"`
    domain_id uint32 `cpp:"uint32_t" ipc:"uint32_t"`
}

type ICOCampaign struct {
    domain_id   uint32 `cpp:"uint32_t" ipc:"uint32_t" is_key:"yes"`
    campaign_id uint32 `cpp:"uint32_t" ipc:"uint32_t"`
}

type Geo struct {
 city string `cpp:"std::string"  ipc:"shared_string" is_key:"yes"`
 country string `cpp:"std::string"  ipc:"shared_string" is_key:"yes"`
 geo_id uint32 `cpp:"uint32_t" ipc:"uint32_t" is_value:"yes"`
}

type GeoCampaign struct {
    geo_id      uint32 `cpp:"uint32_t" ipc:"uint32_t" is_key:"yes"`
    campaign_id uint32 `cpp:"uint32_t" ipc:"uint32_t" is_value:"yes"`
}
```

Currently we generate C++ bidder code based on vanilla-rtb library API and manually code other languages bindings.
Ideally we should generate both the C++ layer and the bindings.
 
5. Experimental generator for vanilla-rtb stand alone C++ project 
```bash
go run  bidder_generator.go --output-dir . --input-template templates/cmake.tmpl -g cmake
```

[![Join the chat at https://gitter.im/vanilla-rtb/Lobby](https://badges.gitter.im/vanilla-rtb/Lobby.svg)](https://gitter.im/vanilla-rtb/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) 
