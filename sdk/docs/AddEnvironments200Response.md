# AddEnvironments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**ListEnvironments200ResponseAllOfEnvironmentsInner**](ListEnvironments200ResponseAllOfEnvironmentsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddEnvironments200Response

`func NewAddEnvironments200Response() *AddEnvironments200Response`

NewAddEnvironments200Response instantiates a new AddEnvironments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddEnvironments200ResponseWithDefaults

`func NewAddEnvironments200ResponseWithDefaults() *AddEnvironments200Response`

NewAddEnvironments200ResponseWithDefaults instantiates a new AddEnvironments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *AddEnvironments200Response) GetEnvironment() ListEnvironments200ResponseAllOfEnvironmentsInner`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AddEnvironments200Response) GetEnvironmentOk() (*ListEnvironments200ResponseAllOfEnvironmentsInner, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AddEnvironments200Response) SetEnvironment(v ListEnvironments200ResponseAllOfEnvironmentsInner)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AddEnvironments200Response) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetSuccess

`func (o *AddEnvironments200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddEnvironments200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddEnvironments200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddEnvironments200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


