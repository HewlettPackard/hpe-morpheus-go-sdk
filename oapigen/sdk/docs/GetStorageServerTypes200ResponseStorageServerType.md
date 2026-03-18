# GetStorageServerTypes200ResponseStorageServerType

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
**IconCode** | Pointer to **NullableString** |  | [optional] 
**HasFileBrowser** | Pointer to **bool** |  | [optional] 
**OptionTypes** | Pointer to [**[]GetStorageServerTypes200ResponseStorageServerTypeOptionTypesInner**](GetStorageServerTypes200ResponseStorageServerTypeOptionTypesInner.md) |  | [optional] 
**GroupOptionTypes** | Pointer to [**[]GetStorageServerTypes200ResponseStorageServerTypeGroupOptionTypesInner**](GetStorageServerTypes200ResponseStorageServerTypeGroupOptionTypesInner.md) |  | [optional] 
**BucketOptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ShareOptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ShareAccessOptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**StorageVolumeTypes** | Pointer to [**[]GetStorageServerTypes200ResponseStorageServerTypeStorageVolumeTypesInner**](GetStorageServerTypes200ResponseStorageServerTypeStorageVolumeTypesInner.md) |  | [optional] 

## Methods

### NewGetStorageServerTypes200ResponseStorageServerType

`func NewGetStorageServerTypes200ResponseStorageServerType() *GetStorageServerTypes200ResponseStorageServerType`

NewGetStorageServerTypes200ResponseStorageServerType instantiates a new GetStorageServerTypes200ResponseStorageServerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStorageServerTypes200ResponseStorageServerTypeWithDefaults

`func NewGetStorageServerTypes200ResponseStorageServerTypeWithDefaults() *GetStorageServerTypes200ResponseStorageServerType`

NewGetStorageServerTypes200ResponseStorageServerTypeWithDefaults instantiates a new GetStorageServerTypes200ResponseStorageServerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatable

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetHasNamespaces

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasNamespaces() bool`

GetHasNamespaces returns the HasNamespaces field if non-nil, zero value otherwise.

### GetHasNamespacesOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasNamespacesOk() (*bool, bool)`

GetHasNamespacesOk returns a tuple with the HasNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNamespaces

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetHasNamespaces(v bool)`

SetHasNamespaces sets HasNamespaces field to given value.

### HasHasNamespaces

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasHasNamespaces() bool`

HasHasNamespaces returns a boolean if a field has been set.

### GetHasGroups

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasGroups() bool`

GetHasGroups returns the HasGroups field if non-nil, zero value otherwise.

### GetHasGroupsOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasGroupsOk() (*bool, bool)`

GetHasGroupsOk returns a tuple with the HasGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGroups

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetHasGroups(v bool)`

SetHasGroups sets HasGroups field to given value.

### HasHasGroups

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasHasGroups() bool`

HasHasGroups returns a boolean if a field has been set.

### GetHasBlock

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasBlock() bool`

GetHasBlock returns the HasBlock field if non-nil, zero value otherwise.

### GetHasBlockOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasBlockOk() (*bool, bool)`

GetHasBlockOk returns a tuple with the HasBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBlock

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetHasBlock(v bool)`

SetHasBlock sets HasBlock field to given value.

### HasHasBlock

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasHasBlock() bool`

HasHasBlock returns a boolean if a field has been set.

### GetHasObject

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasObject() bool`

GetHasObject returns the HasObject field if non-nil, zero value otherwise.

### GetHasObjectOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasObjectOk() (*bool, bool)`

GetHasObjectOk returns a tuple with the HasObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasObject

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetHasObject(v bool)`

SetHasObject sets HasObject field to given value.

### HasHasObject

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasHasObject() bool`

HasHasObject returns a boolean if a field has been set.

### GetHasFile

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasFile() bool`

GetHasFile returns the HasFile field if non-nil, zero value otherwise.

### GetHasFileOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasFileOk() (*bool, bool)`

GetHasFileOk returns a tuple with the HasFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFile

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetHasFile(v bool)`

SetHasFile sets HasFile field to given value.

### HasHasFile

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasHasFile() bool`

HasHasFile returns a boolean if a field has been set.

### GetHasDatastore

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetHasDisks

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasDisks() bool`

GetHasDisks returns the HasDisks field if non-nil, zero value otherwise.

### GetHasDisksOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasDisksOk() (*bool, bool)`

GetHasDisksOk returns a tuple with the HasDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDisks

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetHasDisks(v bool)`

SetHasDisks sets HasDisks field to given value.

### HasHasDisks

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasHasDisks() bool`

HasHasDisks returns a boolean if a field has been set.

### GetHasHosts

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasHosts() bool`

GetHasHosts returns the HasHosts field if non-nil, zero value otherwise.

### GetHasHostsOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasHostsOk() (*bool, bool)`

GetHasHostsOk returns a tuple with the HasHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHosts

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetHasHosts(v bool)`

SetHasHosts sets HasHosts field to given value.

### HasHasHosts

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasHasHosts() bool`

HasHasHosts returns a boolean if a field has been set.

### GetCreateNamespaces

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateNamespaces() bool`

GetCreateNamespaces returns the CreateNamespaces field if non-nil, zero value otherwise.

### GetCreateNamespacesOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateNamespacesOk() (*bool, bool)`

GetCreateNamespacesOk returns a tuple with the CreateNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateNamespaces

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetCreateNamespaces(v bool)`

SetCreateNamespaces sets CreateNamespaces field to given value.

### HasCreateNamespaces

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasCreateNamespaces() bool`

HasCreateNamespaces returns a boolean if a field has been set.

### GetCreateGroup

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateGroup() bool`

GetCreateGroup returns the CreateGroup field if non-nil, zero value otherwise.

### GetCreateGroupOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateGroupOk() (*bool, bool)`

GetCreateGroupOk returns a tuple with the CreateGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateGroup

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetCreateGroup(v bool)`

SetCreateGroup sets CreateGroup field to given value.

### HasCreateGroup

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasCreateGroup() bool`

HasCreateGroup returns a boolean if a field has been set.

### GetCreateBlock

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateBlock() bool`

GetCreateBlock returns the CreateBlock field if non-nil, zero value otherwise.

### GetCreateBlockOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateBlockOk() (*bool, bool)`

GetCreateBlockOk returns a tuple with the CreateBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBlock

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetCreateBlock(v bool)`

SetCreateBlock sets CreateBlock field to given value.

### HasCreateBlock

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasCreateBlock() bool`

HasCreateBlock returns a boolean if a field has been set.

### GetCreateObject

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateObject() bool`

GetCreateObject returns the CreateObject field if non-nil, zero value otherwise.

### GetCreateObjectOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateObjectOk() (*bool, bool)`

GetCreateObjectOk returns a tuple with the CreateObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateObject

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetCreateObject(v bool)`

SetCreateObject sets CreateObject field to given value.

### HasCreateObject

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasCreateObject() bool`

HasCreateObject returns a boolean if a field has been set.

### GetCreateFile

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateFile() bool`

GetCreateFile returns the CreateFile field if non-nil, zero value otherwise.

### GetCreateFileOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateFileOk() (*bool, bool)`

GetCreateFileOk returns a tuple with the CreateFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateFile

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetCreateFile(v bool)`

SetCreateFile sets CreateFile field to given value.

### HasCreateFile

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasCreateFile() bool`

HasCreateFile returns a boolean if a field has been set.

### GetCreateDatastore

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateDatastore() bool`

GetCreateDatastore returns the CreateDatastore field if non-nil, zero value otherwise.

### GetCreateDatastoreOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateDatastoreOk() (*bool, bool)`

GetCreateDatastoreOk returns a tuple with the CreateDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDatastore

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetCreateDatastore(v bool)`

SetCreateDatastore sets CreateDatastore field to given value.

### HasCreateDatastore

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasCreateDatastore() bool`

HasCreateDatastore returns a boolean if a field has been set.

### GetCreateDisk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateDisk() bool`

GetCreateDisk returns the CreateDisk field if non-nil, zero value otherwise.

### GetCreateDiskOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateDiskOk() (*bool, bool)`

GetCreateDiskOk returns a tuple with the CreateDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDisk

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetCreateDisk(v bool)`

SetCreateDisk sets CreateDisk field to given value.

### HasCreateDisk

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasCreateDisk() bool`

HasCreateDisk returns a boolean if a field has been set.

### GetCreateHost

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateHost() bool`

GetCreateHost returns the CreateHost field if non-nil, zero value otherwise.

### GetCreateHostOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetCreateHostOk() (*bool, bool)`

GetCreateHostOk returns a tuple with the CreateHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateHost

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetCreateHost(v bool)`

SetCreateHost sets CreateHost field to given value.

### HasCreateHost

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasCreateHost() bool`

HasCreateHost returns a boolean if a field has been set.

### GetIconCode

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetIconCode() string`

GetIconCode returns the IconCode field if non-nil, zero value otherwise.

### GetIconCodeOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetIconCodeOk() (*string, bool)`

GetIconCodeOk returns a tuple with the IconCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconCode

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetIconCode(v string)`

SetIconCode sets IconCode field to given value.

### HasIconCode

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasIconCode() bool`

HasIconCode returns a boolean if a field has been set.

### SetIconCodeNil

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetIconCodeNil(b bool)`

 SetIconCodeNil sets the value for IconCode to be an explicit nil

### UnsetIconCode
`func (o *GetStorageServerTypes200ResponseStorageServerType) UnsetIconCode()`

UnsetIconCode ensures that no value is present for IconCode, not even an explicit nil
### GetHasFileBrowser

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasFileBrowser() bool`

GetHasFileBrowser returns the HasFileBrowser field if non-nil, zero value otherwise.

### GetHasFileBrowserOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetHasFileBrowserOk() (*bool, bool)`

GetHasFileBrowserOk returns a tuple with the HasFileBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFileBrowser

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetHasFileBrowser(v bool)`

SetHasFileBrowser sets HasFileBrowser field to given value.

### HasHasFileBrowser

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasHasFileBrowser() bool`

HasHasFileBrowser returns a boolean if a field has been set.

### GetOptionTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetOptionTypes() []GetStorageServerTypes200ResponseStorageServerTypeOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetOptionTypesOk() (*[]GetStorageServerTypes200ResponseStorageServerTypeOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetOptionTypes(v []GetStorageServerTypes200ResponseStorageServerTypeOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetGroupOptionTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetGroupOptionTypes() []GetStorageServerTypes200ResponseStorageServerTypeGroupOptionTypesInner`

GetGroupOptionTypes returns the GroupOptionTypes field if non-nil, zero value otherwise.

### GetGroupOptionTypesOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetGroupOptionTypesOk() (*[]GetStorageServerTypes200ResponseStorageServerTypeGroupOptionTypesInner, bool)`

GetGroupOptionTypesOk returns a tuple with the GroupOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupOptionTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetGroupOptionTypes(v []GetStorageServerTypes200ResponseStorageServerTypeGroupOptionTypesInner)`

SetGroupOptionTypes sets GroupOptionTypes field to given value.

### HasGroupOptionTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasGroupOptionTypes() bool`

HasGroupOptionTypes returns a boolean if a field has been set.

### GetBucketOptionTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetBucketOptionTypes() []map[string]interface{}`

GetBucketOptionTypes returns the BucketOptionTypes field if non-nil, zero value otherwise.

### GetBucketOptionTypesOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetBucketOptionTypesOk() (*[]map[string]interface{}, bool)`

GetBucketOptionTypesOk returns a tuple with the BucketOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketOptionTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetBucketOptionTypes(v []map[string]interface{})`

SetBucketOptionTypes sets BucketOptionTypes field to given value.

### HasBucketOptionTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasBucketOptionTypes() bool`

HasBucketOptionTypes returns a boolean if a field has been set.

### SetBucketOptionTypesNil

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetBucketOptionTypesNil(b bool)`

 SetBucketOptionTypesNil sets the value for BucketOptionTypes to be an explicit nil

### UnsetBucketOptionTypes
`func (o *GetStorageServerTypes200ResponseStorageServerType) UnsetBucketOptionTypes()`

UnsetBucketOptionTypes ensures that no value is present for BucketOptionTypes, not even an explicit nil
### GetShareOptionTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetShareOptionTypes() []map[string]interface{}`

GetShareOptionTypes returns the ShareOptionTypes field if non-nil, zero value otherwise.

### GetShareOptionTypesOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetShareOptionTypesOk() (*[]map[string]interface{}, bool)`

GetShareOptionTypesOk returns a tuple with the ShareOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareOptionTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetShareOptionTypes(v []map[string]interface{})`

SetShareOptionTypes sets ShareOptionTypes field to given value.

### HasShareOptionTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasShareOptionTypes() bool`

HasShareOptionTypes returns a boolean if a field has been set.

### SetShareOptionTypesNil

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetShareOptionTypesNil(b bool)`

 SetShareOptionTypesNil sets the value for ShareOptionTypes to be an explicit nil

### UnsetShareOptionTypes
`func (o *GetStorageServerTypes200ResponseStorageServerType) UnsetShareOptionTypes()`

UnsetShareOptionTypes ensures that no value is present for ShareOptionTypes, not even an explicit nil
### GetShareAccessOptionTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetShareAccessOptionTypes() []map[string]interface{}`

GetShareAccessOptionTypes returns the ShareAccessOptionTypes field if non-nil, zero value otherwise.

### GetShareAccessOptionTypesOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetShareAccessOptionTypesOk() (*[]map[string]interface{}, bool)`

GetShareAccessOptionTypesOk returns a tuple with the ShareAccessOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareAccessOptionTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetShareAccessOptionTypes(v []map[string]interface{})`

SetShareAccessOptionTypes sets ShareAccessOptionTypes field to given value.

### HasShareAccessOptionTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasShareAccessOptionTypes() bool`

HasShareAccessOptionTypes returns a boolean if a field has been set.

### SetShareAccessOptionTypesNil

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetShareAccessOptionTypesNil(b bool)`

 SetShareAccessOptionTypesNil sets the value for ShareAccessOptionTypes to be an explicit nil

### UnsetShareAccessOptionTypes
`func (o *GetStorageServerTypes200ResponseStorageServerType) UnsetShareAccessOptionTypes()`

UnsetShareAccessOptionTypes ensures that no value is present for ShareAccessOptionTypes, not even an explicit nil
### GetStorageVolumeTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetStorageVolumeTypes() []GetStorageServerTypes200ResponseStorageServerTypeStorageVolumeTypesInner`

GetStorageVolumeTypes returns the StorageVolumeTypes field if non-nil, zero value otherwise.

### GetStorageVolumeTypesOk

`func (o *GetStorageServerTypes200ResponseStorageServerType) GetStorageVolumeTypesOk() (*[]GetStorageServerTypes200ResponseStorageServerTypeStorageVolumeTypesInner, bool)`

GetStorageVolumeTypesOk returns a tuple with the StorageVolumeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVolumeTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) SetStorageVolumeTypes(v []GetStorageServerTypes200ResponseStorageServerTypeStorageVolumeTypesInner)`

SetStorageVolumeTypes sets StorageVolumeTypes field to given value.

### HasStorageVolumeTypes

`func (o *GetStorageServerTypes200ResponseStorageServerType) HasStorageVolumeTypes() bool`

HasStorageVolumeTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


