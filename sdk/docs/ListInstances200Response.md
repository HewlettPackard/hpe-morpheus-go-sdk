# ListInstances200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | Pointer to [**[]ListInstances200ResponseAllOfInstancesInner**](ListInstances200ResponseAllOfInstancesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListInstances200Response

`func NewListInstances200Response() *ListInstances200Response`

NewListInstances200Response instantiates a new ListInstances200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstances200ResponseWithDefaults

`func NewListInstances200ResponseWithDefaults() *ListInstances200Response`

NewListInstances200ResponseWithDefaults instantiates a new ListInstances200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *ListInstances200Response) GetInstances() []ListInstances200ResponseAllOfInstancesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ListInstances200Response) GetInstancesOk() (*[]ListInstances200ResponseAllOfInstancesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ListInstances200Response) SetInstances(v []ListInstances200ResponseAllOfInstancesInner)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ListInstances200Response) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMeta

`func (o *ListInstances200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListInstances200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListInstances200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListInstances200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


