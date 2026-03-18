# ListClusterStatefulSets200ResponseAllOfStatefulsetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ResourceLevel** | Pointer to **NullableString** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Owner** | Pointer to [**ListClusterStatefulSets200ResponseAllOfStatefulsetsInnerOwner**](ListClusterStatefulSets200ResponseAllOfStatefulsetsInnerOwner.md) |  | [optional] 
**TotalCpuUsage** | Pointer to **int64** |  | [optional] 
**Stats** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewListClusterStatefulSets200ResponseAllOfStatefulsetsInner

`func NewListClusterStatefulSets200ResponseAllOfStatefulsetsInner() *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner`

NewListClusterStatefulSets200ResponseAllOfStatefulsetsInner instantiates a new ListClusterStatefulSets200ResponseAllOfStatefulsetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterStatefulSets200ResponseAllOfStatefulsetsInnerWithDefaults

`func NewListClusterStatefulSets200ResponseAllOfStatefulsetsInnerWithDefaults() *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner`

NewListClusterStatefulSets200ResponseAllOfStatefulsetsInnerWithDefaults instantiates a new ListClusterStatefulSets200ResponseAllOfStatefulsetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategory

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetResourceLevel

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetResourceLevel() string`

GetResourceLevel returns the ResourceLevel field if non-nil, zero value otherwise.

### GetResourceLevelOk

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetResourceLevelOk() (*string, bool)`

GetResourceLevelOk returns a tuple with the ResourceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLevel

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) SetResourceLevel(v string)`

SetResourceLevel sets ResourceLevel field to given value.

### HasResourceLevel

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) HasResourceLevel() bool`

HasResourceLevel returns a boolean if a field has been set.

### SetResourceLevelNil

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) SetResourceLevelNil(b bool)`

 SetResourceLevelNil sets the value for ResourceLevel to be an explicit nil

### UnsetResourceLevel
`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) UnsetResourceLevel()`

UnsetResourceLevel ensures that no value is present for ResourceLevel, not even an explicit nil
### GetResourceType

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetManaged

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetStatus

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetOwner

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetOwner() ListClusterStatefulSets200ResponseAllOfStatefulsetsInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetOwnerOk() (*ListClusterStatefulSets200ResponseAllOfStatefulsetsInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) SetOwner(v ListClusterStatefulSets200ResponseAllOfStatefulsetsInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTotalCpuUsage

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetTotalCpuUsage() int64`

GetTotalCpuUsage returns the TotalCpuUsage field if non-nil, zero value otherwise.

### GetTotalCpuUsageOk

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetTotalCpuUsageOk() (*int64, bool)`

GetTotalCpuUsageOk returns a tuple with the TotalCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCpuUsage

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) SetTotalCpuUsage(v int64)`

SetTotalCpuUsage sets TotalCpuUsage field to given value.

### HasTotalCpuUsage

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) HasTotalCpuUsage() bool`

HasTotalCpuUsage returns a boolean if a field has been set.

### GetStats

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListClusterStatefulSets200ResponseAllOfStatefulsetsInner) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


