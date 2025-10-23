# UpdateClusterNamespaceRequestNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Namespace name | [optional] 
**Description** | Pointer to **string** | Namespace description | [optional] 
**Active** | Pointer to **bool** | Namespace active | [optional] [default to false]
**Permissions** | Pointer to [**UpdateClusterNamespaceRequestNamespacePermissions**](UpdateClusterNamespaceRequestNamespacePermissions.md) |  | [optional] 

## Methods

### NewUpdateClusterNamespaceRequestNamespace

`func NewUpdateClusterNamespaceRequestNamespace() *UpdateClusterNamespaceRequestNamespace`

NewUpdateClusterNamespaceRequestNamespace instantiates a new UpdateClusterNamespaceRequestNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterNamespaceRequestNamespaceWithDefaults

`func NewUpdateClusterNamespaceRequestNamespaceWithDefaults() *UpdateClusterNamespaceRequestNamespace`

NewUpdateClusterNamespaceRequestNamespaceWithDefaults instantiates a new UpdateClusterNamespaceRequestNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateClusterNamespaceRequestNamespace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateClusterNamespaceRequestNamespace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateClusterNamespaceRequestNamespace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateClusterNamespaceRequestNamespace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateClusterNamespaceRequestNamespace) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateClusterNamespaceRequestNamespace) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateClusterNamespaceRequestNamespace) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateClusterNamespaceRequestNamespace) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActive

`func (o *UpdateClusterNamespaceRequestNamespace) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateClusterNamespaceRequestNamespace) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateClusterNamespaceRequestNamespace) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateClusterNamespaceRequestNamespace) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPermissions

`func (o *UpdateClusterNamespaceRequestNamespace) GetPermissions() UpdateClusterNamespaceRequestNamespacePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateClusterNamespaceRequestNamespace) GetPermissionsOk() (*UpdateClusterNamespaceRequestNamespacePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateClusterNamespaceRequestNamespace) SetPermissions(v UpdateClusterNamespaceRequestNamespacePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UpdateClusterNamespaceRequestNamespace) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


