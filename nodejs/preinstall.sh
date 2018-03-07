#!/bin/bash

git clone --recursive https://github.com/venediktov/vanilla-rtb.git fwk 
npm config set cmake_vanilla_rtb_root ${PWD}/fwk

mkdir -p data
mkdir -p generator

GEN_ROOT=${PWD}/fwk/examples/bidder/generator
cp -r ${GEN_ROOT}/../data . 
cd generator && 
python ${GEN_ROOT}/ad.py && 
python ${GEN_ROOT}/ico.py && 
python ${GEN_ROOT}/campaign_budget.py > ../data/campaign_budget 

cd .. && rm -rf  generator

