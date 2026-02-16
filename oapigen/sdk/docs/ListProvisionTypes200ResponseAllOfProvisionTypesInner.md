# ListProvisionTypes200ResponseAllOfProvisionTypesInner

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
**OptionTypes** | Pointer to [**[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerOptionTypesInner**](ListProvisionTypes200ResponseAllOfProvisionTypesInnerOptionTypesInner.md) |  | [optional] 
**CustomOptionTypes** | Pointer to [**[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerCustomOptionTypesInner**](ListProvisionTypes200ResponseAllOfProvisionTypesInnerCustomOptionTypesInner.md) |  | [optional] 
**NetworkTypes** | Pointer to [**[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner**](ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner.md) |  | [optional] 
**StorageTypes** | Pointer to [**[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerStorageTypesInner**](ListProvisionTypes200ResponseAllOfProvisionTypesInnerStorageTypesInner.md) |  | [optional] 
**RootStorageTypes** | Pointer to [**[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerRootStorageTypesInner**](ListProvisionTypes200ResponseAllOfProvisionTypesInnerRootStorageTypesInner.md) |  | [optional] 
**ControllerTypes** | Pointer to [**[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerControllerTypesInner**](ListProvisionTypes200ResponseAllOfProvisionTypesInnerControllerTypesInner.md) |  | [optional] 

## Methods

### NewListProvisionTypes200ResponseAllOfProvisionTypesInner

`func NewListProvisionTypes200ResponseAllOfProvisionTypesInner() *ListProvisionTypes200ResponseAllOfProvisionTypesInner`

NewListProvisionTypes200ResponseAllOfProvisionTypesInner instantiates a new ListProvisionTypes200ResponseAllOfProvisionTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProvisionTypes200ResponseAllOfProvisionTypesInnerWithDefaults

`func NewListProvisionTypes200ResponseAllOfProvisionTypesInnerWithDefaults() *ListProvisionTypes200ResponseAllOfProvisionTypesInner`

NewListProvisionTypes200ResponseAllOfProvisionTypesInnerWithDefaults instantiates a new ListProvisionTypes200ResponseAllOfProvisionTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAclEnabled

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetAclEnabled() bool`

GetAclEnabled returns the AclEnabled field if non-nil, zero value otherwise.

### GetAclEnabledOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetAclEnabledOk() (*bool, bool)`

GetAclEnabledOk returns a tuple with the AclEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclEnabled

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetAclEnabled(v bool)`

SetAclEnabled sets AclEnabled field to given value.

### HasAclEnabled

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasAclEnabled() bool`

HasAclEnabled returns a boolean if a field has been set.

### GetMultiTenant

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetMultiTenant() bool`

GetMultiTenant returns the MultiTenant field if non-nil, zero value otherwise.

### GetMultiTenantOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetMultiTenantOk() (*bool, bool)`

GetMultiTenantOk returns a tuple with the MultiTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenant

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetMultiTenant(v bool)`

SetMultiTenant sets MultiTenant field to given value.

### HasMultiTenant

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasMultiTenant() bool`

HasMultiTenant returns a boolean if a field has been set.

### GetManaged

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetHostNetwork

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHostNetwork() bool`

GetHostNetwork returns the HostNetwork field if non-nil, zero value otherwise.

### GetHostNetworkOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHostNetworkOk() (*bool, bool)`

GetHostNetworkOk returns a tuple with the HostNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNetwork

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetHostNetwork(v bool)`

SetHostNetwork sets HostNetwork field to given value.

### HasHostNetwork

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasHostNetwork() bool`

HasHostNetwork returns a boolean if a field has been set.

### GetCustomSupported

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetCustomSupported() bool`

GetCustomSupported returns the CustomSupported field if non-nil, zero value otherwise.

### GetCustomSupportedOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetCustomSupportedOk() (*bool, bool)`

GetCustomSupportedOk returns a tuple with the CustomSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSupported

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetCustomSupported(v bool)`

SetCustomSupported sets CustomSupported field to given value.

### HasCustomSupported

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasCustomSupported() bool`

HasCustomSupported returns a boolean if a field has been set.

### GetMapPorts

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetMapPorts() bool`

GetMapPorts returns the MapPorts field if non-nil, zero value otherwise.

### GetMapPortsOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetMapPortsOk() (*bool, bool)`

GetMapPortsOk returns a tuple with the MapPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapPorts

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetMapPorts(v bool)`

SetMapPorts sets MapPorts field to given value.

### HasMapPorts

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasMapPorts() bool`

HasMapPorts returns a boolean if a field has been set.

### GetExportServer

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetExportServer() bool`

GetExportServer returns the ExportServer field if non-nil, zero value otherwise.

### GetExportServerOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetExportServerOk() (*bool, bool)`

GetExportServerOk returns a tuple with the ExportServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportServer

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetExportServer(v bool)`

SetExportServer sets ExportServer field to given value.

### HasExportServer

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasExportServer() bool`

HasExportServer returns a boolean if a field has been set.

### GetViewSet

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetViewSet() string`

GetViewSet returns the ViewSet field if non-nil, zero value otherwise.

### GetViewSetOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetViewSetOk() (*string, bool)`

GetViewSetOk returns a tuple with the ViewSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewSet

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetViewSet(v string)`

SetViewSet sets ViewSet field to given value.

### HasViewSet

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasViewSet() bool`

HasViewSet returns a boolean if a field has been set.

### SetViewSetNil

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetViewSetNil(b bool)`

 SetViewSetNil sets the value for ViewSet to be an explicit nil

### UnsetViewSet
`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) UnsetViewSet()`

UnsetViewSet ensures that no value is present for ViewSet, not even an explicit nil
### GetServerType

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetHostType

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### GetAddVolumes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetHasVolumes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasVolumes() bool`

GetHasVolumes returns the HasVolumes field if non-nil, zero value otherwise.

### GetHasVolumesOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasVolumesOk() (*bool, bool)`

GetHasVolumesOk returns a tuple with the HasVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVolumes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetHasVolumes(v bool)`

SetHasVolumes sets HasVolumes field to given value.

### HasHasVolumes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasHasVolumes() bool`

HasHasVolumes returns a boolean if a field has been set.

### GetHasDatastore

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetHasNetworks

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### GetMaxNetworks

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetMaxNetworks() int64`

GetMaxNetworks returns the MaxNetworks field if non-nil, zero value otherwise.

### GetMaxNetworksOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetMaxNetworksOk() (*int64, bool)`

GetMaxNetworksOk returns a tuple with the MaxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworks

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetMaxNetworks(v int64)`

SetMaxNetworks sets MaxNetworks field to given value.

### HasMaxNetworks

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasMaxNetworks() bool`

HasMaxNetworks returns a boolean if a field has been set.

### GetCustomizeVolume

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetCustomizeVolume() bool`

GetCustomizeVolume returns the CustomizeVolume field if non-nil, zero value otherwise.

### GetCustomizeVolumeOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetCustomizeVolumeOk() (*bool, bool)`

GetCustomizeVolumeOk returns a tuple with the CustomizeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizeVolume

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetCustomizeVolume(v bool)`

SetCustomizeVolume sets CustomizeVolume field to given value.

### HasCustomizeVolume

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasCustomizeVolume() bool`

HasCustomizeVolume returns a boolean if a field has been set.

### GetRootDiskCustomizable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetRootDiskCustomizable() bool`

GetRootDiskCustomizable returns the RootDiskCustomizable field if non-nil, zero value otherwise.

### GetRootDiskCustomizableOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetRootDiskCustomizableOk() (*bool, bool)`

GetRootDiskCustomizableOk returns a tuple with the RootDiskCustomizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskCustomizable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetRootDiskCustomizable(v bool)`

SetRootDiskCustomizable sets RootDiskCustomizable field to given value.

### HasRootDiskCustomizable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasRootDiskCustomizable() bool`

HasRootDiskCustomizable returns a boolean if a field has been set.

### GetRootDiskSizeKnown

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetRootDiskSizeKnown() bool`

GetRootDiskSizeKnown returns the RootDiskSizeKnown field if non-nil, zero value otherwise.

### GetRootDiskSizeKnownOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetRootDiskSizeKnownOk() (*bool, bool)`

GetRootDiskSizeKnownOk returns a tuple with the RootDiskSizeKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSizeKnown

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetRootDiskSizeKnown(v bool)`

SetRootDiskSizeKnown sets RootDiskSizeKnown field to given value.

### HasRootDiskSizeKnown

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasRootDiskSizeKnown() bool`

HasRootDiskSizeKnown returns a boolean if a field has been set.

### GetRootDiskResizable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetRootDiskResizable() bool`

GetRootDiskResizable returns the RootDiskResizable field if non-nil, zero value otherwise.

### GetRootDiskResizableOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetRootDiskResizableOk() (*bool, bool)`

GetRootDiskResizableOk returns a tuple with the RootDiskResizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskResizable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetRootDiskResizable(v bool)`

SetRootDiskResizable sets RootDiskResizable field to given value.

### HasRootDiskResizable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasRootDiskResizable() bool`

HasRootDiskResizable returns a boolean if a field has been set.

### GetLvmSupported

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetLvmSupported() bool`

GetLvmSupported returns the LvmSupported field if non-nil, zero value otherwise.

### GetLvmSupportedOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetLvmSupportedOk() (*bool, bool)`

GetLvmSupportedOk returns a tuple with the LvmSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmSupported

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetLvmSupported(v bool)`

SetLvmSupported sets LvmSupported field to given value.

### HasLvmSupported

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasLvmSupported() bool`

HasLvmSupported returns a boolean if a field has been set.

### GetHostDiskMode

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHostDiskMode() string`

GetHostDiskMode returns the HostDiskMode field if non-nil, zero value otherwise.

### GetHostDiskModeOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHostDiskModeOk() (*string, bool)`

GetHostDiskModeOk returns a tuple with the HostDiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDiskMode

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetHostDiskMode(v string)`

SetHostDiskMode sets HostDiskMode field to given value.

### HasHostDiskMode

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasHostDiskMode() bool`

HasHostDiskMode returns a boolean if a field has been set.

### GetMinDisk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetMinDisk() int64`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetMinDiskOk() (*int64, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetMinDisk(v int64)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMaxDisk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetMaxDisk() string`

GetMaxDisk returns the MaxDisk field if non-nil, zero value otherwise.

### GetMaxDiskOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetMaxDiskOk() (*string, bool)`

GetMaxDiskOk returns a tuple with the MaxDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetMaxDisk(v string)`

SetMaxDisk sets MaxDisk field to given value.

### HasMaxDisk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasMaxDisk() bool`

HasMaxDisk returns a boolean if a field has been set.

### SetMaxDiskNil

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetMaxDiskNil(b bool)`

 SetMaxDiskNil sets the value for MaxDisk to be an explicit nil

### UnsetMaxDisk
`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) UnsetMaxDisk()`

UnsetMaxDisk ensures that no value is present for MaxDisk, not even an explicit nil
### GetResizeCopiesVolumes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetResizeCopiesVolumes() bool`

GetResizeCopiesVolumes returns the ResizeCopiesVolumes field if non-nil, zero value otherwise.

### GetResizeCopiesVolumesOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetResizeCopiesVolumesOk() (*bool, bool)`

GetResizeCopiesVolumesOk returns a tuple with the ResizeCopiesVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeCopiesVolumes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetResizeCopiesVolumes(v bool)`

SetResizeCopiesVolumes sets ResizeCopiesVolumes field to given value.

### HasResizeCopiesVolumes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasResizeCopiesVolumes() bool`

HasResizeCopiesVolumes returns a boolean if a field has been set.

### GetSupportsAutoDatastore

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetSupportsAutoDatastore() bool`

GetSupportsAutoDatastore returns the SupportsAutoDatastore field if non-nil, zero value otherwise.

### GetSupportsAutoDatastoreOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetSupportsAutoDatastoreOk() (*bool, bool)`

GetSupportsAutoDatastoreOk returns a tuple with the SupportsAutoDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAutoDatastore

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetSupportsAutoDatastore(v bool)`

SetSupportsAutoDatastore sets SupportsAutoDatastore field to given value.

### HasSupportsAutoDatastore

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasSupportsAutoDatastore() bool`

HasSupportsAutoDatastore returns a boolean if a field has been set.

### GetHasZonePools

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasZonePools() bool`

GetHasZonePools returns the HasZonePools field if non-nil, zero value otherwise.

### GetHasZonePoolsOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasZonePoolsOk() (*bool, bool)`

GetHasZonePoolsOk returns a tuple with the HasZonePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasZonePools

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetHasZonePools(v bool)`

SetHasZonePools sets HasZonePools field to given value.

### HasHasZonePools

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasHasZonePools() bool`

HasHasZonePools returns a boolean if a field has been set.

### GetHasSecurityGroups

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasSecurityGroups() bool`

GetHasSecurityGroups returns the HasSecurityGroups field if non-nil, zero value otherwise.

### GetHasSecurityGroupsOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasSecurityGroupsOk() (*bool, bool)`

GetHasSecurityGroupsOk returns a tuple with the HasSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecurityGroups

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetHasSecurityGroups(v bool)`

SetHasSecurityGroups sets HasSecurityGroups field to given value.

### HasHasSecurityGroups

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasHasSecurityGroups() bool`

HasHasSecurityGroups returns a boolean if a field has been set.

### GetHasParameters

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasParameters() bool`

GetHasParameters returns the HasParameters field if non-nil, zero value otherwise.

### GetHasParametersOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasParametersOk() (*bool, bool)`

GetHasParametersOk returns a tuple with the HasParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasParameters

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetHasParameters(v bool)`

SetHasParameters sets HasParameters field to given value.

### HasHasParameters

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasHasParameters() bool`

HasHasParameters returns a boolean if a field has been set.

### GetCanEnforceTags

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetCanEnforceTags() bool`

GetCanEnforceTags returns the CanEnforceTags field if non-nil, zero value otherwise.

### GetCanEnforceTagsOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetCanEnforceTagsOk() (*bool, bool)`

GetCanEnforceTagsOk returns a tuple with the CanEnforceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEnforceTags

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetCanEnforceTags(v bool)`

SetCanEnforceTags sets CanEnforceTags field to given value.

### HasCanEnforceTags

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasCanEnforceTags() bool`

HasCanEnforceTags returns a boolean if a field has been set.

### SetCanEnforceTagsNil

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetCanEnforceTagsNil(b bool)`

 SetCanEnforceTagsNil sets the value for CanEnforceTags to be an explicit nil

### UnsetCanEnforceTags
`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) UnsetCanEnforceTags()`

UnsetCanEnforceTags ensures that no value is present for CanEnforceTags, not even an explicit nil
### GetDisableRootDatastore

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetDisableRootDatastore() bool`

GetDisableRootDatastore returns the DisableRootDatastore field if non-nil, zero value otherwise.

### GetDisableRootDatastoreOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetDisableRootDatastoreOk() (*bool, bool)`

GetDisableRootDatastoreOk returns a tuple with the DisableRootDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRootDatastore

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetDisableRootDatastore(v bool)`

SetDisableRootDatastore sets DisableRootDatastore field to given value.

### HasDisableRootDatastore

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasDisableRootDatastore() bool`

HasDisableRootDatastore returns a boolean if a field has been set.

### GetHasSnapshots

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasSnapshots() bool`

GetHasSnapshots returns the HasSnapshots field if non-nil, zero value otherwise.

### GetHasSnapshotsOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasSnapshotsOk() (*bool, bool)`

GetHasSnapshotsOk returns a tuple with the HasSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSnapshots

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetHasSnapshots(v bool)`

SetHasSnapshots sets HasSnapshots field to given value.

### HasHasSnapshots

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasHasSnapshots() bool`

HasHasSnapshots returns a boolean if a field has been set.

### GetHasMemorySnapshots

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasMemorySnapshots() bool`

GetHasMemorySnapshots returns the HasMemorySnapshots field if non-nil, zero value otherwise.

### GetHasMemorySnapshotsOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasMemorySnapshotsOk() (*bool, bool)`

GetHasMemorySnapshotsOk returns a tuple with the HasMemorySnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMemorySnapshots

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetHasMemorySnapshots(v bool)`

SetHasMemorySnapshots sets HasMemorySnapshots field to given value.

### HasHasMemorySnapshots

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasHasMemorySnapshots() bool`

HasHasMemorySnapshots returns a boolean if a field has been set.

### GetHasSpecTemplates

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasSpecTemplates() bool`

GetHasSpecTemplates returns the HasSpecTemplates field if non-nil, zero value otherwise.

### GetHasSpecTemplatesOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasSpecTemplatesOk() (*bool, bool)`

GetHasSpecTemplatesOk returns a tuple with the HasSpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSpecTemplates

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetHasSpecTemplates(v bool)`

SetHasSpecTemplates sets HasSpecTemplates field to given value.

### HasHasSpecTemplates

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasHasSpecTemplates() bool`

HasHasSpecTemplates returns a boolean if a field has been set.

### GetHasPreview

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasPreview() bool`

GetHasPreview returns the HasPreview field if non-nil, zero value otherwise.

### GetHasPreviewOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasPreviewOk() (*bool, bool)`

GetHasPreviewOk returns a tuple with the HasPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreview

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetHasPreview(v bool)`

SetHasPreview sets HasPreview field to given value.

### HasHasPreview

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasHasPreview() bool`

HasHasPreview returns a boolean if a field has been set.

### GetZonePoolRequired

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetZonePoolRequired() bool`

GetZonePoolRequired returns the ZonePoolRequired field if non-nil, zero value otherwise.

### GetZonePoolRequiredOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetZonePoolRequiredOk() (*bool, bool)`

GetZonePoolRequiredOk returns a tuple with the ZonePoolRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePoolRequired

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetZonePoolRequired(v bool)`

SetZonePoolRequired sets ZonePoolRequired field to given value.

### HasZonePoolRequired

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasZonePoolRequired() bool`

HasZonePoolRequired returns a boolean if a field has been set.

### GetPlanRequiresPool

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetPlanRequiresPool() bool`

GetPlanRequiresPool returns the PlanRequiresPool field if non-nil, zero value otherwise.

### GetPlanRequiresPoolOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetPlanRequiresPoolOk() (*bool, bool)`

GetPlanRequiresPoolOk returns a tuple with the PlanRequiresPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanRequiresPool

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetPlanRequiresPool(v bool)`

SetPlanRequiresPool sets PlanRequiresPool field to given value.

### HasPlanRequiresPool

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasPlanRequiresPool() bool`

HasPlanRequiresPool returns a boolean if a field has been set.

### GetHasFolders

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasFolders() bool`

GetHasFolders returns the HasFolders field if non-nil, zero value otherwise.

### GetHasFoldersOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetHasFoldersOk() (*bool, bool)`

GetHasFoldersOk returns a tuple with the HasFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFolders

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetHasFolders(v bool)`

SetHasFolders sets HasFolders field to given value.

### HasHasFolders

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasHasFolders() bool`

HasHasFolders returns a boolean if a field has been set.

### SetHasFoldersNil

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetHasFoldersNil(b bool)`

 SetHasFoldersNil sets the value for HasFolders to be an explicit nil

### UnsetHasFolders
`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) UnsetHasFolders()`

UnsetHasFolders ensures that no value is present for HasFolders, not even an explicit nil
### GetOptionTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetOptionTypes() []ListProvisionTypes200ResponseAllOfProvisionTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetOptionTypesOk() (*[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetOptionTypes(v []ListProvisionTypes200ResponseAllOfProvisionTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetCustomOptionTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetCustomOptionTypes() []ListProvisionTypes200ResponseAllOfProvisionTypesInnerCustomOptionTypesInner`

GetCustomOptionTypes returns the CustomOptionTypes field if non-nil, zero value otherwise.

### GetCustomOptionTypesOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetCustomOptionTypesOk() (*[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerCustomOptionTypesInner, bool)`

GetCustomOptionTypesOk returns a tuple with the CustomOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptionTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetCustomOptionTypes(v []ListProvisionTypes200ResponseAllOfProvisionTypesInnerCustomOptionTypesInner)`

SetCustomOptionTypes sets CustomOptionTypes field to given value.

### HasCustomOptionTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasCustomOptionTypes() bool`

HasCustomOptionTypes returns a boolean if a field has been set.

### SetCustomOptionTypesNil

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetCustomOptionTypesNil(b bool)`

 SetCustomOptionTypesNil sets the value for CustomOptionTypes to be an explicit nil

### UnsetCustomOptionTypes
`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) UnsetCustomOptionTypes()`

UnsetCustomOptionTypes ensures that no value is present for CustomOptionTypes, not even an explicit nil
### GetNetworkTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetNetworkTypes() []ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner`

GetNetworkTypes returns the NetworkTypes field if non-nil, zero value otherwise.

### GetNetworkTypesOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetNetworkTypesOk() (*[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner, bool)`

GetNetworkTypesOk returns a tuple with the NetworkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetNetworkTypes(v []ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner)`

SetNetworkTypes sets NetworkTypes field to given value.

### HasNetworkTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasNetworkTypes() bool`

HasNetworkTypes returns a boolean if a field has been set.

### SetNetworkTypesNil

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetNetworkTypesNil(b bool)`

 SetNetworkTypesNil sets the value for NetworkTypes to be an explicit nil

### UnsetNetworkTypes
`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) UnsetNetworkTypes()`

UnsetNetworkTypes ensures that no value is present for NetworkTypes, not even an explicit nil
### GetStorageTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetStorageTypes() []ListProvisionTypes200ResponseAllOfProvisionTypesInnerStorageTypesInner`

GetStorageTypes returns the StorageTypes field if non-nil, zero value otherwise.

### GetStorageTypesOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetStorageTypesOk() (*[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerStorageTypesInner, bool)`

GetStorageTypesOk returns a tuple with the StorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetStorageTypes(v []ListProvisionTypes200ResponseAllOfProvisionTypesInnerStorageTypesInner)`

SetStorageTypes sets StorageTypes field to given value.

### HasStorageTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasStorageTypes() bool`

HasStorageTypes returns a boolean if a field has been set.

### SetStorageTypesNil

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetStorageTypesNil(b bool)`

 SetStorageTypesNil sets the value for StorageTypes to be an explicit nil

### UnsetStorageTypes
`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) UnsetStorageTypes()`

UnsetStorageTypes ensures that no value is present for StorageTypes, not even an explicit nil
### GetRootStorageTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetRootStorageTypes() []ListProvisionTypes200ResponseAllOfProvisionTypesInnerRootStorageTypesInner`

GetRootStorageTypes returns the RootStorageTypes field if non-nil, zero value otherwise.

### GetRootStorageTypesOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetRootStorageTypesOk() (*[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerRootStorageTypesInner, bool)`

GetRootStorageTypesOk returns a tuple with the RootStorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootStorageTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetRootStorageTypes(v []ListProvisionTypes200ResponseAllOfProvisionTypesInnerRootStorageTypesInner)`

SetRootStorageTypes sets RootStorageTypes field to given value.

### HasRootStorageTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasRootStorageTypes() bool`

HasRootStorageTypes returns a boolean if a field has been set.

### SetRootStorageTypesNil

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetRootStorageTypesNil(b bool)`

 SetRootStorageTypesNil sets the value for RootStorageTypes to be an explicit nil

### UnsetRootStorageTypes
`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) UnsetRootStorageTypes()`

UnsetRootStorageTypes ensures that no value is present for RootStorageTypes, not even an explicit nil
### GetControllerTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetControllerTypes() []ListProvisionTypes200ResponseAllOfProvisionTypesInnerControllerTypesInner`

GetControllerTypes returns the ControllerTypes field if non-nil, zero value otherwise.

### GetControllerTypesOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) GetControllerTypesOk() (*[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerControllerTypesInner, bool)`

GetControllerTypesOk returns a tuple with the ControllerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetControllerTypes(v []ListProvisionTypes200ResponseAllOfProvisionTypesInnerControllerTypesInner)`

SetControllerTypes sets ControllerTypes field to given value.

### HasControllerTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) HasControllerTypes() bool`

HasControllerTypes returns a boolean if a field has been set.

### SetControllerTypesNil

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) SetControllerTypesNil(b bool)`

 SetControllerTypesNil sets the value for ControllerTypes to be an explicit nil

### UnsetControllerTypes
`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInner) UnsetControllerTypes()`

UnsetControllerTypes ensures that no value is present for ControllerTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


