# ZoneType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Provision** | Pointer to **bool** |  | [optional] 
**AutoCapacity** | Pointer to **bool** |  | [optional] 
**MigrationTarget** | Pointer to **bool** |  | [optional] 
**HasDatastores** | Pointer to **bool** |  | [optional] 
**HasNetworks** | Pointer to **bool** |  | [optional] 
**HasResourcePools** | Pointer to **bool** |  | [optional] 
**HasSecurityGroups** | Pointer to **bool** |  | [optional] 
**HasContainers** | Pointer to **bool** |  | [optional] 
**HasBareMetal** | Pointer to **bool** |  | [optional] 
**HasServices** | Pointer to **bool** |  | [optional] 
**HasFunctions** | Pointer to **bool** |  | [optional] 
**HasJobs** | Pointer to **bool** |  | [optional] 
**HasDiscovery** | Pointer to **bool** |  | [optional] 
**HasCloudInit** | Pointer to **bool** |  | [optional] 
**HasFolders** | Pointer to **bool** |  | [optional] 
**HasMarketplace** | Pointer to **bool** |  | [optional] 
**HasNativePlans** | Pointer to **bool** |  | [optional] 
**CanCreateResourcePools** | Pointer to **bool** |  | [optional] 
**CanDeleteResourcePools** | Pointer to **bool** |  | [optional] 
**CanCreateDatastores** | Pointer to **bool** |  | [optional] 
**CanCreateNetworks** | Pointer to **bool** |  | [optional] 
**CanChooseContainerMode** | Pointer to **bool** |  | [optional] 
**ProvisionRequiresResourcePool** | Pointer to **bool** |  | [optional] 
**SupportsDistributedWorker** | Pointer to **bool** |  | [optional] 
**Cloud** | Pointer to **string** |  | [optional] 
**ProvisionTypes** | Pointer to **[]int64** |  | [optional] 
**ZoneInstanceTypeLayoutId** | Pointer to **int64** |  | [optional] 
**ServerTypes** | Pointer to [**[]ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner**](ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner.md) |  | [optional] 
**OptionTypes** | Pointer to [**[]ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInner**](ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInner.md) |  | [optional] 

## Methods

### NewZoneType

`func NewZoneType() *ZoneType`

NewZoneType instantiates a new ZoneType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneTypeWithDefaults

`func NewZoneTypeWithDefaults() *ZoneType`

NewZoneTypeWithDefaults instantiates a new ZoneType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ZoneType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ZoneType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ZoneType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ZoneType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ZoneType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ZoneType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEnabled

`func (o *ZoneType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ZoneType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ZoneType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ZoneType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProvision

`func (o *ZoneType) GetProvision() bool`

GetProvision returns the Provision field if non-nil, zero value otherwise.

### GetProvisionOk

`func (o *ZoneType) GetProvisionOk() (*bool, bool)`

GetProvisionOk returns a tuple with the Provision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvision

`func (o *ZoneType) SetProvision(v bool)`

SetProvision sets Provision field to given value.

### HasProvision

`func (o *ZoneType) HasProvision() bool`

HasProvision returns a boolean if a field has been set.

### GetAutoCapacity

`func (o *ZoneType) GetAutoCapacity() bool`

GetAutoCapacity returns the AutoCapacity field if non-nil, zero value otherwise.

### GetAutoCapacityOk

`func (o *ZoneType) GetAutoCapacityOk() (*bool, bool)`

GetAutoCapacityOk returns a tuple with the AutoCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCapacity

`func (o *ZoneType) SetAutoCapacity(v bool)`

SetAutoCapacity sets AutoCapacity field to given value.

### HasAutoCapacity

`func (o *ZoneType) HasAutoCapacity() bool`

HasAutoCapacity returns a boolean if a field has been set.

### GetMigrationTarget

`func (o *ZoneType) GetMigrationTarget() bool`

GetMigrationTarget returns the MigrationTarget field if non-nil, zero value otherwise.

### GetMigrationTargetOk

`func (o *ZoneType) GetMigrationTargetOk() (*bool, bool)`

GetMigrationTargetOk returns a tuple with the MigrationTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationTarget

`func (o *ZoneType) SetMigrationTarget(v bool)`

SetMigrationTarget sets MigrationTarget field to given value.

### HasMigrationTarget

`func (o *ZoneType) HasMigrationTarget() bool`

HasMigrationTarget returns a boolean if a field has been set.

### GetHasDatastores

`func (o *ZoneType) GetHasDatastores() bool`

GetHasDatastores returns the HasDatastores field if non-nil, zero value otherwise.

### GetHasDatastoresOk

`func (o *ZoneType) GetHasDatastoresOk() (*bool, bool)`

GetHasDatastoresOk returns a tuple with the HasDatastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastores

`func (o *ZoneType) SetHasDatastores(v bool)`

SetHasDatastores sets HasDatastores field to given value.

### HasHasDatastores

`func (o *ZoneType) HasHasDatastores() bool`

HasHasDatastores returns a boolean if a field has been set.

### GetHasNetworks

`func (o *ZoneType) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *ZoneType) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *ZoneType) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *ZoneType) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### GetHasResourcePools

`func (o *ZoneType) GetHasResourcePools() bool`

GetHasResourcePools returns the HasResourcePools field if non-nil, zero value otherwise.

### GetHasResourcePoolsOk

`func (o *ZoneType) GetHasResourcePoolsOk() (*bool, bool)`

GetHasResourcePoolsOk returns a tuple with the HasResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasResourcePools

`func (o *ZoneType) SetHasResourcePools(v bool)`

SetHasResourcePools sets HasResourcePools field to given value.

### HasHasResourcePools

`func (o *ZoneType) HasHasResourcePools() bool`

HasHasResourcePools returns a boolean if a field has been set.

### GetHasSecurityGroups

`func (o *ZoneType) GetHasSecurityGroups() bool`

GetHasSecurityGroups returns the HasSecurityGroups field if non-nil, zero value otherwise.

### GetHasSecurityGroupsOk

`func (o *ZoneType) GetHasSecurityGroupsOk() (*bool, bool)`

GetHasSecurityGroupsOk returns a tuple with the HasSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecurityGroups

`func (o *ZoneType) SetHasSecurityGroups(v bool)`

SetHasSecurityGroups sets HasSecurityGroups field to given value.

### HasHasSecurityGroups

`func (o *ZoneType) HasHasSecurityGroups() bool`

HasHasSecurityGroups returns a boolean if a field has been set.

### GetHasContainers

`func (o *ZoneType) GetHasContainers() bool`

GetHasContainers returns the HasContainers field if non-nil, zero value otherwise.

### GetHasContainersOk

`func (o *ZoneType) GetHasContainersOk() (*bool, bool)`

GetHasContainersOk returns a tuple with the HasContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasContainers

`func (o *ZoneType) SetHasContainers(v bool)`

SetHasContainers sets HasContainers field to given value.

### HasHasContainers

`func (o *ZoneType) HasHasContainers() bool`

HasHasContainers returns a boolean if a field has been set.

### GetHasBareMetal

`func (o *ZoneType) GetHasBareMetal() bool`

GetHasBareMetal returns the HasBareMetal field if non-nil, zero value otherwise.

### GetHasBareMetalOk

`func (o *ZoneType) GetHasBareMetalOk() (*bool, bool)`

GetHasBareMetalOk returns a tuple with the HasBareMetal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBareMetal

`func (o *ZoneType) SetHasBareMetal(v bool)`

SetHasBareMetal sets HasBareMetal field to given value.

### HasHasBareMetal

`func (o *ZoneType) HasHasBareMetal() bool`

HasHasBareMetal returns a boolean if a field has been set.

### GetHasServices

`func (o *ZoneType) GetHasServices() bool`

GetHasServices returns the HasServices field if non-nil, zero value otherwise.

### GetHasServicesOk

`func (o *ZoneType) GetHasServicesOk() (*bool, bool)`

GetHasServicesOk returns a tuple with the HasServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasServices

`func (o *ZoneType) SetHasServices(v bool)`

SetHasServices sets HasServices field to given value.

### HasHasServices

`func (o *ZoneType) HasHasServices() bool`

HasHasServices returns a boolean if a field has been set.

### GetHasFunctions

`func (o *ZoneType) GetHasFunctions() bool`

GetHasFunctions returns the HasFunctions field if non-nil, zero value otherwise.

### GetHasFunctionsOk

`func (o *ZoneType) GetHasFunctionsOk() (*bool, bool)`

GetHasFunctionsOk returns a tuple with the HasFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFunctions

`func (o *ZoneType) SetHasFunctions(v bool)`

SetHasFunctions sets HasFunctions field to given value.

### HasHasFunctions

`func (o *ZoneType) HasHasFunctions() bool`

HasHasFunctions returns a boolean if a field has been set.

### GetHasJobs

`func (o *ZoneType) GetHasJobs() bool`

GetHasJobs returns the HasJobs field if non-nil, zero value otherwise.

### GetHasJobsOk

`func (o *ZoneType) GetHasJobsOk() (*bool, bool)`

GetHasJobsOk returns a tuple with the HasJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasJobs

`func (o *ZoneType) SetHasJobs(v bool)`

SetHasJobs sets HasJobs field to given value.

### HasHasJobs

`func (o *ZoneType) HasHasJobs() bool`

HasHasJobs returns a boolean if a field has been set.

### GetHasDiscovery

`func (o *ZoneType) GetHasDiscovery() bool`

GetHasDiscovery returns the HasDiscovery field if non-nil, zero value otherwise.

### GetHasDiscoveryOk

`func (o *ZoneType) GetHasDiscoveryOk() (*bool, bool)`

GetHasDiscoveryOk returns a tuple with the HasDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDiscovery

`func (o *ZoneType) SetHasDiscovery(v bool)`

SetHasDiscovery sets HasDiscovery field to given value.

### HasHasDiscovery

`func (o *ZoneType) HasHasDiscovery() bool`

HasHasDiscovery returns a boolean if a field has been set.

### GetHasCloudInit

`func (o *ZoneType) GetHasCloudInit() bool`

GetHasCloudInit returns the HasCloudInit field if non-nil, zero value otherwise.

### GetHasCloudInitOk

`func (o *ZoneType) GetHasCloudInitOk() (*bool, bool)`

GetHasCloudInitOk returns a tuple with the HasCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCloudInit

`func (o *ZoneType) SetHasCloudInit(v bool)`

SetHasCloudInit sets HasCloudInit field to given value.

### HasHasCloudInit

`func (o *ZoneType) HasHasCloudInit() bool`

HasHasCloudInit returns a boolean if a field has been set.

### GetHasFolders

`func (o *ZoneType) GetHasFolders() bool`

GetHasFolders returns the HasFolders field if non-nil, zero value otherwise.

### GetHasFoldersOk

`func (o *ZoneType) GetHasFoldersOk() (*bool, bool)`

GetHasFoldersOk returns a tuple with the HasFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFolders

`func (o *ZoneType) SetHasFolders(v bool)`

SetHasFolders sets HasFolders field to given value.

### HasHasFolders

`func (o *ZoneType) HasHasFolders() bool`

HasHasFolders returns a boolean if a field has been set.

### GetHasMarketplace

`func (o *ZoneType) GetHasMarketplace() bool`

GetHasMarketplace returns the HasMarketplace field if non-nil, zero value otherwise.

### GetHasMarketplaceOk

`func (o *ZoneType) GetHasMarketplaceOk() (*bool, bool)`

GetHasMarketplaceOk returns a tuple with the HasMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMarketplace

`func (o *ZoneType) SetHasMarketplace(v bool)`

SetHasMarketplace sets HasMarketplace field to given value.

### HasHasMarketplace

`func (o *ZoneType) HasHasMarketplace() bool`

HasHasMarketplace returns a boolean if a field has been set.

### GetHasNativePlans

`func (o *ZoneType) GetHasNativePlans() bool`

GetHasNativePlans returns the HasNativePlans field if non-nil, zero value otherwise.

### GetHasNativePlansOk

`func (o *ZoneType) GetHasNativePlansOk() (*bool, bool)`

GetHasNativePlansOk returns a tuple with the HasNativePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNativePlans

`func (o *ZoneType) SetHasNativePlans(v bool)`

SetHasNativePlans sets HasNativePlans field to given value.

### HasHasNativePlans

`func (o *ZoneType) HasHasNativePlans() bool`

HasHasNativePlans returns a boolean if a field has been set.

### GetCanCreateResourcePools

`func (o *ZoneType) GetCanCreateResourcePools() bool`

GetCanCreateResourcePools returns the CanCreateResourcePools field if non-nil, zero value otherwise.

### GetCanCreateResourcePoolsOk

`func (o *ZoneType) GetCanCreateResourcePoolsOk() (*bool, bool)`

GetCanCreateResourcePoolsOk returns a tuple with the CanCreateResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateResourcePools

`func (o *ZoneType) SetCanCreateResourcePools(v bool)`

SetCanCreateResourcePools sets CanCreateResourcePools field to given value.

### HasCanCreateResourcePools

`func (o *ZoneType) HasCanCreateResourcePools() bool`

HasCanCreateResourcePools returns a boolean if a field has been set.

### GetCanDeleteResourcePools

`func (o *ZoneType) GetCanDeleteResourcePools() bool`

GetCanDeleteResourcePools returns the CanDeleteResourcePools field if non-nil, zero value otherwise.

### GetCanDeleteResourcePoolsOk

`func (o *ZoneType) GetCanDeleteResourcePoolsOk() (*bool, bool)`

GetCanDeleteResourcePoolsOk returns a tuple with the CanDeleteResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteResourcePools

`func (o *ZoneType) SetCanDeleteResourcePools(v bool)`

SetCanDeleteResourcePools sets CanDeleteResourcePools field to given value.

### HasCanDeleteResourcePools

`func (o *ZoneType) HasCanDeleteResourcePools() bool`

HasCanDeleteResourcePools returns a boolean if a field has been set.

### GetCanCreateDatastores

`func (o *ZoneType) GetCanCreateDatastores() bool`

GetCanCreateDatastores returns the CanCreateDatastores field if non-nil, zero value otherwise.

### GetCanCreateDatastoresOk

`func (o *ZoneType) GetCanCreateDatastoresOk() (*bool, bool)`

GetCanCreateDatastoresOk returns a tuple with the CanCreateDatastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateDatastores

`func (o *ZoneType) SetCanCreateDatastores(v bool)`

SetCanCreateDatastores sets CanCreateDatastores field to given value.

### HasCanCreateDatastores

`func (o *ZoneType) HasCanCreateDatastores() bool`

HasCanCreateDatastores returns a boolean if a field has been set.

### GetCanCreateNetworks

`func (o *ZoneType) GetCanCreateNetworks() bool`

GetCanCreateNetworks returns the CanCreateNetworks field if non-nil, zero value otherwise.

### GetCanCreateNetworksOk

`func (o *ZoneType) GetCanCreateNetworksOk() (*bool, bool)`

GetCanCreateNetworksOk returns a tuple with the CanCreateNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateNetworks

`func (o *ZoneType) SetCanCreateNetworks(v bool)`

SetCanCreateNetworks sets CanCreateNetworks field to given value.

### HasCanCreateNetworks

`func (o *ZoneType) HasCanCreateNetworks() bool`

HasCanCreateNetworks returns a boolean if a field has been set.

### GetCanChooseContainerMode

`func (o *ZoneType) GetCanChooseContainerMode() bool`

GetCanChooseContainerMode returns the CanChooseContainerMode field if non-nil, zero value otherwise.

### GetCanChooseContainerModeOk

`func (o *ZoneType) GetCanChooseContainerModeOk() (*bool, bool)`

GetCanChooseContainerModeOk returns a tuple with the CanChooseContainerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanChooseContainerMode

`func (o *ZoneType) SetCanChooseContainerMode(v bool)`

SetCanChooseContainerMode sets CanChooseContainerMode field to given value.

### HasCanChooseContainerMode

`func (o *ZoneType) HasCanChooseContainerMode() bool`

HasCanChooseContainerMode returns a boolean if a field has been set.

### GetProvisionRequiresResourcePool

`func (o *ZoneType) GetProvisionRequiresResourcePool() bool`

GetProvisionRequiresResourcePool returns the ProvisionRequiresResourcePool field if non-nil, zero value otherwise.

### GetProvisionRequiresResourcePoolOk

`func (o *ZoneType) GetProvisionRequiresResourcePoolOk() (*bool, bool)`

GetProvisionRequiresResourcePoolOk returns a tuple with the ProvisionRequiresResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionRequiresResourcePool

`func (o *ZoneType) SetProvisionRequiresResourcePool(v bool)`

SetProvisionRequiresResourcePool sets ProvisionRequiresResourcePool field to given value.

### HasProvisionRequiresResourcePool

`func (o *ZoneType) HasProvisionRequiresResourcePool() bool`

HasProvisionRequiresResourcePool returns a boolean if a field has been set.

### GetSupportsDistributedWorker

`func (o *ZoneType) GetSupportsDistributedWorker() bool`

GetSupportsDistributedWorker returns the SupportsDistributedWorker field if non-nil, zero value otherwise.

### GetSupportsDistributedWorkerOk

`func (o *ZoneType) GetSupportsDistributedWorkerOk() (*bool, bool)`

GetSupportsDistributedWorkerOk returns a tuple with the SupportsDistributedWorker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDistributedWorker

`func (o *ZoneType) SetSupportsDistributedWorker(v bool)`

SetSupportsDistributedWorker sets SupportsDistributedWorker field to given value.

### HasSupportsDistributedWorker

`func (o *ZoneType) HasSupportsDistributedWorker() bool`

HasSupportsDistributedWorker returns a boolean if a field has been set.

### GetCloud

`func (o *ZoneType) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ZoneType) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ZoneType) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ZoneType) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetProvisionTypes

`func (o *ZoneType) GetProvisionTypes() []int64`

GetProvisionTypes returns the ProvisionTypes field if non-nil, zero value otherwise.

### GetProvisionTypesOk

`func (o *ZoneType) GetProvisionTypesOk() (*[]int64, bool)`

GetProvisionTypesOk returns a tuple with the ProvisionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionTypes

`func (o *ZoneType) SetProvisionTypes(v []int64)`

SetProvisionTypes sets ProvisionTypes field to given value.

### HasProvisionTypes

`func (o *ZoneType) HasProvisionTypes() bool`

HasProvisionTypes returns a boolean if a field has been set.

### GetZoneInstanceTypeLayoutId

`func (o *ZoneType) GetZoneInstanceTypeLayoutId() int64`

GetZoneInstanceTypeLayoutId returns the ZoneInstanceTypeLayoutId field if non-nil, zero value otherwise.

### GetZoneInstanceTypeLayoutIdOk

`func (o *ZoneType) GetZoneInstanceTypeLayoutIdOk() (*int64, bool)`

GetZoneInstanceTypeLayoutIdOk returns a tuple with the ZoneInstanceTypeLayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneInstanceTypeLayoutId

`func (o *ZoneType) SetZoneInstanceTypeLayoutId(v int64)`

SetZoneInstanceTypeLayoutId sets ZoneInstanceTypeLayoutId field to given value.

### HasZoneInstanceTypeLayoutId

`func (o *ZoneType) HasZoneInstanceTypeLayoutId() bool`

HasZoneInstanceTypeLayoutId returns a boolean if a field has been set.

### GetServerTypes

`func (o *ZoneType) GetServerTypes() []ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner`

GetServerTypes returns the ServerTypes field if non-nil, zero value otherwise.

### GetServerTypesOk

`func (o *ZoneType) GetServerTypesOk() (*[]ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner, bool)`

GetServerTypesOk returns a tuple with the ServerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypes

`func (o *ZoneType) SetServerTypes(v []ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner)`

SetServerTypes sets ServerTypes field to given value.

### HasServerTypes

`func (o *ZoneType) HasServerTypes() bool`

HasServerTypes returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ZoneType) GetOptionTypes() []ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ZoneType) GetOptionTypesOk() (*[]ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ZoneType) SetOptionTypes(v []ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ZoneType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


