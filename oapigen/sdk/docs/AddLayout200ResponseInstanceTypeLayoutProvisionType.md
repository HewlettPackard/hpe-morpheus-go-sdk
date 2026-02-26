# AddLayout200ResponseInstanceTypeLayoutProvisionType

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
**OptionTypes** | Pointer to [**[]AddLayout200ResponseInstanceTypeLayoutProvisionTypeOptionTypesInner**](AddLayout200ResponseInstanceTypeLayoutProvisionTypeOptionTypesInner.md) |  | [optional] 
**CustomOptionTypes** | Pointer to [**[]AddLayout200ResponseInstanceTypeLayoutProvisionTypeCustomOptionTypesInner**](AddLayout200ResponseInstanceTypeLayoutProvisionTypeCustomOptionTypesInner.md) |  | [optional] 
**NetworkTypes** | Pointer to [**[]AddLayout200ResponseInstanceTypeLayoutProvisionTypeNetworkTypesInner**](AddLayout200ResponseInstanceTypeLayoutProvisionTypeNetworkTypesInner.md) |  | [optional] 
**StorageTypes** | Pointer to [**[]AddLayout200ResponseInstanceTypeLayoutProvisionTypeStorageTypesInner**](AddLayout200ResponseInstanceTypeLayoutProvisionTypeStorageTypesInner.md) |  | [optional] 
**RootStorageTypes** | Pointer to [**[]AddLayout200ResponseInstanceTypeLayoutProvisionTypeRootStorageTypesInner**](AddLayout200ResponseInstanceTypeLayoutProvisionTypeRootStorageTypesInner.md) |  | [optional] 
**ControllerTypes** | Pointer to [**[]AddLayout200ResponseInstanceTypeLayoutProvisionTypeControllerTypesInner**](AddLayout200ResponseInstanceTypeLayoutProvisionTypeControllerTypesInner.md) |  | [optional] 
**StorageProfiles** | Pointer to [**[]AddLayout200ResponseInstanceTypeLayoutProvisionTypeStorageProfilesInner**](AddLayout200ResponseInstanceTypeLayoutProvisionTypeStorageProfilesInner.md) |  | [optional] 

## Methods

### NewAddLayout200ResponseInstanceTypeLayoutProvisionType

`func NewAddLayout200ResponseInstanceTypeLayoutProvisionType() *AddLayout200ResponseInstanceTypeLayoutProvisionType`

NewAddLayout200ResponseInstanceTypeLayoutProvisionType instantiates a new AddLayout200ResponseInstanceTypeLayoutProvisionType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLayout200ResponseInstanceTypeLayoutProvisionTypeWithDefaults

`func NewAddLayout200ResponseInstanceTypeLayoutProvisionTypeWithDefaults() *AddLayout200ResponseInstanceTypeLayoutProvisionType`

NewAddLayout200ResponseInstanceTypeLayoutProvisionTypeWithDefaults instantiates a new AddLayout200ResponseInstanceTypeLayoutProvisionType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAclEnabled

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetAclEnabled() bool`

GetAclEnabled returns the AclEnabled field if non-nil, zero value otherwise.

### GetAclEnabledOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetAclEnabledOk() (*bool, bool)`

GetAclEnabledOk returns a tuple with the AclEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclEnabled

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetAclEnabled(v bool)`

SetAclEnabled sets AclEnabled field to given value.

### HasAclEnabled

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasAclEnabled() bool`

HasAclEnabled returns a boolean if a field has been set.

### GetMultiTenant

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetMultiTenant() bool`

GetMultiTenant returns the MultiTenant field if non-nil, zero value otherwise.

### GetMultiTenantOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetMultiTenantOk() (*bool, bool)`

GetMultiTenantOk returns a tuple with the MultiTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenant

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetMultiTenant(v bool)`

SetMultiTenant sets MultiTenant field to given value.

### HasMultiTenant

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasMultiTenant() bool`

HasMultiTenant returns a boolean if a field has been set.

### GetManaged

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetHostNetwork

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHostNetwork() bool`

GetHostNetwork returns the HostNetwork field if non-nil, zero value otherwise.

### GetHostNetworkOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHostNetworkOk() (*bool, bool)`

GetHostNetworkOk returns a tuple with the HostNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNetwork

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetHostNetwork(v bool)`

SetHostNetwork sets HostNetwork field to given value.

### HasHostNetwork

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasHostNetwork() bool`

HasHostNetwork returns a boolean if a field has been set.

### GetCustomSupported

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetCustomSupported() bool`

GetCustomSupported returns the CustomSupported field if non-nil, zero value otherwise.

### GetCustomSupportedOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetCustomSupportedOk() (*bool, bool)`

GetCustomSupportedOk returns a tuple with the CustomSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSupported

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetCustomSupported(v bool)`

SetCustomSupported sets CustomSupported field to given value.

### HasCustomSupported

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasCustomSupported() bool`

HasCustomSupported returns a boolean if a field has been set.

### GetMapPorts

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetMapPorts() bool`

GetMapPorts returns the MapPorts field if non-nil, zero value otherwise.

### GetMapPortsOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetMapPortsOk() (*bool, bool)`

GetMapPortsOk returns a tuple with the MapPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapPorts

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetMapPorts(v bool)`

SetMapPorts sets MapPorts field to given value.

### HasMapPorts

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasMapPorts() bool`

HasMapPorts returns a boolean if a field has been set.

### GetExportServer

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetExportServer() bool`

GetExportServer returns the ExportServer field if non-nil, zero value otherwise.

### GetExportServerOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetExportServerOk() (*bool, bool)`

GetExportServerOk returns a tuple with the ExportServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportServer

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetExportServer(v bool)`

SetExportServer sets ExportServer field to given value.

### HasExportServer

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasExportServer() bool`

HasExportServer returns a boolean if a field has been set.

### GetViewSet

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetViewSet() string`

GetViewSet returns the ViewSet field if non-nil, zero value otherwise.

### GetViewSetOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetViewSetOk() (*string, bool)`

GetViewSetOk returns a tuple with the ViewSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewSet

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetViewSet(v string)`

SetViewSet sets ViewSet field to given value.

### HasViewSet

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasViewSet() bool`

HasViewSet returns a boolean if a field has been set.

### SetViewSetNil

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetViewSetNil(b bool)`

 SetViewSetNil sets the value for ViewSet to be an explicit nil

### UnsetViewSet
`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) UnsetViewSet()`

UnsetViewSet ensures that no value is present for ViewSet, not even an explicit nil
### GetServerType

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetHostType

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### GetAddVolumes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetHasVolumes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasVolumes() bool`

GetHasVolumes returns the HasVolumes field if non-nil, zero value otherwise.

### GetHasVolumesOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasVolumesOk() (*bool, bool)`

GetHasVolumesOk returns a tuple with the HasVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVolumes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetHasVolumes(v bool)`

SetHasVolumes sets HasVolumes field to given value.

### HasHasVolumes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasHasVolumes() bool`

HasHasVolumes returns a boolean if a field has been set.

### GetHasDatastore

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetHasNetworks

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### GetMaxNetworks

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetMaxNetworks() int64`

GetMaxNetworks returns the MaxNetworks field if non-nil, zero value otherwise.

### GetMaxNetworksOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetMaxNetworksOk() (*int64, bool)`

GetMaxNetworksOk returns a tuple with the MaxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworks

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetMaxNetworks(v int64)`

SetMaxNetworks sets MaxNetworks field to given value.

### HasMaxNetworks

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasMaxNetworks() bool`

HasMaxNetworks returns a boolean if a field has been set.

### GetCustomizeVolume

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetCustomizeVolume() bool`

GetCustomizeVolume returns the CustomizeVolume field if non-nil, zero value otherwise.

### GetCustomizeVolumeOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetCustomizeVolumeOk() (*bool, bool)`

GetCustomizeVolumeOk returns a tuple with the CustomizeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizeVolume

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetCustomizeVolume(v bool)`

SetCustomizeVolume sets CustomizeVolume field to given value.

### HasCustomizeVolume

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasCustomizeVolume() bool`

HasCustomizeVolume returns a boolean if a field has been set.

### GetRootDiskCustomizable

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetRootDiskCustomizable() bool`

GetRootDiskCustomizable returns the RootDiskCustomizable field if non-nil, zero value otherwise.

### GetRootDiskCustomizableOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetRootDiskCustomizableOk() (*bool, bool)`

GetRootDiskCustomizableOk returns a tuple with the RootDiskCustomizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskCustomizable

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetRootDiskCustomizable(v bool)`

SetRootDiskCustomizable sets RootDiskCustomizable field to given value.

### HasRootDiskCustomizable

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasRootDiskCustomizable() bool`

HasRootDiskCustomizable returns a boolean if a field has been set.

### GetRootDiskSizeKnown

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetRootDiskSizeKnown() bool`

GetRootDiskSizeKnown returns the RootDiskSizeKnown field if non-nil, zero value otherwise.

### GetRootDiskSizeKnownOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetRootDiskSizeKnownOk() (*bool, bool)`

GetRootDiskSizeKnownOk returns a tuple with the RootDiskSizeKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSizeKnown

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetRootDiskSizeKnown(v bool)`

SetRootDiskSizeKnown sets RootDiskSizeKnown field to given value.

### HasRootDiskSizeKnown

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasRootDiskSizeKnown() bool`

HasRootDiskSizeKnown returns a boolean if a field has been set.

### GetRootDiskResizable

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetRootDiskResizable() bool`

GetRootDiskResizable returns the RootDiskResizable field if non-nil, zero value otherwise.

### GetRootDiskResizableOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetRootDiskResizableOk() (*bool, bool)`

GetRootDiskResizableOk returns a tuple with the RootDiskResizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskResizable

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetRootDiskResizable(v bool)`

SetRootDiskResizable sets RootDiskResizable field to given value.

### HasRootDiskResizable

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasRootDiskResizable() bool`

HasRootDiskResizable returns a boolean if a field has been set.

### GetLvmSupported

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetLvmSupported() bool`

GetLvmSupported returns the LvmSupported field if non-nil, zero value otherwise.

### GetLvmSupportedOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetLvmSupportedOk() (*bool, bool)`

GetLvmSupportedOk returns a tuple with the LvmSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmSupported

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetLvmSupported(v bool)`

SetLvmSupported sets LvmSupported field to given value.

### HasLvmSupported

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasLvmSupported() bool`

HasLvmSupported returns a boolean if a field has been set.

### GetHostDiskMode

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHostDiskMode() string`

GetHostDiskMode returns the HostDiskMode field if non-nil, zero value otherwise.

### GetHostDiskModeOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHostDiskModeOk() (*string, bool)`

GetHostDiskModeOk returns a tuple with the HostDiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDiskMode

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetHostDiskMode(v string)`

SetHostDiskMode sets HostDiskMode field to given value.

### HasHostDiskMode

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasHostDiskMode() bool`

HasHostDiskMode returns a boolean if a field has been set.

### GetMinDisk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetMinDisk() int64`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetMinDiskOk() (*int64, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetMinDisk(v int64)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMaxDisk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetMaxDisk() string`

GetMaxDisk returns the MaxDisk field if non-nil, zero value otherwise.

### GetMaxDiskOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetMaxDiskOk() (*string, bool)`

GetMaxDiskOk returns a tuple with the MaxDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetMaxDisk(v string)`

SetMaxDisk sets MaxDisk field to given value.

### HasMaxDisk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasMaxDisk() bool`

HasMaxDisk returns a boolean if a field has been set.

### SetMaxDiskNil

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetMaxDiskNil(b bool)`

 SetMaxDiskNil sets the value for MaxDisk to be an explicit nil

### UnsetMaxDisk
`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) UnsetMaxDisk()`

UnsetMaxDisk ensures that no value is present for MaxDisk, not even an explicit nil
### GetResizeCopiesVolumes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetResizeCopiesVolumes() bool`

GetResizeCopiesVolumes returns the ResizeCopiesVolumes field if non-nil, zero value otherwise.

### GetResizeCopiesVolumesOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetResizeCopiesVolumesOk() (*bool, bool)`

GetResizeCopiesVolumesOk returns a tuple with the ResizeCopiesVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeCopiesVolumes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetResizeCopiesVolumes(v bool)`

SetResizeCopiesVolumes sets ResizeCopiesVolumes field to given value.

### HasResizeCopiesVolumes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasResizeCopiesVolumes() bool`

HasResizeCopiesVolumes returns a boolean if a field has been set.

### GetSupportsAutoDatastore

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetSupportsAutoDatastore() bool`

GetSupportsAutoDatastore returns the SupportsAutoDatastore field if non-nil, zero value otherwise.

### GetSupportsAutoDatastoreOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetSupportsAutoDatastoreOk() (*bool, bool)`

GetSupportsAutoDatastoreOk returns a tuple with the SupportsAutoDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAutoDatastore

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetSupportsAutoDatastore(v bool)`

SetSupportsAutoDatastore sets SupportsAutoDatastore field to given value.

### HasSupportsAutoDatastore

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasSupportsAutoDatastore() bool`

HasSupportsAutoDatastore returns a boolean if a field has been set.

### GetHasZonePools

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasZonePools() bool`

GetHasZonePools returns the HasZonePools field if non-nil, zero value otherwise.

### GetHasZonePoolsOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasZonePoolsOk() (*bool, bool)`

GetHasZonePoolsOk returns a tuple with the HasZonePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasZonePools

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetHasZonePools(v bool)`

SetHasZonePools sets HasZonePools field to given value.

### HasHasZonePools

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasHasZonePools() bool`

HasHasZonePools returns a boolean if a field has been set.

### GetHasSecurityGroups

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasSecurityGroups() bool`

GetHasSecurityGroups returns the HasSecurityGroups field if non-nil, zero value otherwise.

### GetHasSecurityGroupsOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasSecurityGroupsOk() (*bool, bool)`

GetHasSecurityGroupsOk returns a tuple with the HasSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecurityGroups

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetHasSecurityGroups(v bool)`

SetHasSecurityGroups sets HasSecurityGroups field to given value.

### HasHasSecurityGroups

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasHasSecurityGroups() bool`

HasHasSecurityGroups returns a boolean if a field has been set.

### GetHasParameters

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasParameters() bool`

GetHasParameters returns the HasParameters field if non-nil, zero value otherwise.

### GetHasParametersOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasParametersOk() (*bool, bool)`

GetHasParametersOk returns a tuple with the HasParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasParameters

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetHasParameters(v bool)`

SetHasParameters sets HasParameters field to given value.

### HasHasParameters

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasHasParameters() bool`

HasHasParameters returns a boolean if a field has been set.

### GetCanEnforceTags

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetCanEnforceTags() bool`

GetCanEnforceTags returns the CanEnforceTags field if non-nil, zero value otherwise.

### GetCanEnforceTagsOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetCanEnforceTagsOk() (*bool, bool)`

GetCanEnforceTagsOk returns a tuple with the CanEnforceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEnforceTags

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetCanEnforceTags(v bool)`

SetCanEnforceTags sets CanEnforceTags field to given value.

### HasCanEnforceTags

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasCanEnforceTags() bool`

HasCanEnforceTags returns a boolean if a field has been set.

### SetCanEnforceTagsNil

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetCanEnforceTagsNil(b bool)`

 SetCanEnforceTagsNil sets the value for CanEnforceTags to be an explicit nil

### UnsetCanEnforceTags
`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) UnsetCanEnforceTags()`

UnsetCanEnforceTags ensures that no value is present for CanEnforceTags, not even an explicit nil
### GetDisableRootDatastore

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetDisableRootDatastore() bool`

GetDisableRootDatastore returns the DisableRootDatastore field if non-nil, zero value otherwise.

### GetDisableRootDatastoreOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetDisableRootDatastoreOk() (*bool, bool)`

GetDisableRootDatastoreOk returns a tuple with the DisableRootDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRootDatastore

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetDisableRootDatastore(v bool)`

SetDisableRootDatastore sets DisableRootDatastore field to given value.

### HasDisableRootDatastore

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasDisableRootDatastore() bool`

HasDisableRootDatastore returns a boolean if a field has been set.

### GetHasSnapshots

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasSnapshots() bool`

GetHasSnapshots returns the HasSnapshots field if non-nil, zero value otherwise.

### GetHasSnapshotsOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasSnapshotsOk() (*bool, bool)`

GetHasSnapshotsOk returns a tuple with the HasSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSnapshots

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetHasSnapshots(v bool)`

SetHasSnapshots sets HasSnapshots field to given value.

### HasHasSnapshots

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasHasSnapshots() bool`

HasHasSnapshots returns a boolean if a field has been set.

### GetHasMemorySnapshots

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasMemorySnapshots() bool`

GetHasMemorySnapshots returns the HasMemorySnapshots field if non-nil, zero value otherwise.

### GetHasMemorySnapshotsOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasMemorySnapshotsOk() (*bool, bool)`

GetHasMemorySnapshotsOk returns a tuple with the HasMemorySnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMemorySnapshots

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetHasMemorySnapshots(v bool)`

SetHasMemorySnapshots sets HasMemorySnapshots field to given value.

### HasHasMemorySnapshots

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasHasMemorySnapshots() bool`

HasHasMemorySnapshots returns a boolean if a field has been set.

### GetHasSpecTemplates

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasSpecTemplates() bool`

GetHasSpecTemplates returns the HasSpecTemplates field if non-nil, zero value otherwise.

### GetHasSpecTemplatesOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasSpecTemplatesOk() (*bool, bool)`

GetHasSpecTemplatesOk returns a tuple with the HasSpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSpecTemplates

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetHasSpecTemplates(v bool)`

SetHasSpecTemplates sets HasSpecTemplates field to given value.

### HasHasSpecTemplates

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasHasSpecTemplates() bool`

HasHasSpecTemplates returns a boolean if a field has been set.

### GetHasPreview

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasPreview() bool`

GetHasPreview returns the HasPreview field if non-nil, zero value otherwise.

### GetHasPreviewOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasPreviewOk() (*bool, bool)`

GetHasPreviewOk returns a tuple with the HasPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreview

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetHasPreview(v bool)`

SetHasPreview sets HasPreview field to given value.

### HasHasPreview

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasHasPreview() bool`

HasHasPreview returns a boolean if a field has been set.

### GetZonePoolRequired

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetZonePoolRequired() bool`

GetZonePoolRequired returns the ZonePoolRequired field if non-nil, zero value otherwise.

### GetZonePoolRequiredOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetZonePoolRequiredOk() (*bool, bool)`

GetZonePoolRequiredOk returns a tuple with the ZonePoolRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePoolRequired

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetZonePoolRequired(v bool)`

SetZonePoolRequired sets ZonePoolRequired field to given value.

### HasZonePoolRequired

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasZonePoolRequired() bool`

HasZonePoolRequired returns a boolean if a field has been set.

### GetPlanRequiresPool

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetPlanRequiresPool() bool`

GetPlanRequiresPool returns the PlanRequiresPool field if non-nil, zero value otherwise.

### GetPlanRequiresPoolOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetPlanRequiresPoolOk() (*bool, bool)`

GetPlanRequiresPoolOk returns a tuple with the PlanRequiresPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanRequiresPool

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetPlanRequiresPool(v bool)`

SetPlanRequiresPool sets PlanRequiresPool field to given value.

### HasPlanRequiresPool

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasPlanRequiresPool() bool`

HasPlanRequiresPool returns a boolean if a field has been set.

### GetHasFolders

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasFolders() bool`

GetHasFolders returns the HasFolders field if non-nil, zero value otherwise.

### GetHasFoldersOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetHasFoldersOk() (*bool, bool)`

GetHasFoldersOk returns a tuple with the HasFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFolders

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetHasFolders(v bool)`

SetHasFolders sets HasFolders field to given value.

### HasHasFolders

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasHasFolders() bool`

HasHasFolders returns a boolean if a field has been set.

### SetHasFoldersNil

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetHasFoldersNil(b bool)`

 SetHasFoldersNil sets the value for HasFolders to be an explicit nil

### UnsetHasFolders
`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) UnsetHasFolders()`

UnsetHasFolders ensures that no value is present for HasFolders, not even an explicit nil
### GetOptionTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetOptionTypes() []AddLayout200ResponseInstanceTypeLayoutProvisionTypeOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetOptionTypesOk() (*[]AddLayout200ResponseInstanceTypeLayoutProvisionTypeOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetOptionTypes(v []AddLayout200ResponseInstanceTypeLayoutProvisionTypeOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetCustomOptionTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetCustomOptionTypes() []AddLayout200ResponseInstanceTypeLayoutProvisionTypeCustomOptionTypesInner`

GetCustomOptionTypes returns the CustomOptionTypes field if non-nil, zero value otherwise.

### GetCustomOptionTypesOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetCustomOptionTypesOk() (*[]AddLayout200ResponseInstanceTypeLayoutProvisionTypeCustomOptionTypesInner, bool)`

GetCustomOptionTypesOk returns a tuple with the CustomOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptionTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetCustomOptionTypes(v []AddLayout200ResponseInstanceTypeLayoutProvisionTypeCustomOptionTypesInner)`

SetCustomOptionTypes sets CustomOptionTypes field to given value.

### HasCustomOptionTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasCustomOptionTypes() bool`

HasCustomOptionTypes returns a boolean if a field has been set.

### SetCustomOptionTypesNil

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetCustomOptionTypesNil(b bool)`

 SetCustomOptionTypesNil sets the value for CustomOptionTypes to be an explicit nil

### UnsetCustomOptionTypes
`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) UnsetCustomOptionTypes()`

UnsetCustomOptionTypes ensures that no value is present for CustomOptionTypes, not even an explicit nil
### GetNetworkTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetNetworkTypes() []AddLayout200ResponseInstanceTypeLayoutProvisionTypeNetworkTypesInner`

GetNetworkTypes returns the NetworkTypes field if non-nil, zero value otherwise.

### GetNetworkTypesOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetNetworkTypesOk() (*[]AddLayout200ResponseInstanceTypeLayoutProvisionTypeNetworkTypesInner, bool)`

GetNetworkTypesOk returns a tuple with the NetworkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetNetworkTypes(v []AddLayout200ResponseInstanceTypeLayoutProvisionTypeNetworkTypesInner)`

SetNetworkTypes sets NetworkTypes field to given value.

### HasNetworkTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasNetworkTypes() bool`

HasNetworkTypes returns a boolean if a field has been set.

### SetNetworkTypesNil

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetNetworkTypesNil(b bool)`

 SetNetworkTypesNil sets the value for NetworkTypes to be an explicit nil

### UnsetNetworkTypes
`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) UnsetNetworkTypes()`

UnsetNetworkTypes ensures that no value is present for NetworkTypes, not even an explicit nil
### GetStorageTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetStorageTypes() []AddLayout200ResponseInstanceTypeLayoutProvisionTypeStorageTypesInner`

GetStorageTypes returns the StorageTypes field if non-nil, zero value otherwise.

### GetStorageTypesOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetStorageTypesOk() (*[]AddLayout200ResponseInstanceTypeLayoutProvisionTypeStorageTypesInner, bool)`

GetStorageTypesOk returns a tuple with the StorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetStorageTypes(v []AddLayout200ResponseInstanceTypeLayoutProvisionTypeStorageTypesInner)`

SetStorageTypes sets StorageTypes field to given value.

### HasStorageTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasStorageTypes() bool`

HasStorageTypes returns a boolean if a field has been set.

### SetStorageTypesNil

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetStorageTypesNil(b bool)`

 SetStorageTypesNil sets the value for StorageTypes to be an explicit nil

### UnsetStorageTypes
`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) UnsetStorageTypes()`

UnsetStorageTypes ensures that no value is present for StorageTypes, not even an explicit nil
### GetRootStorageTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetRootStorageTypes() []AddLayout200ResponseInstanceTypeLayoutProvisionTypeRootStorageTypesInner`

GetRootStorageTypes returns the RootStorageTypes field if non-nil, zero value otherwise.

### GetRootStorageTypesOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetRootStorageTypesOk() (*[]AddLayout200ResponseInstanceTypeLayoutProvisionTypeRootStorageTypesInner, bool)`

GetRootStorageTypesOk returns a tuple with the RootStorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootStorageTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetRootStorageTypes(v []AddLayout200ResponseInstanceTypeLayoutProvisionTypeRootStorageTypesInner)`

SetRootStorageTypes sets RootStorageTypes field to given value.

### HasRootStorageTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasRootStorageTypes() bool`

HasRootStorageTypes returns a boolean if a field has been set.

### SetRootStorageTypesNil

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetRootStorageTypesNil(b bool)`

 SetRootStorageTypesNil sets the value for RootStorageTypes to be an explicit nil

### UnsetRootStorageTypes
`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) UnsetRootStorageTypes()`

UnsetRootStorageTypes ensures that no value is present for RootStorageTypes, not even an explicit nil
### GetControllerTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetControllerTypes() []AddLayout200ResponseInstanceTypeLayoutProvisionTypeControllerTypesInner`

GetControllerTypes returns the ControllerTypes field if non-nil, zero value otherwise.

### GetControllerTypesOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetControllerTypesOk() (*[]AddLayout200ResponseInstanceTypeLayoutProvisionTypeControllerTypesInner, bool)`

GetControllerTypesOk returns a tuple with the ControllerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetControllerTypes(v []AddLayout200ResponseInstanceTypeLayoutProvisionTypeControllerTypesInner)`

SetControllerTypes sets ControllerTypes field to given value.

### HasControllerTypes

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasControllerTypes() bool`

HasControllerTypes returns a boolean if a field has been set.

### SetControllerTypesNil

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetControllerTypesNil(b bool)`

 SetControllerTypesNil sets the value for ControllerTypes to be an explicit nil

### UnsetControllerTypes
`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) UnsetControllerTypes()`

UnsetControllerTypes ensures that no value is present for ControllerTypes, not even an explicit nil
### GetStorageProfiles

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetStorageProfiles() []AddLayout200ResponseInstanceTypeLayoutProvisionTypeStorageProfilesInner`

GetStorageProfiles returns the StorageProfiles field if non-nil, zero value otherwise.

### GetStorageProfilesOk

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) GetStorageProfilesOk() (*[]AddLayout200ResponseInstanceTypeLayoutProvisionTypeStorageProfilesInner, bool)`

GetStorageProfilesOk returns a tuple with the StorageProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfiles

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetStorageProfiles(v []AddLayout200ResponseInstanceTypeLayoutProvisionTypeStorageProfilesInner)`

SetStorageProfiles sets StorageProfiles field to given value.

### HasStorageProfiles

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) HasStorageProfiles() bool`

HasStorageProfiles returns a boolean if a field has been set.

### SetStorageProfilesNil

`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) SetStorageProfilesNil(b bool)`

 SetStorageProfilesNil sets the value for StorageProfiles to be an explicit nil

### UnsetStorageProfiles
`func (o *AddLayout200ResponseInstanceTypeLayoutProvisionType) UnsetStorageProfiles()`

UnsetStorageProfiles ensures that no value is present for StorageProfiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


