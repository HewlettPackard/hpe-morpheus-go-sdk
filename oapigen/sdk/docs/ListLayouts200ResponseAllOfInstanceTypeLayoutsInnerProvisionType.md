# ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType

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
**OptionTypes** | Pointer to [**[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeOptionTypesInner**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeOptionTypesInner.md) |  | [optional] 
**CustomOptionTypes** | Pointer to [**[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeCustomOptionTypesInner**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeCustomOptionTypesInner.md) |  | [optional] 
**NetworkTypes** | Pointer to [**[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner.md) |  | [optional] 
**StorageTypes** | Pointer to [**[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner.md) |  | [optional] 
**RootStorageTypes** | Pointer to [**[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeRootStorageTypesInner**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeRootStorageTypesInner.md) |  | [optional] 
**ControllerTypes** | Pointer to [**[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeControllerTypesInner**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeControllerTypesInner.md) |  | [optional] 

## Methods

### NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType

`func NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType() *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType`

NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType instantiates a new ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeWithDefaults

`func NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeWithDefaults() *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType`

NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeWithDefaults instantiates a new ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAclEnabled

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetAclEnabled() bool`

GetAclEnabled returns the AclEnabled field if non-nil, zero value otherwise.

### GetAclEnabledOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetAclEnabledOk() (*bool, bool)`

GetAclEnabledOk returns a tuple with the AclEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclEnabled

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetAclEnabled(v bool)`

SetAclEnabled sets AclEnabled field to given value.

### HasAclEnabled

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasAclEnabled() bool`

HasAclEnabled returns a boolean if a field has been set.

### GetMultiTenant

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetMultiTenant() bool`

GetMultiTenant returns the MultiTenant field if non-nil, zero value otherwise.

### GetMultiTenantOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetMultiTenantOk() (*bool, bool)`

GetMultiTenantOk returns a tuple with the MultiTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenant

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetMultiTenant(v bool)`

SetMultiTenant sets MultiTenant field to given value.

### HasMultiTenant

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasMultiTenant() bool`

HasMultiTenant returns a boolean if a field has been set.

### GetManaged

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetHostNetwork

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHostNetwork() bool`

GetHostNetwork returns the HostNetwork field if non-nil, zero value otherwise.

### GetHostNetworkOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHostNetworkOk() (*bool, bool)`

GetHostNetworkOk returns a tuple with the HostNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNetwork

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetHostNetwork(v bool)`

SetHostNetwork sets HostNetwork field to given value.

### HasHostNetwork

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasHostNetwork() bool`

HasHostNetwork returns a boolean if a field has been set.

### GetCustomSupported

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetCustomSupported() bool`

GetCustomSupported returns the CustomSupported field if non-nil, zero value otherwise.

### GetCustomSupportedOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetCustomSupportedOk() (*bool, bool)`

GetCustomSupportedOk returns a tuple with the CustomSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSupported

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetCustomSupported(v bool)`

SetCustomSupported sets CustomSupported field to given value.

### HasCustomSupported

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasCustomSupported() bool`

HasCustomSupported returns a boolean if a field has been set.

### GetMapPorts

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetMapPorts() bool`

GetMapPorts returns the MapPorts field if non-nil, zero value otherwise.

### GetMapPortsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetMapPortsOk() (*bool, bool)`

GetMapPortsOk returns a tuple with the MapPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapPorts

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetMapPorts(v bool)`

SetMapPorts sets MapPorts field to given value.

### HasMapPorts

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasMapPorts() bool`

HasMapPorts returns a boolean if a field has been set.

### GetExportServer

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetExportServer() bool`

GetExportServer returns the ExportServer field if non-nil, zero value otherwise.

### GetExportServerOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetExportServerOk() (*bool, bool)`

GetExportServerOk returns a tuple with the ExportServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportServer

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetExportServer(v bool)`

SetExportServer sets ExportServer field to given value.

### HasExportServer

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasExportServer() bool`

HasExportServer returns a boolean if a field has been set.

### GetViewSet

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetViewSet() string`

GetViewSet returns the ViewSet field if non-nil, zero value otherwise.

### GetViewSetOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetViewSetOk() (*string, bool)`

GetViewSetOk returns a tuple with the ViewSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewSet

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetViewSet(v string)`

SetViewSet sets ViewSet field to given value.

### HasViewSet

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasViewSet() bool`

HasViewSet returns a boolean if a field has been set.

### SetViewSetNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetViewSetNil(b bool)`

 SetViewSetNil sets the value for ViewSet to be an explicit nil

### UnsetViewSet
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) UnsetViewSet()`

UnsetViewSet ensures that no value is present for ViewSet, not even an explicit nil
### GetServerType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetHostType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### GetAddVolumes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetHasVolumes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasVolumes() bool`

GetHasVolumes returns the HasVolumes field if non-nil, zero value otherwise.

### GetHasVolumesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasVolumesOk() (*bool, bool)`

GetHasVolumesOk returns a tuple with the HasVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVolumes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetHasVolumes(v bool)`

SetHasVolumes sets HasVolumes field to given value.

### HasHasVolumes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasHasVolumes() bool`

HasHasVolumes returns a boolean if a field has been set.

### GetHasDatastore

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetHasNetworks

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### GetMaxNetworks

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetMaxNetworks() int64`

GetMaxNetworks returns the MaxNetworks field if non-nil, zero value otherwise.

### GetMaxNetworksOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetMaxNetworksOk() (*int64, bool)`

GetMaxNetworksOk returns a tuple with the MaxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworks

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetMaxNetworks(v int64)`

SetMaxNetworks sets MaxNetworks field to given value.

### HasMaxNetworks

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasMaxNetworks() bool`

HasMaxNetworks returns a boolean if a field has been set.

### GetCustomizeVolume

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetCustomizeVolume() bool`

GetCustomizeVolume returns the CustomizeVolume field if non-nil, zero value otherwise.

### GetCustomizeVolumeOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetCustomizeVolumeOk() (*bool, bool)`

GetCustomizeVolumeOk returns a tuple with the CustomizeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizeVolume

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetCustomizeVolume(v bool)`

SetCustomizeVolume sets CustomizeVolume field to given value.

### HasCustomizeVolume

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasCustomizeVolume() bool`

HasCustomizeVolume returns a boolean if a field has been set.

### GetRootDiskCustomizable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetRootDiskCustomizable() bool`

GetRootDiskCustomizable returns the RootDiskCustomizable field if non-nil, zero value otherwise.

### GetRootDiskCustomizableOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetRootDiskCustomizableOk() (*bool, bool)`

GetRootDiskCustomizableOk returns a tuple with the RootDiskCustomizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskCustomizable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetRootDiskCustomizable(v bool)`

SetRootDiskCustomizable sets RootDiskCustomizable field to given value.

### HasRootDiskCustomizable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasRootDiskCustomizable() bool`

HasRootDiskCustomizable returns a boolean if a field has been set.

### GetRootDiskSizeKnown

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetRootDiskSizeKnown() bool`

GetRootDiskSizeKnown returns the RootDiskSizeKnown field if non-nil, zero value otherwise.

### GetRootDiskSizeKnownOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetRootDiskSizeKnownOk() (*bool, bool)`

GetRootDiskSizeKnownOk returns a tuple with the RootDiskSizeKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSizeKnown

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetRootDiskSizeKnown(v bool)`

SetRootDiskSizeKnown sets RootDiskSizeKnown field to given value.

### HasRootDiskSizeKnown

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasRootDiskSizeKnown() bool`

HasRootDiskSizeKnown returns a boolean if a field has been set.

### GetRootDiskResizable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetRootDiskResizable() bool`

GetRootDiskResizable returns the RootDiskResizable field if non-nil, zero value otherwise.

### GetRootDiskResizableOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetRootDiskResizableOk() (*bool, bool)`

GetRootDiskResizableOk returns a tuple with the RootDiskResizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskResizable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetRootDiskResizable(v bool)`

SetRootDiskResizable sets RootDiskResizable field to given value.

### HasRootDiskResizable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasRootDiskResizable() bool`

HasRootDiskResizable returns a boolean if a field has been set.

### GetLvmSupported

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetLvmSupported() bool`

GetLvmSupported returns the LvmSupported field if non-nil, zero value otherwise.

### GetLvmSupportedOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetLvmSupportedOk() (*bool, bool)`

GetLvmSupportedOk returns a tuple with the LvmSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmSupported

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetLvmSupported(v bool)`

SetLvmSupported sets LvmSupported field to given value.

### HasLvmSupported

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasLvmSupported() bool`

HasLvmSupported returns a boolean if a field has been set.

### GetHostDiskMode

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHostDiskMode() string`

GetHostDiskMode returns the HostDiskMode field if non-nil, zero value otherwise.

### GetHostDiskModeOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHostDiskModeOk() (*string, bool)`

GetHostDiskModeOk returns a tuple with the HostDiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDiskMode

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetHostDiskMode(v string)`

SetHostDiskMode sets HostDiskMode field to given value.

### HasHostDiskMode

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasHostDiskMode() bool`

HasHostDiskMode returns a boolean if a field has been set.

### GetMinDisk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetMinDisk() int64`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetMinDiskOk() (*int64, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetMinDisk(v int64)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMaxDisk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetMaxDisk() string`

GetMaxDisk returns the MaxDisk field if non-nil, zero value otherwise.

### GetMaxDiskOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetMaxDiskOk() (*string, bool)`

GetMaxDiskOk returns a tuple with the MaxDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetMaxDisk(v string)`

SetMaxDisk sets MaxDisk field to given value.

### HasMaxDisk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasMaxDisk() bool`

HasMaxDisk returns a boolean if a field has been set.

### SetMaxDiskNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetMaxDiskNil(b bool)`

 SetMaxDiskNil sets the value for MaxDisk to be an explicit nil

### UnsetMaxDisk
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) UnsetMaxDisk()`

UnsetMaxDisk ensures that no value is present for MaxDisk, not even an explicit nil
### GetResizeCopiesVolumes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetResizeCopiesVolumes() bool`

GetResizeCopiesVolumes returns the ResizeCopiesVolumes field if non-nil, zero value otherwise.

### GetResizeCopiesVolumesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetResizeCopiesVolumesOk() (*bool, bool)`

GetResizeCopiesVolumesOk returns a tuple with the ResizeCopiesVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeCopiesVolumes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetResizeCopiesVolumes(v bool)`

SetResizeCopiesVolumes sets ResizeCopiesVolumes field to given value.

### HasResizeCopiesVolumes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasResizeCopiesVolumes() bool`

HasResizeCopiesVolumes returns a boolean if a field has been set.

### GetSupportsAutoDatastore

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetSupportsAutoDatastore() bool`

GetSupportsAutoDatastore returns the SupportsAutoDatastore field if non-nil, zero value otherwise.

### GetSupportsAutoDatastoreOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetSupportsAutoDatastoreOk() (*bool, bool)`

GetSupportsAutoDatastoreOk returns a tuple with the SupportsAutoDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAutoDatastore

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetSupportsAutoDatastore(v bool)`

SetSupportsAutoDatastore sets SupportsAutoDatastore field to given value.

### HasSupportsAutoDatastore

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasSupportsAutoDatastore() bool`

HasSupportsAutoDatastore returns a boolean if a field has been set.

### GetHasZonePools

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasZonePools() bool`

GetHasZonePools returns the HasZonePools field if non-nil, zero value otherwise.

### GetHasZonePoolsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasZonePoolsOk() (*bool, bool)`

GetHasZonePoolsOk returns a tuple with the HasZonePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasZonePools

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetHasZonePools(v bool)`

SetHasZonePools sets HasZonePools field to given value.

### HasHasZonePools

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasHasZonePools() bool`

HasHasZonePools returns a boolean if a field has been set.

### GetHasSecurityGroups

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasSecurityGroups() bool`

GetHasSecurityGroups returns the HasSecurityGroups field if non-nil, zero value otherwise.

### GetHasSecurityGroupsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasSecurityGroupsOk() (*bool, bool)`

GetHasSecurityGroupsOk returns a tuple with the HasSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecurityGroups

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetHasSecurityGroups(v bool)`

SetHasSecurityGroups sets HasSecurityGroups field to given value.

### HasHasSecurityGroups

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasHasSecurityGroups() bool`

HasHasSecurityGroups returns a boolean if a field has been set.

### GetHasParameters

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasParameters() bool`

GetHasParameters returns the HasParameters field if non-nil, zero value otherwise.

### GetHasParametersOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasParametersOk() (*bool, bool)`

GetHasParametersOk returns a tuple with the HasParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasParameters

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetHasParameters(v bool)`

SetHasParameters sets HasParameters field to given value.

### HasHasParameters

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasHasParameters() bool`

HasHasParameters returns a boolean if a field has been set.

### GetCanEnforceTags

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetCanEnforceTags() bool`

GetCanEnforceTags returns the CanEnforceTags field if non-nil, zero value otherwise.

### GetCanEnforceTagsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetCanEnforceTagsOk() (*bool, bool)`

GetCanEnforceTagsOk returns a tuple with the CanEnforceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEnforceTags

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetCanEnforceTags(v bool)`

SetCanEnforceTags sets CanEnforceTags field to given value.

### HasCanEnforceTags

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasCanEnforceTags() bool`

HasCanEnforceTags returns a boolean if a field has been set.

### SetCanEnforceTagsNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetCanEnforceTagsNil(b bool)`

 SetCanEnforceTagsNil sets the value for CanEnforceTags to be an explicit nil

### UnsetCanEnforceTags
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) UnsetCanEnforceTags()`

UnsetCanEnforceTags ensures that no value is present for CanEnforceTags, not even an explicit nil
### GetDisableRootDatastore

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetDisableRootDatastore() bool`

GetDisableRootDatastore returns the DisableRootDatastore field if non-nil, zero value otherwise.

### GetDisableRootDatastoreOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetDisableRootDatastoreOk() (*bool, bool)`

GetDisableRootDatastoreOk returns a tuple with the DisableRootDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRootDatastore

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetDisableRootDatastore(v bool)`

SetDisableRootDatastore sets DisableRootDatastore field to given value.

### HasDisableRootDatastore

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasDisableRootDatastore() bool`

HasDisableRootDatastore returns a boolean if a field has been set.

### GetHasSnapshots

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasSnapshots() bool`

GetHasSnapshots returns the HasSnapshots field if non-nil, zero value otherwise.

### GetHasSnapshotsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasSnapshotsOk() (*bool, bool)`

GetHasSnapshotsOk returns a tuple with the HasSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSnapshots

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetHasSnapshots(v bool)`

SetHasSnapshots sets HasSnapshots field to given value.

### HasHasSnapshots

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasHasSnapshots() bool`

HasHasSnapshots returns a boolean if a field has been set.

### GetHasMemorySnapshots

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasMemorySnapshots() bool`

GetHasMemorySnapshots returns the HasMemorySnapshots field if non-nil, zero value otherwise.

### GetHasMemorySnapshotsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasMemorySnapshotsOk() (*bool, bool)`

GetHasMemorySnapshotsOk returns a tuple with the HasMemorySnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMemorySnapshots

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetHasMemorySnapshots(v bool)`

SetHasMemorySnapshots sets HasMemorySnapshots field to given value.

### HasHasMemorySnapshots

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasHasMemorySnapshots() bool`

HasHasMemorySnapshots returns a boolean if a field has been set.

### GetHasSpecTemplates

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasSpecTemplates() bool`

GetHasSpecTemplates returns the HasSpecTemplates field if non-nil, zero value otherwise.

### GetHasSpecTemplatesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasSpecTemplatesOk() (*bool, bool)`

GetHasSpecTemplatesOk returns a tuple with the HasSpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSpecTemplates

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetHasSpecTemplates(v bool)`

SetHasSpecTemplates sets HasSpecTemplates field to given value.

### HasHasSpecTemplates

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasHasSpecTemplates() bool`

HasHasSpecTemplates returns a boolean if a field has been set.

### GetHasPreview

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasPreview() bool`

GetHasPreview returns the HasPreview field if non-nil, zero value otherwise.

### GetHasPreviewOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasPreviewOk() (*bool, bool)`

GetHasPreviewOk returns a tuple with the HasPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreview

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetHasPreview(v bool)`

SetHasPreview sets HasPreview field to given value.

### HasHasPreview

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasHasPreview() bool`

HasHasPreview returns a boolean if a field has been set.

### GetZonePoolRequired

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetZonePoolRequired() bool`

GetZonePoolRequired returns the ZonePoolRequired field if non-nil, zero value otherwise.

### GetZonePoolRequiredOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetZonePoolRequiredOk() (*bool, bool)`

GetZonePoolRequiredOk returns a tuple with the ZonePoolRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePoolRequired

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetZonePoolRequired(v bool)`

SetZonePoolRequired sets ZonePoolRequired field to given value.

### HasZonePoolRequired

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasZonePoolRequired() bool`

HasZonePoolRequired returns a boolean if a field has been set.

### GetPlanRequiresPool

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetPlanRequiresPool() bool`

GetPlanRequiresPool returns the PlanRequiresPool field if non-nil, zero value otherwise.

### GetPlanRequiresPoolOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetPlanRequiresPoolOk() (*bool, bool)`

GetPlanRequiresPoolOk returns a tuple with the PlanRequiresPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanRequiresPool

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetPlanRequiresPool(v bool)`

SetPlanRequiresPool sets PlanRequiresPool field to given value.

### HasPlanRequiresPool

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasPlanRequiresPool() bool`

HasPlanRequiresPool returns a boolean if a field has been set.

### GetHasFolders

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasFolders() bool`

GetHasFolders returns the HasFolders field if non-nil, zero value otherwise.

### GetHasFoldersOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetHasFoldersOk() (*bool, bool)`

GetHasFoldersOk returns a tuple with the HasFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFolders

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetHasFolders(v bool)`

SetHasFolders sets HasFolders field to given value.

### HasHasFolders

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasHasFolders() bool`

HasHasFolders returns a boolean if a field has been set.

### SetHasFoldersNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetHasFoldersNil(b bool)`

 SetHasFoldersNil sets the value for HasFolders to be an explicit nil

### UnsetHasFolders
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) UnsetHasFolders()`

UnsetHasFolders ensures that no value is present for HasFolders, not even an explicit nil
### GetOptionTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetOptionTypes() []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetOptionTypesOk() (*[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetOptionTypes(v []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetCustomOptionTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetCustomOptionTypes() []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeCustomOptionTypesInner`

GetCustomOptionTypes returns the CustomOptionTypes field if non-nil, zero value otherwise.

### GetCustomOptionTypesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetCustomOptionTypesOk() (*[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeCustomOptionTypesInner, bool)`

GetCustomOptionTypesOk returns a tuple with the CustomOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptionTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetCustomOptionTypes(v []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeCustomOptionTypesInner)`

SetCustomOptionTypes sets CustomOptionTypes field to given value.

### HasCustomOptionTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasCustomOptionTypes() bool`

HasCustomOptionTypes returns a boolean if a field has been set.

### SetCustomOptionTypesNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetCustomOptionTypesNil(b bool)`

 SetCustomOptionTypesNil sets the value for CustomOptionTypes to be an explicit nil

### UnsetCustomOptionTypes
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) UnsetCustomOptionTypes()`

UnsetCustomOptionTypes ensures that no value is present for CustomOptionTypes, not even an explicit nil
### GetNetworkTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetNetworkTypes() []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner`

GetNetworkTypes returns the NetworkTypes field if non-nil, zero value otherwise.

### GetNetworkTypesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetNetworkTypesOk() (*[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner, bool)`

GetNetworkTypesOk returns a tuple with the NetworkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetNetworkTypes(v []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner)`

SetNetworkTypes sets NetworkTypes field to given value.

### HasNetworkTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasNetworkTypes() bool`

HasNetworkTypes returns a boolean if a field has been set.

### SetNetworkTypesNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetNetworkTypesNil(b bool)`

 SetNetworkTypesNil sets the value for NetworkTypes to be an explicit nil

### UnsetNetworkTypes
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) UnsetNetworkTypes()`

UnsetNetworkTypes ensures that no value is present for NetworkTypes, not even an explicit nil
### GetStorageTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetStorageTypes() []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner`

GetStorageTypes returns the StorageTypes field if non-nil, zero value otherwise.

### GetStorageTypesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetStorageTypesOk() (*[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner, bool)`

GetStorageTypesOk returns a tuple with the StorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetStorageTypes(v []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner)`

SetStorageTypes sets StorageTypes field to given value.

### HasStorageTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasStorageTypes() bool`

HasStorageTypes returns a boolean if a field has been set.

### SetStorageTypesNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetStorageTypesNil(b bool)`

 SetStorageTypesNil sets the value for StorageTypes to be an explicit nil

### UnsetStorageTypes
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) UnsetStorageTypes()`

UnsetStorageTypes ensures that no value is present for StorageTypes, not even an explicit nil
### GetRootStorageTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetRootStorageTypes() []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeRootStorageTypesInner`

GetRootStorageTypes returns the RootStorageTypes field if non-nil, zero value otherwise.

### GetRootStorageTypesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetRootStorageTypesOk() (*[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeRootStorageTypesInner, bool)`

GetRootStorageTypesOk returns a tuple with the RootStorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootStorageTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetRootStorageTypes(v []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeRootStorageTypesInner)`

SetRootStorageTypes sets RootStorageTypes field to given value.

### HasRootStorageTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasRootStorageTypes() bool`

HasRootStorageTypes returns a boolean if a field has been set.

### SetRootStorageTypesNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetRootStorageTypesNil(b bool)`

 SetRootStorageTypesNil sets the value for RootStorageTypes to be an explicit nil

### UnsetRootStorageTypes
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) UnsetRootStorageTypes()`

UnsetRootStorageTypes ensures that no value is present for RootStorageTypes, not even an explicit nil
### GetControllerTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetControllerTypes() []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeControllerTypesInner`

GetControllerTypes returns the ControllerTypes field if non-nil, zero value otherwise.

### GetControllerTypesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) GetControllerTypesOk() (*[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeControllerTypesInner, bool)`

GetControllerTypesOk returns a tuple with the ControllerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetControllerTypes(v []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeControllerTypesInner)`

SetControllerTypes sets ControllerTypes field to given value.

### HasControllerTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) HasControllerTypes() bool`

HasControllerTypes returns a boolean if a field has been set.

### SetControllerTypesNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) SetControllerTypesNil(b bool)`

 SetControllerTypesNil sets the value for ControllerTypes to be an explicit nil

### UnsetControllerTypes
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType) UnsetControllerTypes()`

UnsetControllerTypes ensures that no value is present for ControllerTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


