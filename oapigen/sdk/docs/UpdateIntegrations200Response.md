# UpdateIntegrations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integration** | Pointer to [**GetIntegrations200ResponseAllOfIntegration**](GetIntegrations200ResponseAllOfIntegration.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateIntegrations200Response

`func NewUpdateIntegrations200Response() *UpdateIntegrations200Response`

NewUpdateIntegrations200Response instantiates a new UpdateIntegrations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetIntegration

`func (o *UpdateIntegrations200Response) GetIntegration() GetIntegrations200ResponseAllOfIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *UpdateIntegrations200Response) GetIntegrationOk() (*GetIntegrations200ResponseAllOfIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *UpdateIntegrations200Response) SetIntegration(v GetIntegrations200ResponseAllOfIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *UpdateIntegrations200Response) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateIntegrations200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateIntegrations200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateIntegrations200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateIntegrations200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


