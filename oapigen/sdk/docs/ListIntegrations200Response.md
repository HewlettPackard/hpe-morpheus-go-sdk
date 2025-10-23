# ListIntegrations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integrations** | Pointer to [**[]ListIntegrations200ResponseAnyOfIntegrationsInner**](ListIntegrations200ResponseAnyOfIntegrationsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListIntegrations200Response

`func NewListIntegrations200Response() *ListIntegrations200Response`

NewListIntegrations200Response instantiates a new ListIntegrations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIntegrations200ResponseWithDefaults

`func NewListIntegrations200ResponseWithDefaults() *ListIntegrations200Response`

NewListIntegrations200ResponseWithDefaults instantiates a new ListIntegrations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrations

`func (o *ListIntegrations200Response) GetIntegrations() []ListIntegrations200ResponseAnyOfIntegrationsInner`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *ListIntegrations200Response) GetIntegrationsOk() (*[]ListIntegrations200ResponseAnyOfIntegrationsInner, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *ListIntegrations200Response) SetIntegrations(v []ListIntegrations200ResponseAnyOfIntegrationsInner)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *ListIntegrations200Response) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetMeta

`func (o *ListIntegrations200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListIntegrations200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListIntegrations200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListIntegrations200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


