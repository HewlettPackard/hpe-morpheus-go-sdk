# AddIntegrationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integration** | [**AddIntegrationsRequestOneOf6Integration**](AddIntegrationsRequestOneOf6Integration.md) |  | 

## Methods

### NewAddIntegrationsRequest

`func NewAddIntegrationsRequest(integration AddIntegrationsRequestOneOf6Integration, ) *AddIntegrationsRequest`

NewAddIntegrationsRequest instantiates a new AddIntegrationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIntegrationsRequestWithDefaults

`func NewAddIntegrationsRequestWithDefaults() *AddIntegrationsRequest`

NewAddIntegrationsRequestWithDefaults instantiates a new AddIntegrationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegration

`func (o *AddIntegrationsRequest) GetIntegration() AddIntegrationsRequestOneOf6Integration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *AddIntegrationsRequest) GetIntegrationOk() (*AddIntegrationsRequestOneOf6Integration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *AddIntegrationsRequest) SetIntegration(v AddIntegrationsRequestOneOf6Integration)`

SetIntegration sets Integration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


