# ClusterNamespaceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Namespace name | 
**Description** | Pointer to **string** | Namespace description | [optional] 
**Active** | Pointer to **bool** | Namespace active | [optional] [default to false]
**ResourcePermissions** | Pointer to [**AddClusterNamespaceRequestNamespaceResourcePermissions**](AddClusterNamespaceRequestNamespaceResourcePermissions.md) |  | [optional] 

## Methods

### NewClusterNamespaceCreate

`func NewClusterNamespaceCreate(name string, ) *ClusterNamespaceCreate`

NewClusterNamespaceCreate instantiates a new ClusterNamespaceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNamespaceCreateWithDefaults

`func NewClusterNamespaceCreateWithDefaults() *ClusterNamespaceCreate`

NewClusterNamespaceCreateWithDefaults instantiates a new ClusterNamespaceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterNamespaceCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterNamespaceCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterNamespaceCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ClusterNamespaceCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterNamespaceCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterNamespaceCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterNamespaceCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActive

`func (o *ClusterNamespaceCreate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ClusterNamespaceCreate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ClusterNamespaceCreate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ClusterNamespaceCreate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *ClusterNamespaceCreate) GetResourcePermissions() AddClusterNamespaceRequestNamespaceResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ClusterNamespaceCreate) GetResourcePermissionsOk() (*AddClusterNamespaceRequestNamespaceResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ClusterNamespaceCreate) SetResourcePermissions(v AddClusterNamespaceRequestNamespaceResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ClusterNamespaceCreate) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


