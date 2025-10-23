# ListServicePlans200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePlans** | Pointer to [**[]ListServicePlans200ResponseAllOfServicePlansInner**](ListServicePlans200ResponseAllOfServicePlansInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListServicePlans200Response

`func NewListServicePlans200Response() *ListServicePlans200Response`

NewListServicePlans200Response instantiates a new ListServicePlans200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServicePlans200ResponseWithDefaults

`func NewListServicePlans200ResponseWithDefaults() *ListServicePlans200Response`

NewListServicePlans200ResponseWithDefaults instantiates a new ListServicePlans200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePlans

`func (o *ListServicePlans200Response) GetServicePlans() []ListServicePlans200ResponseAllOfServicePlansInner`

GetServicePlans returns the ServicePlans field if non-nil, zero value otherwise.

### GetServicePlansOk

`func (o *ListServicePlans200Response) GetServicePlansOk() (*[]ListServicePlans200ResponseAllOfServicePlansInner, bool)`

GetServicePlansOk returns a tuple with the ServicePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlans

`func (o *ListServicePlans200Response) SetServicePlans(v []ListServicePlans200ResponseAllOfServicePlansInner)`

SetServicePlans sets ServicePlans field to given value.

### HasServicePlans

`func (o *ListServicePlans200Response) HasServicePlans() bool`

HasServicePlans returns a boolean if a field has been set.

### GetMeta

`func (o *ListServicePlans200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListServicePlans200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListServicePlans200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListServicePlans200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


