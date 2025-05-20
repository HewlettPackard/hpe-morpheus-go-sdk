# StorageServerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**HasNamespaces** | Pointer to **bool** |  | [optional] 
**HasGroups** | Pointer to **bool** |  | [optional] 
**HasBlock** | Pointer to **bool** |  | [optional] 
**HasObject** | Pointer to **bool** |  | [optional] 
**HasFile** | Pointer to **bool** |  | [optional] 
**HasDatastore** | Pointer to **bool** |  | [optional] 
**HasDisks** | Pointer to **bool** |  | [optional] 
**HasHosts** | Pointer to **bool** |  | [optional] 
**CreateNamespaces** | Pointer to **bool** |  | [optional] 
**CreateGroup** | Pointer to **bool** |  | [optional] 
**CreateBlock** | Pointer to **bool** |  | [optional] 
**CreateObject** | Pointer to **bool** |  | [optional] 
**CreateFile** | Pointer to **bool** |  | [optional] 
**CreateDatastore** | Pointer to **bool** |  | [optional] 
**CreateDisk** | Pointer to **bool** |  | [optional] 
**CreateHost** | Pointer to **bool** |  | [optional] 
**IconCode** | Pointer to **string** |  | [optional] 
**HasFileBrowser** | Pointer to **bool** |  | [optional] 
**OptionTypes** | Pointer to [**[]ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner**](ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner.md) |  | [optional] 
**GroupOptionTypes** | Pointer to [**[]ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerGroupOptionTypesInner**](ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerGroupOptionTypesInner.md) |  | [optional] 
**BucketOptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ShareOptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ShareAccessOptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**StorageVolumeTypes** | Pointer to [**[]GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner**](GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner.md) |  | [optional] 

## Methods

### NewStorageServerType

`func NewStorageServerType() *StorageServerType`

NewStorageServerType instantiates a new StorageServerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageServerTypeWithDefaults

`func NewStorageServerTypeWithDefaults() *StorageServerType`

NewStorageServerTypeWithDefaults instantiates a new StorageServerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageServerType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageServerType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageServerType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *StorageServerType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *StorageServerType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StorageServerType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StorageServerType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *StorageServerType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *StorageServerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageServerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageServerType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageServerType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *StorageServerType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageServerType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageServerType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageServerType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *StorageServerType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageServerType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageServerType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageServerType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatable

`func (o *StorageServerType) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *StorageServerType) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *StorageServerType) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *StorageServerType) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetHasNamespaces

`func (o *StorageServerType) GetHasNamespaces() bool`

GetHasNamespaces returns the HasNamespaces field if non-nil, zero value otherwise.

### GetHasNamespacesOk

`func (o *StorageServerType) GetHasNamespacesOk() (*bool, bool)`

GetHasNamespacesOk returns a tuple with the HasNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNamespaces

`func (o *StorageServerType) SetHasNamespaces(v bool)`

SetHasNamespaces sets HasNamespaces field to given value.

### HasHasNamespaces

`func (o *StorageServerType) HasHasNamespaces() bool`

HasHasNamespaces returns a boolean if a field has been set.

### GetHasGroups

`func (o *StorageServerType) GetHasGroups() bool`

GetHasGroups returns the HasGroups field if non-nil, zero value otherwise.

### GetHasGroupsOk

`func (o *StorageServerType) GetHasGroupsOk() (*bool, bool)`

GetHasGroupsOk returns a tuple with the HasGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGroups

`func (o *StorageServerType) SetHasGroups(v bool)`

SetHasGroups sets HasGroups field to given value.

### HasHasGroups

`func (o *StorageServerType) HasHasGroups() bool`

HasHasGroups returns a boolean if a field has been set.

### GetHasBlock

`func (o *StorageServerType) GetHasBlock() bool`

GetHasBlock returns the HasBlock field if non-nil, zero value otherwise.

### GetHasBlockOk

`func (o *StorageServerType) GetHasBlockOk() (*bool, bool)`

GetHasBlockOk returns a tuple with the HasBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBlock

`func (o *StorageServerType) SetHasBlock(v bool)`

SetHasBlock sets HasBlock field to given value.

### HasHasBlock

`func (o *StorageServerType) HasHasBlock() bool`

HasHasBlock returns a boolean if a field has been set.

### GetHasObject

`func (o *StorageServerType) GetHasObject() bool`

GetHasObject returns the HasObject field if non-nil, zero value otherwise.

### GetHasObjectOk

`func (o *StorageServerType) GetHasObjectOk() (*bool, bool)`

GetHasObjectOk returns a tuple with the HasObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasObject

`func (o *StorageServerType) SetHasObject(v bool)`

SetHasObject sets HasObject field to given value.

### HasHasObject

`func (o *StorageServerType) HasHasObject() bool`

HasHasObject returns a boolean if a field has been set.

### GetHasFile

`func (o *StorageServerType) GetHasFile() bool`

GetHasFile returns the HasFile field if non-nil, zero value otherwise.

### GetHasFileOk

`func (o *StorageServerType) GetHasFileOk() (*bool, bool)`

GetHasFileOk returns a tuple with the HasFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFile

`func (o *StorageServerType) SetHasFile(v bool)`

SetHasFile sets HasFile field to given value.

### HasHasFile

`func (o *StorageServerType) HasHasFile() bool`

HasHasFile returns a boolean if a field has been set.

### GetHasDatastore

`func (o *StorageServerType) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *StorageServerType) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *StorageServerType) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *StorageServerType) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetHasDisks

`func (o *StorageServerType) GetHasDisks() bool`

GetHasDisks returns the HasDisks field if non-nil, zero value otherwise.

### GetHasDisksOk

`func (o *StorageServerType) GetHasDisksOk() (*bool, bool)`

GetHasDisksOk returns a tuple with the HasDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDisks

`func (o *StorageServerType) SetHasDisks(v bool)`

SetHasDisks sets HasDisks field to given value.

### HasHasDisks

`func (o *StorageServerType) HasHasDisks() bool`

HasHasDisks returns a boolean if a field has been set.

### GetHasHosts

`func (o *StorageServerType) GetHasHosts() bool`

GetHasHosts returns the HasHosts field if non-nil, zero value otherwise.

### GetHasHostsOk

`func (o *StorageServerType) GetHasHostsOk() (*bool, bool)`

GetHasHostsOk returns a tuple with the HasHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHosts

`func (o *StorageServerType) SetHasHosts(v bool)`

SetHasHosts sets HasHosts field to given value.

### HasHasHosts

`func (o *StorageServerType) HasHasHosts() bool`

HasHasHosts returns a boolean if a field has been set.

### GetCreateNamespaces

`func (o *StorageServerType) GetCreateNamespaces() bool`

GetCreateNamespaces returns the CreateNamespaces field if non-nil, zero value otherwise.

### GetCreateNamespacesOk

`func (o *StorageServerType) GetCreateNamespacesOk() (*bool, bool)`

GetCreateNamespacesOk returns a tuple with the CreateNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateNamespaces

`func (o *StorageServerType) SetCreateNamespaces(v bool)`

SetCreateNamespaces sets CreateNamespaces field to given value.

### HasCreateNamespaces

`func (o *StorageServerType) HasCreateNamespaces() bool`

HasCreateNamespaces returns a boolean if a field has been set.

### GetCreateGroup

`func (o *StorageServerType) GetCreateGroup() bool`

GetCreateGroup returns the CreateGroup field if non-nil, zero value otherwise.

### GetCreateGroupOk

`func (o *StorageServerType) GetCreateGroupOk() (*bool, bool)`

GetCreateGroupOk returns a tuple with the CreateGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateGroup

`func (o *StorageServerType) SetCreateGroup(v bool)`

SetCreateGroup sets CreateGroup field to given value.

### HasCreateGroup

`func (o *StorageServerType) HasCreateGroup() bool`

HasCreateGroup returns a boolean if a field has been set.

### GetCreateBlock

`func (o *StorageServerType) GetCreateBlock() bool`

GetCreateBlock returns the CreateBlock field if non-nil, zero value otherwise.

### GetCreateBlockOk

`func (o *StorageServerType) GetCreateBlockOk() (*bool, bool)`

GetCreateBlockOk returns a tuple with the CreateBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBlock

`func (o *StorageServerType) SetCreateBlock(v bool)`

SetCreateBlock sets CreateBlock field to given value.

### HasCreateBlock

`func (o *StorageServerType) HasCreateBlock() bool`

HasCreateBlock returns a boolean if a field has been set.

### GetCreateObject

`func (o *StorageServerType) GetCreateObject() bool`

GetCreateObject returns the CreateObject field if non-nil, zero value otherwise.

### GetCreateObjectOk

`func (o *StorageServerType) GetCreateObjectOk() (*bool, bool)`

GetCreateObjectOk returns a tuple with the CreateObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateObject

`func (o *StorageServerType) SetCreateObject(v bool)`

SetCreateObject sets CreateObject field to given value.

### HasCreateObject

`func (o *StorageServerType) HasCreateObject() bool`

HasCreateObject returns a boolean if a field has been set.

### GetCreateFile

`func (o *StorageServerType) GetCreateFile() bool`

GetCreateFile returns the CreateFile field if non-nil, zero value otherwise.

### GetCreateFileOk

`func (o *StorageServerType) GetCreateFileOk() (*bool, bool)`

GetCreateFileOk returns a tuple with the CreateFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateFile

`func (o *StorageServerType) SetCreateFile(v bool)`

SetCreateFile sets CreateFile field to given value.

### HasCreateFile

`func (o *StorageServerType) HasCreateFile() bool`

HasCreateFile returns a boolean if a field has been set.

### GetCreateDatastore

`func (o *StorageServerType) GetCreateDatastore() bool`

GetCreateDatastore returns the CreateDatastore field if non-nil, zero value otherwise.

### GetCreateDatastoreOk

`func (o *StorageServerType) GetCreateDatastoreOk() (*bool, bool)`

GetCreateDatastoreOk returns a tuple with the CreateDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDatastore

`func (o *StorageServerType) SetCreateDatastore(v bool)`

SetCreateDatastore sets CreateDatastore field to given value.

### HasCreateDatastore

`func (o *StorageServerType) HasCreateDatastore() bool`

HasCreateDatastore returns a boolean if a field has been set.

### GetCreateDisk

`func (o *StorageServerType) GetCreateDisk() bool`

GetCreateDisk returns the CreateDisk field if non-nil, zero value otherwise.

### GetCreateDiskOk

`func (o *StorageServerType) GetCreateDiskOk() (*bool, bool)`

GetCreateDiskOk returns a tuple with the CreateDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDisk

`func (o *StorageServerType) SetCreateDisk(v bool)`

SetCreateDisk sets CreateDisk field to given value.

### HasCreateDisk

`func (o *StorageServerType) HasCreateDisk() bool`

HasCreateDisk returns a boolean if a field has been set.

### GetCreateHost

`func (o *StorageServerType) GetCreateHost() bool`

GetCreateHost returns the CreateHost field if non-nil, zero value otherwise.

### GetCreateHostOk

`func (o *StorageServerType) GetCreateHostOk() (*bool, bool)`

GetCreateHostOk returns a tuple with the CreateHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateHost

`func (o *StorageServerType) SetCreateHost(v bool)`

SetCreateHost sets CreateHost field to given value.

### HasCreateHost

`func (o *StorageServerType) HasCreateHost() bool`

HasCreateHost returns a boolean if a field has been set.

### GetIconCode

`func (o *StorageServerType) GetIconCode() string`

GetIconCode returns the IconCode field if non-nil, zero value otherwise.

### GetIconCodeOk

`func (o *StorageServerType) GetIconCodeOk() (*string, bool)`

GetIconCodeOk returns a tuple with the IconCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconCode

`func (o *StorageServerType) SetIconCode(v string)`

SetIconCode sets IconCode field to given value.

### HasIconCode

`func (o *StorageServerType) HasIconCode() bool`

HasIconCode returns a boolean if a field has been set.

### GetHasFileBrowser

`func (o *StorageServerType) GetHasFileBrowser() bool`

GetHasFileBrowser returns the HasFileBrowser field if non-nil, zero value otherwise.

### GetHasFileBrowserOk

`func (o *StorageServerType) GetHasFileBrowserOk() (*bool, bool)`

GetHasFileBrowserOk returns a tuple with the HasFileBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFileBrowser

`func (o *StorageServerType) SetHasFileBrowser(v bool)`

SetHasFileBrowser sets HasFileBrowser field to given value.

### HasHasFileBrowser

`func (o *StorageServerType) HasHasFileBrowser() bool`

HasHasFileBrowser returns a boolean if a field has been set.

### GetOptionTypes

`func (o *StorageServerType) GetOptionTypes() []ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *StorageServerType) GetOptionTypesOk() (*[]ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *StorageServerType) SetOptionTypes(v []ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *StorageServerType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetGroupOptionTypes

`func (o *StorageServerType) GetGroupOptionTypes() []ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerGroupOptionTypesInner`

GetGroupOptionTypes returns the GroupOptionTypes field if non-nil, zero value otherwise.

### GetGroupOptionTypesOk

`func (o *StorageServerType) GetGroupOptionTypesOk() (*[]ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerGroupOptionTypesInner, bool)`

GetGroupOptionTypesOk returns a tuple with the GroupOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupOptionTypes

`func (o *StorageServerType) SetGroupOptionTypes(v []ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerGroupOptionTypesInner)`

SetGroupOptionTypes sets GroupOptionTypes field to given value.

### HasGroupOptionTypes

`func (o *StorageServerType) HasGroupOptionTypes() bool`

HasGroupOptionTypes returns a boolean if a field has been set.

### GetBucketOptionTypes

`func (o *StorageServerType) GetBucketOptionTypes() []map[string]interface{}`

GetBucketOptionTypes returns the BucketOptionTypes field if non-nil, zero value otherwise.

### GetBucketOptionTypesOk

`func (o *StorageServerType) GetBucketOptionTypesOk() (*[]map[string]interface{}, bool)`

GetBucketOptionTypesOk returns a tuple with the BucketOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketOptionTypes

`func (o *StorageServerType) SetBucketOptionTypes(v []map[string]interface{})`

SetBucketOptionTypes sets BucketOptionTypes field to given value.

### HasBucketOptionTypes

`func (o *StorageServerType) HasBucketOptionTypes() bool`

HasBucketOptionTypes returns a boolean if a field has been set.

### GetShareOptionTypes

`func (o *StorageServerType) GetShareOptionTypes() []map[string]interface{}`

GetShareOptionTypes returns the ShareOptionTypes field if non-nil, zero value otherwise.

### GetShareOptionTypesOk

`func (o *StorageServerType) GetShareOptionTypesOk() (*[]map[string]interface{}, bool)`

GetShareOptionTypesOk returns a tuple with the ShareOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareOptionTypes

`func (o *StorageServerType) SetShareOptionTypes(v []map[string]interface{})`

SetShareOptionTypes sets ShareOptionTypes field to given value.

### HasShareOptionTypes

`func (o *StorageServerType) HasShareOptionTypes() bool`

HasShareOptionTypes returns a boolean if a field has been set.

### GetShareAccessOptionTypes

`func (o *StorageServerType) GetShareAccessOptionTypes() []map[string]interface{}`

GetShareAccessOptionTypes returns the ShareAccessOptionTypes field if non-nil, zero value otherwise.

### GetShareAccessOptionTypesOk

`func (o *StorageServerType) GetShareAccessOptionTypesOk() (*[]map[string]interface{}, bool)`

GetShareAccessOptionTypesOk returns a tuple with the ShareAccessOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareAccessOptionTypes

`func (o *StorageServerType) SetShareAccessOptionTypes(v []map[string]interface{})`

SetShareAccessOptionTypes sets ShareAccessOptionTypes field to given value.

### HasShareAccessOptionTypes

`func (o *StorageServerType) HasShareAccessOptionTypes() bool`

HasShareAccessOptionTypes returns a boolean if a field has been set.

### GetStorageVolumeTypes

`func (o *StorageServerType) GetStorageVolumeTypes() []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner`

GetStorageVolumeTypes returns the StorageVolumeTypes field if non-nil, zero value otherwise.

### GetStorageVolumeTypesOk

`func (o *StorageServerType) GetStorageVolumeTypesOk() (*[]GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner, bool)`

GetStorageVolumeTypesOk returns a tuple with the StorageVolumeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVolumeTypes

`func (o *StorageServerType) SetStorageVolumeTypes(v []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner)`

SetStorageVolumeTypes sets StorageVolumeTypes field to given value.

### HasStorageVolumeTypes

`func (o *StorageServerType) HasStorageVolumeTypes() bool`

HasStorageVolumeTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


