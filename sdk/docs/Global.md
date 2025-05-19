# Global

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Global** | Pointer to **bool** | Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users. | [optional] [default to false]

## Methods

### NewGlobal

`func NewGlobal() *Global`

NewGlobal instantiates a new Global object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalWithDefaults

`func NewGlobalWithDefaults() *Global`

NewGlobalWithDefaults instantiates a new Global object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobal

`func (o *Global) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *Global) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *Global) SetGlobal(v bool)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *Global) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


