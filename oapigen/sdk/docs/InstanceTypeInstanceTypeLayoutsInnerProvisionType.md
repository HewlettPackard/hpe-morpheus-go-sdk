# InstanceTypeInstanceTypeLayoutsInnerProvisionType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**AclEnabled** | Pointer to **bool** |  | [optional] 
**MultiTenant** | Pointer to **bool** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**HostNetwork** | Pointer to **bool** |  | [optional] 
**CustomSupported** | Pointer to **bool** |  | [optional] 
**MapPorts** | Pointer to **bool** |  | [optional] 
**ExportServer** | Pointer to **bool** |  | [optional] 
**ViewSet** | Pointer to **NullableString** |  | [optional] 
**ServerType** | Pointer to **string** |  | [optional] 
**HostType** | Pointer to **string** |  | [optional] 
**AddVolumes** | Pointer to **bool** |  | [optional] 
**HasVolumes** | Pointer to **bool** |  | [optional] 
**HasDatastore** | Pointer to **bool** |  | [optional] 
**HasNetworks** | Pointer to **bool** |  | [optional] 
**MaxNetworks** | Pointer to **int64** |  | [optional] 
**CustomizeVolume** | Pointer to **bool** |  | [optional] 
**RootDiskCustomizable** | Pointer to **bool** |  | [optional] 
**RootDiskSizeKnown** | Pointer to **bool** |  | [optional] 
**RootDiskResizable** | Pointer to **bool** |  | [optional] 
**LvmSupported** | Pointer to **bool** |  | [optional] 
**HostDiskMode** | Pointer to **string** |  | [optional] 
**MinDisk** | Pointer to **int64** |  | [optional] 
**MaxDisk** | Pointer to **NullableString** |  | [optional] 
**ResizeCopiesVolumes** | Pointer to **bool** |  | [optional] 
**SupportsAutoDatastore** | Pointer to **bool** |  | [optional] 
**HasZonePools** | Pointer to **bool** |  | [optional] 
**HasSecurityGroups** | Pointer to **bool** |  | [optional] 
**HasParameters** | Pointer to **bool** |  | [optional] 
**CanEnforceTags** | Pointer to **NullableBool** |  | [optional] 
**DisableRootDatastore** | Pointer to **bool** |  | [optional] 
**HasSnapshots** | Pointer to **bool** |  | [optional] 
**HasMemorySnapshots** | Pointer to **bool** |  | [optional] 
**HasSpecTemplates** | Pointer to **bool** |  | [optional] 
**HasPreview** | Pointer to **bool** |  | [optional] 
**ZonePoolRequired** | Pointer to **bool** |  | [optional] 
**PlanRequiresPool** | Pointer to **bool** |  | [optional] 
**HasFolders** | Pointer to **NullableBool** |  | [optional] 
**OptionTypes** | Pointer to [**[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeOptionTypesInner**](InstanceTypeInstanceTypeLayoutsInnerProvisionTypeOptionTypesInner.md) |  | [optional] 
**CustomOptionTypes** | Pointer to [**[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeCustomOptionTypesInner**](InstanceTypeInstanceTypeLayoutsInnerProvisionTypeCustomOptionTypesInner.md) |  | [optional] 
**NetworkTypes** | Pointer to [**[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner**](InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner.md) |  | [optional] 
**StorageTypes** | Pointer to [**[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner**](InstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner.md) |  | [optional] 
**RootStorageTypes** | Pointer to [**[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeRootStorageTypesInner**](InstanceTypeInstanceTypeLayoutsInnerProvisionTypeRootStorageTypesInner.md) |  | [optional] 
**ControllerTypes** | Pointer to [**[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeControllerTypesInner**](InstanceTypeInstanceTypeLayoutsInnerProvisionTypeControllerTypesInner.md) |  | [optional] 
**StorageProfiles** | Pointer to [**[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageProfilesInner**](InstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageProfilesInner.md) |  | [optional] 

## Methods

### NewInstanceTypeInstanceTypeLayoutsInnerProvisionType

`func NewInstanceTypeInstanceTypeLayoutsInnerProvisionType() *InstanceTypeInstanceTypeLayoutsInnerProvisionType`

NewInstanceTypeInstanceTypeLayoutsInnerProvisionType instantiates a new InstanceTypeInstanceTypeLayoutsInnerProvisionType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeInstanceTypeLayoutsInnerProvisionTypeWithDefaults

`func NewInstanceTypeInstanceTypeLayoutsInnerProvisionTypeWithDefaults() *InstanceTypeInstanceTypeLayoutsInnerProvisionType`

NewInstanceTypeInstanceTypeLayoutsInnerProvisionTypeWithDefaults instantiates a new InstanceTypeInstanceTypeLayoutsInnerProvisionType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAclEnabled

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetAclEnabled() bool`

GetAclEnabled returns the AclEnabled field if non-nil, zero value otherwise.

### GetAclEnabledOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetAclEnabledOk() (*bool, bool)`

GetAclEnabledOk returns a tuple with the AclEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclEnabled

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetAclEnabled(v bool)`

SetAclEnabled sets AclEnabled field to given value.

### HasAclEnabled

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasAclEnabled() bool`

HasAclEnabled returns a boolean if a field has been set.

### GetMultiTenant

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMultiTenant() bool`

GetMultiTenant returns the MultiTenant field if non-nil, zero value otherwise.

### GetMultiTenantOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMultiTenantOk() (*bool, bool)`

GetMultiTenantOk returns a tuple with the MultiTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenant

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetMultiTenant(v bool)`

SetMultiTenant sets MultiTenant field to given value.

### HasMultiTenant

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasMultiTenant() bool`

HasMultiTenant returns a boolean if a field has been set.

### GetManaged

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetHostNetwork

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHostNetwork() bool`

GetHostNetwork returns the HostNetwork field if non-nil, zero value otherwise.

### GetHostNetworkOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHostNetworkOk() (*bool, bool)`

GetHostNetworkOk returns a tuple with the HostNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNetwork

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHostNetwork(v bool)`

SetHostNetwork sets HostNetwork field to given value.

### HasHostNetwork

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHostNetwork() bool`

HasHostNetwork returns a boolean if a field has been set.

### GetCustomSupported

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCustomSupported() bool`

GetCustomSupported returns the CustomSupported field if non-nil, zero value otherwise.

### GetCustomSupportedOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCustomSupportedOk() (*bool, bool)`

GetCustomSupportedOk returns a tuple with the CustomSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSupported

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetCustomSupported(v bool)`

SetCustomSupported sets CustomSupported field to given value.

### HasCustomSupported

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasCustomSupported() bool`

HasCustomSupported returns a boolean if a field has been set.

### GetMapPorts

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMapPorts() bool`

GetMapPorts returns the MapPorts field if non-nil, zero value otherwise.

### GetMapPortsOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMapPortsOk() (*bool, bool)`

GetMapPortsOk returns a tuple with the MapPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapPorts

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetMapPorts(v bool)`

SetMapPorts sets MapPorts field to given value.

### HasMapPorts

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasMapPorts() bool`

HasMapPorts returns a boolean if a field has been set.

### GetExportServer

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetExportServer() bool`

GetExportServer returns the ExportServer field if non-nil, zero value otherwise.

### GetExportServerOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetExportServerOk() (*bool, bool)`

GetExportServerOk returns a tuple with the ExportServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportServer

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetExportServer(v bool)`

SetExportServer sets ExportServer field to given value.

### HasExportServer

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasExportServer() bool`

HasExportServer returns a boolean if a field has been set.

### GetViewSet

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetViewSet() string`

GetViewSet returns the ViewSet field if non-nil, zero value otherwise.

### GetViewSetOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetViewSetOk() (*string, bool)`

GetViewSetOk returns a tuple with the ViewSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewSet

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetViewSet(v string)`

SetViewSet sets ViewSet field to given value.

### HasViewSet

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasViewSet() bool`

HasViewSet returns a boolean if a field has been set.

### SetViewSetNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetViewSetNil(b bool)`

 SetViewSetNil sets the value for ViewSet to be an explicit nil

### UnsetViewSet
`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetViewSet()`

UnsetViewSet ensures that no value is present for ViewSet, not even an explicit nil
### GetServerType

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetHostType

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### GetAddVolumes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetHasVolumes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasVolumes() bool`

GetHasVolumes returns the HasVolumes field if non-nil, zero value otherwise.

### GetHasVolumesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasVolumesOk() (*bool, bool)`

GetHasVolumesOk returns a tuple with the HasVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVolumes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasVolumes(v bool)`

SetHasVolumes sets HasVolumes field to given value.

### HasHasVolumes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasVolumes() bool`

HasHasVolumes returns a boolean if a field has been set.

### GetHasDatastore

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetHasNetworks

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### GetMaxNetworks

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMaxNetworks() int64`

GetMaxNetworks returns the MaxNetworks field if non-nil, zero value otherwise.

### GetMaxNetworksOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMaxNetworksOk() (*int64, bool)`

GetMaxNetworksOk returns a tuple with the MaxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworks

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetMaxNetworks(v int64)`

SetMaxNetworks sets MaxNetworks field to given value.

### HasMaxNetworks

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasMaxNetworks() bool`

HasMaxNetworks returns a boolean if a field has been set.

### GetCustomizeVolume

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCustomizeVolume() bool`

GetCustomizeVolume returns the CustomizeVolume field if non-nil, zero value otherwise.

### GetCustomizeVolumeOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCustomizeVolumeOk() (*bool, bool)`

GetCustomizeVolumeOk returns a tuple with the CustomizeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizeVolume

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetCustomizeVolume(v bool)`

SetCustomizeVolume sets CustomizeVolume field to given value.

### HasCustomizeVolume

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasCustomizeVolume() bool`

HasCustomizeVolume returns a boolean if a field has been set.

### GetRootDiskCustomizable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootDiskCustomizable() bool`

GetRootDiskCustomizable returns the RootDiskCustomizable field if non-nil, zero value otherwise.

### GetRootDiskCustomizableOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootDiskCustomizableOk() (*bool, bool)`

GetRootDiskCustomizableOk returns a tuple with the RootDiskCustomizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskCustomizable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetRootDiskCustomizable(v bool)`

SetRootDiskCustomizable sets RootDiskCustomizable field to given value.

### HasRootDiskCustomizable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasRootDiskCustomizable() bool`

HasRootDiskCustomizable returns a boolean if a field has been set.

### GetRootDiskSizeKnown

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootDiskSizeKnown() bool`

GetRootDiskSizeKnown returns the RootDiskSizeKnown field if non-nil, zero value otherwise.

### GetRootDiskSizeKnownOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootDiskSizeKnownOk() (*bool, bool)`

GetRootDiskSizeKnownOk returns a tuple with the RootDiskSizeKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSizeKnown

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetRootDiskSizeKnown(v bool)`

SetRootDiskSizeKnown sets RootDiskSizeKnown field to given value.

### HasRootDiskSizeKnown

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasRootDiskSizeKnown() bool`

HasRootDiskSizeKnown returns a boolean if a field has been set.

### GetRootDiskResizable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootDiskResizable() bool`

GetRootDiskResizable returns the RootDiskResizable field if non-nil, zero value otherwise.

### GetRootDiskResizableOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootDiskResizableOk() (*bool, bool)`

GetRootDiskResizableOk returns a tuple with the RootDiskResizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskResizable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetRootDiskResizable(v bool)`

SetRootDiskResizable sets RootDiskResizable field to given value.

### HasRootDiskResizable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasRootDiskResizable() bool`

HasRootDiskResizable returns a boolean if a field has been set.

### GetLvmSupported

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetLvmSupported() bool`

GetLvmSupported returns the LvmSupported field if non-nil, zero value otherwise.

### GetLvmSupportedOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetLvmSupportedOk() (*bool, bool)`

GetLvmSupportedOk returns a tuple with the LvmSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmSupported

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetLvmSupported(v bool)`

SetLvmSupported sets LvmSupported field to given value.

### HasLvmSupported

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasLvmSupported() bool`

HasLvmSupported returns a boolean if a field has been set.

### GetHostDiskMode

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHostDiskMode() string`

GetHostDiskMode returns the HostDiskMode field if non-nil, zero value otherwise.

### GetHostDiskModeOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHostDiskModeOk() (*string, bool)`

GetHostDiskModeOk returns a tuple with the HostDiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDiskMode

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHostDiskMode(v string)`

SetHostDiskMode sets HostDiskMode field to given value.

### HasHostDiskMode

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHostDiskMode() bool`

HasHostDiskMode returns a boolean if a field has been set.

### GetMinDisk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMinDisk() int64`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMinDiskOk() (*int64, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetMinDisk(v int64)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMaxDisk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMaxDisk() string`

GetMaxDisk returns the MaxDisk field if non-nil, zero value otherwise.

### GetMaxDiskOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMaxDiskOk() (*string, bool)`

GetMaxDiskOk returns a tuple with the MaxDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetMaxDisk(v string)`

SetMaxDisk sets MaxDisk field to given value.

### HasMaxDisk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasMaxDisk() bool`

HasMaxDisk returns a boolean if a field has been set.

### SetMaxDiskNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetMaxDiskNil(b bool)`

 SetMaxDiskNil sets the value for MaxDisk to be an explicit nil

### UnsetMaxDisk
`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetMaxDisk()`

UnsetMaxDisk ensures that no value is present for MaxDisk, not even an explicit nil
### GetResizeCopiesVolumes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetResizeCopiesVolumes() bool`

GetResizeCopiesVolumes returns the ResizeCopiesVolumes field if non-nil, zero value otherwise.

### GetResizeCopiesVolumesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetResizeCopiesVolumesOk() (*bool, bool)`

GetResizeCopiesVolumesOk returns a tuple with the ResizeCopiesVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeCopiesVolumes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetResizeCopiesVolumes(v bool)`

SetResizeCopiesVolumes sets ResizeCopiesVolumes field to given value.

### HasResizeCopiesVolumes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasResizeCopiesVolumes() bool`

HasResizeCopiesVolumes returns a boolean if a field has been set.

### GetSupportsAutoDatastore

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetSupportsAutoDatastore() bool`

GetSupportsAutoDatastore returns the SupportsAutoDatastore field if non-nil, zero value otherwise.

### GetSupportsAutoDatastoreOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetSupportsAutoDatastoreOk() (*bool, bool)`

GetSupportsAutoDatastoreOk returns a tuple with the SupportsAutoDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAutoDatastore

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetSupportsAutoDatastore(v bool)`

SetSupportsAutoDatastore sets SupportsAutoDatastore field to given value.

### HasSupportsAutoDatastore

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasSupportsAutoDatastore() bool`

HasSupportsAutoDatastore returns a boolean if a field has been set.

### GetHasZonePools

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasZonePools() bool`

GetHasZonePools returns the HasZonePools field if non-nil, zero value otherwise.

### GetHasZonePoolsOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasZonePoolsOk() (*bool, bool)`

GetHasZonePoolsOk returns a tuple with the HasZonePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasZonePools

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasZonePools(v bool)`

SetHasZonePools sets HasZonePools field to given value.

### HasHasZonePools

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasZonePools() bool`

HasHasZonePools returns a boolean if a field has been set.

### GetHasSecurityGroups

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasSecurityGroups() bool`

GetHasSecurityGroups returns the HasSecurityGroups field if non-nil, zero value otherwise.

### GetHasSecurityGroupsOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasSecurityGroupsOk() (*bool, bool)`

GetHasSecurityGroupsOk returns a tuple with the HasSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecurityGroups

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasSecurityGroups(v bool)`

SetHasSecurityGroups sets HasSecurityGroups field to given value.

### HasHasSecurityGroups

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasSecurityGroups() bool`

HasHasSecurityGroups returns a boolean if a field has been set.

### GetHasParameters

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasParameters() bool`

GetHasParameters returns the HasParameters field if non-nil, zero value otherwise.

### GetHasParametersOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasParametersOk() (*bool, bool)`

GetHasParametersOk returns a tuple with the HasParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasParameters

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasParameters(v bool)`

SetHasParameters sets HasParameters field to given value.

### HasHasParameters

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasParameters() bool`

HasHasParameters returns a boolean if a field has been set.

### GetCanEnforceTags

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCanEnforceTags() bool`

GetCanEnforceTags returns the CanEnforceTags field if non-nil, zero value otherwise.

### GetCanEnforceTagsOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCanEnforceTagsOk() (*bool, bool)`

GetCanEnforceTagsOk returns a tuple with the CanEnforceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEnforceTags

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetCanEnforceTags(v bool)`

SetCanEnforceTags sets CanEnforceTags field to given value.

### HasCanEnforceTags

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasCanEnforceTags() bool`

HasCanEnforceTags returns a boolean if a field has been set.

### SetCanEnforceTagsNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetCanEnforceTagsNil(b bool)`

 SetCanEnforceTagsNil sets the value for CanEnforceTags to be an explicit nil

### UnsetCanEnforceTags
`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetCanEnforceTags()`

UnsetCanEnforceTags ensures that no value is present for CanEnforceTags, not even an explicit nil
### GetDisableRootDatastore

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetDisableRootDatastore() bool`

GetDisableRootDatastore returns the DisableRootDatastore field if non-nil, zero value otherwise.

### GetDisableRootDatastoreOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetDisableRootDatastoreOk() (*bool, bool)`

GetDisableRootDatastoreOk returns a tuple with the DisableRootDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRootDatastore

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetDisableRootDatastore(v bool)`

SetDisableRootDatastore sets DisableRootDatastore field to given value.

### HasDisableRootDatastore

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasDisableRootDatastore() bool`

HasDisableRootDatastore returns a boolean if a field has been set.

### GetHasSnapshots

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasSnapshots() bool`

GetHasSnapshots returns the HasSnapshots field if non-nil, zero value otherwise.

### GetHasSnapshotsOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasSnapshotsOk() (*bool, bool)`

GetHasSnapshotsOk returns a tuple with the HasSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSnapshots

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasSnapshots(v bool)`

SetHasSnapshots sets HasSnapshots field to given value.

### HasHasSnapshots

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasSnapshots() bool`

HasHasSnapshots returns a boolean if a field has been set.

### GetHasMemorySnapshots

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasMemorySnapshots() bool`

GetHasMemorySnapshots returns the HasMemorySnapshots field if non-nil, zero value otherwise.

### GetHasMemorySnapshotsOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasMemorySnapshotsOk() (*bool, bool)`

GetHasMemorySnapshotsOk returns a tuple with the HasMemorySnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMemorySnapshots

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasMemorySnapshots(v bool)`

SetHasMemorySnapshots sets HasMemorySnapshots field to given value.

### HasHasMemorySnapshots

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasMemorySnapshots() bool`

HasHasMemorySnapshots returns a boolean if a field has been set.

### GetHasSpecTemplates

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasSpecTemplates() bool`

GetHasSpecTemplates returns the HasSpecTemplates field if non-nil, zero value otherwise.

### GetHasSpecTemplatesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasSpecTemplatesOk() (*bool, bool)`

GetHasSpecTemplatesOk returns a tuple with the HasSpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSpecTemplates

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasSpecTemplates(v bool)`

SetHasSpecTemplates sets HasSpecTemplates field to given value.

### HasHasSpecTemplates

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasSpecTemplates() bool`

HasHasSpecTemplates returns a boolean if a field has been set.

### GetHasPreview

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasPreview() bool`

GetHasPreview returns the HasPreview field if non-nil, zero value otherwise.

### GetHasPreviewOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasPreviewOk() (*bool, bool)`

GetHasPreviewOk returns a tuple with the HasPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreview

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasPreview(v bool)`

SetHasPreview sets HasPreview field to given value.

### HasHasPreview

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasPreview() bool`

HasHasPreview returns a boolean if a field has been set.

### GetZonePoolRequired

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetZonePoolRequired() bool`

GetZonePoolRequired returns the ZonePoolRequired field if non-nil, zero value otherwise.

### GetZonePoolRequiredOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetZonePoolRequiredOk() (*bool, bool)`

GetZonePoolRequiredOk returns a tuple with the ZonePoolRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePoolRequired

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetZonePoolRequired(v bool)`

SetZonePoolRequired sets ZonePoolRequired field to given value.

### HasZonePoolRequired

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasZonePoolRequired() bool`

HasZonePoolRequired returns a boolean if a field has been set.

### GetPlanRequiresPool

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetPlanRequiresPool() bool`

GetPlanRequiresPool returns the PlanRequiresPool field if non-nil, zero value otherwise.

### GetPlanRequiresPoolOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetPlanRequiresPoolOk() (*bool, bool)`

GetPlanRequiresPoolOk returns a tuple with the PlanRequiresPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanRequiresPool

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetPlanRequiresPool(v bool)`

SetPlanRequiresPool sets PlanRequiresPool field to given value.

### HasPlanRequiresPool

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasPlanRequiresPool() bool`

HasPlanRequiresPool returns a boolean if a field has been set.

### GetHasFolders

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasFolders() bool`

GetHasFolders returns the HasFolders field if non-nil, zero value otherwise.

### GetHasFoldersOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasFoldersOk() (*bool, bool)`

GetHasFoldersOk returns a tuple with the HasFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFolders

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasFolders(v bool)`

SetHasFolders sets HasFolders field to given value.

### HasHasFolders

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasFolders() bool`

HasHasFolders returns a boolean if a field has been set.

### SetHasFoldersNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasFoldersNil(b bool)`

 SetHasFoldersNil sets the value for HasFolders to be an explicit nil

### UnsetHasFolders
`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetHasFolders()`

UnsetHasFolders ensures that no value is present for HasFolders, not even an explicit nil
### GetOptionTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetOptionTypes() []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetOptionTypesOk() (*[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetOptionTypes(v []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetCustomOptionTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCustomOptionTypes() []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeCustomOptionTypesInner`

GetCustomOptionTypes returns the CustomOptionTypes field if non-nil, zero value otherwise.

### GetCustomOptionTypesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCustomOptionTypesOk() (*[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeCustomOptionTypesInner, bool)`

GetCustomOptionTypesOk returns a tuple with the CustomOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptionTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetCustomOptionTypes(v []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeCustomOptionTypesInner)`

SetCustomOptionTypes sets CustomOptionTypes field to given value.

### HasCustomOptionTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasCustomOptionTypes() bool`

HasCustomOptionTypes returns a boolean if a field has been set.

### SetCustomOptionTypesNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetCustomOptionTypesNil(b bool)`

 SetCustomOptionTypesNil sets the value for CustomOptionTypes to be an explicit nil

### UnsetCustomOptionTypes
`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetCustomOptionTypes()`

UnsetCustomOptionTypes ensures that no value is present for CustomOptionTypes, not even an explicit nil
### GetNetworkTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetNetworkTypes() []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner`

GetNetworkTypes returns the NetworkTypes field if non-nil, zero value otherwise.

### GetNetworkTypesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetNetworkTypesOk() (*[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner, bool)`

GetNetworkTypesOk returns a tuple with the NetworkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetNetworkTypes(v []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner)`

SetNetworkTypes sets NetworkTypes field to given value.

### HasNetworkTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasNetworkTypes() bool`

HasNetworkTypes returns a boolean if a field has been set.

### SetNetworkTypesNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetNetworkTypesNil(b bool)`

 SetNetworkTypesNil sets the value for NetworkTypes to be an explicit nil

### UnsetNetworkTypes
`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetNetworkTypes()`

UnsetNetworkTypes ensures that no value is present for NetworkTypes, not even an explicit nil
### GetStorageTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetStorageTypes() []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner`

GetStorageTypes returns the StorageTypes field if non-nil, zero value otherwise.

### GetStorageTypesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetStorageTypesOk() (*[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner, bool)`

GetStorageTypesOk returns a tuple with the StorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetStorageTypes(v []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner)`

SetStorageTypes sets StorageTypes field to given value.

### HasStorageTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasStorageTypes() bool`

HasStorageTypes returns a boolean if a field has been set.

### SetStorageTypesNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetStorageTypesNil(b bool)`

 SetStorageTypesNil sets the value for StorageTypes to be an explicit nil

### UnsetStorageTypes
`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetStorageTypes()`

UnsetStorageTypes ensures that no value is present for StorageTypes, not even an explicit nil
### GetRootStorageTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootStorageTypes() []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeRootStorageTypesInner`

GetRootStorageTypes returns the RootStorageTypes field if non-nil, zero value otherwise.

### GetRootStorageTypesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootStorageTypesOk() (*[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeRootStorageTypesInner, bool)`

GetRootStorageTypesOk returns a tuple with the RootStorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootStorageTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetRootStorageTypes(v []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeRootStorageTypesInner)`

SetRootStorageTypes sets RootStorageTypes field to given value.

### HasRootStorageTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasRootStorageTypes() bool`

HasRootStorageTypes returns a boolean if a field has been set.

### SetRootStorageTypesNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetRootStorageTypesNil(b bool)`

 SetRootStorageTypesNil sets the value for RootStorageTypes to be an explicit nil

### UnsetRootStorageTypes
`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetRootStorageTypes()`

UnsetRootStorageTypes ensures that no value is present for RootStorageTypes, not even an explicit nil
### GetControllerTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetControllerTypes() []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeControllerTypesInner`

GetControllerTypes returns the ControllerTypes field if non-nil, zero value otherwise.

### GetControllerTypesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetControllerTypesOk() (*[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeControllerTypesInner, bool)`

GetControllerTypesOk returns a tuple with the ControllerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetControllerTypes(v []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeControllerTypesInner)`

SetControllerTypes sets ControllerTypes field to given value.

### HasControllerTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasControllerTypes() bool`

HasControllerTypes returns a boolean if a field has been set.

### SetControllerTypesNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetControllerTypesNil(b bool)`

 SetControllerTypesNil sets the value for ControllerTypes to be an explicit nil

### UnsetControllerTypes
`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetControllerTypes()`

UnsetControllerTypes ensures that no value is present for ControllerTypes, not even an explicit nil
### GetStorageProfiles

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetStorageProfiles() []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageProfilesInner`

GetStorageProfiles returns the StorageProfiles field if non-nil, zero value otherwise.

### GetStorageProfilesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) GetStorageProfilesOk() (*[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageProfilesInner, bool)`

GetStorageProfilesOk returns a tuple with the StorageProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfiles

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetStorageProfiles(v []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageProfilesInner)`

SetStorageProfiles sets StorageProfiles field to given value.

### HasStorageProfiles

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) HasStorageProfiles() bool`

HasStorageProfiles returns a boolean if a field has been set.

### SetStorageProfilesNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) SetStorageProfilesNil(b bool)`

 SetStorageProfilesNil sets the value for StorageProfiles to be an explicit nil

### UnsetStorageProfiles
`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetStorageProfiles()`

UnsetStorageProfiles ensures that no value is present for StorageProfiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


