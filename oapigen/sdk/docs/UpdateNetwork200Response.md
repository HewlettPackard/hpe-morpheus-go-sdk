# UpdateNetwork200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**UpdateNetwork200ResponseAllOfNetwork**](UpdateNetwork200ResponseAllOfNetwork.md) |  | [optional] 
**Errors** | Pointer to **map[string]interface{}** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Msg** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateNetwork200Response

`func NewUpdateNetwork200Response() *UpdateNetwork200Response`

NewUpdateNetwork200Response instantiates a new UpdateNetwork200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetwork200ResponseWithDefaults

`func NewUpdateNetwork200ResponseWithDefaults() *UpdateNetwork200Response`

NewUpdateNetwork200ResponseWithDefaults instantiates a new UpdateNetwork200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *UpdateNetwork200Response) GetNetwork() UpdateNetwork200ResponseAllOfNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *UpdateNetwork200Response) GetNetworkOk() (*UpdateNetwork200ResponseAllOfNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *UpdateNetwork200Response) SetNetwork(v UpdateNetwork200ResponseAllOfNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *UpdateNetwork200Response) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetErrors

`func (o *UpdateNetwork200Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UpdateNetwork200Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UpdateNetwork200Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *UpdateNetwork200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateNetwork200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateNetwork200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateNetwork200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateNetwork200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *UpdateNetwork200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UpdateNetwork200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UpdateNetwork200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *UpdateNetwork200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *UpdateNetwork200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *UpdateNetwork200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


