# ClusterNamespaceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Namespace name | [optional] 
**Description** | Pointer to **string** | Namespace description | [optional] 
**Active** | Pointer to **bool** | Namespace active | [optional] [default to false]
**Permissions** | Pointer to [**UpdateClusterNamespaceRequestNamespacePermissions**](UpdateClusterNamespaceRequestNamespacePermissions.md) |  | [optional] 

## Methods

### NewClusterNamespaceUpdate

`func NewClusterNamespaceUpdate() *ClusterNamespaceUpdate`

NewClusterNamespaceUpdate instantiates a new ClusterNamespaceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNamespaceUpdateWithDefaults

`func NewClusterNamespaceUpdateWithDefaults() *ClusterNamespaceUpdate`

NewClusterNamespaceUpdateWithDefaults instantiates a new ClusterNamespaceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterNamespaceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterNamespaceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterNamespaceUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterNamespaceUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterNamespaceUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterNamespaceUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterNamespaceUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterNamespaceUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActive

`func (o *ClusterNamespaceUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ClusterNamespaceUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ClusterNamespaceUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ClusterNamespaceUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPermissions

`func (o *ClusterNamespaceUpdate) GetPermissions() UpdateClusterNamespaceRequestNamespacePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ClusterNamespaceUpdate) GetPermissionsOk() (*UpdateClusterNamespaceRequestNamespacePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ClusterNamespaceUpdate) SetPermissions(v UpdateClusterNamespaceRequestNamespacePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ClusterNamespaceUpdate) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


