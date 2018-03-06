# Node.js bindings  

Those are optional and should not be used if package.json is available 
```bash
npm install nan --save
npm install cmake-js --save
npm install boost-lib --save
npm install bindings --save
```

2.(TODO) Generate bindings for Node.js - in progress may not be needed 
```
go run  bidder_generator.go  --output-dir . --input-template nodejs.tmpl --bindings 
```

3. (TODO) Generate CMakeLists.txt 
```
go run  bidder_generator.go --output-dir . --input-template ../templates/cmakejs.tmpl -g cmake
```

4. Build vanilla-rtb.nodejs bidder  
```
npm config set cmake_VANILLA_RTB_ROOT /path/to/framework
npm install
###npm run bidder --config config.cfg ###  --config does not get passed 
node bidder.js --config config.cfg
```


