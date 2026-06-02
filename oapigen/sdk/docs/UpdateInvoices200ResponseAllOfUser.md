# UpdateInvoices200ResponseAllOfUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**OwnerId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetInvoices200ResponseAllOfInvoiceAccount**](GetInvoices200ResponseAllOfInvoiceAccount.md) |  | [optional] 
**Group** | Pointer to **map[string]interface{}** |  | [optional] 
**Cloud** | Pointer to [**GetInvoices200ResponseAllOfInvoiceCloud**](GetInvoices200ResponseAllOfInvoiceCloud.md) |  | [optional] 
**Instance** | Pointer to **map[string]interface{}** |  | [optional] 
**Server** | Pointer to **NullableString** |  | [optional] 
**Cluster** | Pointer to **NullableString** |  | [optional] 
**User** | Pointer to **map[string]interface{}** |  | [optional] 
**Plan** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Project** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**RefUuid** | Pointer to **NullableString** |  | [optional] 
**RefName** | Pointer to **string** |  | [optional] 
**RefCategory** | Pointer to **string** |  | [optional] 
**ResourceId** | Pointer to **NullableString** |  | [optional] 
**ResourceUuid** | Pointer to **NullableString** |  | [optional] 
**ResourceType** | Pointer to **NullableString** |  | [optional] 
**ResourceName** | Pointer to **NullableString** |  | [optional] 
**ResourceExternalId** | Pointer to **NullableString** |  | [optional] 
**ResourceInternalId** | Pointer to **NullableString** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**Estimate** | Pointer to **bool** |  | [optional] 
**SummaryInvoice** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**RefStart** | Pointer to **time.Time** |  | [optional] 
**RefEnd** | Pointer to **time.Time** |  | [optional] 
**EstimatedComputePrice** | Pointer to **float32** |  | [optional] 
**EstimatedComputeCost** | Pointer to **float32** |  | [optional] 
**EstimatedMemoryPrice** | Pointer to **float32** |  | [optional] 
**EstimatedMemoryCost** | Pointer to **float32** |  | [optional] 
**EstimatedStoragePrice** | Pointer to **float32** |  | [optional] 
**EstimatedStorageCost** | Pointer to **float32** |  | [optional] 
**EstimatedNetworkPrice** | Pointer to **float32** |  | [optional] 
**EstimatedNetworkCost** | Pointer to **float32** |  | [optional] 
**EstimatedLicensePrice** | Pointer to **float32** |  | [optional] 
**EstimatedLicenseCost** | Pointer to **float32** |  | [optional] 
**EstimatedExtraPrice** | Pointer to **float32** |  | [optional] 
**EstimatedExtraCost** | Pointer to **float32** |  | [optional] 
**EstimatedTotalPrice** | Pointer to **float32** |  | [optional] 
**EstimatedTotalCost** | Pointer to **float32** |  | [optional] 
**EstimatedRunningPrice** | Pointer to **float32** |  | [optional] 
**EstimatedRunningCost** | Pointer to **float32** |  | [optional] 
**EstimatedCurrency** | Pointer to **string** |  | [optional] 
**EstimatedConversionRate** | Pointer to **float32** |  | [optional] 
**ActualComputePrice** | Pointer to **float32** |  | [optional] 
**ActualComputeCost** | Pointer to **float32** |  | [optional] 
**ActualMemoryPrice** | Pointer to **float32** |  | [optional] 
**ActualMemoryCost** | Pointer to **float32** |  | [optional] 
**ActualStoragePrice** | Pointer to **float32** |  | [optional] 
**ActualStorageCost** | Pointer to **float32** |  | [optional] 
**ActualNetworkPrice** | Pointer to **float32** |  | [optional] 
**ActualNetworkCost** | Pointer to **float32** |  | [optional] 
**ActualLicensePrice** | Pointer to **float32** |  | [optional] 
**ActualLicenseCost** | Pointer to **float32** |  | [optional] 
**ActualExtraPrice** | Pointer to **float32** |  | [optional] 
**ActualExtraCost** | Pointer to **float32** |  | [optional] 
**ActualTotalPrice** | Pointer to **float32** |  | [optional] 
**ActualTotalCost** | Pointer to **float32** |  | [optional] 
**ActualRunningPrice** | Pointer to **float32** |  | [optional] 
**ActualRunningCost** | Pointer to **float32** |  | [optional] 
**ActualCurrency** | Pointer to **string** |  | [optional] 
**ActualConversionRate** | Pointer to **float32** |  | [optional] 
**ComputePrice** | Pointer to **float32** |  | [optional] 
**ComputeCost** | Pointer to **float32** |  | [optional] 
**MemoryPrice** | Pointer to **float32** |  | [optional] 
**MemoryCost** | Pointer to **float32** |  | [optional] 
**StoragePrice** | Pointer to **float32** |  | [optional] 
**StorageCost** | Pointer to **float32** |  | [optional] 
**NetworkPrice** | Pointer to **float32** |  | [optional] 
**NetworkCost** | Pointer to **float32** |  | [optional] 
**LicensePrice** | Pointer to **float32** |  | [optional] 
**LicenseCost** | Pointer to **float32** |  | [optional] 
**ExtraPrice** | Pointer to **float32** |  | [optional] 
**ExtraCost** | Pointer to **float32** |  | [optional] 
**TotalPrice** | Pointer to **float32** |  | [optional] 
**TotalCost** | Pointer to **float32** |  | [optional] 
**RunningPrice** | Pointer to **float32** |  | [optional] 
**RunningCost** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ConversionRate** | Pointer to **float32** |  | [optional] 
**CostType** | Pointer to **string** |  | [optional] 
**OffTime** | Pointer to **int64** |  | [optional] 
**PowerState** | Pointer to **NullableString** |  | [optional] 
**PowerDate** | Pointer to **time.Time** |  | [optional] 
**RunningMultiplier** | Pointer to **float32** |  | [optional] 
**UsageType** | Pointer to **NullableString** |  | [optional] 
**UsageCategory** | Pointer to **NullableString** |  | [optional] 
**LastCostDate** | Pointer to **NullableTime** |  | [optional] 
**LastActualDate** | Pointer to **NullableTime** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LineItemCount** | Pointer to **int64** |  | [optional] 
**LineItems** | Pointer to [**[]GetInvoices200ResponseAllOfInvoiceLineItemsInner**](GetInvoices200ResponseAllOfInvoiceLineItemsInner.md) |  | [optional] 

## Methods

### NewUpdateInvoices200ResponseAllOfUser

`func NewUpdateInvoices200ResponseAllOfUser() *UpdateInvoices200ResponseAllOfUser`

NewUpdateInvoices200ResponseAllOfUser instantiates a new UpdateInvoices200ResponseAllOfUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateInvoices200ResponseAllOfUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateInvoices200ResponseAllOfUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateInvoices200ResponseAllOfUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwnerId

`func (o *UpdateInvoices200ResponseAllOfUser) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *UpdateInvoices200ResponseAllOfUser) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *UpdateInvoices200ResponseAllOfUser) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateInvoices200ResponseAllOfUser) GetAccount() GetInvoices200ResponseAllOfInvoiceAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetAccountOk() (*GetInvoices200ResponseAllOfInvoiceAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateInvoices200ResponseAllOfUser) SetAccount(v GetInvoices200ResponseAllOfInvoiceAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateInvoices200ResponseAllOfUser) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetGroup

`func (o *UpdateInvoices200ResponseAllOfUser) GetGroup() map[string]interface{}`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetGroupOk() (*map[string]interface{}, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UpdateInvoices200ResponseAllOfUser) SetGroup(v map[string]interface{})`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UpdateInvoices200ResponseAllOfUser) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetCloud

`func (o *UpdateInvoices200ResponseAllOfUser) GetCloud() GetInvoices200ResponseAllOfInvoiceCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetCloudOk() (*GetInvoices200ResponseAllOfInvoiceCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *UpdateInvoices200ResponseAllOfUser) SetCloud(v GetInvoices200ResponseAllOfInvoiceCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *UpdateInvoices200ResponseAllOfUser) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetInstance

`func (o *UpdateInvoices200ResponseAllOfUser) GetInstance() map[string]interface{}`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetInstanceOk() (*map[string]interface{}, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *UpdateInvoices200ResponseAllOfUser) SetInstance(v map[string]interface{})`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *UpdateInvoices200ResponseAllOfUser) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### SetInstanceNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil
### GetServer

`func (o *UpdateInvoices200ResponseAllOfUser) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *UpdateInvoices200ResponseAllOfUser) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *UpdateInvoices200ResponseAllOfUser) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetCluster

`func (o *UpdateInvoices200ResponseAllOfUser) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *UpdateInvoices200ResponseAllOfUser) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *UpdateInvoices200ResponseAllOfUser) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetUser

`func (o *UpdateInvoices200ResponseAllOfUser) GetUser() map[string]interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetUserOk() (*map[string]interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdateInvoices200ResponseAllOfUser) SetUser(v map[string]interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *UpdateInvoices200ResponseAllOfUser) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPlan

`func (o *UpdateInvoices200ResponseAllOfUser) GetPlan() map[string]interface{}`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetPlanOk() (*map[string]interface{}, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UpdateInvoices200ResponseAllOfUser) SetPlan(v map[string]interface{})`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UpdateInvoices200ResponseAllOfUser) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### SetPlanNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetPlanNil(b bool)`

 SetPlanNil sets the value for Plan to be an explicit nil

### UnsetPlan
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetPlan()`

UnsetPlan ensures that no value is present for Plan, not even an explicit nil
### GetTags

`func (o *UpdateInvoices200ResponseAllOfUser) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateInvoices200ResponseAllOfUser) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateInvoices200ResponseAllOfUser) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetProject

`func (o *UpdateInvoices200ResponseAllOfUser) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *UpdateInvoices200ResponseAllOfUser) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *UpdateInvoices200ResponseAllOfUser) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetRefType

`func (o *UpdateInvoices200ResponseAllOfUser) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdateInvoices200ResponseAllOfUser) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdateInvoices200ResponseAllOfUser) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *UpdateInvoices200ResponseAllOfUser) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdateInvoices200ResponseAllOfUser) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdateInvoices200ResponseAllOfUser) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefUuid

`func (o *UpdateInvoices200ResponseAllOfUser) GetRefUuid() string`

GetRefUuid returns the RefUuid field if non-nil, zero value otherwise.

### GetRefUuidOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetRefUuidOk() (*string, bool)`

GetRefUuidOk returns a tuple with the RefUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUuid

`func (o *UpdateInvoices200ResponseAllOfUser) SetRefUuid(v string)`

SetRefUuid sets RefUuid field to given value.

### HasRefUuid

`func (o *UpdateInvoices200ResponseAllOfUser) HasRefUuid() bool`

HasRefUuid returns a boolean if a field has been set.

### SetRefUuidNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetRefUuidNil(b bool)`

 SetRefUuidNil sets the value for RefUuid to be an explicit nil

### UnsetRefUuid
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetRefUuid()`

UnsetRefUuid ensures that no value is present for RefUuid, not even an explicit nil
### GetRefName

`func (o *UpdateInvoices200ResponseAllOfUser) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *UpdateInvoices200ResponseAllOfUser) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *UpdateInvoices200ResponseAllOfUser) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetRefCategory

`func (o *UpdateInvoices200ResponseAllOfUser) GetRefCategory() string`

GetRefCategory returns the RefCategory field if non-nil, zero value otherwise.

### GetRefCategoryOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetRefCategoryOk() (*string, bool)`

GetRefCategoryOk returns a tuple with the RefCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCategory

`func (o *UpdateInvoices200ResponseAllOfUser) SetRefCategory(v string)`

SetRefCategory sets RefCategory field to given value.

### HasRefCategory

`func (o *UpdateInvoices200ResponseAllOfUser) HasRefCategory() bool`

HasRefCategory returns a boolean if a field has been set.

### GetResourceId

`func (o *UpdateInvoices200ResponseAllOfUser) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *UpdateInvoices200ResponseAllOfUser) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *UpdateInvoices200ResponseAllOfUser) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceUuid

`func (o *UpdateInvoices200ResponseAllOfUser) GetResourceUuid() string`

GetResourceUuid returns the ResourceUuid field if non-nil, zero value otherwise.

### GetResourceUuidOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetResourceUuidOk() (*string, bool)`

GetResourceUuidOk returns a tuple with the ResourceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUuid

`func (o *UpdateInvoices200ResponseAllOfUser) SetResourceUuid(v string)`

SetResourceUuid sets ResourceUuid field to given value.

### HasResourceUuid

`func (o *UpdateInvoices200ResponseAllOfUser) HasResourceUuid() bool`

HasResourceUuid returns a boolean if a field has been set.

### SetResourceUuidNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetResourceUuidNil(b bool)`

 SetResourceUuidNil sets the value for ResourceUuid to be an explicit nil

### UnsetResourceUuid
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetResourceUuid()`

UnsetResourceUuid ensures that no value is present for ResourceUuid, not even an explicit nil
### GetResourceType

`func (o *UpdateInvoices200ResponseAllOfUser) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *UpdateInvoices200ResponseAllOfUser) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *UpdateInvoices200ResponseAllOfUser) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetResourceName

`func (o *UpdateInvoices200ResponseAllOfUser) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *UpdateInvoices200ResponseAllOfUser) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *UpdateInvoices200ResponseAllOfUser) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetResourceExternalId

`func (o *UpdateInvoices200ResponseAllOfUser) GetResourceExternalId() string`

GetResourceExternalId returns the ResourceExternalId field if non-nil, zero value otherwise.

### GetResourceExternalIdOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetResourceExternalIdOk() (*string, bool)`

GetResourceExternalIdOk returns a tuple with the ResourceExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceExternalId

`func (o *UpdateInvoices200ResponseAllOfUser) SetResourceExternalId(v string)`

SetResourceExternalId sets ResourceExternalId field to given value.

### HasResourceExternalId

`func (o *UpdateInvoices200ResponseAllOfUser) HasResourceExternalId() bool`

HasResourceExternalId returns a boolean if a field has been set.

### SetResourceExternalIdNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetResourceExternalIdNil(b bool)`

 SetResourceExternalIdNil sets the value for ResourceExternalId to be an explicit nil

### UnsetResourceExternalId
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetResourceExternalId()`

UnsetResourceExternalId ensures that no value is present for ResourceExternalId, not even an explicit nil
### GetResourceInternalId

`func (o *UpdateInvoices200ResponseAllOfUser) GetResourceInternalId() string`

GetResourceInternalId returns the ResourceInternalId field if non-nil, zero value otherwise.

### GetResourceInternalIdOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetResourceInternalIdOk() (*string, bool)`

GetResourceInternalIdOk returns a tuple with the ResourceInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInternalId

`func (o *UpdateInvoices200ResponseAllOfUser) SetResourceInternalId(v string)`

SetResourceInternalId sets ResourceInternalId field to given value.

### HasResourceInternalId

`func (o *UpdateInvoices200ResponseAllOfUser) HasResourceInternalId() bool`

HasResourceInternalId returns a boolean if a field has been set.

### SetResourceInternalIdNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetResourceInternalIdNil(b bool)`

 SetResourceInternalIdNil sets the value for ResourceInternalId to be an explicit nil

### UnsetResourceInternalId
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetResourceInternalId()`

UnsetResourceInternalId ensures that no value is present for ResourceInternalId, not even an explicit nil
### GetInterval

`func (o *UpdateInvoices200ResponseAllOfUser) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UpdateInvoices200ResponseAllOfUser) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *UpdateInvoices200ResponseAllOfUser) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetPeriod

`func (o *UpdateInvoices200ResponseAllOfUser) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *UpdateInvoices200ResponseAllOfUser) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *UpdateInvoices200ResponseAllOfUser) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetEstimate

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimate() bool`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimateOk() (*bool, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimate(v bool)`

SetEstimate sets Estimate field to given value.

### HasEstimate

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimate() bool`

HasEstimate returns a boolean if a field has been set.

### GetSummaryInvoice

`func (o *UpdateInvoices200ResponseAllOfUser) GetSummaryInvoice() bool`

GetSummaryInvoice returns the SummaryInvoice field if non-nil, zero value otherwise.

### GetSummaryInvoiceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetSummaryInvoiceOk() (*bool, bool)`

GetSummaryInvoiceOk returns a tuple with the SummaryInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryInvoice

`func (o *UpdateInvoices200ResponseAllOfUser) SetSummaryInvoice(v bool)`

SetSummaryInvoice sets SummaryInvoice field to given value.

### HasSummaryInvoice

`func (o *UpdateInvoices200ResponseAllOfUser) HasSummaryInvoice() bool`

HasSummaryInvoice returns a boolean if a field has been set.

### GetActive

`func (o *UpdateInvoices200ResponseAllOfUser) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateInvoices200ResponseAllOfUser) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateInvoices200ResponseAllOfUser) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateInvoices200ResponseAllOfUser) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateInvoices200ResponseAllOfUser) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateInvoices200ResponseAllOfUser) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UpdateInvoices200ResponseAllOfUser) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateInvoices200ResponseAllOfUser) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateInvoices200ResponseAllOfUser) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetRefStart

`func (o *UpdateInvoices200ResponseAllOfUser) GetRefStart() time.Time`

GetRefStart returns the RefStart field if non-nil, zero value otherwise.

### GetRefStartOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetRefStartOk() (*time.Time, bool)`

GetRefStartOk returns a tuple with the RefStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefStart

`func (o *UpdateInvoices200ResponseAllOfUser) SetRefStart(v time.Time)`

SetRefStart sets RefStart field to given value.

### HasRefStart

`func (o *UpdateInvoices200ResponseAllOfUser) HasRefStart() bool`

HasRefStart returns a boolean if a field has been set.

### GetRefEnd

`func (o *UpdateInvoices200ResponseAllOfUser) GetRefEnd() time.Time`

GetRefEnd returns the RefEnd field if non-nil, zero value otherwise.

### GetRefEndOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetRefEndOk() (*time.Time, bool)`

GetRefEndOk returns a tuple with the RefEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefEnd

`func (o *UpdateInvoices200ResponseAllOfUser) SetRefEnd(v time.Time)`

SetRefEnd sets RefEnd field to given value.

### HasRefEnd

`func (o *UpdateInvoices200ResponseAllOfUser) HasRefEnd() bool`

HasRefEnd returns a boolean if a field has been set.

### GetEstimatedComputePrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedComputePrice() float32`

GetEstimatedComputePrice returns the EstimatedComputePrice field if non-nil, zero value otherwise.

### GetEstimatedComputePriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedComputePriceOk() (*float32, bool)`

GetEstimatedComputePriceOk returns a tuple with the EstimatedComputePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedComputePrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedComputePrice(v float32)`

SetEstimatedComputePrice sets EstimatedComputePrice field to given value.

### HasEstimatedComputePrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedComputePrice() bool`

HasEstimatedComputePrice returns a boolean if a field has been set.

### GetEstimatedComputeCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedComputeCost() float32`

GetEstimatedComputeCost returns the EstimatedComputeCost field if non-nil, zero value otherwise.

### GetEstimatedComputeCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedComputeCostOk() (*float32, bool)`

GetEstimatedComputeCostOk returns a tuple with the EstimatedComputeCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedComputeCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedComputeCost(v float32)`

SetEstimatedComputeCost sets EstimatedComputeCost field to given value.

### HasEstimatedComputeCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedComputeCost() bool`

HasEstimatedComputeCost returns a boolean if a field has been set.

### GetEstimatedMemoryPrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedMemoryPrice() float32`

GetEstimatedMemoryPrice returns the EstimatedMemoryPrice field if non-nil, zero value otherwise.

### GetEstimatedMemoryPriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedMemoryPriceOk() (*float32, bool)`

GetEstimatedMemoryPriceOk returns a tuple with the EstimatedMemoryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMemoryPrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedMemoryPrice(v float32)`

SetEstimatedMemoryPrice sets EstimatedMemoryPrice field to given value.

### HasEstimatedMemoryPrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedMemoryPrice() bool`

HasEstimatedMemoryPrice returns a boolean if a field has been set.

### GetEstimatedMemoryCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedMemoryCost() float32`

GetEstimatedMemoryCost returns the EstimatedMemoryCost field if non-nil, zero value otherwise.

### GetEstimatedMemoryCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedMemoryCostOk() (*float32, bool)`

GetEstimatedMemoryCostOk returns a tuple with the EstimatedMemoryCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMemoryCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedMemoryCost(v float32)`

SetEstimatedMemoryCost sets EstimatedMemoryCost field to given value.

### HasEstimatedMemoryCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedMemoryCost() bool`

HasEstimatedMemoryCost returns a boolean if a field has been set.

### GetEstimatedStoragePrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedStoragePrice() float32`

GetEstimatedStoragePrice returns the EstimatedStoragePrice field if non-nil, zero value otherwise.

### GetEstimatedStoragePriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedStoragePriceOk() (*float32, bool)`

GetEstimatedStoragePriceOk returns a tuple with the EstimatedStoragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStoragePrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedStoragePrice(v float32)`

SetEstimatedStoragePrice sets EstimatedStoragePrice field to given value.

### HasEstimatedStoragePrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedStoragePrice() bool`

HasEstimatedStoragePrice returns a boolean if a field has been set.

### GetEstimatedStorageCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedStorageCost() float32`

GetEstimatedStorageCost returns the EstimatedStorageCost field if non-nil, zero value otherwise.

### GetEstimatedStorageCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedStorageCostOk() (*float32, bool)`

GetEstimatedStorageCostOk returns a tuple with the EstimatedStorageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStorageCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedStorageCost(v float32)`

SetEstimatedStorageCost sets EstimatedStorageCost field to given value.

### HasEstimatedStorageCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedStorageCost() bool`

HasEstimatedStorageCost returns a boolean if a field has been set.

### GetEstimatedNetworkPrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedNetworkPrice() float32`

GetEstimatedNetworkPrice returns the EstimatedNetworkPrice field if non-nil, zero value otherwise.

### GetEstimatedNetworkPriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedNetworkPriceOk() (*float32, bool)`

GetEstimatedNetworkPriceOk returns a tuple with the EstimatedNetworkPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedNetworkPrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedNetworkPrice(v float32)`

SetEstimatedNetworkPrice sets EstimatedNetworkPrice field to given value.

### HasEstimatedNetworkPrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedNetworkPrice() bool`

HasEstimatedNetworkPrice returns a boolean if a field has been set.

### GetEstimatedNetworkCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedNetworkCost() float32`

GetEstimatedNetworkCost returns the EstimatedNetworkCost field if non-nil, zero value otherwise.

### GetEstimatedNetworkCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedNetworkCostOk() (*float32, bool)`

GetEstimatedNetworkCostOk returns a tuple with the EstimatedNetworkCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedNetworkCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedNetworkCost(v float32)`

SetEstimatedNetworkCost sets EstimatedNetworkCost field to given value.

### HasEstimatedNetworkCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedNetworkCost() bool`

HasEstimatedNetworkCost returns a boolean if a field has been set.

### GetEstimatedLicensePrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedLicensePrice() float32`

GetEstimatedLicensePrice returns the EstimatedLicensePrice field if non-nil, zero value otherwise.

### GetEstimatedLicensePriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedLicensePriceOk() (*float32, bool)`

GetEstimatedLicensePriceOk returns a tuple with the EstimatedLicensePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedLicensePrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedLicensePrice(v float32)`

SetEstimatedLicensePrice sets EstimatedLicensePrice field to given value.

### HasEstimatedLicensePrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedLicensePrice() bool`

HasEstimatedLicensePrice returns a boolean if a field has been set.

### GetEstimatedLicenseCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedLicenseCost() float32`

GetEstimatedLicenseCost returns the EstimatedLicenseCost field if non-nil, zero value otherwise.

### GetEstimatedLicenseCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedLicenseCostOk() (*float32, bool)`

GetEstimatedLicenseCostOk returns a tuple with the EstimatedLicenseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedLicenseCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedLicenseCost(v float32)`

SetEstimatedLicenseCost sets EstimatedLicenseCost field to given value.

### HasEstimatedLicenseCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedLicenseCost() bool`

HasEstimatedLicenseCost returns a boolean if a field has been set.

### GetEstimatedExtraPrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedExtraPrice() float32`

GetEstimatedExtraPrice returns the EstimatedExtraPrice field if non-nil, zero value otherwise.

### GetEstimatedExtraPriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedExtraPriceOk() (*float32, bool)`

GetEstimatedExtraPriceOk returns a tuple with the EstimatedExtraPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedExtraPrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedExtraPrice(v float32)`

SetEstimatedExtraPrice sets EstimatedExtraPrice field to given value.

### HasEstimatedExtraPrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedExtraPrice() bool`

HasEstimatedExtraPrice returns a boolean if a field has been set.

### GetEstimatedExtraCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedExtraCost() float32`

GetEstimatedExtraCost returns the EstimatedExtraCost field if non-nil, zero value otherwise.

### GetEstimatedExtraCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedExtraCostOk() (*float32, bool)`

GetEstimatedExtraCostOk returns a tuple with the EstimatedExtraCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedExtraCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedExtraCost(v float32)`

SetEstimatedExtraCost sets EstimatedExtraCost field to given value.

### HasEstimatedExtraCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedExtraCost() bool`

HasEstimatedExtraCost returns a boolean if a field has been set.

### GetEstimatedTotalPrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedTotalPrice() float32`

GetEstimatedTotalPrice returns the EstimatedTotalPrice field if non-nil, zero value otherwise.

### GetEstimatedTotalPriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedTotalPriceOk() (*float32, bool)`

GetEstimatedTotalPriceOk returns a tuple with the EstimatedTotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTotalPrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedTotalPrice(v float32)`

SetEstimatedTotalPrice sets EstimatedTotalPrice field to given value.

### HasEstimatedTotalPrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedTotalPrice() bool`

HasEstimatedTotalPrice returns a boolean if a field has been set.

### GetEstimatedTotalCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedTotalCost() float32`

GetEstimatedTotalCost returns the EstimatedTotalCost field if non-nil, zero value otherwise.

### GetEstimatedTotalCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedTotalCostOk() (*float32, bool)`

GetEstimatedTotalCostOk returns a tuple with the EstimatedTotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTotalCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedTotalCost(v float32)`

SetEstimatedTotalCost sets EstimatedTotalCost field to given value.

### HasEstimatedTotalCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedTotalCost() bool`

HasEstimatedTotalCost returns a boolean if a field has been set.

### GetEstimatedRunningPrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedRunningPrice() float32`

GetEstimatedRunningPrice returns the EstimatedRunningPrice field if non-nil, zero value otherwise.

### GetEstimatedRunningPriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedRunningPriceOk() (*float32, bool)`

GetEstimatedRunningPriceOk returns a tuple with the EstimatedRunningPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRunningPrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedRunningPrice(v float32)`

SetEstimatedRunningPrice sets EstimatedRunningPrice field to given value.

### HasEstimatedRunningPrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedRunningPrice() bool`

HasEstimatedRunningPrice returns a boolean if a field has been set.

### GetEstimatedRunningCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedRunningCost() float32`

GetEstimatedRunningCost returns the EstimatedRunningCost field if non-nil, zero value otherwise.

### GetEstimatedRunningCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedRunningCostOk() (*float32, bool)`

GetEstimatedRunningCostOk returns a tuple with the EstimatedRunningCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRunningCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedRunningCost(v float32)`

SetEstimatedRunningCost sets EstimatedRunningCost field to given value.

### HasEstimatedRunningCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedRunningCost() bool`

HasEstimatedRunningCost returns a boolean if a field has been set.

### GetEstimatedCurrency

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedCurrency() string`

GetEstimatedCurrency returns the EstimatedCurrency field if non-nil, zero value otherwise.

### GetEstimatedCurrencyOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedCurrencyOk() (*string, bool)`

GetEstimatedCurrencyOk returns a tuple with the EstimatedCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCurrency

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedCurrency(v string)`

SetEstimatedCurrency sets EstimatedCurrency field to given value.

### HasEstimatedCurrency

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedCurrency() bool`

HasEstimatedCurrency returns a boolean if a field has been set.

### GetEstimatedConversionRate

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedConversionRate() float32`

GetEstimatedConversionRate returns the EstimatedConversionRate field if non-nil, zero value otherwise.

### GetEstimatedConversionRateOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetEstimatedConversionRateOk() (*float32, bool)`

GetEstimatedConversionRateOk returns a tuple with the EstimatedConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedConversionRate

`func (o *UpdateInvoices200ResponseAllOfUser) SetEstimatedConversionRate(v float32)`

SetEstimatedConversionRate sets EstimatedConversionRate field to given value.

### HasEstimatedConversionRate

`func (o *UpdateInvoices200ResponseAllOfUser) HasEstimatedConversionRate() bool`

HasEstimatedConversionRate returns a boolean if a field has been set.

### GetActualComputePrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualComputePrice() float32`

GetActualComputePrice returns the ActualComputePrice field if non-nil, zero value otherwise.

### GetActualComputePriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualComputePriceOk() (*float32, bool)`

GetActualComputePriceOk returns a tuple with the ActualComputePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualComputePrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualComputePrice(v float32)`

SetActualComputePrice sets ActualComputePrice field to given value.

### HasActualComputePrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualComputePrice() bool`

HasActualComputePrice returns a boolean if a field has been set.

### GetActualComputeCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualComputeCost() float32`

GetActualComputeCost returns the ActualComputeCost field if non-nil, zero value otherwise.

### GetActualComputeCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualComputeCostOk() (*float32, bool)`

GetActualComputeCostOk returns a tuple with the ActualComputeCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualComputeCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualComputeCost(v float32)`

SetActualComputeCost sets ActualComputeCost field to given value.

### HasActualComputeCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualComputeCost() bool`

HasActualComputeCost returns a boolean if a field has been set.

### GetActualMemoryPrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualMemoryPrice() float32`

GetActualMemoryPrice returns the ActualMemoryPrice field if non-nil, zero value otherwise.

### GetActualMemoryPriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualMemoryPriceOk() (*float32, bool)`

GetActualMemoryPriceOk returns a tuple with the ActualMemoryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualMemoryPrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualMemoryPrice(v float32)`

SetActualMemoryPrice sets ActualMemoryPrice field to given value.

### HasActualMemoryPrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualMemoryPrice() bool`

HasActualMemoryPrice returns a boolean if a field has been set.

### GetActualMemoryCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualMemoryCost() float32`

GetActualMemoryCost returns the ActualMemoryCost field if non-nil, zero value otherwise.

### GetActualMemoryCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualMemoryCostOk() (*float32, bool)`

GetActualMemoryCostOk returns a tuple with the ActualMemoryCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualMemoryCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualMemoryCost(v float32)`

SetActualMemoryCost sets ActualMemoryCost field to given value.

### HasActualMemoryCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualMemoryCost() bool`

HasActualMemoryCost returns a boolean if a field has been set.

### GetActualStoragePrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualStoragePrice() float32`

GetActualStoragePrice returns the ActualStoragePrice field if non-nil, zero value otherwise.

### GetActualStoragePriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualStoragePriceOk() (*float32, bool)`

GetActualStoragePriceOk returns a tuple with the ActualStoragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualStoragePrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualStoragePrice(v float32)`

SetActualStoragePrice sets ActualStoragePrice field to given value.

### HasActualStoragePrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualStoragePrice() bool`

HasActualStoragePrice returns a boolean if a field has been set.

### GetActualStorageCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualStorageCost() float32`

GetActualStorageCost returns the ActualStorageCost field if non-nil, zero value otherwise.

### GetActualStorageCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualStorageCostOk() (*float32, bool)`

GetActualStorageCostOk returns a tuple with the ActualStorageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualStorageCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualStorageCost(v float32)`

SetActualStorageCost sets ActualStorageCost field to given value.

### HasActualStorageCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualStorageCost() bool`

HasActualStorageCost returns a boolean if a field has been set.

### GetActualNetworkPrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualNetworkPrice() float32`

GetActualNetworkPrice returns the ActualNetworkPrice field if non-nil, zero value otherwise.

### GetActualNetworkPriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualNetworkPriceOk() (*float32, bool)`

GetActualNetworkPriceOk returns a tuple with the ActualNetworkPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualNetworkPrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualNetworkPrice(v float32)`

SetActualNetworkPrice sets ActualNetworkPrice field to given value.

### HasActualNetworkPrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualNetworkPrice() bool`

HasActualNetworkPrice returns a boolean if a field has been set.

### GetActualNetworkCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualNetworkCost() float32`

GetActualNetworkCost returns the ActualNetworkCost field if non-nil, zero value otherwise.

### GetActualNetworkCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualNetworkCostOk() (*float32, bool)`

GetActualNetworkCostOk returns a tuple with the ActualNetworkCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualNetworkCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualNetworkCost(v float32)`

SetActualNetworkCost sets ActualNetworkCost field to given value.

### HasActualNetworkCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualNetworkCost() bool`

HasActualNetworkCost returns a boolean if a field has been set.

### GetActualLicensePrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualLicensePrice() float32`

GetActualLicensePrice returns the ActualLicensePrice field if non-nil, zero value otherwise.

### GetActualLicensePriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualLicensePriceOk() (*float32, bool)`

GetActualLicensePriceOk returns a tuple with the ActualLicensePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualLicensePrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualLicensePrice(v float32)`

SetActualLicensePrice sets ActualLicensePrice field to given value.

### HasActualLicensePrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualLicensePrice() bool`

HasActualLicensePrice returns a boolean if a field has been set.

### GetActualLicenseCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualLicenseCost() float32`

GetActualLicenseCost returns the ActualLicenseCost field if non-nil, zero value otherwise.

### GetActualLicenseCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualLicenseCostOk() (*float32, bool)`

GetActualLicenseCostOk returns a tuple with the ActualLicenseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualLicenseCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualLicenseCost(v float32)`

SetActualLicenseCost sets ActualLicenseCost field to given value.

### HasActualLicenseCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualLicenseCost() bool`

HasActualLicenseCost returns a boolean if a field has been set.

### GetActualExtraPrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualExtraPrice() float32`

GetActualExtraPrice returns the ActualExtraPrice field if non-nil, zero value otherwise.

### GetActualExtraPriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualExtraPriceOk() (*float32, bool)`

GetActualExtraPriceOk returns a tuple with the ActualExtraPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualExtraPrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualExtraPrice(v float32)`

SetActualExtraPrice sets ActualExtraPrice field to given value.

### HasActualExtraPrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualExtraPrice() bool`

HasActualExtraPrice returns a boolean if a field has been set.

### GetActualExtraCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualExtraCost() float32`

GetActualExtraCost returns the ActualExtraCost field if non-nil, zero value otherwise.

### GetActualExtraCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualExtraCostOk() (*float32, bool)`

GetActualExtraCostOk returns a tuple with the ActualExtraCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualExtraCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualExtraCost(v float32)`

SetActualExtraCost sets ActualExtraCost field to given value.

### HasActualExtraCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualExtraCost() bool`

HasActualExtraCost returns a boolean if a field has been set.

### GetActualTotalPrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualTotalPrice() float32`

GetActualTotalPrice returns the ActualTotalPrice field if non-nil, zero value otherwise.

### GetActualTotalPriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualTotalPriceOk() (*float32, bool)`

GetActualTotalPriceOk returns a tuple with the ActualTotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualTotalPrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualTotalPrice(v float32)`

SetActualTotalPrice sets ActualTotalPrice field to given value.

### HasActualTotalPrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualTotalPrice() bool`

HasActualTotalPrice returns a boolean if a field has been set.

### GetActualTotalCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualTotalCost() float32`

GetActualTotalCost returns the ActualTotalCost field if non-nil, zero value otherwise.

### GetActualTotalCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualTotalCostOk() (*float32, bool)`

GetActualTotalCostOk returns a tuple with the ActualTotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualTotalCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualTotalCost(v float32)`

SetActualTotalCost sets ActualTotalCost field to given value.

### HasActualTotalCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualTotalCost() bool`

HasActualTotalCost returns a boolean if a field has been set.

### GetActualRunningPrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualRunningPrice() float32`

GetActualRunningPrice returns the ActualRunningPrice field if non-nil, zero value otherwise.

### GetActualRunningPriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualRunningPriceOk() (*float32, bool)`

GetActualRunningPriceOk returns a tuple with the ActualRunningPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualRunningPrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualRunningPrice(v float32)`

SetActualRunningPrice sets ActualRunningPrice field to given value.

### HasActualRunningPrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualRunningPrice() bool`

HasActualRunningPrice returns a boolean if a field has been set.

### GetActualRunningCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualRunningCost() float32`

GetActualRunningCost returns the ActualRunningCost field if non-nil, zero value otherwise.

### GetActualRunningCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualRunningCostOk() (*float32, bool)`

GetActualRunningCostOk returns a tuple with the ActualRunningCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualRunningCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualRunningCost(v float32)`

SetActualRunningCost sets ActualRunningCost field to given value.

### HasActualRunningCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualRunningCost() bool`

HasActualRunningCost returns a boolean if a field has been set.

### GetActualCurrency

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualCurrency() string`

GetActualCurrency returns the ActualCurrency field if non-nil, zero value otherwise.

### GetActualCurrencyOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualCurrencyOk() (*string, bool)`

GetActualCurrencyOk returns a tuple with the ActualCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualCurrency

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualCurrency(v string)`

SetActualCurrency sets ActualCurrency field to given value.

### HasActualCurrency

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualCurrency() bool`

HasActualCurrency returns a boolean if a field has been set.

### GetActualConversionRate

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualConversionRate() float32`

GetActualConversionRate returns the ActualConversionRate field if non-nil, zero value otherwise.

### GetActualConversionRateOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetActualConversionRateOk() (*float32, bool)`

GetActualConversionRateOk returns a tuple with the ActualConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualConversionRate

`func (o *UpdateInvoices200ResponseAllOfUser) SetActualConversionRate(v float32)`

SetActualConversionRate sets ActualConversionRate field to given value.

### HasActualConversionRate

`func (o *UpdateInvoices200ResponseAllOfUser) HasActualConversionRate() bool`

HasActualConversionRate returns a boolean if a field has been set.

### GetComputePrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetComputePrice() float32`

GetComputePrice returns the ComputePrice field if non-nil, zero value otherwise.

### GetComputePriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetComputePriceOk() (*float32, bool)`

GetComputePriceOk returns a tuple with the ComputePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetComputePrice(v float32)`

SetComputePrice sets ComputePrice field to given value.

### HasComputePrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasComputePrice() bool`

HasComputePrice returns a boolean if a field has been set.

### GetComputeCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetComputeCost() float32`

GetComputeCost returns the ComputeCost field if non-nil, zero value otherwise.

### GetComputeCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetComputeCostOk() (*float32, bool)`

GetComputeCostOk returns a tuple with the ComputeCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetComputeCost(v float32)`

SetComputeCost sets ComputeCost field to given value.

### HasComputeCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasComputeCost() bool`

HasComputeCost returns a boolean if a field has been set.

### GetMemoryPrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetMemoryPrice() float32`

GetMemoryPrice returns the MemoryPrice field if non-nil, zero value otherwise.

### GetMemoryPriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetMemoryPriceOk() (*float32, bool)`

GetMemoryPriceOk returns a tuple with the MemoryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetMemoryPrice(v float32)`

SetMemoryPrice sets MemoryPrice field to given value.

### HasMemoryPrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasMemoryPrice() bool`

HasMemoryPrice returns a boolean if a field has been set.

### GetMemoryCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetMemoryCost() float32`

GetMemoryCost returns the MemoryCost field if non-nil, zero value otherwise.

### GetMemoryCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetMemoryCostOk() (*float32, bool)`

GetMemoryCostOk returns a tuple with the MemoryCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetMemoryCost(v float32)`

SetMemoryCost sets MemoryCost field to given value.

### HasMemoryCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasMemoryCost() bool`

HasMemoryCost returns a boolean if a field has been set.

### GetStoragePrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetStoragePrice() float32`

GetStoragePrice returns the StoragePrice field if non-nil, zero value otherwise.

### GetStoragePriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetStoragePriceOk() (*float32, bool)`

GetStoragePriceOk returns a tuple with the StoragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetStoragePrice(v float32)`

SetStoragePrice sets StoragePrice field to given value.

### HasStoragePrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasStoragePrice() bool`

HasStoragePrice returns a boolean if a field has been set.

### GetStorageCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetStorageCost() float32`

GetStorageCost returns the StorageCost field if non-nil, zero value otherwise.

### GetStorageCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetStorageCostOk() (*float32, bool)`

GetStorageCostOk returns a tuple with the StorageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetStorageCost(v float32)`

SetStorageCost sets StorageCost field to given value.

### HasStorageCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasStorageCost() bool`

HasStorageCost returns a boolean if a field has been set.

### GetNetworkPrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetNetworkPrice() float32`

GetNetworkPrice returns the NetworkPrice field if non-nil, zero value otherwise.

### GetNetworkPriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetNetworkPriceOk() (*float32, bool)`

GetNetworkPriceOk returns a tuple with the NetworkPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetNetworkPrice(v float32)`

SetNetworkPrice sets NetworkPrice field to given value.

### HasNetworkPrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasNetworkPrice() bool`

HasNetworkPrice returns a boolean if a field has been set.

### GetNetworkCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetNetworkCost() float32`

GetNetworkCost returns the NetworkCost field if non-nil, zero value otherwise.

### GetNetworkCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetNetworkCostOk() (*float32, bool)`

GetNetworkCostOk returns a tuple with the NetworkCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetNetworkCost(v float32)`

SetNetworkCost sets NetworkCost field to given value.

### HasNetworkCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasNetworkCost() bool`

HasNetworkCost returns a boolean if a field has been set.

### GetLicensePrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetLicensePrice() float32`

GetLicensePrice returns the LicensePrice field if non-nil, zero value otherwise.

### GetLicensePriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetLicensePriceOk() (*float32, bool)`

GetLicensePriceOk returns a tuple with the LicensePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetLicensePrice(v float32)`

SetLicensePrice sets LicensePrice field to given value.

### HasLicensePrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasLicensePrice() bool`

HasLicensePrice returns a boolean if a field has been set.

### GetLicenseCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetLicenseCost() float32`

GetLicenseCost returns the LicenseCost field if non-nil, zero value otherwise.

### GetLicenseCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetLicenseCostOk() (*float32, bool)`

GetLicenseCostOk returns a tuple with the LicenseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetLicenseCost(v float32)`

SetLicenseCost sets LicenseCost field to given value.

### HasLicenseCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasLicenseCost() bool`

HasLicenseCost returns a boolean if a field has been set.

### GetExtraPrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetExtraPrice() float32`

GetExtraPrice returns the ExtraPrice field if non-nil, zero value otherwise.

### GetExtraPriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetExtraPriceOk() (*float32, bool)`

GetExtraPriceOk returns a tuple with the ExtraPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraPrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetExtraPrice(v float32)`

SetExtraPrice sets ExtraPrice field to given value.

### HasExtraPrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasExtraPrice() bool`

HasExtraPrice returns a boolean if a field has been set.

### GetExtraCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetExtraCost() float32`

GetExtraCost returns the ExtraCost field if non-nil, zero value otherwise.

### GetExtraCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetExtraCostOk() (*float32, bool)`

GetExtraCostOk returns a tuple with the ExtraCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetExtraCost(v float32)`

SetExtraCost sets ExtraCost field to given value.

### HasExtraCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasExtraCost() bool`

HasExtraCost returns a boolean if a field has been set.

### GetTotalPrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetTotalCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetRunningPrice

`func (o *UpdateInvoices200ResponseAllOfUser) GetRunningPrice() float32`

GetRunningPrice returns the RunningPrice field if non-nil, zero value otherwise.

### GetRunningPriceOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetRunningPriceOk() (*float32, bool)`

GetRunningPriceOk returns a tuple with the RunningPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningPrice

`func (o *UpdateInvoices200ResponseAllOfUser) SetRunningPrice(v float32)`

SetRunningPrice sets RunningPrice field to given value.

### HasRunningPrice

`func (o *UpdateInvoices200ResponseAllOfUser) HasRunningPrice() bool`

HasRunningPrice returns a boolean if a field has been set.

### GetRunningCost

`func (o *UpdateInvoices200ResponseAllOfUser) GetRunningCost() float32`

GetRunningCost returns the RunningCost field if non-nil, zero value otherwise.

### GetRunningCostOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetRunningCostOk() (*float32, bool)`

GetRunningCostOk returns a tuple with the RunningCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningCost

`func (o *UpdateInvoices200ResponseAllOfUser) SetRunningCost(v float32)`

SetRunningCost sets RunningCost field to given value.

### HasRunningCost

`func (o *UpdateInvoices200ResponseAllOfUser) HasRunningCost() bool`

HasRunningCost returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdateInvoices200ResponseAllOfUser) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateInvoices200ResponseAllOfUser) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdateInvoices200ResponseAllOfUser) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetConversionRate

`func (o *UpdateInvoices200ResponseAllOfUser) GetConversionRate() float32`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetConversionRateOk() (*float32, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *UpdateInvoices200ResponseAllOfUser) SetConversionRate(v float32)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *UpdateInvoices200ResponseAllOfUser) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetCostType

`func (o *UpdateInvoices200ResponseAllOfUser) GetCostType() string`

GetCostType returns the CostType field if non-nil, zero value otherwise.

### GetCostTypeOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetCostTypeOk() (*string, bool)`

GetCostTypeOk returns a tuple with the CostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostType

`func (o *UpdateInvoices200ResponseAllOfUser) SetCostType(v string)`

SetCostType sets CostType field to given value.

### HasCostType

`func (o *UpdateInvoices200ResponseAllOfUser) HasCostType() bool`

HasCostType returns a boolean if a field has been set.

### GetOffTime

`func (o *UpdateInvoices200ResponseAllOfUser) GetOffTime() int64`

GetOffTime returns the OffTime field if non-nil, zero value otherwise.

### GetOffTimeOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetOffTimeOk() (*int64, bool)`

GetOffTimeOk returns a tuple with the OffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffTime

`func (o *UpdateInvoices200ResponseAllOfUser) SetOffTime(v int64)`

SetOffTime sets OffTime field to given value.

### HasOffTime

`func (o *UpdateInvoices200ResponseAllOfUser) HasOffTime() bool`

HasOffTime returns a boolean if a field has been set.

### GetPowerState

`func (o *UpdateInvoices200ResponseAllOfUser) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *UpdateInvoices200ResponseAllOfUser) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *UpdateInvoices200ResponseAllOfUser) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### SetPowerStateNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetPowerStateNil(b bool)`

 SetPowerStateNil sets the value for PowerState to be an explicit nil

### UnsetPowerState
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetPowerState()`

UnsetPowerState ensures that no value is present for PowerState, not even an explicit nil
### GetPowerDate

`func (o *UpdateInvoices200ResponseAllOfUser) GetPowerDate() time.Time`

GetPowerDate returns the PowerDate field if non-nil, zero value otherwise.

### GetPowerDateOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetPowerDateOk() (*time.Time, bool)`

GetPowerDateOk returns a tuple with the PowerDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerDate

`func (o *UpdateInvoices200ResponseAllOfUser) SetPowerDate(v time.Time)`

SetPowerDate sets PowerDate field to given value.

### HasPowerDate

`func (o *UpdateInvoices200ResponseAllOfUser) HasPowerDate() bool`

HasPowerDate returns a boolean if a field has been set.

### GetRunningMultiplier

`func (o *UpdateInvoices200ResponseAllOfUser) GetRunningMultiplier() float32`

GetRunningMultiplier returns the RunningMultiplier field if non-nil, zero value otherwise.

### GetRunningMultiplierOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetRunningMultiplierOk() (*float32, bool)`

GetRunningMultiplierOk returns a tuple with the RunningMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningMultiplier

`func (o *UpdateInvoices200ResponseAllOfUser) SetRunningMultiplier(v float32)`

SetRunningMultiplier sets RunningMultiplier field to given value.

### HasRunningMultiplier

`func (o *UpdateInvoices200ResponseAllOfUser) HasRunningMultiplier() bool`

HasRunningMultiplier returns a boolean if a field has been set.

### GetUsageType

`func (o *UpdateInvoices200ResponseAllOfUser) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *UpdateInvoices200ResponseAllOfUser) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *UpdateInvoices200ResponseAllOfUser) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### SetUsageTypeNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetUsageTypeNil(b bool)`

 SetUsageTypeNil sets the value for UsageType to be an explicit nil

### UnsetUsageType
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetUsageType()`

UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil
### GetUsageCategory

`func (o *UpdateInvoices200ResponseAllOfUser) GetUsageCategory() string`

GetUsageCategory returns the UsageCategory field if non-nil, zero value otherwise.

### GetUsageCategoryOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetUsageCategoryOk() (*string, bool)`

GetUsageCategoryOk returns a tuple with the UsageCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCategory

`func (o *UpdateInvoices200ResponseAllOfUser) SetUsageCategory(v string)`

SetUsageCategory sets UsageCategory field to given value.

### HasUsageCategory

`func (o *UpdateInvoices200ResponseAllOfUser) HasUsageCategory() bool`

HasUsageCategory returns a boolean if a field has been set.

### SetUsageCategoryNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetUsageCategoryNil(b bool)`

 SetUsageCategoryNil sets the value for UsageCategory to be an explicit nil

### UnsetUsageCategory
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetUsageCategory()`

UnsetUsageCategory ensures that no value is present for UsageCategory, not even an explicit nil
### GetLastCostDate

`func (o *UpdateInvoices200ResponseAllOfUser) GetLastCostDate() time.Time`

GetLastCostDate returns the LastCostDate field if non-nil, zero value otherwise.

### GetLastCostDateOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetLastCostDateOk() (*time.Time, bool)`

GetLastCostDateOk returns a tuple with the LastCostDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCostDate

`func (o *UpdateInvoices200ResponseAllOfUser) SetLastCostDate(v time.Time)`

SetLastCostDate sets LastCostDate field to given value.

### HasLastCostDate

`func (o *UpdateInvoices200ResponseAllOfUser) HasLastCostDate() bool`

HasLastCostDate returns a boolean if a field has been set.

### SetLastCostDateNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetLastCostDateNil(b bool)`

 SetLastCostDateNil sets the value for LastCostDate to be an explicit nil

### UnsetLastCostDate
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetLastCostDate()`

UnsetLastCostDate ensures that no value is present for LastCostDate, not even an explicit nil
### GetLastActualDate

`func (o *UpdateInvoices200ResponseAllOfUser) GetLastActualDate() time.Time`

GetLastActualDate returns the LastActualDate field if non-nil, zero value otherwise.

### GetLastActualDateOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetLastActualDateOk() (*time.Time, bool)`

GetLastActualDateOk returns a tuple with the LastActualDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActualDate

`func (o *UpdateInvoices200ResponseAllOfUser) SetLastActualDate(v time.Time)`

SetLastActualDate sets LastActualDate field to given value.

### HasLastActualDate

`func (o *UpdateInvoices200ResponseAllOfUser) HasLastActualDate() bool`

HasLastActualDate returns a boolean if a field has been set.

### SetLastActualDateNil

`func (o *UpdateInvoices200ResponseAllOfUser) SetLastActualDateNil(b bool)`

 SetLastActualDateNil sets the value for LastActualDate to be an explicit nil

### UnsetLastActualDate
`func (o *UpdateInvoices200ResponseAllOfUser) UnsetLastActualDate()`

UnsetLastActualDate ensures that no value is present for LastActualDate, not even an explicit nil
### GetDateCreated

`func (o *UpdateInvoices200ResponseAllOfUser) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateInvoices200ResponseAllOfUser) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateInvoices200ResponseAllOfUser) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateInvoices200ResponseAllOfUser) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateInvoices200ResponseAllOfUser) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateInvoices200ResponseAllOfUser) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLineItemCount

`func (o *UpdateInvoices200ResponseAllOfUser) GetLineItemCount() int64`

GetLineItemCount returns the LineItemCount field if non-nil, zero value otherwise.

### GetLineItemCountOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetLineItemCountOk() (*int64, bool)`

GetLineItemCountOk returns a tuple with the LineItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemCount

`func (o *UpdateInvoices200ResponseAllOfUser) SetLineItemCount(v int64)`

SetLineItemCount sets LineItemCount field to given value.

### HasLineItemCount

`func (o *UpdateInvoices200ResponseAllOfUser) HasLineItemCount() bool`

HasLineItemCount returns a boolean if a field has been set.

### GetLineItems

`func (o *UpdateInvoices200ResponseAllOfUser) GetLineItems() []GetInvoices200ResponseAllOfInvoiceLineItemsInner`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *UpdateInvoices200ResponseAllOfUser) GetLineItemsOk() (*[]GetInvoices200ResponseAllOfInvoiceLineItemsInner, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *UpdateInvoices200ResponseAllOfUser) SetLineItems(v []GetInvoices200ResponseAllOfInvoiceLineItemsInner)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *UpdateInvoices200ResponseAllOfUser) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


