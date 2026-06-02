# GetPriceSets200ResponsePriceSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**PriceUnit** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**RegionCode** | Pointer to **string** |  | [optional] 
**SystemCreated** | Pointer to **bool** |  | [optional] 
**Zone** | Pointer to **NullableString** |  | [optional] 
**ZonePool** | Pointer to **NullableString** |  | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 
**Prices** | Pointer to [**[]AddPriceSets200ResponseAllOfBudgetPricesInner**](AddPriceSets200ResponseAllOfBudgetPricesInner.md) |  | [optional] 

## Methods

### NewGetPriceSets200ResponsePriceSet

`func NewGetPriceSets200ResponsePriceSet() *GetPriceSets200ResponsePriceSet`

NewGetPriceSets200ResponsePriceSet instantiates a new GetPriceSets200ResponsePriceSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetPriceSets200ResponsePriceSet) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPriceSets200ResponsePriceSet) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPriceSets200ResponsePriceSet) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetPriceSets200ResponsePriceSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetPriceSets200ResponsePriceSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetPriceSets200ResponsePriceSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetPriceSets200ResponsePriceSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetPriceSets200ResponsePriceSet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetPriceSets200ResponsePriceSet) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetPriceSets200ResponsePriceSet) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetPriceSets200ResponsePriceSet) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetPriceSets200ResponsePriceSet) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *GetPriceSets200ResponsePriceSet) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetPriceSets200ResponsePriceSet) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetPriceSets200ResponsePriceSet) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetPriceSets200ResponsePriceSet) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPriceUnit

`func (o *GetPriceSets200ResponsePriceSet) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *GetPriceSets200ResponsePriceSet) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *GetPriceSets200ResponsePriceSet) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *GetPriceSets200ResponsePriceSet) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetType

`func (o *GetPriceSets200ResponsePriceSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetPriceSets200ResponsePriceSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetPriceSets200ResponsePriceSet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetPriceSets200ResponsePriceSet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegionCode

`func (o *GetPriceSets200ResponsePriceSet) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *GetPriceSets200ResponsePriceSet) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *GetPriceSets200ResponsePriceSet) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *GetPriceSets200ResponsePriceSet) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetSystemCreated

`func (o *GetPriceSets200ResponsePriceSet) GetSystemCreated() bool`

GetSystemCreated returns the SystemCreated field if non-nil, zero value otherwise.

### GetSystemCreatedOk

`func (o *GetPriceSets200ResponsePriceSet) GetSystemCreatedOk() (*bool, bool)`

GetSystemCreatedOk returns a tuple with the SystemCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemCreated

`func (o *GetPriceSets200ResponsePriceSet) SetSystemCreated(v bool)`

SetSystemCreated sets SystemCreated field to given value.

### HasSystemCreated

`func (o *GetPriceSets200ResponsePriceSet) HasSystemCreated() bool`

HasSystemCreated returns a boolean if a field has been set.

### GetZone

`func (o *GetPriceSets200ResponsePriceSet) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetPriceSets200ResponsePriceSet) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetPriceSets200ResponsePriceSet) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetPriceSets200ResponsePriceSet) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *GetPriceSets200ResponsePriceSet) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *GetPriceSets200ResponsePriceSet) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetZonePool

`func (o *GetPriceSets200ResponsePriceSet) GetZonePool() string`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *GetPriceSets200ResponsePriceSet) GetZonePoolOk() (*string, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *GetPriceSets200ResponsePriceSet) SetZonePool(v string)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *GetPriceSets200ResponsePriceSet) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### SetZonePoolNil

`func (o *GetPriceSets200ResponsePriceSet) SetZonePoolNil(b bool)`

 SetZonePoolNil sets the value for ZonePool to be an explicit nil

### UnsetZonePool
`func (o *GetPriceSets200ResponsePriceSet) UnsetZonePool()`

UnsetZonePool ensures that no value is present for ZonePool, not even an explicit nil
### GetAccount

`func (o *GetPriceSets200ResponsePriceSet) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetPriceSets200ResponsePriceSet) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetPriceSets200ResponsePriceSet) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetPriceSets200ResponsePriceSet) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *GetPriceSets200ResponsePriceSet) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *GetPriceSets200ResponsePriceSet) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetPrices

`func (o *GetPriceSets200ResponsePriceSet) GetPrices() []AddPriceSets200ResponseAllOfBudgetPricesInner`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *GetPriceSets200ResponsePriceSet) GetPricesOk() (*[]AddPriceSets200ResponseAllOfBudgetPricesInner, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *GetPriceSets200ResponsePriceSet) SetPrices(v []AddPriceSets200ResponseAllOfBudgetPricesInner)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *GetPriceSets200ResponsePriceSet) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


