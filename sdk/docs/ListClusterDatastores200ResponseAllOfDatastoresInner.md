# ListClusterDatastores200ResponseAllOfDatastoresInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**StorageSize** | Pointer to **int64** |  | [optional] 
**FreeSpace** | Pointer to **int64** |  | [optional] 
**DrsEnabled** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**AllowWrite** | Pointer to **bool** |  | [optional] 
**DefaultStore** | Pointer to **bool** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**AllowRead** | Pointer to **bool** |  | [optional] 
**AllowProvision** | Pointer to **bool** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**ZonePool** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Tenants** | Pointer to [**[]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner**](ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner.md) |  | [optional] 
**Datastores** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Locations** | Pointer to [**[]ListClusterDatastores200ResponseAllOfDatastoresInnerLocationsInner**](ListClusterDatastores200ResponseAllOfDatastoresInnerLocationsInner.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions**](ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions.md) |  | [optional] 

## Methods

### NewListClusterDatastores200ResponseAllOfDatastoresInner

`func NewListClusterDatastores200ResponseAllOfDatastoresInner() *ListClusterDatastores200ResponseAllOfDatastoresInner`

NewListClusterDatastores200ResponseAllOfDatastoresInner instantiates a new ListClusterDatastores200ResponseAllOfDatastoresInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterDatastores200ResponseAllOfDatastoresInnerWithDefaults

`func NewListClusterDatastores200ResponseAllOfDatastoresInnerWithDefaults() *ListClusterDatastores200ResponseAllOfDatastoresInner`

NewListClusterDatastores200ResponseAllOfDatastoresInnerWithDefaults instantiates a new ListClusterDatastores200ResponseAllOfDatastoresInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetStorageSize

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetFreeSpace

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### GetDrsEnabled

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetDrsEnabled() bool`

GetDrsEnabled returns the DrsEnabled field if non-nil, zero value otherwise.

### GetDrsEnabledOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetDrsEnabledOk() (*bool, bool)`

GetDrsEnabledOk returns a tuple with the DrsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrsEnabled

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetDrsEnabled(v bool)`

SetDrsEnabled sets DrsEnabled field to given value.

### HasDrsEnabled

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasDrsEnabled() bool`

HasDrsEnabled returns a boolean if a field has been set.

### GetActive

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAllowWrite

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.

### HasAllowWrite

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasAllowWrite() bool`

HasAllowWrite returns a boolean if a field has been set.

### GetDefaultStore

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetOnline

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetAllowRead

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetAllowRead() bool`

GetAllowRead returns the AllowRead field if non-nil, zero value otherwise.

### GetAllowReadOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetAllowReadOk() (*bool, bool)`

GetAllowReadOk returns a tuple with the AllowRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRead

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetAllowRead(v bool)`

SetAllowRead sets AllowRead field to given value.

### HasAllowRead

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasAllowRead() bool`

HasAllowRead returns a boolean if a field has been set.

### GetAllowProvision

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetAllowProvision() bool`

GetAllowProvision returns the AllowProvision field if non-nil, zero value otherwise.

### GetAllowProvisionOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetAllowProvisionOk() (*bool, bool)`

GetAllowProvisionOk returns a tuple with the AllowProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProvision

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetAllowProvision(v bool)`

SetAllowProvision sets AllowProvision field to given value.

### HasAllowProvision

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasAllowProvision() bool`

HasAllowProvision returns a boolean if a field has been set.

### GetRefType

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetExternalId

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetStatus

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetZone

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetZone() GetAlerts200ResponseAllOfChecksInnerAccount`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetZoneOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetZone(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetZonePool() GetAlerts200ResponseAllOfChecksInnerAccount`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetZonePoolOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetZonePool(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetOwner

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetOwner() GetAlerts200ResponseAllOfChecksInnerAccount`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetOwnerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetOwner(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenants

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetTenants() []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetTenantsOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetTenants(v []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetDatastores

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetDatastores() []map[string]interface{}`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetDatastoresOk() (*[]map[string]interface{}, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetDatastores(v []map[string]interface{})`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetLocations

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetLocations() []ListClusterDatastores200ResponseAllOfDatastoresInnerLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetLocationsOk() (*[]ListClusterDatastores200ResponseAllOfDatastoresInnerLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetLocations(v []ListClusterDatastores200ResponseAllOfDatastoresInnerLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetResourcePermissions() ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) GetResourcePermissionsOk() (*ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) SetResourcePermissions(v ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInner) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


