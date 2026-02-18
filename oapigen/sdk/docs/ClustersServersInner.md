# ClustersServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TypeSet** | Pointer to [**ClustersServersInnerTypeSet**](ClustersServersInnerTypeSet.md) |  | [optional] 
**ComputeServerType** | Pointer to [**ClustersServersInnerComputeServerType**](ClustersServersInnerComputeServerType.md) |  | [optional] 

## Methods

### NewClustersServersInner

`func NewClustersServersInner() *ClustersServersInner`

NewClustersServersInner instantiates a new ClustersServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClustersServersInnerWithDefaults

`func NewClustersServersInnerWithDefaults() *ClustersServersInner`

NewClustersServersInnerWithDefaults instantiates a new ClustersServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClustersServersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClustersServersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClustersServersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClustersServersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClustersServersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClustersServersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClustersServersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClustersServersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTypeSet

`func (o *ClustersServersInner) GetTypeSet() ClustersServersInnerTypeSet`

GetTypeSet returns the TypeSet field if non-nil, zero value otherwise.

### GetTypeSetOk

`func (o *ClustersServersInner) GetTypeSetOk() (*ClustersServersInnerTypeSet, bool)`

GetTypeSetOk returns a tuple with the TypeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeSet

`func (o *ClustersServersInner) SetTypeSet(v ClustersServersInnerTypeSet)`

SetTypeSet sets TypeSet field to given value.

### HasTypeSet

`func (o *ClustersServersInner) HasTypeSet() bool`

HasTypeSet returns a boolean if a field has been set.

### GetComputeServerType

`func (o *ClustersServersInner) GetComputeServerType() ClustersServersInnerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *ClustersServersInner) GetComputeServerTypeOk() (*ClustersServersInnerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *ClustersServersInner) SetComputeServerType(v ClustersServersInnerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *ClustersServersInner) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


