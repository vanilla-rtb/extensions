#!/bin/bash

echo $1

rm -rf $1

git clone --recursive https://github.com/venediktov/vanilla-rtb.git $1 
npm config set cmake_vanilla_rtb_root ${PWD}/$1

#generate generic bidder code withouot callback to node.js
go run  ../bidder_generator.go --output-dir . --input-template ../templates/bidder.tmpl -g app -T ico -B LIB
go run  ../bidder_generator.go --output-dir . --input-template ../templates/matcher.tmpl -g matchers

mkdir -p data
mkdir -p generator

GEN_ROOT=${PWD}/fwk/examples/bidder/generator
cp -r ${GEN_ROOT}/../data . 
cd generator && 
python ${GEN_ROOT}/ad.py && 
python ${GEN_ROOT}/ico.py && 
python ${GEN_ROOT}/campaign_budget.py > ../data/campaign_budget 

#TODO generate all data format correcly reflection can be used to determine order of fields
#current format of ico_domain is reverse domain_id\tdomain_name it should be in reverse order
#generator can automatically inspect and generate data in correct order
#TODO: remove these 3 lines when Go generator for data is implemented
mv ../data/ico_domains ../data/domain
mv ../data/ico_campaign ../data/icocampaign
mv ../data/ico_ads ../data/ads
##################################

cd .. && rm -rf  generator

