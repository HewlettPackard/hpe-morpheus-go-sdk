# ListEnvironments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | Pointer to [**[]ListEnvironments200ResponseAllOfEnvironmentsInner**](ListEnvironments200ResponseAllOfEnvironmentsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListEnvironments200Response

`func NewListEnvironments200Response() *ListEnvironments200Response`

NewListEnvironments200Response instantiates a new ListEnvironments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEnvironments200ResponseWithDefaults

`func NewListEnvironments200ResponseWithDefaults() *ListEnvironments200Response`

NewListEnvironments200ResponseWithDefaults instantiates a new ListEnvironments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *ListEnvironments200Response) GetEnvironments() []ListEnvironments200ResponseAllOfEnvironmentsInner`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ListEnvironments200Response) GetEnvironmentsOk() (*[]ListEnvironments200ResponseAllOfEnvironmentsInner, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ListEnvironments200Response) SetEnvironments(v []ListEnvironments200ResponseAllOfEnvironmentsInner)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ListEnvironments200Response) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetMeta

`func (o *ListEnvironments200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListEnvironments200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListEnvironments200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListEnvironments200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


