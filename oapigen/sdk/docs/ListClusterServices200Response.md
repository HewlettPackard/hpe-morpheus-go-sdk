# ListClusterServices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to [**[]ListClusterServices200ResponseAllOfServicesInner**](ListClusterServices200ResponseAllOfServicesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterServices200Response

`func NewListClusterServices200Response() *ListClusterServices200Response`

NewListClusterServices200Response instantiates a new ListClusterServices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterServices200ResponseWithDefaults

`func NewListClusterServices200ResponseWithDefaults() *ListClusterServices200Response`

NewListClusterServices200ResponseWithDefaults instantiates a new ListClusterServices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *ListClusterServices200Response) GetServices() []ListClusterServices200ResponseAllOfServicesInner`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ListClusterServices200Response) GetServicesOk() (*[]ListClusterServices200ResponseAllOfServicesInner, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ListClusterServices200Response) SetServices(v []ListClusterServices200ResponseAllOfServicesInner)`

SetServices sets Services field to given value.

### HasServices

`func (o *ListClusterServices200Response) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterServices200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterServices200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterServices200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterServices200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


