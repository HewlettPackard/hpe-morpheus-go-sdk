# ListCloudTypes200ResponseAllOfZoneTypesInner

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

### NewListCloudTypes200ResponseAllOfZoneTypesInner

`func NewListCloudTypes200ResponseAllOfZoneTypesInner() *ListCloudTypes200ResponseAllOfZoneTypesInner`

NewListCloudTypes200ResponseAllOfZoneTypesInner instantiates a new ListCloudTypes200ResponseAllOfZoneTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCloudTypes200ResponseAllOfZoneTypesInnerWithDefaults

`func NewListCloudTypes200ResponseAllOfZoneTypesInnerWithDefaults() *ListCloudTypes200ResponseAllOfZoneTypesInner`

NewListCloudTypes200ResponseAllOfZoneTypesInnerWithDefaults instantiates a new ListCloudTypes200ResponseAllOfZoneTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEnabled

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProvision

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetProvision() bool`

GetProvision returns the Provision field if non-nil, zero value otherwise.

### GetProvisionOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetProvisionOk() (*bool, bool)`

GetProvisionOk returns a tuple with the Provision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvision

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetProvision(v bool)`

SetProvision sets Provision field to given value.

### HasProvision

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasProvision() bool`

HasProvision returns a boolean if a field has been set.

### GetAutoCapacity

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetAutoCapacity() bool`

GetAutoCapacity returns the AutoCapacity field if non-nil, zero value otherwise.

### GetAutoCapacityOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetAutoCapacityOk() (*bool, bool)`

GetAutoCapacityOk returns a tuple with the AutoCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCapacity

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetAutoCapacity(v bool)`

SetAutoCapacity sets AutoCapacity field to given value.

### HasAutoCapacity

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasAutoCapacity() bool`

HasAutoCapacity returns a boolean if a field has been set.

### GetMigrationTarget

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetMigrationTarget() bool`

GetMigrationTarget returns the MigrationTarget field if non-nil, zero value otherwise.

### GetMigrationTargetOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetMigrationTargetOk() (*bool, bool)`

GetMigrationTargetOk returns a tuple with the MigrationTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationTarget

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetMigrationTarget(v bool)`

SetMigrationTarget sets MigrationTarget field to given value.

### HasMigrationTarget

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasMigrationTarget() bool`

HasMigrationTarget returns a boolean if a field has been set.

### GetHasDatastores

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasDatastores() bool`

GetHasDatastores returns the HasDatastores field if non-nil, zero value otherwise.

### GetHasDatastoresOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasDatastoresOk() (*bool, bool)`

GetHasDatastoresOk returns a tuple with the HasDatastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastores

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetHasDatastores(v bool)`

SetHasDatastores sets HasDatastores field to given value.

### HasHasDatastores

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasHasDatastores() bool`

HasHasDatastores returns a boolean if a field has been set.

### GetHasNetworks

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### GetHasResourcePools

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasResourcePools() bool`

GetHasResourcePools returns the HasResourcePools field if non-nil, zero value otherwise.

### GetHasResourcePoolsOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasResourcePoolsOk() (*bool, bool)`

GetHasResourcePoolsOk returns a tuple with the HasResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasResourcePools

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetHasResourcePools(v bool)`

SetHasResourcePools sets HasResourcePools field to given value.

### HasHasResourcePools

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasHasResourcePools() bool`

HasHasResourcePools returns a boolean if a field has been set.

### GetHasSecurityGroups

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasSecurityGroups() bool`

GetHasSecurityGroups returns the HasSecurityGroups field if non-nil, zero value otherwise.

### GetHasSecurityGroupsOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasSecurityGroupsOk() (*bool, bool)`

GetHasSecurityGroupsOk returns a tuple with the HasSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecurityGroups

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetHasSecurityGroups(v bool)`

SetHasSecurityGroups sets HasSecurityGroups field to given value.

### HasHasSecurityGroups

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasHasSecurityGroups() bool`

HasHasSecurityGroups returns a boolean if a field has been set.

### GetHasContainers

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasContainers() bool`

GetHasContainers returns the HasContainers field if non-nil, zero value otherwise.

### GetHasContainersOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasContainersOk() (*bool, bool)`

GetHasContainersOk returns a tuple with the HasContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasContainers

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetHasContainers(v bool)`

SetHasContainers sets HasContainers field to given value.

### HasHasContainers

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasHasContainers() bool`

HasHasContainers returns a boolean if a field has been set.

### GetHasBareMetal

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasBareMetal() bool`

GetHasBareMetal returns the HasBareMetal field if non-nil, zero value otherwise.

### GetHasBareMetalOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasBareMetalOk() (*bool, bool)`

GetHasBareMetalOk returns a tuple with the HasBareMetal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBareMetal

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetHasBareMetal(v bool)`

SetHasBareMetal sets HasBareMetal field to given value.

### HasHasBareMetal

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasHasBareMetal() bool`

HasHasBareMetal returns a boolean if a field has been set.

### GetHasServices

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasServices() bool`

GetHasServices returns the HasServices field if non-nil, zero value otherwise.

### GetHasServicesOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasServicesOk() (*bool, bool)`

GetHasServicesOk returns a tuple with the HasServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasServices

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetHasServices(v bool)`

SetHasServices sets HasServices field to given value.

### HasHasServices

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasHasServices() bool`

HasHasServices returns a boolean if a field has been set.

### GetHasFunctions

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasFunctions() bool`

GetHasFunctions returns the HasFunctions field if non-nil, zero value otherwise.

### GetHasFunctionsOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasFunctionsOk() (*bool, bool)`

GetHasFunctionsOk returns a tuple with the HasFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFunctions

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetHasFunctions(v bool)`

SetHasFunctions sets HasFunctions field to given value.

### HasHasFunctions

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasHasFunctions() bool`

HasHasFunctions returns a boolean if a field has been set.

### GetHasJobs

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasJobs() bool`

GetHasJobs returns the HasJobs field if non-nil, zero value otherwise.

### GetHasJobsOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasJobsOk() (*bool, bool)`

GetHasJobsOk returns a tuple with the HasJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasJobs

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetHasJobs(v bool)`

SetHasJobs sets HasJobs field to given value.

### HasHasJobs

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasHasJobs() bool`

HasHasJobs returns a boolean if a field has been set.

### GetHasDiscovery

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasDiscovery() bool`

GetHasDiscovery returns the HasDiscovery field if non-nil, zero value otherwise.

### GetHasDiscoveryOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasDiscoveryOk() (*bool, bool)`

GetHasDiscoveryOk returns a tuple with the HasDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDiscovery

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetHasDiscovery(v bool)`

SetHasDiscovery sets HasDiscovery field to given value.

### HasHasDiscovery

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasHasDiscovery() bool`

HasHasDiscovery returns a boolean if a field has been set.

### GetHasCloudInit

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasCloudInit() bool`

GetHasCloudInit returns the HasCloudInit field if non-nil, zero value otherwise.

### GetHasCloudInitOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasCloudInitOk() (*bool, bool)`

GetHasCloudInitOk returns a tuple with the HasCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCloudInit

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetHasCloudInit(v bool)`

SetHasCloudInit sets HasCloudInit field to given value.

### HasHasCloudInit

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasHasCloudInit() bool`

HasHasCloudInit returns a boolean if a field has been set.

### GetHasFolders

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasFolders() bool`

GetHasFolders returns the HasFolders field if non-nil, zero value otherwise.

### GetHasFoldersOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasFoldersOk() (*bool, bool)`

GetHasFoldersOk returns a tuple with the HasFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFolders

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetHasFolders(v bool)`

SetHasFolders sets HasFolders field to given value.

### HasHasFolders

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasHasFolders() bool`

HasHasFolders returns a boolean if a field has been set.

### GetHasMarketplace

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasMarketplace() bool`

GetHasMarketplace returns the HasMarketplace field if non-nil, zero value otherwise.

### GetHasMarketplaceOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasMarketplaceOk() (*bool, bool)`

GetHasMarketplaceOk returns a tuple with the HasMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMarketplace

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetHasMarketplace(v bool)`

SetHasMarketplace sets HasMarketplace field to given value.

### HasHasMarketplace

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasHasMarketplace() bool`

HasHasMarketplace returns a boolean if a field has been set.

### GetHasNativePlans

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasNativePlans() bool`

GetHasNativePlans returns the HasNativePlans field if non-nil, zero value otherwise.

### GetHasNativePlansOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetHasNativePlansOk() (*bool, bool)`

GetHasNativePlansOk returns a tuple with the HasNativePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNativePlans

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetHasNativePlans(v bool)`

SetHasNativePlans sets HasNativePlans field to given value.

### HasHasNativePlans

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasHasNativePlans() bool`

HasHasNativePlans returns a boolean if a field has been set.

### GetCanCreateResourcePools

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetCanCreateResourcePools() bool`

GetCanCreateResourcePools returns the CanCreateResourcePools field if non-nil, zero value otherwise.

### GetCanCreateResourcePoolsOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetCanCreateResourcePoolsOk() (*bool, bool)`

GetCanCreateResourcePoolsOk returns a tuple with the CanCreateResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateResourcePools

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetCanCreateResourcePools(v bool)`

SetCanCreateResourcePools sets CanCreateResourcePools field to given value.

### HasCanCreateResourcePools

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasCanCreateResourcePools() bool`

HasCanCreateResourcePools returns a boolean if a field has been set.

### GetCanDeleteResourcePools

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetCanDeleteResourcePools() bool`

GetCanDeleteResourcePools returns the CanDeleteResourcePools field if non-nil, zero value otherwise.

### GetCanDeleteResourcePoolsOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetCanDeleteResourcePoolsOk() (*bool, bool)`

GetCanDeleteResourcePoolsOk returns a tuple with the CanDeleteResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeleteResourcePools

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetCanDeleteResourcePools(v bool)`

SetCanDeleteResourcePools sets CanDeleteResourcePools field to given value.

### HasCanDeleteResourcePools

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasCanDeleteResourcePools() bool`

HasCanDeleteResourcePools returns a boolean if a field has been set.

### GetCanCreateDatastores

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetCanCreateDatastores() bool`

GetCanCreateDatastores returns the CanCreateDatastores field if non-nil, zero value otherwise.

### GetCanCreateDatastoresOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetCanCreateDatastoresOk() (*bool, bool)`

GetCanCreateDatastoresOk returns a tuple with the CanCreateDatastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateDatastores

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetCanCreateDatastores(v bool)`

SetCanCreateDatastores sets CanCreateDatastores field to given value.

### HasCanCreateDatastores

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasCanCreateDatastores() bool`

HasCanCreateDatastores returns a boolean if a field has been set.

### GetCanCreateNetworks

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetCanCreateNetworks() bool`

GetCanCreateNetworks returns the CanCreateNetworks field if non-nil, zero value otherwise.

### GetCanCreateNetworksOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetCanCreateNetworksOk() (*bool, bool)`

GetCanCreateNetworksOk returns a tuple with the CanCreateNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateNetworks

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetCanCreateNetworks(v bool)`

SetCanCreateNetworks sets CanCreateNetworks field to given value.

### HasCanCreateNetworks

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasCanCreateNetworks() bool`

HasCanCreateNetworks returns a boolean if a field has been set.

### GetCanChooseContainerMode

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetCanChooseContainerMode() bool`

GetCanChooseContainerMode returns the CanChooseContainerMode field if non-nil, zero value otherwise.

### GetCanChooseContainerModeOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetCanChooseContainerModeOk() (*bool, bool)`

GetCanChooseContainerModeOk returns a tuple with the CanChooseContainerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanChooseContainerMode

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetCanChooseContainerMode(v bool)`

SetCanChooseContainerMode sets CanChooseContainerMode field to given value.

### HasCanChooseContainerMode

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasCanChooseContainerMode() bool`

HasCanChooseContainerMode returns a boolean if a field has been set.

### GetProvisionRequiresResourcePool

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetProvisionRequiresResourcePool() bool`

GetProvisionRequiresResourcePool returns the ProvisionRequiresResourcePool field if non-nil, zero value otherwise.

### GetProvisionRequiresResourcePoolOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetProvisionRequiresResourcePoolOk() (*bool, bool)`

GetProvisionRequiresResourcePoolOk returns a tuple with the ProvisionRequiresResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionRequiresResourcePool

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetProvisionRequiresResourcePool(v bool)`

SetProvisionRequiresResourcePool sets ProvisionRequiresResourcePool field to given value.

### HasProvisionRequiresResourcePool

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasProvisionRequiresResourcePool() bool`

HasProvisionRequiresResourcePool returns a boolean if a field has been set.

### GetSupportsDistributedWorker

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetSupportsDistributedWorker() bool`

GetSupportsDistributedWorker returns the SupportsDistributedWorker field if non-nil, zero value otherwise.

### GetSupportsDistributedWorkerOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetSupportsDistributedWorkerOk() (*bool, bool)`

GetSupportsDistributedWorkerOk returns a tuple with the SupportsDistributedWorker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDistributedWorker

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetSupportsDistributedWorker(v bool)`

SetSupportsDistributedWorker sets SupportsDistributedWorker field to given value.

### HasSupportsDistributedWorker

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasSupportsDistributedWorker() bool`

HasSupportsDistributedWorker returns a boolean if a field has been set.

### GetCloud

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetProvisionTypes

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetProvisionTypes() []int64`

GetProvisionTypes returns the ProvisionTypes field if non-nil, zero value otherwise.

### GetProvisionTypesOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetProvisionTypesOk() (*[]int64, bool)`

GetProvisionTypesOk returns a tuple with the ProvisionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionTypes

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetProvisionTypes(v []int64)`

SetProvisionTypes sets ProvisionTypes field to given value.

### HasProvisionTypes

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasProvisionTypes() bool`

HasProvisionTypes returns a boolean if a field has been set.

### GetZoneInstanceTypeLayoutId

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetZoneInstanceTypeLayoutId() int64`

GetZoneInstanceTypeLayoutId returns the ZoneInstanceTypeLayoutId field if non-nil, zero value otherwise.

### GetZoneInstanceTypeLayoutIdOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetZoneInstanceTypeLayoutIdOk() (*int64, bool)`

GetZoneInstanceTypeLayoutIdOk returns a tuple with the ZoneInstanceTypeLayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneInstanceTypeLayoutId

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetZoneInstanceTypeLayoutId(v int64)`

SetZoneInstanceTypeLayoutId sets ZoneInstanceTypeLayoutId field to given value.

### HasZoneInstanceTypeLayoutId

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasZoneInstanceTypeLayoutId() bool`

HasZoneInstanceTypeLayoutId returns a boolean if a field has been set.

### GetServerTypes

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetServerTypes() []ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner`

GetServerTypes returns the ServerTypes field if non-nil, zero value otherwise.

### GetServerTypesOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetServerTypesOk() (*[]ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner, bool)`

GetServerTypesOk returns a tuple with the ServerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypes

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetServerTypes(v []ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner)`

SetServerTypes sets ServerTypes field to given value.

### HasServerTypes

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasServerTypes() bool`

HasServerTypes returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetOptionTypes() []ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) GetOptionTypesOk() (*[]ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) SetOptionTypes(v []ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


