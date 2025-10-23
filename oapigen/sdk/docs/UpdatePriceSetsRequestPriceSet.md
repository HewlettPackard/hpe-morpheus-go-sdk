# UpdatePriceSetsRequestPriceSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Price set name | [optional] 
**Code** | Pointer to **string** | Price set code. Must be unique. | [optional] 
**RegionCode** | Pointer to **string** | Price set region code | [optional] 
**Zone** | Pointer to [**AddPriceSetsRequestPriceSetZone**](AddPriceSetsRequestPriceSetZone.md) |  | [optional] 
**ZonePool** | Pointer to [**AddPriceSetsRequestPriceSetZonePool**](AddPriceSetsRequestPriceSetZonePool.md) |  | [optional] 
**PriceUnit** | Pointer to **string** | Price Unit | [optional] 
**Type** | Pointer to **string** | Price set type | [optional] 
**Prices** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewUpdatePriceSetsRequestPriceSet

`func NewUpdatePriceSetsRequestPriceSet() *UpdatePriceSetsRequestPriceSet`

NewUpdatePriceSetsRequestPriceSet instantiates a new UpdatePriceSetsRequestPriceSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePriceSetsRequestPriceSetWithDefaults

`func NewUpdatePriceSetsRequestPriceSetWithDefaults() *UpdatePriceSetsRequestPriceSet`

NewUpdatePriceSetsRequestPriceSetWithDefaults instantiates a new UpdatePriceSetsRequestPriceSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePriceSetsRequestPriceSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePriceSetsRequestPriceSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePriceSetsRequestPriceSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePriceSetsRequestPriceSet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdatePriceSetsRequestPriceSet) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdatePriceSetsRequestPriceSet) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdatePriceSetsRequestPriceSet) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdatePriceSetsRequestPriceSet) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRegionCode

`func (o *UpdatePriceSetsRequestPriceSet) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *UpdatePriceSetsRequestPriceSet) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *UpdatePriceSetsRequestPriceSet) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *UpdatePriceSetsRequestPriceSet) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetZone

`func (o *UpdatePriceSetsRequestPriceSet) GetZone() AddPriceSetsRequestPriceSetZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdatePriceSetsRequestPriceSet) GetZoneOk() (*AddPriceSetsRequestPriceSetZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdatePriceSetsRequestPriceSet) SetZone(v AddPriceSetsRequestPriceSetZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdatePriceSetsRequestPriceSet) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *UpdatePriceSetsRequestPriceSet) GetZonePool() AddPriceSetsRequestPriceSetZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *UpdatePriceSetsRequestPriceSet) GetZonePoolOk() (*AddPriceSetsRequestPriceSetZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *UpdatePriceSetsRequestPriceSet) SetZonePool(v AddPriceSetsRequestPriceSetZonePool)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *UpdatePriceSetsRequestPriceSet) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetPriceUnit

`func (o *UpdatePriceSetsRequestPriceSet) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *UpdatePriceSetsRequestPriceSet) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *UpdatePriceSetsRequestPriceSet) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *UpdatePriceSetsRequestPriceSet) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetType

`func (o *UpdatePriceSetsRequestPriceSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdatePriceSetsRequestPriceSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdatePriceSetsRequestPriceSet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdatePriceSetsRequestPriceSet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrices

`func (o *UpdatePriceSetsRequestPriceSet) GetPrices() []int64`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *UpdatePriceSetsRequestPriceSet) GetPricesOk() (*[]int64, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *UpdatePriceSetsRequestPriceSet) SetPrices(v []int64)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *UpdatePriceSetsRequestPriceSet) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


