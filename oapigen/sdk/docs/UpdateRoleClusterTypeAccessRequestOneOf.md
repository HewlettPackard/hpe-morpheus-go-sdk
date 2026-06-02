# UpdateRoleClusterTypeAccessRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterTypeId** | **int32** | &#x60;id&#x60; of the cluster type | 
**Access** | **string** | The new access level. | 

## Methods

### NewUpdateRoleClusterTypeAccessRequestOneOf

`func NewUpdateRoleClusterTypeAccessRequestOneOf(clusterTypeId int32, access string, ) *UpdateRoleClusterTypeAccessRequestOneOf`

NewUpdateRoleClusterTypeAccessRequestOneOf instantiates a new UpdateRoleClusterTypeAccessRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetClusterTypeId

`func (o *UpdateRoleClusterTypeAccessRequestOneOf) GetClusterTypeId() int32`

GetClusterTypeId returns the ClusterTypeId field if non-nil, zero value otherwise.

### GetClusterTypeIdOk

`func (o *UpdateRoleClusterTypeAccessRequestOneOf) GetClusterTypeIdOk() (*int32, bool)`

GetClusterTypeIdOk returns a tuple with the ClusterTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterTypeId

`func (o *UpdateRoleClusterTypeAccessRequestOneOf) SetClusterTypeId(v int32)`

SetClusterTypeId sets ClusterTypeId field to given value.


### GetAccess

`func (o *UpdateRoleClusterTypeAccessRequestOneOf) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleClusterTypeAccessRequestOneOf) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleClusterTypeAccessRequestOneOf) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


