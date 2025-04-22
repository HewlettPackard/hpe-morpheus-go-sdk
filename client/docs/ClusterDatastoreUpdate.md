# ClusterDatastoreUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Datastore active | [optional] [default to true]
**Permissions** | Pointer to [**UpdateClusterDatastoreRequestDatastorePermissions**](UpdateClusterDatastoreRequestDatastorePermissions.md) |  | [optional] 
**Visibility** | Pointer to **string** | Visibility for datastore | [optional] [default to "private"]

## Methods

### NewClusterDatastoreUpdate

`func NewClusterDatastoreUpdate() *ClusterDatastoreUpdate`

NewClusterDatastoreUpdate instantiates a new ClusterDatastoreUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDatastoreUpdateWithDefaults

`func NewClusterDatastoreUpdateWithDefaults() *ClusterDatastoreUpdate`

NewClusterDatastoreUpdateWithDefaults instantiates a new ClusterDatastoreUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ClusterDatastoreUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ClusterDatastoreUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ClusterDatastoreUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ClusterDatastoreUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPermissions

`func (o *ClusterDatastoreUpdate) GetPermissions() UpdateClusterDatastoreRequestDatastorePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ClusterDatastoreUpdate) GetPermissionsOk() (*UpdateClusterDatastoreRequestDatastorePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ClusterDatastoreUpdate) SetPermissions(v UpdateClusterDatastoreRequestDatastorePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ClusterDatastoreUpdate) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetVisibility

`func (o *ClusterDatastoreUpdate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ClusterDatastoreUpdate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ClusterDatastoreUpdate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ClusterDatastoreUpdate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


