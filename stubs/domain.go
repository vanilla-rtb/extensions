package stubs

import (
    "reflect"
)

type Domain struct {
    name      string `cpp:"std::string"  ipc:"shared_string" is_key:"yes"`
    domain_id uint32 `cpp:"uint32_t" ipc:"uint32_t" is_value:"yes"`
}

type ICOCampaign struct {
    domain_id   uint32 `cpp:"uint32_t" ipc:"uint32_t" is_key:"yes"`
    campaign_id uint32 `cpp:"uint32_t" ipc:"uint32_t"  is_value:"yes"`
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


//register all types for generator even unrelated can go all in one registry
var TypeRegistry = []reflect.Type{
    reflect.ValueOf(Domain{}).Type(),
    reflect.ValueOf(ICOCampaign{}).Type(),
    reflect.ValueOf(Geo{}).Type(),
    reflect.ValueOf(GeoCampaign{}).Type(),
}

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
