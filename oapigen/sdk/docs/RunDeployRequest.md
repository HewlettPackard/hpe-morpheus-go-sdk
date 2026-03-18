# RunDeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDeploy** | Pointer to [**RunDeployRequestAppDeploy**](RunDeployRequestAppDeploy.md) |  | [optional] 

## Methods

### NewRunDeployRequest

`func NewRunDeployRequest() *RunDeployRequest`

NewRunDeployRequest instantiates a new RunDeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunDeployRequestWithDefaults

`func NewRunDeployRequestWithDefaults() *RunDeployRequest`

NewRunDeployRequestWithDefaults instantiates a new RunDeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppDeploy

`func (o *RunDeployRequest) GetAppDeploy() RunDeployRequestAppDeploy`

GetAppDeploy returns the AppDeploy field if non-nil, zero value otherwise.

### GetAppDeployOk

`func (o *RunDeployRequest) GetAppDeployOk() (*RunDeployRequestAppDeploy, bool)`

GetAppDeployOk returns a tuple with the AppDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDeploy

`func (o *RunDeployRequest) SetAppDeploy(v RunDeployRequestAppDeploy)`

SetAppDeploy sets AppDeploy field to given value.

### HasAppDeploy

`func (o *RunDeployRequest) HasAppDeploy() bool`

HasAppDeploy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


