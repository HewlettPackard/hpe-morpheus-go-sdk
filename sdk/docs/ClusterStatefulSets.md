# ClusterStatefulSets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ResourceLevel** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**TotalCpuUsage** | Pointer to **int64** |  | [optional] 
**Stats** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewClusterStatefulSets

`func NewClusterStatefulSets() *ClusterStatefulSets`

NewClusterStatefulSets instantiates a new ClusterStatefulSets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatefulSetsWithDefaults

`func NewClusterStatefulSetsWithDefaults() *ClusterStatefulSets`

NewClusterStatefulSetsWithDefaults instantiates a new ClusterStatefulSets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterStatefulSets) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterStatefulSets) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterStatefulSets) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterStatefulSets) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClusterStatefulSets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterStatefulSets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterStatefulSets) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterStatefulSets) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ClusterStatefulSets) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ClusterStatefulSets) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ClusterStatefulSets) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ClusterStatefulSets) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterStatefulSets) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterStatefulSets) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterStatefulSets) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterStatefulSets) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *ClusterStatefulSets) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ClusterStatefulSets) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ClusterStatefulSets) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ClusterStatefulSets) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetResourceLevel

`func (o *ClusterStatefulSets) GetResourceLevel() string`

GetResourceLevel returns the ResourceLevel field if non-nil, zero value otherwise.

### GetResourceLevelOk

`func (o *ClusterStatefulSets) GetResourceLevelOk() (*string, bool)`

GetResourceLevelOk returns a tuple with the ResourceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLevel

`func (o *ClusterStatefulSets) SetResourceLevel(v string)`

SetResourceLevel sets ResourceLevel field to given value.

### HasResourceLevel

`func (o *ClusterStatefulSets) HasResourceLevel() bool`

HasResourceLevel returns a boolean if a field has been set.

### GetResourceType

`func (o *ClusterStatefulSets) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ClusterStatefulSets) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ClusterStatefulSets) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ClusterStatefulSets) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetManaged

`func (o *ClusterStatefulSets) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ClusterStatefulSets) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ClusterStatefulSets) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ClusterStatefulSets) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterStatefulSets) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterStatefulSets) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterStatefulSets) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterStatefulSets) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ClusterStatefulSets) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ClusterStatefulSets) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ClusterStatefulSets) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ClusterStatefulSets) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetOwner

`func (o *ClusterStatefulSets) GetOwner() GetAlerts200ResponseAllOfChecksInnerAccount`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ClusterStatefulSets) GetOwnerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ClusterStatefulSets) SetOwner(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ClusterStatefulSets) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTotalCpuUsage

`func (o *ClusterStatefulSets) GetTotalCpuUsage() int64`

GetTotalCpuUsage returns the TotalCpuUsage field if non-nil, zero value otherwise.

### GetTotalCpuUsageOk

`func (o *ClusterStatefulSets) GetTotalCpuUsageOk() (*int64, bool)`

GetTotalCpuUsageOk returns a tuple with the TotalCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCpuUsage

`func (o *ClusterStatefulSets) SetTotalCpuUsage(v int64)`

SetTotalCpuUsage sets TotalCpuUsage field to given value.

### HasTotalCpuUsage

`func (o *ClusterStatefulSets) HasTotalCpuUsage() bool`

HasTotalCpuUsage returns a boolean if a field has been set.

### GetStats

`func (o *ClusterStatefulSets) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ClusterStatefulSets) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ClusterStatefulSets) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *ClusterStatefulSets) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


