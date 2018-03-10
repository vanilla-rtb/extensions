
#Generate code for NodeJS
go run  bidder_generator.go --output-dir nodejs/ --input-template templates/bidder.tmpl -g app -T ico -B LIB
go run  bidder_generator.go --output-dir nodejs/ --input-template templates/matcher.tmpl -g matchers

