#!/bin/bash

echo $1

git clone --recursive https://github.com/venediktov/vanilla-rtb.git $1 
npm config set cmake_vanilla_rtb_root ${PWD}/$1

mkdir -p data
mkdir -p generator

GEN_ROOT=${PWD}/fwk/examples/bidder/generator
cp -r ${GEN_ROOT}/../data . 
pushd generator && 
python ${GEN_ROOT}/ad.py && 
python ${GEN_ROOT}/ico.py && 
python ${GEN_ROOT}/campaign_budget.py > ../data/campaign_budget 

popd && rm -rf  generator

go build -buildmode=c-archive bid_handler.go


