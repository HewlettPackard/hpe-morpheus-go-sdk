# AddPriceSetsRequestPriceSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Price set name | 
**Code** | **string** | Price set code. Must be unique. | 
**RegionCode** | Pointer to **string** | Price set region code | [optional] 
**Zone** | Pointer to [**AddPriceSetsRequestPriceSetZone**](AddPriceSetsRequestPriceSetZone.md) |  | [optional] 
**ZonePool** | Pointer to [**AddPriceSetsRequestPriceSetZonePool**](AddPriceSetsRequestPriceSetZonePool.md) |  | [optional] 
**PriceUnit** | **string** | Price Unit | 
**Type** | **string** | Price set type | 
**Prices** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewAddPriceSetsRequestPriceSet

`func NewAddPriceSetsRequestPriceSet(name string, code string, priceUnit string, type_ string, ) *AddPriceSetsRequestPriceSet`

NewAddPriceSetsRequestPriceSet instantiates a new AddPriceSetsRequestPriceSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPriceSetsRequestPriceSetWithDefaults

`func NewAddPriceSetsRequestPriceSetWithDefaults() *AddPriceSetsRequestPriceSet`

NewAddPriceSetsRequestPriceSetWithDefaults instantiates a new AddPriceSetsRequestPriceSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddPriceSetsRequestPriceSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddPriceSetsRequestPriceSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddPriceSetsRequestPriceSet) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *AddPriceSetsRequestPriceSet) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddPriceSetsRequestPriceSet) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddPriceSetsRequestPriceSet) SetCode(v string)`

SetCode sets Code field to given value.


### GetRegionCode

`func (o *AddPriceSetsRequestPriceSet) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *AddPriceSetsRequestPriceSet) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *AddPriceSetsRequestPriceSet) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *AddPriceSetsRequestPriceSet) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetZone

`func (o *AddPriceSetsRequestPriceSet) GetZone() AddPriceSetsRequestPriceSetZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddPriceSetsRequestPriceSet) GetZoneOk() (*AddPriceSetsRequestPriceSetZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddPriceSetsRequestPriceSet) SetZone(v AddPriceSetsRequestPriceSetZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddPriceSetsRequestPriceSet) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *AddPriceSetsRequestPriceSet) GetZonePool() AddPriceSetsRequestPriceSetZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *AddPriceSetsRequestPriceSet) GetZonePoolOk() (*AddPriceSetsRequestPriceSetZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *AddPriceSetsRequestPriceSet) SetZonePool(v AddPriceSetsRequestPriceSetZonePool)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *AddPriceSetsRequestPriceSet) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetPriceUnit

`func (o *AddPriceSetsRequestPriceSet) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *AddPriceSetsRequestPriceSet) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *AddPriceSetsRequestPriceSet) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.


### GetType

`func (o *AddPriceSetsRequestPriceSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddPriceSetsRequestPriceSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddPriceSetsRequestPriceSet) SetType(v string)`

SetType sets Type field to given value.


### GetPrices

`func (o *AddPriceSetsRequestPriceSet) GetPrices() []int64`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *AddPriceSetsRequestPriceSet) GetPricesOk() (*[]int64, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *AddPriceSetsRequestPriceSet) SetPrices(v []int64)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *AddPriceSetsRequestPriceSet) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


