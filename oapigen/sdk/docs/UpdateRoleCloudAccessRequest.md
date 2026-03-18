# UpdateRoleCloudAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudId** | **int32** | &#x60;id&#x60; of the cloud (zone) | 
**Access** | **string** | The new access level. | 
**AllClouds** | **bool** | Apply to all clouds | 

## Methods

### NewUpdateRoleCloudAccessRequest

`func NewUpdateRoleCloudAccessRequest(cloudId int32, access string, allClouds bool, ) *UpdateRoleCloudAccessRequest`

NewUpdateRoleCloudAccessRequest instantiates a new UpdateRoleCloudAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleCloudAccessRequestWithDefaults

`func NewUpdateRoleCloudAccessRequestWithDefaults() *UpdateRoleCloudAccessRequest`

NewUpdateRoleCloudAccessRequestWithDefaults instantiates a new UpdateRoleCloudAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudId

`func (o *UpdateRoleCloudAccessRequest) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *UpdateRoleCloudAccessRequest) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *UpdateRoleCloudAccessRequest) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.


### GetAccess

`func (o *UpdateRoleCloudAccessRequest) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleCloudAccessRequest) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleCloudAccessRequest) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetAllClouds

`func (o *UpdateRoleCloudAccessRequest) GetAllClouds() bool`

GetAllClouds returns the AllClouds field if non-nil, zero value otherwise.

### GetAllCloudsOk

`func (o *UpdateRoleCloudAccessRequest) GetAllCloudsOk() (*bool, bool)`

GetAllCloudsOk returns a tuple with the AllClouds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllClouds

`func (o *UpdateRoleCloudAccessRequest) SetAllClouds(v bool)`

SetAllClouds sets AllClouds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


