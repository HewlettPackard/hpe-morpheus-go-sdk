# ListUsages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to [**[]ListUsages200ResponseActivityInner**](ListUsages200ResponseActivityInner.md) |  | [optional] 
**Meta** | Pointer to [**ListUsages200ResponseMeta**](ListUsages200ResponseMeta.md) |  | [optional] 

## Methods

### NewListUsages200Response

`func NewListUsages200Response() *ListUsages200Response`

NewListUsages200Response instantiates a new ListUsages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsages200ResponseWithDefaults

`func NewListUsages200ResponseWithDefaults() *ListUsages200Response`

NewListUsages200ResponseWithDefaults instantiates a new ListUsages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *ListUsages200Response) GetActivity() []ListUsages200ResponseActivityInner`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *ListUsages200Response) GetActivityOk() (*[]ListUsages200ResponseActivityInner, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *ListUsages200Response) SetActivity(v []ListUsages200ResponseActivityInner)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *ListUsages200Response) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetMeta

`func (o *ListUsages200Response) GetMeta() ListUsages200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListUsages200Response) GetMetaOk() (*ListUsages200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListUsages200Response) SetMeta(v ListUsages200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListUsages200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


