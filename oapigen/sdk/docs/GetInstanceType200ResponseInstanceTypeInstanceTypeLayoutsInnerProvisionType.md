# GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType

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
**OptionTypes** | Pointer to [**[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeOptionTypesInner**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeOptionTypesInner.md) |  | [optional] 
**CustomOptionTypes** | Pointer to [**[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeCustomOptionTypesInner**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeCustomOptionTypesInner.md) |  | [optional] 
**NetworkTypes** | Pointer to [**[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner.md) |  | [optional] 
**StorageTypes** | Pointer to [**[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner.md) |  | [optional] 
**RootStorageTypes** | Pointer to [**[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeRootStorageTypesInner**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeRootStorageTypesInner.md) |  | [optional] 
**ControllerTypes** | Pointer to [**[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeControllerTypesInner**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeControllerTypesInner.md) |  | [optional] 
**StorageProfiles** | Pointer to [**[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageProfilesInner**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageProfilesInner.md) |  | [optional] 

## Methods

### NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType

`func NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType() *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType`

NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType instantiates a new GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeWithDefaults

`func NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeWithDefaults() *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType`

NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeWithDefaults instantiates a new GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAclEnabled

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetAclEnabled() bool`

GetAclEnabled returns the AclEnabled field if non-nil, zero value otherwise.

### GetAclEnabledOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetAclEnabledOk() (*bool, bool)`

GetAclEnabledOk returns a tuple with the AclEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclEnabled

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetAclEnabled(v bool)`

SetAclEnabled sets AclEnabled field to given value.

### HasAclEnabled

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasAclEnabled() bool`

HasAclEnabled returns a boolean if a field has been set.

### GetMultiTenant

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMultiTenant() bool`

GetMultiTenant returns the MultiTenant field if non-nil, zero value otherwise.

### GetMultiTenantOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMultiTenantOk() (*bool, bool)`

GetMultiTenantOk returns a tuple with the MultiTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenant

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetMultiTenant(v bool)`

SetMultiTenant sets MultiTenant field to given value.

### HasMultiTenant

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasMultiTenant() bool`

HasMultiTenant returns a boolean if a field has been set.

### GetManaged

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetHostNetwork

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHostNetwork() bool`

GetHostNetwork returns the HostNetwork field if non-nil, zero value otherwise.

### GetHostNetworkOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHostNetworkOk() (*bool, bool)`

GetHostNetworkOk returns a tuple with the HostNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNetwork

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHostNetwork(v bool)`

SetHostNetwork sets HostNetwork field to given value.

### HasHostNetwork

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHostNetwork() bool`

HasHostNetwork returns a boolean if a field has been set.

### GetCustomSupported

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCustomSupported() bool`

GetCustomSupported returns the CustomSupported field if non-nil, zero value otherwise.

### GetCustomSupportedOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCustomSupportedOk() (*bool, bool)`

GetCustomSupportedOk returns a tuple with the CustomSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSupported

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetCustomSupported(v bool)`

SetCustomSupported sets CustomSupported field to given value.

### HasCustomSupported

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasCustomSupported() bool`

HasCustomSupported returns a boolean if a field has been set.

### GetMapPorts

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMapPorts() bool`

GetMapPorts returns the MapPorts field if non-nil, zero value otherwise.

### GetMapPortsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMapPortsOk() (*bool, bool)`

GetMapPortsOk returns a tuple with the MapPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapPorts

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetMapPorts(v bool)`

SetMapPorts sets MapPorts field to given value.

### HasMapPorts

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasMapPorts() bool`

HasMapPorts returns a boolean if a field has been set.

### GetExportServer

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetExportServer() bool`

GetExportServer returns the ExportServer field if non-nil, zero value otherwise.

### GetExportServerOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetExportServerOk() (*bool, bool)`

GetExportServerOk returns a tuple with the ExportServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportServer

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetExportServer(v bool)`

SetExportServer sets ExportServer field to given value.

### HasExportServer

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasExportServer() bool`

HasExportServer returns a boolean if a field has been set.

### GetViewSet

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetViewSet() string`

GetViewSet returns the ViewSet field if non-nil, zero value otherwise.

### GetViewSetOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetViewSetOk() (*string, bool)`

GetViewSetOk returns a tuple with the ViewSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewSet

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetViewSet(v string)`

SetViewSet sets ViewSet field to given value.

### HasViewSet

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasViewSet() bool`

HasViewSet returns a boolean if a field has been set.

### SetViewSetNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetViewSetNil(b bool)`

 SetViewSetNil sets the value for ViewSet to be an explicit nil

### UnsetViewSet
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetViewSet()`

UnsetViewSet ensures that no value is present for ViewSet, not even an explicit nil
### GetServerType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetHostType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### GetAddVolumes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetHasVolumes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasVolumes() bool`

GetHasVolumes returns the HasVolumes field if non-nil, zero value otherwise.

### GetHasVolumesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasVolumesOk() (*bool, bool)`

GetHasVolumesOk returns a tuple with the HasVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVolumes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasVolumes(v bool)`

SetHasVolumes sets HasVolumes field to given value.

### HasHasVolumes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasVolumes() bool`

HasHasVolumes returns a boolean if a field has been set.

### GetHasDatastore

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetHasNetworks

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### GetMaxNetworks

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMaxNetworks() int64`

GetMaxNetworks returns the MaxNetworks field if non-nil, zero value otherwise.

### GetMaxNetworksOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMaxNetworksOk() (*int64, bool)`

GetMaxNetworksOk returns a tuple with the MaxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworks

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetMaxNetworks(v int64)`

SetMaxNetworks sets MaxNetworks field to given value.

### HasMaxNetworks

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasMaxNetworks() bool`

HasMaxNetworks returns a boolean if a field has been set.

### GetCustomizeVolume

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCustomizeVolume() bool`

GetCustomizeVolume returns the CustomizeVolume field if non-nil, zero value otherwise.

### GetCustomizeVolumeOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCustomizeVolumeOk() (*bool, bool)`

GetCustomizeVolumeOk returns a tuple with the CustomizeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizeVolume

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetCustomizeVolume(v bool)`

SetCustomizeVolume sets CustomizeVolume field to given value.

### HasCustomizeVolume

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasCustomizeVolume() bool`

HasCustomizeVolume returns a boolean if a field has been set.

### GetRootDiskCustomizable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootDiskCustomizable() bool`

GetRootDiskCustomizable returns the RootDiskCustomizable field if non-nil, zero value otherwise.

### GetRootDiskCustomizableOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootDiskCustomizableOk() (*bool, bool)`

GetRootDiskCustomizableOk returns a tuple with the RootDiskCustomizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskCustomizable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetRootDiskCustomizable(v bool)`

SetRootDiskCustomizable sets RootDiskCustomizable field to given value.

### HasRootDiskCustomizable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasRootDiskCustomizable() bool`

HasRootDiskCustomizable returns a boolean if a field has been set.

### GetRootDiskSizeKnown

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootDiskSizeKnown() bool`

GetRootDiskSizeKnown returns the RootDiskSizeKnown field if non-nil, zero value otherwise.

### GetRootDiskSizeKnownOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootDiskSizeKnownOk() (*bool, bool)`

GetRootDiskSizeKnownOk returns a tuple with the RootDiskSizeKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSizeKnown

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetRootDiskSizeKnown(v bool)`

SetRootDiskSizeKnown sets RootDiskSizeKnown field to given value.

### HasRootDiskSizeKnown

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasRootDiskSizeKnown() bool`

HasRootDiskSizeKnown returns a boolean if a field has been set.

### GetRootDiskResizable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootDiskResizable() bool`

GetRootDiskResizable returns the RootDiskResizable field if non-nil, zero value otherwise.

### GetRootDiskResizableOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootDiskResizableOk() (*bool, bool)`

GetRootDiskResizableOk returns a tuple with the RootDiskResizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskResizable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetRootDiskResizable(v bool)`

SetRootDiskResizable sets RootDiskResizable field to given value.

### HasRootDiskResizable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasRootDiskResizable() bool`

HasRootDiskResizable returns a boolean if a field has been set.

### GetLvmSupported

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetLvmSupported() bool`

GetLvmSupported returns the LvmSupported field if non-nil, zero value otherwise.

### GetLvmSupportedOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetLvmSupportedOk() (*bool, bool)`

GetLvmSupportedOk returns a tuple with the LvmSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmSupported

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetLvmSupported(v bool)`

SetLvmSupported sets LvmSupported field to given value.

### HasLvmSupported

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasLvmSupported() bool`

HasLvmSupported returns a boolean if a field has been set.

### GetHostDiskMode

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHostDiskMode() string`

GetHostDiskMode returns the HostDiskMode field if non-nil, zero value otherwise.

### GetHostDiskModeOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHostDiskModeOk() (*string, bool)`

GetHostDiskModeOk returns a tuple with the HostDiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDiskMode

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHostDiskMode(v string)`

SetHostDiskMode sets HostDiskMode field to given value.

### HasHostDiskMode

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHostDiskMode() bool`

HasHostDiskMode returns a boolean if a field has been set.

### GetMinDisk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMinDisk() int64`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMinDiskOk() (*int64, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetMinDisk(v int64)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMaxDisk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMaxDisk() string`

GetMaxDisk returns the MaxDisk field if non-nil, zero value otherwise.

### GetMaxDiskOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetMaxDiskOk() (*string, bool)`

GetMaxDiskOk returns a tuple with the MaxDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetMaxDisk(v string)`

SetMaxDisk sets MaxDisk field to given value.

### HasMaxDisk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasMaxDisk() bool`

HasMaxDisk returns a boolean if a field has been set.

### SetMaxDiskNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetMaxDiskNil(b bool)`

 SetMaxDiskNil sets the value for MaxDisk to be an explicit nil

### UnsetMaxDisk
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetMaxDisk()`

UnsetMaxDisk ensures that no value is present for MaxDisk, not even an explicit nil
### GetResizeCopiesVolumes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetResizeCopiesVolumes() bool`

GetResizeCopiesVolumes returns the ResizeCopiesVolumes field if non-nil, zero value otherwise.

### GetResizeCopiesVolumesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetResizeCopiesVolumesOk() (*bool, bool)`

GetResizeCopiesVolumesOk returns a tuple with the ResizeCopiesVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeCopiesVolumes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetResizeCopiesVolumes(v bool)`

SetResizeCopiesVolumes sets ResizeCopiesVolumes field to given value.

### HasResizeCopiesVolumes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasResizeCopiesVolumes() bool`

HasResizeCopiesVolumes returns a boolean if a field has been set.

### GetSupportsAutoDatastore

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetSupportsAutoDatastore() bool`

GetSupportsAutoDatastore returns the SupportsAutoDatastore field if non-nil, zero value otherwise.

### GetSupportsAutoDatastoreOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetSupportsAutoDatastoreOk() (*bool, bool)`

GetSupportsAutoDatastoreOk returns a tuple with the SupportsAutoDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAutoDatastore

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetSupportsAutoDatastore(v bool)`

SetSupportsAutoDatastore sets SupportsAutoDatastore field to given value.

### HasSupportsAutoDatastore

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasSupportsAutoDatastore() bool`

HasSupportsAutoDatastore returns a boolean if a field has been set.

### GetHasZonePools

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasZonePools() bool`

GetHasZonePools returns the HasZonePools field if non-nil, zero value otherwise.

### GetHasZonePoolsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasZonePoolsOk() (*bool, bool)`

GetHasZonePoolsOk returns a tuple with the HasZonePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasZonePools

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasZonePools(v bool)`

SetHasZonePools sets HasZonePools field to given value.

### HasHasZonePools

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasZonePools() bool`

HasHasZonePools returns a boolean if a field has been set.

### GetHasSecurityGroups

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasSecurityGroups() bool`

GetHasSecurityGroups returns the HasSecurityGroups field if non-nil, zero value otherwise.

### GetHasSecurityGroupsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasSecurityGroupsOk() (*bool, bool)`

GetHasSecurityGroupsOk returns a tuple with the HasSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecurityGroups

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasSecurityGroups(v bool)`

SetHasSecurityGroups sets HasSecurityGroups field to given value.

### HasHasSecurityGroups

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasSecurityGroups() bool`

HasHasSecurityGroups returns a boolean if a field has been set.

### GetHasParameters

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasParameters() bool`

GetHasParameters returns the HasParameters field if non-nil, zero value otherwise.

### GetHasParametersOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasParametersOk() (*bool, bool)`

GetHasParametersOk returns a tuple with the HasParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasParameters

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasParameters(v bool)`

SetHasParameters sets HasParameters field to given value.

### HasHasParameters

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasParameters() bool`

HasHasParameters returns a boolean if a field has been set.

### GetCanEnforceTags

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCanEnforceTags() bool`

GetCanEnforceTags returns the CanEnforceTags field if non-nil, zero value otherwise.

### GetCanEnforceTagsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCanEnforceTagsOk() (*bool, bool)`

GetCanEnforceTagsOk returns a tuple with the CanEnforceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEnforceTags

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetCanEnforceTags(v bool)`

SetCanEnforceTags sets CanEnforceTags field to given value.

### HasCanEnforceTags

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasCanEnforceTags() bool`

HasCanEnforceTags returns a boolean if a field has been set.

### SetCanEnforceTagsNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetCanEnforceTagsNil(b bool)`

 SetCanEnforceTagsNil sets the value for CanEnforceTags to be an explicit nil

### UnsetCanEnforceTags
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetCanEnforceTags()`

UnsetCanEnforceTags ensures that no value is present for CanEnforceTags, not even an explicit nil
### GetDisableRootDatastore

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetDisableRootDatastore() bool`

GetDisableRootDatastore returns the DisableRootDatastore field if non-nil, zero value otherwise.

### GetDisableRootDatastoreOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetDisableRootDatastoreOk() (*bool, bool)`

GetDisableRootDatastoreOk returns a tuple with the DisableRootDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRootDatastore

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetDisableRootDatastore(v bool)`

SetDisableRootDatastore sets DisableRootDatastore field to given value.

### HasDisableRootDatastore

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasDisableRootDatastore() bool`

HasDisableRootDatastore returns a boolean if a field has been set.

### GetHasSnapshots

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasSnapshots() bool`

GetHasSnapshots returns the HasSnapshots field if non-nil, zero value otherwise.

### GetHasSnapshotsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasSnapshotsOk() (*bool, bool)`

GetHasSnapshotsOk returns a tuple with the HasSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSnapshots

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasSnapshots(v bool)`

SetHasSnapshots sets HasSnapshots field to given value.

### HasHasSnapshots

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasSnapshots() bool`

HasHasSnapshots returns a boolean if a field has been set.

### GetHasMemorySnapshots

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasMemorySnapshots() bool`

GetHasMemorySnapshots returns the HasMemorySnapshots field if non-nil, zero value otherwise.

### GetHasMemorySnapshotsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasMemorySnapshotsOk() (*bool, bool)`

GetHasMemorySnapshotsOk returns a tuple with the HasMemorySnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMemorySnapshots

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasMemorySnapshots(v bool)`

SetHasMemorySnapshots sets HasMemorySnapshots field to given value.

### HasHasMemorySnapshots

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasMemorySnapshots() bool`

HasHasMemorySnapshots returns a boolean if a field has been set.

### GetHasSpecTemplates

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasSpecTemplates() bool`

GetHasSpecTemplates returns the HasSpecTemplates field if non-nil, zero value otherwise.

### GetHasSpecTemplatesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasSpecTemplatesOk() (*bool, bool)`

GetHasSpecTemplatesOk returns a tuple with the HasSpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSpecTemplates

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasSpecTemplates(v bool)`

SetHasSpecTemplates sets HasSpecTemplates field to given value.

### HasHasSpecTemplates

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasSpecTemplates() bool`

HasHasSpecTemplates returns a boolean if a field has been set.

### GetHasPreview

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasPreview() bool`

GetHasPreview returns the HasPreview field if non-nil, zero value otherwise.

### GetHasPreviewOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasPreviewOk() (*bool, bool)`

GetHasPreviewOk returns a tuple with the HasPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreview

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasPreview(v bool)`

SetHasPreview sets HasPreview field to given value.

### HasHasPreview

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasPreview() bool`

HasHasPreview returns a boolean if a field has been set.

### GetZonePoolRequired

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetZonePoolRequired() bool`

GetZonePoolRequired returns the ZonePoolRequired field if non-nil, zero value otherwise.

### GetZonePoolRequiredOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetZonePoolRequiredOk() (*bool, bool)`

GetZonePoolRequiredOk returns a tuple with the ZonePoolRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePoolRequired

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetZonePoolRequired(v bool)`

SetZonePoolRequired sets ZonePoolRequired field to given value.

### HasZonePoolRequired

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasZonePoolRequired() bool`

HasZonePoolRequired returns a boolean if a field has been set.

### GetPlanRequiresPool

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetPlanRequiresPool() bool`

GetPlanRequiresPool returns the PlanRequiresPool field if non-nil, zero value otherwise.

### GetPlanRequiresPoolOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetPlanRequiresPoolOk() (*bool, bool)`

GetPlanRequiresPoolOk returns a tuple with the PlanRequiresPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanRequiresPool

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetPlanRequiresPool(v bool)`

SetPlanRequiresPool sets PlanRequiresPool field to given value.

### HasPlanRequiresPool

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasPlanRequiresPool() bool`

HasPlanRequiresPool returns a boolean if a field has been set.

### GetHasFolders

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasFolders() bool`

GetHasFolders returns the HasFolders field if non-nil, zero value otherwise.

### GetHasFoldersOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetHasFoldersOk() (*bool, bool)`

GetHasFoldersOk returns a tuple with the HasFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFolders

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasFolders(v bool)`

SetHasFolders sets HasFolders field to given value.

### HasHasFolders

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasHasFolders() bool`

HasHasFolders returns a boolean if a field has been set.

### SetHasFoldersNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetHasFoldersNil(b bool)`

 SetHasFoldersNil sets the value for HasFolders to be an explicit nil

### UnsetHasFolders
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetHasFolders()`

UnsetHasFolders ensures that no value is present for HasFolders, not even an explicit nil
### GetOptionTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetOptionTypes() []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetOptionTypesOk() (*[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetOptionTypes(v []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetCustomOptionTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCustomOptionTypes() []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeCustomOptionTypesInner`

GetCustomOptionTypes returns the CustomOptionTypes field if non-nil, zero value otherwise.

### GetCustomOptionTypesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetCustomOptionTypesOk() (*[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeCustomOptionTypesInner, bool)`

GetCustomOptionTypesOk returns a tuple with the CustomOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptionTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetCustomOptionTypes(v []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeCustomOptionTypesInner)`

SetCustomOptionTypes sets CustomOptionTypes field to given value.

### HasCustomOptionTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasCustomOptionTypes() bool`

HasCustomOptionTypes returns a boolean if a field has been set.

### SetCustomOptionTypesNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetCustomOptionTypesNil(b bool)`

 SetCustomOptionTypesNil sets the value for CustomOptionTypes to be an explicit nil

### UnsetCustomOptionTypes
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetCustomOptionTypes()`

UnsetCustomOptionTypes ensures that no value is present for CustomOptionTypes, not even an explicit nil
### GetNetworkTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetNetworkTypes() []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner`

GetNetworkTypes returns the NetworkTypes field if non-nil, zero value otherwise.

### GetNetworkTypesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetNetworkTypesOk() (*[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner, bool)`

GetNetworkTypesOk returns a tuple with the NetworkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetNetworkTypes(v []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner)`

SetNetworkTypes sets NetworkTypes field to given value.

### HasNetworkTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasNetworkTypes() bool`

HasNetworkTypes returns a boolean if a field has been set.

### SetNetworkTypesNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetNetworkTypesNil(b bool)`

 SetNetworkTypesNil sets the value for NetworkTypes to be an explicit nil

### UnsetNetworkTypes
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetNetworkTypes()`

UnsetNetworkTypes ensures that no value is present for NetworkTypes, not even an explicit nil
### GetStorageTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetStorageTypes() []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner`

GetStorageTypes returns the StorageTypes field if non-nil, zero value otherwise.

### GetStorageTypesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetStorageTypesOk() (*[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner, bool)`

GetStorageTypesOk returns a tuple with the StorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetStorageTypes(v []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner)`

SetStorageTypes sets StorageTypes field to given value.

### HasStorageTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasStorageTypes() bool`

HasStorageTypes returns a boolean if a field has been set.

### SetStorageTypesNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetStorageTypesNil(b bool)`

 SetStorageTypesNil sets the value for StorageTypes to be an explicit nil

### UnsetStorageTypes
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetStorageTypes()`

UnsetStorageTypes ensures that no value is present for StorageTypes, not even an explicit nil
### GetRootStorageTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootStorageTypes() []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeRootStorageTypesInner`

GetRootStorageTypes returns the RootStorageTypes field if non-nil, zero value otherwise.

### GetRootStorageTypesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetRootStorageTypesOk() (*[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeRootStorageTypesInner, bool)`

GetRootStorageTypesOk returns a tuple with the RootStorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootStorageTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetRootStorageTypes(v []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeRootStorageTypesInner)`

SetRootStorageTypes sets RootStorageTypes field to given value.

### HasRootStorageTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasRootStorageTypes() bool`

HasRootStorageTypes returns a boolean if a field has been set.

### SetRootStorageTypesNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetRootStorageTypesNil(b bool)`

 SetRootStorageTypesNil sets the value for RootStorageTypes to be an explicit nil

### UnsetRootStorageTypes
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetRootStorageTypes()`

UnsetRootStorageTypes ensures that no value is present for RootStorageTypes, not even an explicit nil
### GetControllerTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetControllerTypes() []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeControllerTypesInner`

GetControllerTypes returns the ControllerTypes field if non-nil, zero value otherwise.

### GetControllerTypesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetControllerTypesOk() (*[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeControllerTypesInner, bool)`

GetControllerTypesOk returns a tuple with the ControllerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetControllerTypes(v []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeControllerTypesInner)`

SetControllerTypes sets ControllerTypes field to given value.

### HasControllerTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasControllerTypes() bool`

HasControllerTypes returns a boolean if a field has been set.

### SetControllerTypesNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetControllerTypesNil(b bool)`

 SetControllerTypesNil sets the value for ControllerTypes to be an explicit nil

### UnsetControllerTypes
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetControllerTypes()`

UnsetControllerTypes ensures that no value is present for ControllerTypes, not even an explicit nil
### GetStorageProfiles

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetStorageProfiles() []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageProfilesInner`

GetStorageProfiles returns the StorageProfiles field if non-nil, zero value otherwise.

### GetStorageProfilesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) GetStorageProfilesOk() (*[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageProfilesInner, bool)`

GetStorageProfilesOk returns a tuple with the StorageProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfiles

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetStorageProfiles(v []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageProfilesInner)`

SetStorageProfiles sets StorageProfiles field to given value.

### HasStorageProfiles

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) HasStorageProfiles() bool`

HasStorageProfiles returns a boolean if a field has been set.

### SetStorageProfilesNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) SetStorageProfilesNil(b bool)`

 SetStorageProfilesNil sets the value for StorageProfiles to be an explicit nil

### UnsetStorageProfiles
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType) UnsetStorageProfiles()`

UnsetStorageProfiles ensures that no value is present for StorageProfiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


