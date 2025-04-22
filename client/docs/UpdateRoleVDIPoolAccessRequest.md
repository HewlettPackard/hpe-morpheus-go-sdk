# UpdateRoleVDIPoolAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VdiPoolId** | **int32** | &#x60;id&#x60; of the VDI Pool | 
**Access** | **string** | The new access level. | 

## Methods

### NewUpdateRoleVDIPoolAccessRequest

`func NewUpdateRoleVDIPoolAccessRequest(vdiPoolId int32, access string, ) *UpdateRoleVDIPoolAccessRequest`

NewUpdateRoleVDIPoolAccessRequest instantiates a new UpdateRoleVDIPoolAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleVDIPoolAccessRequestWithDefaults

`func NewUpdateRoleVDIPoolAccessRequestWithDefaults() *UpdateRoleVDIPoolAccessRequest`

NewUpdateRoleVDIPoolAccessRequestWithDefaults instantiates a new UpdateRoleVDIPoolAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVdiPoolId

`func (o *UpdateRoleVDIPoolAccessRequest) GetVdiPoolId() int32`

GetVdiPoolId returns the VdiPoolId field if non-nil, zero value otherwise.

### GetVdiPoolIdOk

`func (o *UpdateRoleVDIPoolAccessRequest) GetVdiPoolIdOk() (*int32, bool)`

GetVdiPoolIdOk returns a tuple with the VdiPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPoolId

`func (o *UpdateRoleVDIPoolAccessRequest) SetVdiPoolId(v int32)`

SetVdiPoolId sets VdiPoolId field to given value.


### GetAccess

`func (o *UpdateRoleVDIPoolAccessRequest) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleVDIPoolAccessRequest) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleVDIPoolAccessRequest) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


