# ListPriceSets200ResponseAllOfPriceSetsInner

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
**Zone** | Pointer to **string** |  | [optional] 
**ZonePool** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **string** |  | [optional] 
**Prices** | Pointer to [**[]ListPriceSets200ResponseAllOfPriceSetsInnerPricesInner**](ListPriceSets200ResponseAllOfPriceSetsInnerPricesInner.md) |  | [optional] 

## Methods

### NewListPriceSets200ResponseAllOfPriceSetsInner

`func NewListPriceSets200ResponseAllOfPriceSetsInner() *ListPriceSets200ResponseAllOfPriceSetsInner`

NewListPriceSets200ResponseAllOfPriceSetsInner instantiates a new ListPriceSets200ResponseAllOfPriceSetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPriceSets200ResponseAllOfPriceSetsInnerWithDefaults

`func NewListPriceSets200ResponseAllOfPriceSetsInnerWithDefaults() *ListPriceSets200ResponseAllOfPriceSetsInner`

NewListPriceSets200ResponseAllOfPriceSetsInnerWithDefaults instantiates a new ListPriceSets200ResponseAllOfPriceSetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPriceUnit

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetType

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegionCode

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetSystemCreated

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetSystemCreated() bool`

GetSystemCreated returns the SystemCreated field if non-nil, zero value otherwise.

### GetSystemCreatedOk

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetSystemCreatedOk() (*bool, bool)`

GetSystemCreatedOk returns a tuple with the SystemCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemCreated

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) SetSystemCreated(v bool)`

SetSystemCreated sets SystemCreated field to given value.

### HasSystemCreated

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) HasSystemCreated() bool`

HasSystemCreated returns a boolean if a field has been set.

### GetZone

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetZonePool() string`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetZonePoolOk() (*string, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) SetZonePool(v string)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetAccount

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPrices

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetPrices() []ListPriceSets200ResponseAllOfPriceSetsInnerPricesInner`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) GetPricesOk() (*[]ListPriceSets200ResponseAllOfPriceSetsInnerPricesInner, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) SetPrices(v []ListPriceSets200ResponseAllOfPriceSetsInnerPricesInner)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *ListPriceSets200ResponseAllOfPriceSetsInner) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


