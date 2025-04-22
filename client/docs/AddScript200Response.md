# AddScript200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ContainerScript** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 

## Methods

### NewAddScript200Response

`func NewAddScript200Response() *AddScript200Response`

NewAddScript200Response instantiates a new AddScript200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddScript200ResponseWithDefaults

`func NewAddScript200ResponseWithDefaults() *AddScript200Response`

NewAddScript200ResponseWithDefaults instantiates a new AddScript200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AddScript200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddScript200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddScript200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddScript200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetContainerScript

`func (o *AddScript200Response) GetContainerScript() GetAlerts200ResponseAllOfChecksInnerAccount`

GetContainerScript returns the ContainerScript field if non-nil, zero value otherwise.

### GetContainerScriptOk

`func (o *AddScript200Response) GetContainerScriptOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetContainerScriptOk returns a tuple with the ContainerScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScript

`func (o *AddScript200Response) SetContainerScript(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetContainerScript sets ContainerScript field to given value.

### HasContainerScript

`func (o *AddScript200Response) HasContainerScript() bool`

HasContainerScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


