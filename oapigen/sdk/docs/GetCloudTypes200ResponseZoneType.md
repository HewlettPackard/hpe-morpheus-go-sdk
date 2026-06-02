# GetCloudTypes200ResponseZoneType

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
**HasAffinityGroups** | Pointer to **bool** |  | [optional] 
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
**ServerTypes** | Pointer to [**[]GetCloudTypes200ResponseZoneTypeServerTypesInner**](GetCloudTypes200ResponseZoneTypeServerTypesInner.md) |  | [optional] 
**OptionTypes** | Pointer to [**[]GetCloudTypes200ResponseZoneTypeOptionTypesInner**](GetCloudTypes200ResponseZoneTypeOptionTypesInner.md) |  | [optional] 

## Methods

### NewGetCloudTypes200ResponseZoneType

`func NewGetCloudTypes200ResponseZoneType() *GetCloudTypes200ResponseZoneType`

NewGetCloudTypes200ResponseZoneType instantiates a new GetCloudTypes200ResponseZoneType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetCloudTypes200ResponseZoneType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCloudTypes200ResponseZoneType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCloudTypes200ResponseZoneType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCloudTypes200ResponseZoneType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetCloudTypes200ResponseZoneType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCloudTypes200ResponseZoneType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCloudTypes200ResponseZoneType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCloudTypes200ResponseZoneType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetCloudTypes200ResponseZoneType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetCloudTypes200ResponseZoneType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetCloudTypes200ResponseZoneType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetCloudTypes200ResponseZoneType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEnabled

`func (o *GetCloudTypes200ResponseZoneType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetCloudTypes200ResponseZoneType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetCloudTypes200ResponseZoneType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetCloudTypes200ResponseZoneType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProvision

`func (o *GetCloudTypes200ResponseZoneType) GetProvision() bool`

GetProvision returns the Provision field if non-nil, zero value otherwise.

### GetProvisionOk

`func (o *GetCloudTypes200ResponseZoneType) GetProvisionOk() (*bool, bool)`

GetProvisionOk returns a tuple with the Provision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvision

`func (o *GetCloudTypes200ResponseZoneType) SetProvision(v bool)`

SetProvision sets Provision field to given value.

### HasProvision

`func (o *GetCloudTypes200ResponseZoneType) HasProvision() bool`

HasProvision returns a boolean if a field has been set.

### GetAutoCapacity

`func (o *GetCloudTypes200ResponseZoneType) GetAutoCapacity() bool`

GetAutoCapacity returns the AutoCapacity field if non-nil, zero value otherwise.

### GetAutoCapacityOk

`func (o *GetCloudTypes200ResponseZoneType) GetAutoCapacityOk() (*bool, bool)`

GetAutoCapacityOk returns a tuple with the AutoCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCapacity

`func (o *GetCloudTypes200ResponseZoneType) SetAutoCapacity(v bool)`

SetAutoCapacity sets AutoCapacity field to given value.

### HasAutoCapacity

`func (o *GetCloudTypes200ResponseZoneType) HasAutoCapacity() bool`

HasAutoCapacity returns a boolean if a field has been set.

### GetMigrationTarget

`func (o *GetCloudTypes200ResponseZoneType) GetMigrationTarget() bool`

GetMigrationTarget returns the MigrationTarget field if non-nil, zero value otherwise.

### GetMigrationTargetOk

`func (o *GetCloudTypes200ResponseZoneType) GetMigrationTargetOk() (*bool, bool)`

GetMigrationTargetOk returns a tuple with the MigrationTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationTarget

`func (o *GetCloudTypes200ResponseZoneType) SetMigrationTarget(v bool)`

SetMigrationTarget sets MigrationTarget field to given value.

### HasMigrationTarget

`func (o *GetCloudTypes200ResponseZoneType) HasMigrationTarget() bool`

HasMigrationTarget returns a boolean if a field has been set.

### GetHasAffinityGroups

`func (o *GetCloudTypes200ResponseZoneType) GetHasAffinityGroups() bool`

GetHasAffinityGroups returns the HasAffinityGroups field if non-nil, zero value otherwise.

### GetHasAffinityGroupsOk

`func (o *GetCloudTypes200ResponseZoneType) GetHasAffinityGroupsOk() (*bool, bool)`

GetHasAffinityGroupsOk returns a tuple with the HasAffinityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAffinityGroups

`func (o *GetCloudTypes200ResponseZoneType) SetHasAffinityGroups(v bool)`

SetHasAffinityGroups sets HasAffinityGroups field to given value.

### HasHasAffinityGroups

`func (o *GetCloudTypes200ResponseZoneType) HasHasAffinityGroups() bool`

HasHasAffinityGroups returns a boolean if a field has been set.

### GetHasDatastores

`func (o *GetCloudTypes200ResponseZoneType) GetHasDatastores() bool`

GetHasDatastores returns the HasDatastores field if non-nil, zero value otherwise.

### GetHasDatastoresOk

`func (o *GetCloudTypes200ResponseZoneType) GetHasDatastoresOk() (*bool, bool)`

GetHasDatastoresOk returns a tuple with the HasDatastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastores

`func (o *GetCloudTypes200ResponseZoneType) SetHasDatastores(v bool)`

SetHasDatastores sets HasDatastores field to given value.

### HasHasDatastores

`func (o *GetCloudTypes200ResponseZoneType) HasHasDatastores() bool`

HasHasDatastores returns a boolean if a field has been set.

### GetHasNetworks

`func (o *GetCloudTypes200ResponseZoneType) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *GetCloudTypes200ResponseZoneType) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *GetCloudTypes200ResponseZoneType) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *GetCloudTypes200ResponseZoneType) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### GetHasResourcePools

`func (o *GetCloudTypes200ResponseZoneType) GetHasResourcePools() bool`

GetHasResourcePools returns the HasResourcePools field if non-nil, zero value otherwise.

### GetHasResourcePoolsOk

`func (o *GetCloudTypes200ResponseZoneType) GetHasResourcePoolsOk() (*bool, bool)`

GetHasResourcePoolsOk returns a tuple with the HasResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasResourcePools

`func (o *GetCloudTypes200ResponseZoneType) SetHasResourcePools(v bool)`

SetHasResourcePools sets HasResourcePools field to given value.

### HasHasResourcePools

`func (o *GetCloudTypes200ResponseZoneType) HasHasResourcePools() bool`

HasHasResourcePools returns a boolean if a field has been set.

### GetHasSecurityGroups

`func (o *GetCloudTypes200ResponseZoneType) GetHasSecurityGroups() bool`

GetHasSecurityGroups returns the HasSecurityGroups field if non-nil, zero value otherwise.

### GetHasSecurityGroupsOk

`func (o *GetCloudTypes200ResponseZoneType) GetHasSecurityGroupsOk() (*bool, bool)`

GetHasSecurityGroupsOk returns a tuple with the HasSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecurityGroups

`func (o *GetCloudTypes200ResponseZoneType) SetHasSecurityGroups(v bool)`

SetHasSecurityGroups sets HasSecurityGroups field to given value.

### HasHasSecurityGroups

`func (o *GetCloudTypes200ResponseZoneType) HasHasSecurityGroups() bool`

HasHasSecurityGroups returns a boolean if a field has been set.

### GetHasContainers

`func (o *GetCloudTypes200ResponseZoneType) GetHasContainers() bool`

GetHasContainers returns the HasContainers field if non-nil, zero value otherwise.

### GetHasContainersOk

`func (o *GetCloudTypes200ResponseZoneType) GetHasContainersOk() (*bool, bool)`

GetHasContainersOk returns a tuple with the HasContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasContainers

`func (o *GetCloudTypes200ResponseZoneType) SetHasContainers(v bool)`

SetHasContainers sets HasContainers field to given value.

### HasHasContainers

`func (o *GetCloudTypes200ResponseZoneType) HasHasContainers() bool`

HasHasContainers returns a boolean if a field has been set.

### GetHasBareMetal

`func (o *GetCloudTypes200ResponseZoneType) GetHasBareMetal() bool`

GetHasBareMetal returns the HasBareMetal field if non-nil, zero value otherwise.

### GetHasBareMetalOk

`func (o *GetCloudTypes200ResponseZoneType) GetHasBareMetalOk() (*bool, bool)`

GetHasBareMetalOk returns a tuple with the HasBareMetal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBareMetal

`func (o *GetCloudTypes200ResponseZoneType) SetHasBareMetal(v bool)`

SetHasBareMetal sets HasBareMetal field to given value.

### HasHasBareMetal

`func (o *GetCloudTypes200ResponseZoneType) HasHasBareMetal() bool`

HasHasBareMetal returns a boolean if a field has been set.

### GetHasServices

`func (o *GetCloudTypes200ResponseZoneType) GetHasServices() bool`

GetHasServices returns the HasServices field if non-nil, zero value otherwise.

### GetHasServicesOk

`func (o *GetCloudTypes200ResponseZoneType) GetHasServicesOk() (*bool, bool)`

GetHasServicesOk returns a tuple with the HasServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasServices

`func (o *GetCloudTypes200ResponseZoneType) SetHasServices(v bool)`

SetHasServices sets HasServices field to given value.

### HasHasServices

`func (o *GetCloudTypes200ResponseZoneType) HasHasServices() bool`

HasHasServices returns a boolean if a field has been set.

### GetHasFunctions

`func (o *GetCloudTypes200ResponseZoneType) GetHasFunctions() bool`

GetHasFunctions returns the HasFunctions field if non-nil, zero value otherwise.

### GetHasFunctionsOk

`func (o *GetCloudTypes200ResponseZoneType) GetHasFunctionsOk() (*bool, bool)`

GetHasFunctionsOk returns a tuple with the HasFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFunctions

`func (o *GetCloudTypes200ResponseZoneType) SetHasFunctions(v bool)`

SetHasFunctions sets HasFunctions field to given value.

### HasHasFunctions

`func (o *GetCloudTypes200ResponseZoneType) HasHasFunctions() bool`

HasHasFunctions returns a boolean if a field has been set.

### GetHasJobs

`func (o *GetCloudTypes200ResponseZoneType) GetHasJobs() bool`

GetHasJobs returns the HasJobs field if non-nil, zero value otherwise.

### GetHasJobsOk

`func (o *GetCloudTypes200ResponseZoneType) GetHasJobsOk() (*bool, bool)`

GetHasJobsOk returns a tuple with the HasJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasJobs

`func (o *GetCloudTypes200ResponseZoneType) SetHasJobs(v bool)`

SetHasJobs sets HasJobs field to given value.

### HasHasJobs

`func (o *GetCloudTypes200ResponseZoneType) HasHasJobs() bool`

HasHasJobs returns a boolean if a field has been set.

### GetHasDiscovery

`func (o *GetCloudTypes200ResponseZoneType) GetHasDiscovery() bool`

GetHasDiscovery returns the HasDiscovery field if non-nil, zero value otherwise.

### GetHasDiscoveryOk

`func (o *GetCloudTypes200ResponseZoneType) GetHasDiscoveryOk() (*bool, bool)`

GetHasDiscoveryOk returns a tuple with the HasDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDiscovery

`func (o *GetCloudTypes200ResponseZoneType) SetHasDiscovery(v bool)`

SetHasDiscovery sets HasDiscovery field to given value.

### HasHasDiscovery

`func (o *GetCloudTypes200ResponseZoneType) HasHasDiscovery() bool`

HasHasDiscovery returns a boolean if a field has been set.

### GetHasCloudInit

`func (o *GetCloudTypes200ResponseZoneType) GetHasCloudInit() bool`

GetHasCloudInit returns the HasCloudInit field if non-nil, zero value otherwise.

### GetHasCloudInitOk

`func (o *GetCloudTypes200ResponseZoneType) GetHasCloudInitOk() (*bool, bool)`

GetHasCloudInitOk returns a tuple with the HasCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCloudInit

`func (o *GetCloudTypes200ResponseZoneType) SetHasCloudInit(v bool)`

SetHasCloudInit sets HasCloudInit field to given value.

### HasHasCloudInit

`func (o *GetCloudTypes200ResponseZoneType) HasHasCloudInit() bool`

HasHasCloudInit returns a boolean if a field has been set.

### GetHasFolders

`func (o *GetCloudTypes200ResponseZoneType) GetHasFolders() bool`

GetHasFolders returns the HasFolders field if non-nil, zero value otherwise.

### GetHasFoldersOk

`func (o *GetCloudTypes200ResponseZoneType) GetHasFoldersOk() (*bool, bool)`

GetHasFoldersOk returns a tuple with the HasFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFolders

`func (o *GetCloudTypes200ResponseZoneType) SetHasFolders(v bool)`

SetHasFolders sets HasFolders field to given value.

### HasHasFolders

`func (o *GetCloudTypes200ResponseZoneType) HasHasFolders() bool`

HasHasFolders returns a boolean if a field has been set.

### GetHasMarketplace

`func (o *GetCloudTypes200ResponseZoneType) GetHasMarketplace() bool`

GetHasMarketplace returns the HasMarketplace field if non-nil, zero value otherwise.

### GetHasMarketplaceOk

`func (o *GetCloudTypes200ResponseZoneType) GetHasMarketplaceOk() (*bool, bool)`

GetHasMarketplaceOk returns a tuple with the HasMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMarketplace

`func (o *GetCloudTypes200ResponseZoneType) SetHasMarketplace(v bool)`

SetHasMarketplace sets HasMarketplace field to given value.

### HasHasMarketplace

`func (o *GetCloudTypes200ResponseZoneType) HasHasMarketplace() bool`

HasHasMarketplace returns a boolean if a field has been set.

### GetHasNativePlans

`func (o *GetCloudTypes200ResponseZoneType) GetHasNativePlans() bool`

GetHasNativePlans returns the HasNativePlans field if non-nil, zero value otherwise.

### GetHasNativePlansOk

`func (o *GetCloudTypes200ResponseZoneType) GetHasNativePlansOk() (*bool, bool)`

GetHasNativePlansOk returns a tuple with the HasNativePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNativePlans

`func (o *GetCloudTypes200ResponseZoneType) SetHasNativePlans(v bool)`

SetHasNativePlans sets HasNativePlans field to given value.

### HasHasNativePlans

`func (o *GetCloudTypes200ResponseZoneType) HasHasNativePlans() bool`

HasHasNativePlans returns a boolean if a field has been set.

### GetCanCreateResourcePools

`func (o *GetCloudTypes200ResponseZoneType) GetCanCreateResourcePools() bool`

GetCanCreateResourcePools returns the CanCreateResourcePools field if non-nil, zero value otherwise.

### GetCanCreateResourcePoolsOk

`func (o *GetCloudTypes200ResponseZoneType) GetCanCreateResourcePoolsOk() (*bool, bool)`

GetCanCreateResourcePoolsOk returns a tuple with the CanCreateResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateResourcePools

`func (o *GetCloudTypes200ResponseZoneType) SetCanCreateResourcePools(v bool)`

SetCanCreateResourcePools sets CanCreateResourcePools field to given value.

### HasCanCreateResourcePools

`func (o *GetCloudTypes200ResponseZoneType) HasCanCreateResourcePools() bool`

HasCanCreateResourcePools returns a boolean if a field has been set.

### GetCanDeleteResourcePools

`func (o *GetCloudTypes200ResponseZoneType) GetCanDeleteResourcePools() bool`

GetCanDeleteResourcePools returns the CanDeleteResourcePools field if non-nil, zero value otherwise.

### GetCanDeleteResourcePoolsOk

`func (o *GetCloudTypes200ResponseZoneType) GetCanDeleteResourcePoolsOk() (*bool, bool)`

GetCanDeleteResourcePoolsOk returns a tuple with the CanDeleteResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteResourcePools

`func (o *GetCloudTypes200ResponseZoneType) SetCanDeleteResourcePools(v bool)`

SetCanDeleteResourcePools sets CanDeleteResourcePools field to given value.

### HasCanDeleteResourcePools

`func (o *GetCloudTypes200ResponseZoneType) HasCanDeleteResourcePools() bool`

HasCanDeleteResourcePools returns a boolean if a field has been set.

### GetCanCreateDatastores

`func (o *GetCloudTypes200ResponseZoneType) GetCanCreateDatastores() bool`

GetCanCreateDatastores returns the CanCreateDatastores field if non-nil, zero value otherwise.

### GetCanCreateDatastoresOk

`func (o *GetCloudTypes200ResponseZoneType) GetCanCreateDatastoresOk() (*bool, bool)`

GetCanCreateDatastoresOk returns a tuple with the CanCreateDatastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateDatastores

`func (o *GetCloudTypes200ResponseZoneType) SetCanCreateDatastores(v bool)`

SetCanCreateDatastores sets CanCreateDatastores field to given value.

### HasCanCreateDatastores

`func (o *GetCloudTypes200ResponseZoneType) HasCanCreateDatastores() bool`

HasCanCreateDatastores returns a boolean if a field has been set.

### GetCanCreateNetworks

`func (o *GetCloudTypes200ResponseZoneType) GetCanCreateNetworks() bool`

GetCanCreateNetworks returns the CanCreateNetworks field if non-nil, zero value otherwise.

### GetCanCreateNetworksOk

`func (o *GetCloudTypes200ResponseZoneType) GetCanCreateNetworksOk() (*bool, bool)`

GetCanCreateNetworksOk returns a tuple with the CanCreateNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateNetworks

`func (o *GetCloudTypes200ResponseZoneType) SetCanCreateNetworks(v bool)`

SetCanCreateNetworks sets CanCreateNetworks field to given value.

### HasCanCreateNetworks

`func (o *GetCloudTypes200ResponseZoneType) HasCanCreateNetworks() bool`

HasCanCreateNetworks returns a boolean if a field has been set.

### GetCanChooseContainerMode

`func (o *GetCloudTypes200ResponseZoneType) GetCanChooseContainerMode() bool`

GetCanChooseContainerMode returns the CanChooseContainerMode field if non-nil, zero value otherwise.

### GetCanChooseContainerModeOk

`func (o *GetCloudTypes200ResponseZoneType) GetCanChooseContainerModeOk() (*bool, bool)`

GetCanChooseContainerModeOk returns a tuple with the CanChooseContainerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanChooseContainerMode

`func (o *GetCloudTypes200ResponseZoneType) SetCanChooseContainerMode(v bool)`

SetCanChooseContainerMode sets CanChooseContainerMode field to given value.

### HasCanChooseContainerMode

`func (o *GetCloudTypes200ResponseZoneType) HasCanChooseContainerMode() bool`

HasCanChooseContainerMode returns a boolean if a field has been set.

### GetProvisionRequiresResourcePool

`func (o *GetCloudTypes200ResponseZoneType) GetProvisionRequiresResourcePool() bool`

GetProvisionRequiresResourcePool returns the ProvisionRequiresResourcePool field if non-nil, zero value otherwise.

### GetProvisionRequiresResourcePoolOk

`func (o *GetCloudTypes200ResponseZoneType) GetProvisionRequiresResourcePoolOk() (*bool, bool)`

GetProvisionRequiresResourcePoolOk returns a tuple with the ProvisionRequiresResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionRequiresResourcePool

`func (o *GetCloudTypes200ResponseZoneType) SetProvisionRequiresResourcePool(v bool)`

SetProvisionRequiresResourcePool sets ProvisionRequiresResourcePool field to given value.

### HasProvisionRequiresResourcePool

`func (o *GetCloudTypes200ResponseZoneType) HasProvisionRequiresResourcePool() bool`

HasProvisionRequiresResourcePool returns a boolean if a field has been set.

### GetSupportsDistributedWorker

`func (o *GetCloudTypes200ResponseZoneType) GetSupportsDistributedWorker() bool`

GetSupportsDistributedWorker returns the SupportsDistributedWorker field if non-nil, zero value otherwise.

### GetSupportsDistributedWorkerOk

`func (o *GetCloudTypes200ResponseZoneType) GetSupportsDistributedWorkerOk() (*bool, bool)`

GetSupportsDistributedWorkerOk returns a tuple with the SupportsDistributedWorker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDistributedWorker

`func (o *GetCloudTypes200ResponseZoneType) SetSupportsDistributedWorker(v bool)`

SetSupportsDistributedWorker sets SupportsDistributedWorker field to given value.

### HasSupportsDistributedWorker

`func (o *GetCloudTypes200ResponseZoneType) HasSupportsDistributedWorker() bool`

HasSupportsDistributedWorker returns a boolean if a field has been set.

### GetCloud

`func (o *GetCloudTypes200ResponseZoneType) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *GetCloudTypes200ResponseZoneType) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *GetCloudTypes200ResponseZoneType) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *GetCloudTypes200ResponseZoneType) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetProvisionTypes

`func (o *GetCloudTypes200ResponseZoneType) GetProvisionTypes() []int64`

GetProvisionTypes returns the ProvisionTypes field if non-nil, zero value otherwise.

### GetProvisionTypesOk

`func (o *GetCloudTypes200ResponseZoneType) GetProvisionTypesOk() (*[]int64, bool)`

GetProvisionTypesOk returns a tuple with the ProvisionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionTypes

`func (o *GetCloudTypes200ResponseZoneType) SetProvisionTypes(v []int64)`

SetProvisionTypes sets ProvisionTypes field to given value.

### HasProvisionTypes

`func (o *GetCloudTypes200ResponseZoneType) HasProvisionTypes() bool`

HasProvisionTypes returns a boolean if a field has been set.

### GetZoneInstanceTypeLayoutId

`func (o *GetCloudTypes200ResponseZoneType) GetZoneInstanceTypeLayoutId() int64`

GetZoneInstanceTypeLayoutId returns the ZoneInstanceTypeLayoutId field if non-nil, zero value otherwise.

### GetZoneInstanceTypeLayoutIdOk

`func (o *GetCloudTypes200ResponseZoneType) GetZoneInstanceTypeLayoutIdOk() (*int64, bool)`

GetZoneInstanceTypeLayoutIdOk returns a tuple with the ZoneInstanceTypeLayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneInstanceTypeLayoutId

`func (o *GetCloudTypes200ResponseZoneType) SetZoneInstanceTypeLayoutId(v int64)`

SetZoneInstanceTypeLayoutId sets ZoneInstanceTypeLayoutId field to given value.

### HasZoneInstanceTypeLayoutId

`func (o *GetCloudTypes200ResponseZoneType) HasZoneInstanceTypeLayoutId() bool`

HasZoneInstanceTypeLayoutId returns a boolean if a field has been set.

### GetServerTypes

`func (o *GetCloudTypes200ResponseZoneType) GetServerTypes() []GetCloudTypes200ResponseZoneTypeServerTypesInner`

GetServerTypes returns the ServerTypes field if non-nil, zero value otherwise.

### GetServerTypesOk

`func (o *GetCloudTypes200ResponseZoneType) GetServerTypesOk() (*[]GetCloudTypes200ResponseZoneTypeServerTypesInner, bool)`

GetServerTypesOk returns a tuple with the ServerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypes

`func (o *GetCloudTypes200ResponseZoneType) SetServerTypes(v []GetCloudTypes200ResponseZoneTypeServerTypesInner)`

SetServerTypes sets ServerTypes field to given value.

### HasServerTypes

`func (o *GetCloudTypes200ResponseZoneType) HasServerTypes() bool`

HasServerTypes returns a boolean if a field has been set.

### GetOptionTypes

`func (o *GetCloudTypes200ResponseZoneType) GetOptionTypes() []GetCloudTypes200ResponseZoneTypeOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetCloudTypes200ResponseZoneType) GetOptionTypesOk() (*[]GetCloudTypes200ResponseZoneTypeOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetCloudTypes200ResponseZoneType) SetOptionTypes(v []GetCloudTypes200ResponseZoneTypeOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetCloudTypes200ResponseZoneType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


