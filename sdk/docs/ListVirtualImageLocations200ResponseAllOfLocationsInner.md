# ListVirtualImageLocations200ResponseAllOfLocationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Cloud** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ExternalDiskId** | Pointer to **string** |  | [optional] 
**RemotePath** | Pointer to **string** |  | [optional] 
**ImagePath** | Pointer to **string** |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**ImageRegion** | Pointer to **string** |  | [optional] 
**ImageFolder** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**NodeRefType** | Pointer to **string** |  | [optional] 
**NodeRefId** | Pointer to **string** |  | [optional] 
**SubRefType** | Pointer to **string** |  | [optional] 
**SubRefId** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**SystemImage** | Pointer to **bool** |  | [optional] 
**DiskIndex** | Pointer to **int64** |  | [optional] 
**PricePlan** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**StorageControllers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**NetworkInterfaces** | Pointer to **[]map[string]interface{}** |  | [optional] 
**VirtualImage** | Pointer to [**ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage**](ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage.md) |  | [optional] 

## Methods

### NewListVirtualImageLocations200ResponseAllOfLocationsInner

`func NewListVirtualImageLocations200ResponseAllOfLocationsInner() *ListVirtualImageLocations200ResponseAllOfLocationsInner`

NewListVirtualImageLocations200ResponseAllOfLocationsInner instantiates a new ListVirtualImageLocations200ResponseAllOfLocationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVirtualImageLocations200ResponseAllOfLocationsInnerWithDefaults

`func NewListVirtualImageLocations200ResponseAllOfLocationsInnerWithDefaults() *ListVirtualImageLocations200ResponseAllOfLocationsInner`

NewListVirtualImageLocations200ResponseAllOfLocationsInnerWithDefaults instantiates a new ListVirtualImageLocations200ResponseAllOfLocationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCloud

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetCloud() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetCloudOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetCloud(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetCode

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetInternalId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalDiskId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetExternalDiskId() string`

GetExternalDiskId returns the ExternalDiskId field if non-nil, zero value otherwise.

### GetExternalDiskIdOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetExternalDiskIdOk() (*string, bool)`

GetExternalDiskIdOk returns a tuple with the ExternalDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDiskId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetExternalDiskId(v string)`

SetExternalDiskId sets ExternalDiskId field to given value.

### HasExternalDiskId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasExternalDiskId() bool`

HasExternalDiskId returns a boolean if a field has been set.

### GetRemotePath

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetRemotePath() string`

GetRemotePath returns the RemotePath field if non-nil, zero value otherwise.

### GetRemotePathOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetRemotePathOk() (*string, bool)`

GetRemotePathOk returns a tuple with the RemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePath

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetRemotePath(v string)`

SetRemotePath sets RemotePath field to given value.

### HasRemotePath

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasRemotePath() bool`

HasRemotePath returns a boolean if a field has been set.

### GetImagePath

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### GetImageName

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageRegion

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetImageRegion() string`

GetImageRegion returns the ImageRegion field if non-nil, zero value otherwise.

### GetImageRegionOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetImageRegionOk() (*string, bool)`

GetImageRegionOk returns a tuple with the ImageRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegion

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetImageRegion(v string)`

SetImageRegion sets ImageRegion field to given value.

### HasImageRegion

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasImageRegion() bool`

HasImageRegion returns a boolean if a field has been set.

### GetImageFolder

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetImageFolder() string`

GetImageFolder returns the ImageFolder field if non-nil, zero value otherwise.

### GetImageFolderOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetImageFolderOk() (*string, bool)`

GetImageFolderOk returns a tuple with the ImageFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFolder

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetImageFolder(v string)`

SetImageFolder sets ImageFolder field to given value.

### HasImageFolder

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasImageFolder() bool`

HasImageFolder returns a boolean if a field has been set.

### GetRefType

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetNodeRefType

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetNodeRefType() string`

GetNodeRefType returns the NodeRefType field if non-nil, zero value otherwise.

### GetNodeRefTypeOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetNodeRefTypeOk() (*string, bool)`

GetNodeRefTypeOk returns a tuple with the NodeRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeRefType

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetNodeRefType(v string)`

SetNodeRefType sets NodeRefType field to given value.

### HasNodeRefType

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasNodeRefType() bool`

HasNodeRefType returns a boolean if a field has been set.

### GetNodeRefId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetNodeRefId() string`

GetNodeRefId returns the NodeRefId field if non-nil, zero value otherwise.

### GetNodeRefIdOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetNodeRefIdOk() (*string, bool)`

GetNodeRefIdOk returns a tuple with the NodeRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeRefId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetNodeRefId(v string)`

SetNodeRefId sets NodeRefId field to given value.

### HasNodeRefId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasNodeRefId() bool`

HasNodeRefId returns a boolean if a field has been set.

### GetSubRefType

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetSubRefType() string`

GetSubRefType returns the SubRefType field if non-nil, zero value otherwise.

### GetSubRefTypeOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetSubRefTypeOk() (*string, bool)`

GetSubRefTypeOk returns a tuple with the SubRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRefType

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetSubRefType(v string)`

SetSubRefType sets SubRefType field to given value.

### HasSubRefType

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasSubRefType() bool`

HasSubRefType returns a boolean if a field has been set.

### GetSubRefId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetSubRefId() string`

GetSubRefId returns the SubRefId field if non-nil, zero value otherwise.

### GetSubRefIdOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetSubRefIdOk() (*string, bool)`

GetSubRefIdOk returns a tuple with the SubRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRefId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetSubRefId(v string)`

SetSubRefId sets SubRefId field to given value.

### HasSubRefId

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasSubRefId() bool`

HasSubRefId returns a boolean if a field has been set.

### GetIsPublic

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetSystemImage

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetSystemImage() bool`

GetSystemImage returns the SystemImage field if non-nil, zero value otherwise.

### GetSystemImageOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetSystemImageOk() (*bool, bool)`

GetSystemImageOk returns a tuple with the SystemImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemImage

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetSystemImage(v bool)`

SetSystemImage sets SystemImage field to given value.

### HasSystemImage

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasSystemImage() bool`

HasSystemImage returns a boolean if a field has been set.

### GetDiskIndex

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetDiskIndex() int64`

GetDiskIndex returns the DiskIndex field if non-nil, zero value otherwise.

### GetDiskIndexOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetDiskIndexOk() (*int64, bool)`

GetDiskIndexOk returns a tuple with the DiskIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIndex

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetDiskIndex(v int64)`

SetDiskIndex sets DiskIndex field to given value.

### HasDiskIndex

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasDiskIndex() bool`

HasDiskIndex returns a boolean if a field has been set.

### GetPricePlan

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetPricePlan() string`

GetPricePlan returns the PricePlan field if non-nil, zero value otherwise.

### GetPricePlanOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetPricePlanOk() (*string, bool)`

GetPricePlanOk returns a tuple with the PricePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePlan

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetPricePlan(v string)`

SetPricePlan sets PricePlan field to given value.

### HasPricePlan

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasPricePlan() bool`

HasPricePlan returns a boolean if a field has been set.

### GetVolumes

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetVolumes() []map[string]interface{}`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetVolumesOk() (*[]map[string]interface{}, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetVolumes(v []map[string]interface{})`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetStorageControllers

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetStorageControllers() []map[string]interface{}`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetStorageControllersOk() (*[]map[string]interface{}, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetStorageControllers(v []map[string]interface{})`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetNetworkInterfaces() []map[string]interface{}`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetNetworkInterfacesOk() (*[]map[string]interface{}, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetNetworkInterfaces(v []map[string]interface{})`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetVirtualImage

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetVirtualImage() ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) GetVirtualImageOk() (*ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) SetVirtualImage(v ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage)`

SetVirtualImage sets VirtualImage field to given value.

### HasVirtualImage

`func (o *ListVirtualImageLocations200ResponseAllOfLocationsInner) HasVirtualImage() bool`

HasVirtualImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


