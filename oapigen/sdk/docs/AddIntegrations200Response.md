# AddIntegrations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integration** | Pointer to [**AddIntegrations200ResponseAllOfIntegration**](AddIntegrations200ResponseAllOfIntegration.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddIntegrations200Response

`func NewAddIntegrations200Response() *AddIntegrations200Response`

NewAddIntegrations200Response instantiates a new AddIntegrations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIntegrations200ResponseWithDefaults

`func NewAddIntegrations200ResponseWithDefaults() *AddIntegrations200Response`

NewAddIntegrations200ResponseWithDefaults instantiates a new AddIntegrations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegration

`func (o *AddIntegrations200Response) GetIntegration() AddIntegrations200ResponseAllOfIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *AddIntegrations200Response) GetIntegrationOk() (*AddIntegrations200ResponseAllOfIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *AddIntegrations200Response) SetIntegration(v AddIntegrations200ResponseAllOfIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *AddIntegrations200Response) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetSuccess

`func (o *AddIntegrations200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddIntegrations200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddIntegrations200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddIntegrations200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


