# UpdateIntegrationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integration** | [**UpdateIntegrationsRequestOneOf6Integration**](UpdateIntegrationsRequestOneOf6Integration.md) |  | 

## Methods

### NewUpdateIntegrationsRequest

`func NewUpdateIntegrationsRequest(integration UpdateIntegrationsRequestOneOf6Integration, ) *UpdateIntegrationsRequest`

NewUpdateIntegrationsRequest instantiates a new UpdateIntegrationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIntegrationsRequestWithDefaults

`func NewUpdateIntegrationsRequestWithDefaults() *UpdateIntegrationsRequest`

NewUpdateIntegrationsRequestWithDefaults instantiates a new UpdateIntegrationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegration

`func (o *UpdateIntegrationsRequest) GetIntegration() UpdateIntegrationsRequestOneOf6Integration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *UpdateIntegrationsRequest) GetIntegrationOk() (*UpdateIntegrationsRequestOneOf6Integration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *UpdateIntegrationsRequest) SetIntegration(v UpdateIntegrationsRequestOneOf6Integration)`

SetIntegration sets Integration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


