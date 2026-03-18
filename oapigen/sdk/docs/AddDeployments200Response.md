# AddDeployments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployment** | Pointer to [**AddDeployments200ResponseAllOfDeployment**](AddDeployments200ResponseAllOfDeployment.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddDeployments200Response

`func NewAddDeployments200Response() *AddDeployments200Response`

NewAddDeployments200Response instantiates a new AddDeployments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDeployments200ResponseWithDefaults

`func NewAddDeployments200ResponseWithDefaults() *AddDeployments200Response`

NewAddDeployments200ResponseWithDefaults instantiates a new AddDeployments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployment

`func (o *AddDeployments200Response) GetDeployment() AddDeployments200ResponseAllOfDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *AddDeployments200Response) GetDeploymentOk() (*AddDeployments200ResponseAllOfDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *AddDeployments200Response) SetDeployment(v AddDeployments200ResponseAllOfDeployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *AddDeployments200Response) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetSuccess

`func (o *AddDeployments200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddDeployments200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddDeployments200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddDeployments200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


