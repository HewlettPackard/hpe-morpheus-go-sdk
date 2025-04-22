# ClusterVolumes

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

### NewClusterVolumes

`func NewClusterVolumes() *ClusterVolumes`

NewClusterVolumes instantiates a new ClusterVolumes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterVolumesWithDefaults

`func NewClusterVolumesWithDefaults() *ClusterVolumes`

NewClusterVolumesWithDefaults instantiates a new ClusterVolumes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterVolumes) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterVolumes) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterVolumes) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterVolumes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalId

`func (o *ClusterVolumes) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ClusterVolumes) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ClusterVolumes) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ClusterVolumes) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ClusterVolumes) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ClusterVolumes) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ClusterVolumes) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ClusterVolumes) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetActive

`func (o *ClusterVolumes) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ClusterVolumes) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ClusterVolumes) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ClusterVolumes) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUsedStorage

`func (o *ClusterVolumes) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *ClusterVolumes) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *ClusterVolumes) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *ClusterVolumes) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetProvisionType

`func (o *ClusterVolumes) GetProvisionType() string`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ClusterVolumes) GetProvisionTypeOk() (*string, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ClusterVolumes) SetProvisionType(v string)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ClusterVolumes) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetResizeable

`func (o *ClusterVolumes) GetResizeable() bool`

GetResizeable returns the Resizeable field if non-nil, zero value otherwise.

### GetResizeableOk

`func (o *ClusterVolumes) GetResizeableOk() (*bool, bool)`

GetResizeableOk returns a tuple with the Resizeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeable

`func (o *ClusterVolumes) SetResizeable(v bool)`

SetResizeable sets Resizeable field to given value.

### HasResizeable

`func (o *ClusterVolumes) HasResizeable() bool`

HasResizeable returns a boolean if a field has been set.

### GetOnline

`func (o *ClusterVolumes) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ClusterVolumes) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ClusterVolumes) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ClusterVolumes) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetDeviceDisplayName

`func (o *ClusterVolumes) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *ClusterVolumes) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *ClusterVolumes) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.

### HasDeviceDisplayName

`func (o *ClusterVolumes) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### GetRefType

`func (o *ClusterVolumes) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ClusterVolumes) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ClusterVolumes) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ClusterVolumes) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetName

`func (o *ClusterVolumes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterVolumes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterVolumes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterVolumes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *ClusterVolumes) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ClusterVolumes) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ClusterVolumes) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ClusterVolumes) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDatastoreOption

`func (o *ClusterVolumes) GetDatastoreOption() string`

GetDatastoreOption returns the DatastoreOption field if non-nil, zero value otherwise.

### GetDatastoreOptionOk

`func (o *ClusterVolumes) GetDatastoreOptionOk() (*string, bool)`

GetDatastoreOptionOk returns a tuple with the DatastoreOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreOption

`func (o *ClusterVolumes) SetDatastoreOption(v string)`

SetDatastoreOption sets DatastoreOption field to given value.

### HasDatastoreOption

`func (o *ClusterVolumes) HasDatastoreOption() bool`

HasDatastoreOption returns a boolean if a field has been set.

### GetClaimName

`func (o *ClusterVolumes) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *ClusterVolumes) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *ClusterVolumes) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.

### HasClaimName

`func (o *ClusterVolumes) HasClaimName() bool`

HasClaimName returns a boolean if a field has been set.

### GetVolumeType

`func (o *ClusterVolumes) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *ClusterVolumes) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *ClusterVolumes) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *ClusterVolumes) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetDeviceName

`func (o *ClusterVolumes) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ClusterVolumes) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ClusterVolumes) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ClusterVolumes) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetRemovable

`func (o *ClusterVolumes) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *ClusterVolumes) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *ClusterVolumes) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *ClusterVolumes) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.

### GetPoolName

`func (o *ClusterVolumes) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *ClusterVolumes) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *ClusterVolumes) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *ClusterVolumes) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetReadOnly

`func (o *ClusterVolumes) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ClusterVolumes) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ClusterVolumes) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ClusterVolumes) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSourceId

`func (o *ClusterVolumes) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ClusterVolumes) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ClusterVolumes) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ClusterVolumes) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetZoneId

`func (o *ClusterVolumes) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ClusterVolumes) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ClusterVolumes) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ClusterVolumes) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetRootVolume

`func (o *ClusterVolumes) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *ClusterVolumes) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *ClusterVolumes) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *ClusterVolumes) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetRefId

`func (o *ClusterVolumes) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ClusterVolumes) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ClusterVolumes) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ClusterVolumes) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetCategory

`func (o *ClusterVolumes) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ClusterVolumes) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ClusterVolumes) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ClusterVolumes) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterVolumes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterVolumes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterVolumes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterVolumes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRawData

`func (o *ClusterVolumes) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *ClusterVolumes) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *ClusterVolumes) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *ClusterVolumes) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ClusterVolumes) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ClusterVolumes) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ClusterVolumes) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ClusterVolumes) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetAccount

`func (o *ClusterVolumes) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ClusterVolumes) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ClusterVolumes) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ClusterVolumes) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetType

`func (o *ClusterVolumes) GetType() GetAlerts200ResponseAllOfChecksInnerAccount`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterVolumes) GetTypeOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterVolumes) SetType(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterVolumes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDatastore

`func (o *ClusterVolumes) GetDatastore() GetAlerts200ResponseAllOfChecksInnerAccount`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *ClusterVolumes) GetDatastoreOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *ClusterVolumes) SetDatastore(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *ClusterVolumes) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


