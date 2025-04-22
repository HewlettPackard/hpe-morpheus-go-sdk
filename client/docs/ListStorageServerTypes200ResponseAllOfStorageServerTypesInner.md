# ListStorageServerTypes200ResponseAllOfStorageServerTypesInner

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

### NewListStorageServerTypes200ResponseAllOfStorageServerTypesInner

`func NewListStorageServerTypes200ResponseAllOfStorageServerTypesInner() *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner`

NewListStorageServerTypes200ResponseAllOfStorageServerTypesInner instantiates a new ListStorageServerTypes200ResponseAllOfStorageServerTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStorageServerTypes200ResponseAllOfStorageServerTypesInnerWithDefaults

`func NewListStorageServerTypes200ResponseAllOfStorageServerTypesInnerWithDefaults() *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner`

NewListStorageServerTypes200ResponseAllOfStorageServerTypesInnerWithDefaults instantiates a new ListStorageServerTypes200ResponseAllOfStorageServerTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatable

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetHasNamespaces

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasNamespaces() bool`

GetHasNamespaces returns the HasNamespaces field if non-nil, zero value otherwise.

### GetHasNamespacesOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasNamespacesOk() (*bool, bool)`

GetHasNamespacesOk returns a tuple with the HasNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNamespaces

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetHasNamespaces(v bool)`

SetHasNamespaces sets HasNamespaces field to given value.

### HasHasNamespaces

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasHasNamespaces() bool`

HasHasNamespaces returns a boolean if a field has been set.

### GetHasGroups

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasGroups() bool`

GetHasGroups returns the HasGroups field if non-nil, zero value otherwise.

### GetHasGroupsOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasGroupsOk() (*bool, bool)`

GetHasGroupsOk returns a tuple with the HasGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGroups

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetHasGroups(v bool)`

SetHasGroups sets HasGroups field to given value.

### HasHasGroups

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasHasGroups() bool`

HasHasGroups returns a boolean if a field has been set.

### GetHasBlock

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasBlock() bool`

GetHasBlock returns the HasBlock field if non-nil, zero value otherwise.

### GetHasBlockOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasBlockOk() (*bool, bool)`

GetHasBlockOk returns a tuple with the HasBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBlock

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetHasBlock(v bool)`

SetHasBlock sets HasBlock field to given value.

### HasHasBlock

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasHasBlock() bool`

HasHasBlock returns a boolean if a field has been set.

### GetHasObject

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasObject() bool`

GetHasObject returns the HasObject field if non-nil, zero value otherwise.

### GetHasObjectOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasObjectOk() (*bool, bool)`

GetHasObjectOk returns a tuple with the HasObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasObject

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetHasObject(v bool)`

SetHasObject sets HasObject field to given value.

### HasHasObject

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasHasObject() bool`

HasHasObject returns a boolean if a field has been set.

### GetHasFile

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasFile() bool`

GetHasFile returns the HasFile field if non-nil, zero value otherwise.

### GetHasFileOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasFileOk() (*bool, bool)`

GetHasFileOk returns a tuple with the HasFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFile

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetHasFile(v bool)`

SetHasFile sets HasFile field to given value.

### HasHasFile

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasHasFile() bool`

HasHasFile returns a boolean if a field has been set.

### GetHasDatastore

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetHasDisks

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasDisks() bool`

GetHasDisks returns the HasDisks field if non-nil, zero value otherwise.

### GetHasDisksOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasDisksOk() (*bool, bool)`

GetHasDisksOk returns a tuple with the HasDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDisks

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetHasDisks(v bool)`

SetHasDisks sets HasDisks field to given value.

### HasHasDisks

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasHasDisks() bool`

HasHasDisks returns a boolean if a field has been set.

### GetHasHosts

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasHosts() bool`

GetHasHosts returns the HasHosts field if non-nil, zero value otherwise.

### GetHasHostsOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasHostsOk() (*bool, bool)`

GetHasHostsOk returns a tuple with the HasHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHosts

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetHasHosts(v bool)`

SetHasHosts sets HasHosts field to given value.

### HasHasHosts

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasHasHosts() bool`

HasHasHosts returns a boolean if a field has been set.

### GetCreateNamespaces

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateNamespaces() bool`

GetCreateNamespaces returns the CreateNamespaces field if non-nil, zero value otherwise.

### GetCreateNamespacesOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateNamespacesOk() (*bool, bool)`

GetCreateNamespacesOk returns a tuple with the CreateNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateNamespaces

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetCreateNamespaces(v bool)`

SetCreateNamespaces sets CreateNamespaces field to given value.

### HasCreateNamespaces

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasCreateNamespaces() bool`

HasCreateNamespaces returns a boolean if a field has been set.

### GetCreateGroup

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateGroup() bool`

GetCreateGroup returns the CreateGroup field if non-nil, zero value otherwise.

### GetCreateGroupOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateGroupOk() (*bool, bool)`

GetCreateGroupOk returns a tuple with the CreateGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateGroup

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetCreateGroup(v bool)`

SetCreateGroup sets CreateGroup field to given value.

### HasCreateGroup

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasCreateGroup() bool`

HasCreateGroup returns a boolean if a field has been set.

### GetCreateBlock

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateBlock() bool`

GetCreateBlock returns the CreateBlock field if non-nil, zero value otherwise.

### GetCreateBlockOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateBlockOk() (*bool, bool)`

GetCreateBlockOk returns a tuple with the CreateBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBlock

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetCreateBlock(v bool)`

SetCreateBlock sets CreateBlock field to given value.

### HasCreateBlock

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasCreateBlock() bool`

HasCreateBlock returns a boolean if a field has been set.

### GetCreateObject

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateObject() bool`

GetCreateObject returns the CreateObject field if non-nil, zero value otherwise.

### GetCreateObjectOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateObjectOk() (*bool, bool)`

GetCreateObjectOk returns a tuple with the CreateObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateObject

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetCreateObject(v bool)`

SetCreateObject sets CreateObject field to given value.

### HasCreateObject

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasCreateObject() bool`

HasCreateObject returns a boolean if a field has been set.

### GetCreateFile

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateFile() bool`

GetCreateFile returns the CreateFile field if non-nil, zero value otherwise.

### GetCreateFileOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateFileOk() (*bool, bool)`

GetCreateFileOk returns a tuple with the CreateFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateFile

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetCreateFile(v bool)`

SetCreateFile sets CreateFile field to given value.

### HasCreateFile

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasCreateFile() bool`

HasCreateFile returns a boolean if a field has been set.

### GetCreateDatastore

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateDatastore() bool`

GetCreateDatastore returns the CreateDatastore field if non-nil, zero value otherwise.

### GetCreateDatastoreOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateDatastoreOk() (*bool, bool)`

GetCreateDatastoreOk returns a tuple with the CreateDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDatastore

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetCreateDatastore(v bool)`

SetCreateDatastore sets CreateDatastore field to given value.

### HasCreateDatastore

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasCreateDatastore() bool`

HasCreateDatastore returns a boolean if a field has been set.

### GetCreateDisk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateDisk() bool`

GetCreateDisk returns the CreateDisk field if non-nil, zero value otherwise.

### GetCreateDiskOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateDiskOk() (*bool, bool)`

GetCreateDiskOk returns a tuple with the CreateDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDisk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetCreateDisk(v bool)`

SetCreateDisk sets CreateDisk field to given value.

### HasCreateDisk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasCreateDisk() bool`

HasCreateDisk returns a boolean if a field has been set.

### GetCreateHost

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateHost() bool`

GetCreateHost returns the CreateHost field if non-nil, zero value otherwise.

### GetCreateHostOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetCreateHostOk() (*bool, bool)`

GetCreateHostOk returns a tuple with the CreateHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateHost

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetCreateHost(v bool)`

SetCreateHost sets CreateHost field to given value.

### HasCreateHost

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasCreateHost() bool`

HasCreateHost returns a boolean if a field has been set.

### GetIconCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetIconCode() string`

GetIconCode returns the IconCode field if non-nil, zero value otherwise.

### GetIconCodeOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetIconCodeOk() (*string, bool)`

GetIconCodeOk returns a tuple with the IconCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetIconCode(v string)`

SetIconCode sets IconCode field to given value.

### HasIconCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasIconCode() bool`

HasIconCode returns a boolean if a field has been set.

### GetHasFileBrowser

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasFileBrowser() bool`

GetHasFileBrowser returns the HasFileBrowser field if non-nil, zero value otherwise.

### GetHasFileBrowserOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetHasFileBrowserOk() (*bool, bool)`

GetHasFileBrowserOk returns a tuple with the HasFileBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFileBrowser

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetHasFileBrowser(v bool)`

SetHasFileBrowser sets HasFileBrowser field to given value.

### HasHasFileBrowser

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasHasFileBrowser() bool`

HasHasFileBrowser returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetOptionTypes() []ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetOptionTypesOk() (*[]ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetOptionTypes(v []ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetGroupOptionTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetGroupOptionTypes() []ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerGroupOptionTypesInner`

GetGroupOptionTypes returns the GroupOptionTypes field if non-nil, zero value otherwise.

### GetGroupOptionTypesOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetGroupOptionTypesOk() (*[]ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerGroupOptionTypesInner, bool)`

GetGroupOptionTypesOk returns a tuple with the GroupOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupOptionTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetGroupOptionTypes(v []ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerGroupOptionTypesInner)`

SetGroupOptionTypes sets GroupOptionTypes field to given value.

### HasGroupOptionTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasGroupOptionTypes() bool`

HasGroupOptionTypes returns a boolean if a field has been set.

### GetBucketOptionTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetBucketOptionTypes() []map[string]interface{}`

GetBucketOptionTypes returns the BucketOptionTypes field if non-nil, zero value otherwise.

### GetBucketOptionTypesOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetBucketOptionTypesOk() (*[]map[string]interface{}, bool)`

GetBucketOptionTypesOk returns a tuple with the BucketOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketOptionTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetBucketOptionTypes(v []map[string]interface{})`

SetBucketOptionTypes sets BucketOptionTypes field to given value.

### HasBucketOptionTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasBucketOptionTypes() bool`

HasBucketOptionTypes returns a boolean if a field has been set.

### GetShareOptionTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetShareOptionTypes() []map[string]interface{}`

GetShareOptionTypes returns the ShareOptionTypes field if non-nil, zero value otherwise.

### GetShareOptionTypesOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetShareOptionTypesOk() (*[]map[string]interface{}, bool)`

GetShareOptionTypesOk returns a tuple with the ShareOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareOptionTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetShareOptionTypes(v []map[string]interface{})`

SetShareOptionTypes sets ShareOptionTypes field to given value.

### HasShareOptionTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasShareOptionTypes() bool`

HasShareOptionTypes returns a boolean if a field has been set.

### GetShareAccessOptionTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetShareAccessOptionTypes() []map[string]interface{}`

GetShareAccessOptionTypes returns the ShareAccessOptionTypes field if non-nil, zero value otherwise.

### GetShareAccessOptionTypesOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetShareAccessOptionTypesOk() (*[]map[string]interface{}, bool)`

GetShareAccessOptionTypesOk returns a tuple with the ShareAccessOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareAccessOptionTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetShareAccessOptionTypes(v []map[string]interface{})`

SetShareAccessOptionTypes sets ShareAccessOptionTypes field to given value.

### HasShareAccessOptionTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasShareAccessOptionTypes() bool`

HasShareAccessOptionTypes returns a boolean if a field has been set.

### GetStorageVolumeTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetStorageVolumeTypes() []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner`

GetStorageVolumeTypes returns the StorageVolumeTypes field if non-nil, zero value otherwise.

### GetStorageVolumeTypesOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) GetStorageVolumeTypesOk() (*[]GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner, bool)`

GetStorageVolumeTypesOk returns a tuple with the StorageVolumeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVolumeTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) SetStorageVolumeTypes(v []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionTypeStorageTypesInner)`

SetStorageVolumeTypes sets StorageVolumeTypes field to given value.

### HasStorageVolumeTypes

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInner) HasStorageVolumeTypes() bool`

HasStorageVolumeTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


