# AddPriceSets200ResponseAllOfBudget

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

### NewAddPriceSets200ResponseAllOfBudget

`func NewAddPriceSets200ResponseAllOfBudget() *AddPriceSets200ResponseAllOfBudget`

NewAddPriceSets200ResponseAllOfBudget instantiates a new AddPriceSets200ResponseAllOfBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddPriceSets200ResponseAllOfBudget) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddPriceSets200ResponseAllOfBudget) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddPriceSets200ResponseAllOfBudget) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddPriceSets200ResponseAllOfBudget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddPriceSets200ResponseAllOfBudget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddPriceSets200ResponseAllOfBudget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddPriceSets200ResponseAllOfBudget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddPriceSets200ResponseAllOfBudget) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *AddPriceSets200ResponseAllOfBudget) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddPriceSets200ResponseAllOfBudget) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddPriceSets200ResponseAllOfBudget) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddPriceSets200ResponseAllOfBudget) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *AddPriceSets200ResponseAllOfBudget) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddPriceSets200ResponseAllOfBudget) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddPriceSets200ResponseAllOfBudget) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddPriceSets200ResponseAllOfBudget) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPriceUnit

`func (o *AddPriceSets200ResponseAllOfBudget) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *AddPriceSets200ResponseAllOfBudget) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *AddPriceSets200ResponseAllOfBudget) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *AddPriceSets200ResponseAllOfBudget) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetType

`func (o *AddPriceSets200ResponseAllOfBudget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddPriceSets200ResponseAllOfBudget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddPriceSets200ResponseAllOfBudget) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddPriceSets200ResponseAllOfBudget) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegionCode

`func (o *AddPriceSets200ResponseAllOfBudget) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *AddPriceSets200ResponseAllOfBudget) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *AddPriceSets200ResponseAllOfBudget) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *AddPriceSets200ResponseAllOfBudget) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetSystemCreated

`func (o *AddPriceSets200ResponseAllOfBudget) GetSystemCreated() bool`

GetSystemCreated returns the SystemCreated field if non-nil, zero value otherwise.

### GetSystemCreatedOk

`func (o *AddPriceSets200ResponseAllOfBudget) GetSystemCreatedOk() (*bool, bool)`

GetSystemCreatedOk returns a tuple with the SystemCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemCreated

`func (o *AddPriceSets200ResponseAllOfBudget) SetSystemCreated(v bool)`

SetSystemCreated sets SystemCreated field to given value.

### HasSystemCreated

`func (o *AddPriceSets200ResponseAllOfBudget) HasSystemCreated() bool`

HasSystemCreated returns a boolean if a field has been set.

### GetZone

`func (o *AddPriceSets200ResponseAllOfBudget) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddPriceSets200ResponseAllOfBudget) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddPriceSets200ResponseAllOfBudget) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddPriceSets200ResponseAllOfBudget) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *AddPriceSets200ResponseAllOfBudget) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *AddPriceSets200ResponseAllOfBudget) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetZonePool

`func (o *AddPriceSets200ResponseAllOfBudget) GetZonePool() string`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *AddPriceSets200ResponseAllOfBudget) GetZonePoolOk() (*string, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *AddPriceSets200ResponseAllOfBudget) SetZonePool(v string)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *AddPriceSets200ResponseAllOfBudget) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### SetZonePoolNil

`func (o *AddPriceSets200ResponseAllOfBudget) SetZonePoolNil(b bool)`

 SetZonePoolNil sets the value for ZonePool to be an explicit nil

### UnsetZonePool
`func (o *AddPriceSets200ResponseAllOfBudget) UnsetZonePool()`

UnsetZonePool ensures that no value is present for ZonePool, not even an explicit nil
### GetAccount

`func (o *AddPriceSets200ResponseAllOfBudget) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddPriceSets200ResponseAllOfBudget) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddPriceSets200ResponseAllOfBudget) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddPriceSets200ResponseAllOfBudget) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *AddPriceSets200ResponseAllOfBudget) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *AddPriceSets200ResponseAllOfBudget) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetPrices

`func (o *AddPriceSets200ResponseAllOfBudget) GetPrices() []AddPriceSets200ResponseAllOfBudgetPricesInner`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *AddPriceSets200ResponseAllOfBudget) GetPricesOk() (*[]AddPriceSets200ResponseAllOfBudgetPricesInner, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *AddPriceSets200ResponseAllOfBudget) SetPrices(v []AddPriceSets200ResponseAllOfBudgetPricesInner)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *AddPriceSets200ResponseAllOfBudget) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


