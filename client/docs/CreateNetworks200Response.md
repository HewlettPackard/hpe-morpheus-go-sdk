# CreateNetworks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**ListNetworks200ResponseAllOfNetworksInner**](ListNetworks200ResponseAllOfNetworksInner.md) |  | [optional] 
**Errors** | Pointer to **map[string]interface{}** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateNetworks200Response

`func NewCreateNetworks200Response() *CreateNetworks200Response`

NewCreateNetworks200Response instantiates a new CreateNetworks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworks200ResponseWithDefaults

`func NewCreateNetworks200ResponseWithDefaults() *CreateNetworks200Response`

NewCreateNetworks200ResponseWithDefaults instantiates a new CreateNetworks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *CreateNetworks200Response) GetNetwork() ListNetworks200ResponseAllOfNetworksInner`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateNetworks200Response) GetNetworkOk() (*ListNetworks200ResponseAllOfNetworksInner, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateNetworks200Response) SetNetwork(v ListNetworks200ResponseAllOfNetworksInner)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CreateNetworks200Response) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetErrors

`func (o *CreateNetworks200Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateNetworks200Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateNetworks200Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateNetworks200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccess

`func (o *CreateNetworks200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateNetworks200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateNetworks200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateNetworks200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *CreateNetworks200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateNetworks200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateNetworks200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateNetworks200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


