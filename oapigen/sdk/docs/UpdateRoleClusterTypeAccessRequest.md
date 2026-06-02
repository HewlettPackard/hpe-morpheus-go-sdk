# UpdateRoleClusterTypeAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterTypeId** | **int32** | &#x60;id&#x60; of the cluster type | 
**Access** | **string** | The new access level. | 
**AllClusterTypes** | **bool** | Apply to all cluster types | 

## Methods

### NewUpdateRoleClusterTypeAccessRequest

`func NewUpdateRoleClusterTypeAccessRequest(clusterTypeId int32, access string, allClusterTypes bool, ) *UpdateRoleClusterTypeAccessRequest`

NewUpdateRoleClusterTypeAccessRequest instantiates a new UpdateRoleClusterTypeAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetClusterTypeId

`func (o *UpdateRoleClusterTypeAccessRequest) GetClusterTypeId() int32`

GetClusterTypeId returns the ClusterTypeId field if non-nil, zero value otherwise.

### GetClusterTypeIdOk

`func (o *UpdateRoleClusterTypeAccessRequest) GetClusterTypeIdOk() (*int32, bool)`

GetClusterTypeIdOk returns a tuple with the ClusterTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterTypeId

`func (o *UpdateRoleClusterTypeAccessRequest) SetClusterTypeId(v int32)`

SetClusterTypeId sets ClusterTypeId field to given value.


### GetAccess

`func (o *UpdateRoleClusterTypeAccessRequest) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleClusterTypeAccessRequest) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleClusterTypeAccessRequest) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetAllClusterTypes

`func (o *UpdateRoleClusterTypeAccessRequest) GetAllClusterTypes() bool`

GetAllClusterTypes returns the AllClusterTypes field if non-nil, zero value otherwise.

### GetAllClusterTypesOk

`func (o *UpdateRoleClusterTypeAccessRequest) GetAllClusterTypesOk() (*bool, bool)`

GetAllClusterTypesOk returns a tuple with the AllClusterTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllClusterTypes

`func (o *UpdateRoleClusterTypeAccessRequest) SetAllClusterTypes(v bool)`

SetAllClusterTypes sets AllClusterTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


