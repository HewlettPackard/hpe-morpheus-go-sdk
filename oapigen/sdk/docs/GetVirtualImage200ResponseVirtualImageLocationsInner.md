# GetVirtualImage200ResponseVirtualImageLocationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Cloud** | Pointer to **map[string]interface{}** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalDiskId** | Pointer to **NullableString** |  | [optional] 
**RemotePath** | Pointer to **NullableString** |  | [optional] 
**ImagePath** | Pointer to **NullableString** |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**ImageRegion** | Pointer to **NullableString** |  | [optional] 
**ImageFolder** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **NullableInt64** |  | [optional] 
**NodeRefType** | Pointer to **NullableString** |  | [optional] 
**SubRefType** | Pointer to **NullableString** |  | [optional] 
**SubRefId** | Pointer to **NullableInt64** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**SystemImage** | Pointer to **bool** |  | [optional] 
**DiskIndex** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetVirtualImage200ResponseVirtualImageLocationsInner

`func NewGetVirtualImage200ResponseVirtualImageLocationsInner() *GetVirtualImage200ResponseVirtualImageLocationsInner`

NewGetVirtualImage200ResponseVirtualImageLocationsInner instantiates a new GetVirtualImage200ResponseVirtualImageLocationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVirtualImage200ResponseVirtualImageLocationsInnerWithDefaults

`func NewGetVirtualImage200ResponseVirtualImageLocationsInnerWithDefaults() *GetVirtualImage200ResponseVirtualImageLocationsInner`

NewGetVirtualImage200ResponseVirtualImageLocationsInnerWithDefaults instantiates a new GetVirtualImage200ResponseVirtualImageLocationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCloud

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetCloud() map[string]interface{}`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetCloudOk() (*map[string]interface{}, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetCloud(v map[string]interface{})`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetCode

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetInternalId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalDiskId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetExternalDiskId() string`

GetExternalDiskId returns the ExternalDiskId field if non-nil, zero value otherwise.

### GetExternalDiskIdOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetExternalDiskIdOk() (*string, bool)`

GetExternalDiskIdOk returns a tuple with the ExternalDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDiskId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetExternalDiskId(v string)`

SetExternalDiskId sets ExternalDiskId field to given value.

### HasExternalDiskId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasExternalDiskId() bool`

HasExternalDiskId returns a boolean if a field has been set.

### SetExternalDiskIdNil

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetExternalDiskIdNil(b bool)`

 SetExternalDiskIdNil sets the value for ExternalDiskId to be an explicit nil

### UnsetExternalDiskId
`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) UnsetExternalDiskId()`

UnsetExternalDiskId ensures that no value is present for ExternalDiskId, not even an explicit nil
### GetRemotePath

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetRemotePath() string`

GetRemotePath returns the RemotePath field if non-nil, zero value otherwise.

### GetRemotePathOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetRemotePathOk() (*string, bool)`

GetRemotePathOk returns a tuple with the RemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePath

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetRemotePath(v string)`

SetRemotePath sets RemotePath field to given value.

### HasRemotePath

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasRemotePath() bool`

HasRemotePath returns a boolean if a field has been set.

### SetRemotePathNil

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetRemotePathNil(b bool)`

 SetRemotePathNil sets the value for RemotePath to be an explicit nil

### UnsetRemotePath
`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) UnsetRemotePath()`

UnsetRemotePath ensures that no value is present for RemotePath, not even an explicit nil
### GetImagePath

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### SetImagePathNil

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetImagePathNil(b bool)`

 SetImagePathNil sets the value for ImagePath to be an explicit nil

### UnsetImagePath
`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) UnsetImagePath()`

UnsetImagePath ensures that no value is present for ImagePath, not even an explicit nil
### GetImageName

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageRegion

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetImageRegion() string`

GetImageRegion returns the ImageRegion field if non-nil, zero value otherwise.

### GetImageRegionOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetImageRegionOk() (*string, bool)`

GetImageRegionOk returns a tuple with the ImageRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegion

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetImageRegion(v string)`

SetImageRegion sets ImageRegion field to given value.

### HasImageRegion

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasImageRegion() bool`

HasImageRegion returns a boolean if a field has been set.

### SetImageRegionNil

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetImageRegionNil(b bool)`

 SetImageRegionNil sets the value for ImageRegion to be an explicit nil

### UnsetImageRegion
`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) UnsetImageRegion()`

UnsetImageRegion ensures that no value is present for ImageRegion, not even an explicit nil
### GetImageFolder

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetImageFolder() string`

GetImageFolder returns the ImageFolder field if non-nil, zero value otherwise.

### GetImageFolderOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetImageFolderOk() (*string, bool)`

GetImageFolderOk returns a tuple with the ImageFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFolder

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetImageFolder(v string)`

SetImageFolder sets ImageFolder field to given value.

### HasImageFolder

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasImageFolder() bool`

HasImageFolder returns a boolean if a field has been set.

### SetImageFolderNil

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetImageFolderNil(b bool)`

 SetImageFolderNil sets the value for ImageFolder to be an explicit nil

### UnsetImageFolder
`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) UnsetImageFolder()`

UnsetImageFolder ensures that no value is present for ImageFolder, not even an explicit nil
### GetRefType

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetNodeRefType

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetNodeRefType() string`

GetNodeRefType returns the NodeRefType field if non-nil, zero value otherwise.

### GetNodeRefTypeOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetNodeRefTypeOk() (*string, bool)`

GetNodeRefTypeOk returns a tuple with the NodeRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeRefType

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetNodeRefType(v string)`

SetNodeRefType sets NodeRefType field to given value.

### HasNodeRefType

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasNodeRefType() bool`

HasNodeRefType returns a boolean if a field has been set.

### SetNodeRefTypeNil

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetNodeRefTypeNil(b bool)`

 SetNodeRefTypeNil sets the value for NodeRefType to be an explicit nil

### UnsetNodeRefType
`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) UnsetNodeRefType()`

UnsetNodeRefType ensures that no value is present for NodeRefType, not even an explicit nil
### GetSubRefType

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetSubRefType() string`

GetSubRefType returns the SubRefType field if non-nil, zero value otherwise.

### GetSubRefTypeOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetSubRefTypeOk() (*string, bool)`

GetSubRefTypeOk returns a tuple with the SubRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRefType

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetSubRefType(v string)`

SetSubRefType sets SubRefType field to given value.

### HasSubRefType

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasSubRefType() bool`

HasSubRefType returns a boolean if a field has been set.

### SetSubRefTypeNil

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetSubRefTypeNil(b bool)`

 SetSubRefTypeNil sets the value for SubRefType to be an explicit nil

### UnsetSubRefType
`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) UnsetSubRefType()`

UnsetSubRefType ensures that no value is present for SubRefType, not even an explicit nil
### GetSubRefId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetSubRefId() int64`

GetSubRefId returns the SubRefId field if non-nil, zero value otherwise.

### GetSubRefIdOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetSubRefIdOk() (*int64, bool)`

GetSubRefIdOk returns a tuple with the SubRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRefId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetSubRefId(v int64)`

SetSubRefId sets SubRefId field to given value.

### HasSubRefId

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasSubRefId() bool`

HasSubRefId returns a boolean if a field has been set.

### SetSubRefIdNil

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetSubRefIdNil(b bool)`

 SetSubRefIdNil sets the value for SubRefId to be an explicit nil

### UnsetSubRefId
`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) UnsetSubRefId()`

UnsetSubRefId ensures that no value is present for SubRefId, not even an explicit nil
### GetIsPublic

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetSystemImage

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetSystemImage() bool`

GetSystemImage returns the SystemImage field if non-nil, zero value otherwise.

### GetSystemImageOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetSystemImageOk() (*bool, bool)`

GetSystemImageOk returns a tuple with the SystemImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemImage

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetSystemImage(v bool)`

SetSystemImage sets SystemImage field to given value.

### HasSystemImage

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasSystemImage() bool`

HasSystemImage returns a boolean if a field has been set.

### GetDiskIndex

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetDiskIndex() int64`

GetDiskIndex returns the DiskIndex field if non-nil, zero value otherwise.

### GetDiskIndexOk

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) GetDiskIndexOk() (*int64, bool)`

GetDiskIndexOk returns a tuple with the DiskIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIndex

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) SetDiskIndex(v int64)`

SetDiskIndex sets DiskIndex field to given value.

### HasDiskIndex

`func (o *GetVirtualImage200ResponseVirtualImageLocationsInner) HasDiskIndex() bool`

HasDiskIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


