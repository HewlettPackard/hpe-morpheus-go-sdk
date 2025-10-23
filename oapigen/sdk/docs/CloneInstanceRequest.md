# CloneInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the new cloned instance. If none is specified the existing name will be duplicated with the &#39;clone&#39; suffix added. | [optional] 
**Group** | Pointer to [**CloneInstanceRequestGroup**](CloneInstanceRequestGroup.md) |  | [optional] 

## Methods

### NewCloneInstanceRequest

`func NewCloneInstanceRequest() *CloneInstanceRequest`

NewCloneInstanceRequest instantiates a new CloneInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneInstanceRequestWithDefaults

`func NewCloneInstanceRequestWithDefaults() *CloneInstanceRequest`

NewCloneInstanceRequestWithDefaults instantiates a new CloneInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloneInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloneInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloneInstanceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloneInstanceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroup

`func (o *CloneInstanceRequest) GetGroup() CloneInstanceRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CloneInstanceRequest) GetGroupOk() (*CloneInstanceRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CloneInstanceRequest) SetGroup(v CloneInstanceRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CloneInstanceRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


