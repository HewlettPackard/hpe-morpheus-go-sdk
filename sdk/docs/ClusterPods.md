# ClusterPods

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

### NewClusterPods

`func NewClusterPods() *ClusterPods`

NewClusterPods instantiates a new ClusterPods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPodsWithDefaults

`func NewClusterPodsWithDefaults() *ClusterPods`

NewClusterPodsWithDefaults instantiates a new ClusterPods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterPods) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterPods) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterPods) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterPods) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClusterPods) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterPods) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterPods) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterPods) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ClusterPods) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ClusterPods) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ClusterPods) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ClusterPods) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterPods) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterPods) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterPods) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterPods) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *ClusterPods) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ClusterPods) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ClusterPods) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ClusterPods) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetResourceLevel

`func (o *ClusterPods) GetResourceLevel() string`

GetResourceLevel returns the ResourceLevel field if non-nil, zero value otherwise.

### GetResourceLevelOk

`func (o *ClusterPods) GetResourceLevelOk() (*string, bool)`

GetResourceLevelOk returns a tuple with the ResourceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLevel

`func (o *ClusterPods) SetResourceLevel(v string)`

SetResourceLevel sets ResourceLevel field to given value.

### HasResourceLevel

`func (o *ClusterPods) HasResourceLevel() bool`

HasResourceLevel returns a boolean if a field has been set.

### GetResourceType

`func (o *ClusterPods) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ClusterPods) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ClusterPods) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ClusterPods) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetManaged

`func (o *ClusterPods) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ClusterPods) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ClusterPods) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ClusterPods) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterPods) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterPods) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterPods) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterPods) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ClusterPods) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ClusterPods) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ClusterPods) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ClusterPods) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetOwner

`func (o *ClusterPods) GetOwner() GetAlerts200ResponseAllOfChecksInnerAccount`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ClusterPods) GetOwnerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ClusterPods) SetOwner(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ClusterPods) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTotalCpuUsage

`func (o *ClusterPods) GetTotalCpuUsage() int64`

GetTotalCpuUsage returns the TotalCpuUsage field if non-nil, zero value otherwise.

### GetTotalCpuUsageOk

`func (o *ClusterPods) GetTotalCpuUsageOk() (*int64, bool)`

GetTotalCpuUsageOk returns a tuple with the TotalCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCpuUsage

`func (o *ClusterPods) SetTotalCpuUsage(v int64)`

SetTotalCpuUsage sets TotalCpuUsage field to given value.

### HasTotalCpuUsage

`func (o *ClusterPods) HasTotalCpuUsage() bool`

HasTotalCpuUsage returns a boolean if a field has been set.

### GetStats

`func (o *ClusterPods) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ClusterPods) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ClusterPods) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *ClusterPods) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


