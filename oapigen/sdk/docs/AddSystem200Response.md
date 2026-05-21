# AddSystem200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** | ID of the newly created system. | [optional] 

## Methods

### NewAddSystem200Response

`func NewAddSystem200Response() *AddSystem200Response`

NewAddSystem200Response instantiates a new AddSystem200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSystem200ResponseWithDefaults

`func NewAddSystem200ResponseWithDefaults() *AddSystem200Response`

NewAddSystem200ResponseWithDefaults instantiates a new AddSystem200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AddSystem200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddSystem200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddSystem200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddSystem200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetId

`func (o *AddSystem200Response) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddSystem200Response) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddSystem200Response) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddSystem200Response) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


