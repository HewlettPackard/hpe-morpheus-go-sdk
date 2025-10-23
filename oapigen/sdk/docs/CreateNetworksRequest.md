# CreateNetworksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**CreateNetworksRequestNetwork**](CreateNetworksRequestNetwork.md) |  | [optional] 

## Methods

### NewCreateNetworksRequest

`func NewCreateNetworksRequest() *CreateNetworksRequest`

NewCreateNetworksRequest instantiates a new CreateNetworksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworksRequestWithDefaults

`func NewCreateNetworksRequestWithDefaults() *CreateNetworksRequest`

NewCreateNetworksRequestWithDefaults instantiates a new CreateNetworksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *CreateNetworksRequest) GetNetwork() CreateNetworksRequestNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateNetworksRequest) GetNetworkOk() (*CreateNetworksRequestNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateNetworksRequest) SetNetwork(v CreateNetworksRequestNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CreateNetworksRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


