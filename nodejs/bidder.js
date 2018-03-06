// TODO: autogenerate this file with bootstraping code for bidder  
"use strict";

module.exports = require("bindings")("vanilla-rtb");

var addon = require(".");
var assert = require("assert");

if (process.argv.length <= 3) {
    console.log("Usage: " + __filename + " requires 3 parameters");
    process.exit(-1);
}
//console.log(process.argv.length);
//console.log(__filename);
console.log(process.argv[0]);
console.log(process.argv[1]);
console.log(process.argv[2]);
console.log(process.argv[3]);


var result = addon.runBidder(__filename , process.argv[2] , process.argv[3]);

console.log("Bidder Exited");

