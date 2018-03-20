#!/bin/bash

echo $1

git clone --recursive https://github.com/venediktov/vanilla-rtb.git $1 
npm config set cmake_vanilla_rtb_root ${PWD}/$1

rm -rf build/

case "$BUILDTYPE" in
   cpp2go)
   npm config delete cmake_go2cpp
   npm config set cmake_${BUILDTYPE}=BUILDTYPE
   go run  ../bidder_generator.go --output-dir . --input-template ../templates/biddergo.tmpl -g app -T ico -B APP
   go run  ../bidder_generator.go --output-dir . --input-template ../templates/matcher.tmpl -g matchers
   go build -buildmode=c-archive bid_handler.go
   ;;
   go2cpp)
   npm config delete cmake_cpp2go
   npm config set cmake_${BUILDTYPE}=BUILDTYPE
   go run  ../bidder_generator.go --output-dir . --input-template ../templates/biddergo.tmpl -g app -T ico -B LIB
   go run  ../bidder_generator.go --output-dir . --input-template ../templates/matcher.tmpl -g matchers
   go build -buildmode=c-archive bidder.go
   ;;
   *)
   echo "Usage: BUILDTYPE=<go2cpp|cpp2go> npm install"
   exit -1
   ;;
esac

mkdir -p data
mkdir -p generator

GEN_ROOT=${PWD}/fwk/examples/bidder/generator
cp -r ${GEN_ROOT}/../data . 
pushd generator && 
python ${GEN_ROOT}/ad.py && 
python ${GEN_ROOT}/ico.py && 
python ${GEN_ROOT}/campaign_budget.py > ../data/campaign_budget 

popd && rm -rf  generator

