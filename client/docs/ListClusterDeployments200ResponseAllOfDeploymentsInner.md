# ListClusterDeployments200ResponseAllOfDeploymentsInner

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

### NewListClusterDeployments200ResponseAllOfDeploymentsInner

`func NewListClusterDeployments200ResponseAllOfDeploymentsInner() *ListClusterDeployments200ResponseAllOfDeploymentsInner`

NewListClusterDeployments200ResponseAllOfDeploymentsInner instantiates a new ListClusterDeployments200ResponseAllOfDeploymentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterDeployments200ResponseAllOfDeploymentsInnerWithDefaults

`func NewListClusterDeployments200ResponseAllOfDeploymentsInnerWithDefaults() *ListClusterDeployments200ResponseAllOfDeploymentsInner`

NewListClusterDeployments200ResponseAllOfDeploymentsInnerWithDefaults instantiates a new ListClusterDeployments200ResponseAllOfDeploymentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetResourceLevel

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetResourceLevel() string`

GetResourceLevel returns the ResourceLevel field if non-nil, zero value otherwise.

### GetResourceLevelOk

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetResourceLevelOk() (*string, bool)`

GetResourceLevelOk returns a tuple with the ResourceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLevel

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) SetResourceLevel(v string)`

SetResourceLevel sets ResourceLevel field to given value.

### HasResourceLevel

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) HasResourceLevel() bool`

HasResourceLevel returns a boolean if a field has been set.

### GetResourceType

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetManaged

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetStatus

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetOwner

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetOwner() GetAlerts200ResponseAllOfChecksInnerAccount`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetOwnerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) SetOwner(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTotalCpuUsage

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetTotalCpuUsage() int64`

GetTotalCpuUsage returns the TotalCpuUsage field if non-nil, zero value otherwise.

### GetTotalCpuUsageOk

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetTotalCpuUsageOk() (*int64, bool)`

GetTotalCpuUsageOk returns a tuple with the TotalCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCpuUsage

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) SetTotalCpuUsage(v int64)`

SetTotalCpuUsage sets TotalCpuUsage field to given value.

### HasTotalCpuUsage

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) HasTotalCpuUsage() bool`

HasTotalCpuUsage returns a boolean if a field has been set.

### GetStats

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListClusterDeployments200ResponseAllOfDeploymentsInner) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


