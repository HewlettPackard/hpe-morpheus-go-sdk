# ListClusterPods200ResponseAllOfPodsInner

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

### NewListClusterPods200ResponseAllOfPodsInner

`func NewListClusterPods200ResponseAllOfPodsInner() *ListClusterPods200ResponseAllOfPodsInner`

NewListClusterPods200ResponseAllOfPodsInner instantiates a new ListClusterPods200ResponseAllOfPodsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterPods200ResponseAllOfPodsInnerWithDefaults

`func NewListClusterPods200ResponseAllOfPodsInnerWithDefaults() *ListClusterPods200ResponseAllOfPodsInner`

NewListClusterPods200ResponseAllOfPodsInnerWithDefaults instantiates a new ListClusterPods200ResponseAllOfPodsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterPods200ResponseAllOfPodsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterPods200ResponseAllOfPodsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterPods200ResponseAllOfPodsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterPods200ResponseAllOfPodsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListClusterPods200ResponseAllOfPodsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListClusterPods200ResponseAllOfPodsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListClusterPods200ResponseAllOfPodsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListClusterPods200ResponseAllOfPodsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListClusterPods200ResponseAllOfPodsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListClusterPods200ResponseAllOfPodsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetResourceLevel

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetResourceLevel() string`

GetResourceLevel returns the ResourceLevel field if non-nil, zero value otherwise.

### GetResourceLevelOk

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetResourceLevelOk() (*string, bool)`

GetResourceLevelOk returns a tuple with the ResourceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLevel

`func (o *ListClusterPods200ResponseAllOfPodsInner) SetResourceLevel(v string)`

SetResourceLevel sets ResourceLevel field to given value.

### HasResourceLevel

`func (o *ListClusterPods200ResponseAllOfPodsInner) HasResourceLevel() bool`

HasResourceLevel returns a boolean if a field has been set.

### GetResourceType

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ListClusterPods200ResponseAllOfPodsInner) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ListClusterPods200ResponseAllOfPodsInner) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetManaged

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ListClusterPods200ResponseAllOfPodsInner) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ListClusterPods200ResponseAllOfPodsInner) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetStatus

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListClusterPods200ResponseAllOfPodsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListClusterPods200ResponseAllOfPodsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListClusterPods200ResponseAllOfPodsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListClusterPods200ResponseAllOfPodsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetOwner

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetOwner() GetAlerts200ResponseAllOfChecksInnerAccount`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetOwnerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListClusterPods200ResponseAllOfPodsInner) SetOwner(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListClusterPods200ResponseAllOfPodsInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTotalCpuUsage

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetTotalCpuUsage() int64`

GetTotalCpuUsage returns the TotalCpuUsage field if non-nil, zero value otherwise.

### GetTotalCpuUsageOk

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetTotalCpuUsageOk() (*int64, bool)`

GetTotalCpuUsageOk returns a tuple with the TotalCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCpuUsage

`func (o *ListClusterPods200ResponseAllOfPodsInner) SetTotalCpuUsage(v int64)`

SetTotalCpuUsage sets TotalCpuUsage field to given value.

### HasTotalCpuUsage

`func (o *ListClusterPods200ResponseAllOfPodsInner) HasTotalCpuUsage() bool`

HasTotalCpuUsage returns a boolean if a field has been set.

### GetStats

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListClusterPods200ResponseAllOfPodsInner) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListClusterPods200ResponseAllOfPodsInner) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListClusterPods200ResponseAllOfPodsInner) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


