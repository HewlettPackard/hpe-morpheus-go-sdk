# VirtualImageLocation

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

### NewVirtualImageLocation

`func NewVirtualImageLocation() *VirtualImageLocation`

NewVirtualImageLocation instantiates a new VirtualImageLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualImageLocationWithDefaults

`func NewVirtualImageLocationWithDefaults() *VirtualImageLocation`

NewVirtualImageLocationWithDefaults instantiates a new VirtualImageLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualImageLocation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualImageLocation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualImageLocation) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualImageLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCloud

`func (o *VirtualImageLocation) GetCloud() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *VirtualImageLocation) GetCloudOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *VirtualImageLocation) SetCloud(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *VirtualImageLocation) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetCode

`func (o *VirtualImageLocation) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VirtualImageLocation) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VirtualImageLocation) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VirtualImageLocation) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetInternalId

`func (o *VirtualImageLocation) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *VirtualImageLocation) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *VirtualImageLocation) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *VirtualImageLocation) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *VirtualImageLocation) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *VirtualImageLocation) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *VirtualImageLocation) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *VirtualImageLocation) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalDiskId

`func (o *VirtualImageLocation) GetExternalDiskId() string`

GetExternalDiskId returns the ExternalDiskId field if non-nil, zero value otherwise.

### GetExternalDiskIdOk

`func (o *VirtualImageLocation) GetExternalDiskIdOk() (*string, bool)`

GetExternalDiskIdOk returns a tuple with the ExternalDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDiskId

`func (o *VirtualImageLocation) SetExternalDiskId(v string)`

SetExternalDiskId sets ExternalDiskId field to given value.

### HasExternalDiskId

`func (o *VirtualImageLocation) HasExternalDiskId() bool`

HasExternalDiskId returns a boolean if a field has been set.

### GetRemotePath

`func (o *VirtualImageLocation) GetRemotePath() string`

GetRemotePath returns the RemotePath field if non-nil, zero value otherwise.

### GetRemotePathOk

`func (o *VirtualImageLocation) GetRemotePathOk() (*string, bool)`

GetRemotePathOk returns a tuple with the RemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePath

`func (o *VirtualImageLocation) SetRemotePath(v string)`

SetRemotePath sets RemotePath field to given value.

### HasRemotePath

`func (o *VirtualImageLocation) HasRemotePath() bool`

HasRemotePath returns a boolean if a field has been set.

### GetImagePath

`func (o *VirtualImageLocation) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *VirtualImageLocation) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *VirtualImageLocation) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *VirtualImageLocation) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### GetImageName

`func (o *VirtualImageLocation) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *VirtualImageLocation) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *VirtualImageLocation) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *VirtualImageLocation) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageRegion

`func (o *VirtualImageLocation) GetImageRegion() string`

GetImageRegion returns the ImageRegion field if non-nil, zero value otherwise.

### GetImageRegionOk

`func (o *VirtualImageLocation) GetImageRegionOk() (*string, bool)`

GetImageRegionOk returns a tuple with the ImageRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegion

`func (o *VirtualImageLocation) SetImageRegion(v string)`

SetImageRegion sets ImageRegion field to given value.

### HasImageRegion

`func (o *VirtualImageLocation) HasImageRegion() bool`

HasImageRegion returns a boolean if a field has been set.

### GetImageFolder

`func (o *VirtualImageLocation) GetImageFolder() string`

GetImageFolder returns the ImageFolder field if non-nil, zero value otherwise.

### GetImageFolderOk

`func (o *VirtualImageLocation) GetImageFolderOk() (*string, bool)`

GetImageFolderOk returns a tuple with the ImageFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFolder

`func (o *VirtualImageLocation) SetImageFolder(v string)`

SetImageFolder sets ImageFolder field to given value.

### HasImageFolder

`func (o *VirtualImageLocation) HasImageFolder() bool`

HasImageFolder returns a boolean if a field has been set.

### GetRefType

`func (o *VirtualImageLocation) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *VirtualImageLocation) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *VirtualImageLocation) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *VirtualImageLocation) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *VirtualImageLocation) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *VirtualImageLocation) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *VirtualImageLocation) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *VirtualImageLocation) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetNodeRefType

`func (o *VirtualImageLocation) GetNodeRefType() string`

GetNodeRefType returns the NodeRefType field if non-nil, zero value otherwise.

### GetNodeRefTypeOk

`func (o *VirtualImageLocation) GetNodeRefTypeOk() (*string, bool)`

GetNodeRefTypeOk returns a tuple with the NodeRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeRefType

`func (o *VirtualImageLocation) SetNodeRefType(v string)`

SetNodeRefType sets NodeRefType field to given value.

### HasNodeRefType

`func (o *VirtualImageLocation) HasNodeRefType() bool`

HasNodeRefType returns a boolean if a field has been set.

### GetNodeRefId

`func (o *VirtualImageLocation) GetNodeRefId() string`

GetNodeRefId returns the NodeRefId field if non-nil, zero value otherwise.

### GetNodeRefIdOk

`func (o *VirtualImageLocation) GetNodeRefIdOk() (*string, bool)`

GetNodeRefIdOk returns a tuple with the NodeRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeRefId

`func (o *VirtualImageLocation) SetNodeRefId(v string)`

SetNodeRefId sets NodeRefId field to given value.

### HasNodeRefId

`func (o *VirtualImageLocation) HasNodeRefId() bool`

HasNodeRefId returns a boolean if a field has been set.

### GetSubRefType

`func (o *VirtualImageLocation) GetSubRefType() string`

GetSubRefType returns the SubRefType field if non-nil, zero value otherwise.

### GetSubRefTypeOk

`func (o *VirtualImageLocation) GetSubRefTypeOk() (*string, bool)`

GetSubRefTypeOk returns a tuple with the SubRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRefType

`func (o *VirtualImageLocation) SetSubRefType(v string)`

SetSubRefType sets SubRefType field to given value.

### HasSubRefType

`func (o *VirtualImageLocation) HasSubRefType() bool`

HasSubRefType returns a boolean if a field has been set.

### GetSubRefId

`func (o *VirtualImageLocation) GetSubRefId() string`

GetSubRefId returns the SubRefId field if non-nil, zero value otherwise.

### GetSubRefIdOk

`func (o *VirtualImageLocation) GetSubRefIdOk() (*string, bool)`

GetSubRefIdOk returns a tuple with the SubRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRefId

`func (o *VirtualImageLocation) SetSubRefId(v string)`

SetSubRefId sets SubRefId field to given value.

### HasSubRefId

`func (o *VirtualImageLocation) HasSubRefId() bool`

HasSubRefId returns a boolean if a field has been set.

### GetIsPublic

`func (o *VirtualImageLocation) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *VirtualImageLocation) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *VirtualImageLocation) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *VirtualImageLocation) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetSystemImage

`func (o *VirtualImageLocation) GetSystemImage() bool`

GetSystemImage returns the SystemImage field if non-nil, zero value otherwise.

### GetSystemImageOk

`func (o *VirtualImageLocation) GetSystemImageOk() (*bool, bool)`

GetSystemImageOk returns a tuple with the SystemImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemImage

`func (o *VirtualImageLocation) SetSystemImage(v bool)`

SetSystemImage sets SystemImage field to given value.

### HasSystemImage

`func (o *VirtualImageLocation) HasSystemImage() bool`

HasSystemImage returns a boolean if a field has been set.

### GetDiskIndex

`func (o *VirtualImageLocation) GetDiskIndex() int64`

GetDiskIndex returns the DiskIndex field if non-nil, zero value otherwise.

### GetDiskIndexOk

`func (o *VirtualImageLocation) GetDiskIndexOk() (*int64, bool)`

GetDiskIndexOk returns a tuple with the DiskIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIndex

`func (o *VirtualImageLocation) SetDiskIndex(v int64)`

SetDiskIndex sets DiskIndex field to given value.

### HasDiskIndex

`func (o *VirtualImageLocation) HasDiskIndex() bool`

HasDiskIndex returns a boolean if a field has been set.

### GetPricePlan

`func (o *VirtualImageLocation) GetPricePlan() string`

GetPricePlan returns the PricePlan field if non-nil, zero value otherwise.

### GetPricePlanOk

`func (o *VirtualImageLocation) GetPricePlanOk() (*string, bool)`

GetPricePlanOk returns a tuple with the PricePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePlan

`func (o *VirtualImageLocation) SetPricePlan(v string)`

SetPricePlan sets PricePlan field to given value.

### HasPricePlan

`func (o *VirtualImageLocation) HasPricePlan() bool`

HasPricePlan returns a boolean if a field has been set.

### GetVolumes

`func (o *VirtualImageLocation) GetVolumes() []map[string]interface{}`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *VirtualImageLocation) GetVolumesOk() (*[]map[string]interface{}, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *VirtualImageLocation) SetVolumes(v []map[string]interface{})`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *VirtualImageLocation) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetStorageControllers

`func (o *VirtualImageLocation) GetStorageControllers() []map[string]interface{}`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *VirtualImageLocation) GetStorageControllersOk() (*[]map[string]interface{}, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *VirtualImageLocation) SetStorageControllers(v []map[string]interface{})`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *VirtualImageLocation) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *VirtualImageLocation) GetNetworkInterfaces() []map[string]interface{}`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *VirtualImageLocation) GetNetworkInterfacesOk() (*[]map[string]interface{}, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *VirtualImageLocation) SetNetworkInterfaces(v []map[string]interface{})`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *VirtualImageLocation) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetVirtualImage

`func (o *VirtualImageLocation) GetVirtualImage() ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *VirtualImageLocation) GetVirtualImageOk() (*ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *VirtualImageLocation) SetVirtualImage(v ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage)`

SetVirtualImage sets VirtualImage field to given value.

### HasVirtualImage

`func (o *VirtualImageLocation) HasVirtualImage() bool`

HasVirtualImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


