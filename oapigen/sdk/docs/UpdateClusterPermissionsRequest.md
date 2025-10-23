# UpdateClusterPermissionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | [**UpdateClusterDatastoreRequestDatastorePermissions**](UpdateClusterDatastoreRequestDatastorePermissions.md) |  | 

## Methods

### NewUpdateClusterPermissionsRequest

`func NewUpdateClusterPermissionsRequest(permissions UpdateClusterDatastoreRequestDatastorePermissions, ) *UpdateClusterPermissionsRequest`

NewUpdateClusterPermissionsRequest instantiates a new UpdateClusterPermissionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterPermissionsRequestWithDefaults

`func NewUpdateClusterPermissionsRequestWithDefaults() *UpdateClusterPermissionsRequest`

NewUpdateClusterPermissionsRequestWithDefaults instantiates a new UpdateClusterPermissionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *UpdateClusterPermissionsRequest) GetPermissions() UpdateClusterDatastoreRequestDatastorePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateClusterPermissionsRequest) GetPermissionsOk() (*UpdateClusterDatastoreRequestDatastorePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateClusterPermissionsRequest) SetPermissions(v UpdateClusterDatastoreRequestDatastorePermissions)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


