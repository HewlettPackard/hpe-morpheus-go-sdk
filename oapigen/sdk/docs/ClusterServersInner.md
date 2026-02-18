# ClusterServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TypeSet** | Pointer to [**ClusterServersInnerTypeSet**](ClusterServersInnerTypeSet.md) |  | [optional] 
**ComputeServerType** | Pointer to [**ClusterServersInnerComputeServerType**](ClusterServersInnerComputeServerType.md) |  | [optional] 

## Methods

### NewClusterServersInner

`func NewClusterServersInner() *ClusterServersInner`

NewClusterServersInner instantiates a new ClusterServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServersInnerWithDefaults

`func NewClusterServersInnerWithDefaults() *ClusterServersInner`

NewClusterServersInnerWithDefaults instantiates a new ClusterServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterServersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterServersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterServersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterServersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClusterServersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterServersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterServersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterServersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTypeSet

`func (o *ClusterServersInner) GetTypeSet() ClusterServersInnerTypeSet`

GetTypeSet returns the TypeSet field if non-nil, zero value otherwise.

### GetTypeSetOk

`func (o *ClusterServersInner) GetTypeSetOk() (*ClusterServersInnerTypeSet, bool)`

GetTypeSetOk returns a tuple with the TypeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeSet

`func (o *ClusterServersInner) SetTypeSet(v ClusterServersInnerTypeSet)`

SetTypeSet sets TypeSet field to given value.

### HasTypeSet

`func (o *ClusterServersInner) HasTypeSet() bool`

HasTypeSet returns a boolean if a field has been set.

### GetComputeServerType

`func (o *ClusterServersInner) GetComputeServerType() ClusterServersInnerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *ClusterServersInner) GetComputeServerTypeOk() (*ClusterServersInnerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *ClusterServersInner) SetComputeServerType(v ClusterServersInnerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *ClusterServersInner) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


