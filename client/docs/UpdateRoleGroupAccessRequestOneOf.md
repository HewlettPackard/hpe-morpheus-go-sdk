# UpdateRoleGroupAccessRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **int32** | &#x60;id&#x60; of the group (site) | 
**Access** | **string** | The new access level. | 

## Methods

### NewUpdateRoleGroupAccessRequestOneOf

`func NewUpdateRoleGroupAccessRequestOneOf(groupId int32, access string, ) *UpdateRoleGroupAccessRequestOneOf`

NewUpdateRoleGroupAccessRequestOneOf instantiates a new UpdateRoleGroupAccessRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleGroupAccessRequestOneOfWithDefaults

`func NewUpdateRoleGroupAccessRequestOneOfWithDefaults() *UpdateRoleGroupAccessRequestOneOf`

NewUpdateRoleGroupAccessRequestOneOfWithDefaults instantiates a new UpdateRoleGroupAccessRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *UpdateRoleGroupAccessRequestOneOf) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UpdateRoleGroupAccessRequestOneOf) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UpdateRoleGroupAccessRequestOneOf) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetAccess

`func (o *UpdateRoleGroupAccessRequestOneOf) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleGroupAccessRequestOneOf) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleGroupAccessRequestOneOf) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


