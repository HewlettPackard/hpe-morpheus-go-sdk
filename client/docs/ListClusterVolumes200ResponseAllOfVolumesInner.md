# ListClusterVolumes200ResponseAllOfVolumesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**UsedStorage** | Pointer to **int64** |  | [optional] 
**ProvisionType** | Pointer to **string** |  | [optional] 
**Resizeable** | Pointer to **bool** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**DeviceDisplayName** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**DatastoreOption** | Pointer to **string** |  | [optional] 
**ClaimName** | Pointer to **string** |  | [optional] 
**VolumeType** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**Removable** | Pointer to **bool** |  | [optional] 
**PoolName** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**RootVolume** | Pointer to **bool** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RawData** | Pointer to **string** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Type** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Datastore** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 

## Methods

### NewListClusterVolumes200ResponseAllOfVolumesInner

`func NewListClusterVolumes200ResponseAllOfVolumesInner() *ListClusterVolumes200ResponseAllOfVolumesInner`

NewListClusterVolumes200ResponseAllOfVolumesInner instantiates a new ListClusterVolumes200ResponseAllOfVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterVolumes200ResponseAllOfVolumesInnerWithDefaults

`func NewListClusterVolumes200ResponseAllOfVolumesInnerWithDefaults() *ListClusterVolumes200ResponseAllOfVolumesInner`

NewListClusterVolumes200ResponseAllOfVolumesInnerWithDefaults instantiates a new ListClusterVolumes200ResponseAllOfVolumesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetActive

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUsedStorage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetProvisionType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetProvisionType() string`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetProvisionTypeOk() (*string, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetProvisionType(v string)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetResizeable

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetResizeable() bool`

GetResizeable returns the Resizeable field if non-nil, zero value otherwise.

### GetResizeableOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetResizeableOk() (*bool, bool)`

GetResizeableOk returns a tuple with the Resizeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeable

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetResizeable(v bool)`

SetResizeable sets Resizeable field to given value.

### HasResizeable

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasResizeable() bool`

HasResizeable returns a boolean if a field has been set.

### GetOnline

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetDeviceDisplayName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.

### HasDeviceDisplayName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### GetRefType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDatastoreOption

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDatastoreOption() string`

GetDatastoreOption returns the DatastoreOption field if non-nil, zero value otherwise.

### GetDatastoreOptionOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDatastoreOptionOk() (*string, bool)`

GetDatastoreOptionOk returns a tuple with the DatastoreOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreOption

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetDatastoreOption(v string)`

SetDatastoreOption sets DatastoreOption field to given value.

### HasDatastoreOption

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasDatastoreOption() bool`

HasDatastoreOption returns a boolean if a field has been set.

### GetClaimName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.

### HasClaimName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasClaimName() bool`

HasClaimName returns a boolean if a field has been set.

### GetVolumeType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetDeviceName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetRemovable

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.

### GetPoolName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetReadOnly

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSourceId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetZoneId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetRootVolume

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetRefId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetCategory

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetStatus

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRawData

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetAccount

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetType() GetAlerts200ResponseAllOfChecksInnerAccount`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetTypeOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetType(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetType sets Type field to given value.

### HasType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDatastore

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDatastore() GetAlerts200ResponseAllOfChecksInnerAccount`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDatastoreOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetDatastore(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


