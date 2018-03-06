// TODO: autogenerate this file with bootstraping code for bidder  
"use strict";

module.exports = require("bindings")("vanilla-rtb");

var addon = require(".");
var assert = require("assert");

if (process.argv.length <= 3) {
    console.log("Usage: " + __filename + " requires 3 parameters");
    process.exit(-1);
}

var args = process.argv.slice(1);

var result = addon.runBidder(...args);

console.log("Bidder Exited");

