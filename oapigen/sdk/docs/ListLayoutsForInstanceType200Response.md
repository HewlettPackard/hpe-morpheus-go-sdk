# ListLayoutsForInstanceType200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceTypeLayouts** | Pointer to [**[]ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner**](ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListClouds200ResponseAllOfMeta**](ListClouds200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListLayoutsForInstanceType200Response

`func NewListLayoutsForInstanceType200Response() *ListLayoutsForInstanceType200Response`

NewListLayoutsForInstanceType200Response instantiates a new ListLayoutsForInstanceType200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLayoutsForInstanceType200ResponseWithDefaults

`func NewListLayoutsForInstanceType200ResponseWithDefaults() *ListLayoutsForInstanceType200Response`

NewListLayoutsForInstanceType200ResponseWithDefaults instantiates a new ListLayoutsForInstanceType200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceTypeLayouts

`func (o *ListLayoutsForInstanceType200Response) GetInstanceTypeLayouts() []ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner`

GetInstanceTypeLayouts returns the InstanceTypeLayouts field if non-nil, zero value otherwise.

### GetInstanceTypeLayoutsOk

`func (o *ListLayoutsForInstanceType200Response) GetInstanceTypeLayoutsOk() (*[]ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner, bool)`

GetInstanceTypeLayoutsOk returns a tuple with the InstanceTypeLayouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeLayouts

`func (o *ListLayoutsForInstanceType200Response) SetInstanceTypeLayouts(v []ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner)`

SetInstanceTypeLayouts sets InstanceTypeLayouts field to given value.

### HasInstanceTypeLayouts

`func (o *ListLayoutsForInstanceType200Response) HasInstanceTypeLayouts() bool`

HasInstanceTypeLayouts returns a boolean if a field has been set.

### GetMeta

`func (o *ListLayoutsForInstanceType200Response) GetMeta() ListClouds200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListLayoutsForInstanceType200Response) GetMetaOk() (*ListClouds200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListLayoutsForInstanceType200Response) SetMeta(v ListClouds200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListLayoutsForInstanceType200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


