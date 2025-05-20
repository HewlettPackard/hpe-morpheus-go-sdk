# AddCypherKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | Pointer to [**AddCypherKeyRequestTtl**](AddCypherKeyRequestTtl.md) |  | [optional] 
**Value** | Pointer to **string** | The secret value to be stored. This is only parsed if type is passed as &#x60;string&#x60;. | [optional] 

## Methods

### NewAddCypherKeyRequest

`func NewAddCypherKeyRequest() *AddCypherKeyRequest`

NewAddCypherKeyRequest instantiates a new AddCypherKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCypherKeyRequestWithDefaults

`func NewAddCypherKeyRequestWithDefaults() *AddCypherKeyRequest`

NewAddCypherKeyRequestWithDefaults instantiates a new AddCypherKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTtl

`func (o *AddCypherKeyRequest) GetTtl() AddCypherKeyRequestTtl`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AddCypherKeyRequest) GetTtlOk() (*AddCypherKeyRequestTtl, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AddCypherKeyRequest) SetTtl(v AddCypherKeyRequestTtl)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *AddCypherKeyRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetValue

`func (o *AddCypherKeyRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddCypherKeyRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddCypherKeyRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AddCypherKeyRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


