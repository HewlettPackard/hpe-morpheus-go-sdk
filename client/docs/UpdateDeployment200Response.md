# UpdateDeployment200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployment** | Pointer to [**GetDeployment200ResponseDeployment**](GetDeployment200ResponseDeployment.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateDeployment200Response

`func NewUpdateDeployment200Response() *UpdateDeployment200Response`

NewUpdateDeployment200Response instantiates a new UpdateDeployment200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeployment200ResponseWithDefaults

`func NewUpdateDeployment200ResponseWithDefaults() *UpdateDeployment200Response`

NewUpdateDeployment200ResponseWithDefaults instantiates a new UpdateDeployment200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployment

`func (o *UpdateDeployment200Response) GetDeployment() GetDeployment200ResponseDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *UpdateDeployment200Response) GetDeploymentOk() (*GetDeployment200ResponseDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *UpdateDeployment200Response) SetDeployment(v GetDeployment200ResponseDeployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *UpdateDeployment200Response) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateDeployment200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateDeployment200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateDeployment200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateDeployment200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


