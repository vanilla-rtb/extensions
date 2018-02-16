# golang extensions for vanilla-rtb

1. Generated vanilla-rtb bidder targeting caches from ```templates/matcher.tmpl``` by utilizing golang structures in ```stubs```  is used 
by vanilla-rtb bidders - the generator replaces a need for manual programming as all caches have the same pattern.

2. Generators should be capable of producing  other glue code for interfaceing with our bidder, campaign manager and  other parts of vanilla-rtb infrastructure written in C++.

3. Users should be able to either just use generators and run C++ code directly or generate both C++ and CPPGO layer to integrate with their existing DSP written in Go.

