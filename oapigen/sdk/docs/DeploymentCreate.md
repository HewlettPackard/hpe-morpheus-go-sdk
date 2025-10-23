# DeploymentCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name, a unique identifier for the deployment | [optional] 
**Description** | Pointer to **string** | Description | [optional] 

## Methods

### NewDeploymentCreate

`func NewDeploymentCreate() *DeploymentCreate`

NewDeploymentCreate instantiates a new DeploymentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCreateWithDefaults

`func NewDeploymentCreateWithDefaults() *DeploymentCreate`

NewDeploymentCreateWithDefaults instantiates a new DeploymentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DeploymentCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeploymentCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeploymentCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeploymentCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


