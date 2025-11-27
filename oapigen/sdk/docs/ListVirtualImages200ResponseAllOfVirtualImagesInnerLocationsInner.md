# ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner

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

### NewListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner

`func NewListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner() *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner`

NewListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner instantiates a new ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInnerWithDefaults

`func NewListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInnerWithDefaults() *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner`

NewListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInnerWithDefaults instantiates a new ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCloud

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetCloud() map[string]interface{}`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetCloudOk() (*map[string]interface{}, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetCloud(v map[string]interface{})`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetCode

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetInternalId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalDiskId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetExternalDiskId() string`

GetExternalDiskId returns the ExternalDiskId field if non-nil, zero value otherwise.

### GetExternalDiskIdOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetExternalDiskIdOk() (*string, bool)`

GetExternalDiskIdOk returns a tuple with the ExternalDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDiskId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetExternalDiskId(v string)`

SetExternalDiskId sets ExternalDiskId field to given value.

### HasExternalDiskId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasExternalDiskId() bool`

HasExternalDiskId returns a boolean if a field has been set.

### SetExternalDiskIdNil

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetExternalDiskIdNil(b bool)`

 SetExternalDiskIdNil sets the value for ExternalDiskId to be an explicit nil

### UnsetExternalDiskId
`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) UnsetExternalDiskId()`

UnsetExternalDiskId ensures that no value is present for ExternalDiskId, not even an explicit nil
### GetRemotePath

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetRemotePath() string`

GetRemotePath returns the RemotePath field if non-nil, zero value otherwise.

### GetRemotePathOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetRemotePathOk() (*string, bool)`

GetRemotePathOk returns a tuple with the RemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePath

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetRemotePath(v string)`

SetRemotePath sets RemotePath field to given value.

### HasRemotePath

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasRemotePath() bool`

HasRemotePath returns a boolean if a field has been set.

### SetRemotePathNil

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetRemotePathNil(b bool)`

 SetRemotePathNil sets the value for RemotePath to be an explicit nil

### UnsetRemotePath
`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) UnsetRemotePath()`

UnsetRemotePath ensures that no value is present for RemotePath, not even an explicit nil
### GetImagePath

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### SetImagePathNil

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetImagePathNil(b bool)`

 SetImagePathNil sets the value for ImagePath to be an explicit nil

### UnsetImagePath
`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) UnsetImagePath()`

UnsetImagePath ensures that no value is present for ImagePath, not even an explicit nil
### GetImageName

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageRegion

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetImageRegion() string`

GetImageRegion returns the ImageRegion field if non-nil, zero value otherwise.

### GetImageRegionOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetImageRegionOk() (*string, bool)`

GetImageRegionOk returns a tuple with the ImageRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegion

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetImageRegion(v string)`

SetImageRegion sets ImageRegion field to given value.

### HasImageRegion

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasImageRegion() bool`

HasImageRegion returns a boolean if a field has been set.

### SetImageRegionNil

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetImageRegionNil(b bool)`

 SetImageRegionNil sets the value for ImageRegion to be an explicit nil

### UnsetImageRegion
`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) UnsetImageRegion()`

UnsetImageRegion ensures that no value is present for ImageRegion, not even an explicit nil
### GetImageFolder

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetImageFolder() string`

GetImageFolder returns the ImageFolder field if non-nil, zero value otherwise.

### GetImageFolderOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetImageFolderOk() (*string, bool)`

GetImageFolderOk returns a tuple with the ImageFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFolder

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetImageFolder(v string)`

SetImageFolder sets ImageFolder field to given value.

### HasImageFolder

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasImageFolder() bool`

HasImageFolder returns a boolean if a field has been set.

### SetImageFolderNil

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetImageFolderNil(b bool)`

 SetImageFolderNil sets the value for ImageFolder to be an explicit nil

### UnsetImageFolder
`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) UnsetImageFolder()`

UnsetImageFolder ensures that no value is present for ImageFolder, not even an explicit nil
### GetRefType

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetNodeRefType

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetNodeRefType() string`

GetNodeRefType returns the NodeRefType field if non-nil, zero value otherwise.

### GetNodeRefTypeOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetNodeRefTypeOk() (*string, bool)`

GetNodeRefTypeOk returns a tuple with the NodeRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeRefType

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetNodeRefType(v string)`

SetNodeRefType sets NodeRefType field to given value.

### HasNodeRefType

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasNodeRefType() bool`

HasNodeRefType returns a boolean if a field has been set.

### SetNodeRefTypeNil

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetNodeRefTypeNil(b bool)`

 SetNodeRefTypeNil sets the value for NodeRefType to be an explicit nil

### UnsetNodeRefType
`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) UnsetNodeRefType()`

UnsetNodeRefType ensures that no value is present for NodeRefType, not even an explicit nil
### GetSubRefType

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetSubRefType() string`

GetSubRefType returns the SubRefType field if non-nil, zero value otherwise.

### GetSubRefTypeOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetSubRefTypeOk() (*string, bool)`

GetSubRefTypeOk returns a tuple with the SubRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRefType

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetSubRefType(v string)`

SetSubRefType sets SubRefType field to given value.

### HasSubRefType

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasSubRefType() bool`

HasSubRefType returns a boolean if a field has been set.

### SetSubRefTypeNil

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetSubRefTypeNil(b bool)`

 SetSubRefTypeNil sets the value for SubRefType to be an explicit nil

### UnsetSubRefType
`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) UnsetSubRefType()`

UnsetSubRefType ensures that no value is present for SubRefType, not even an explicit nil
### GetSubRefId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetSubRefId() int64`

GetSubRefId returns the SubRefId field if non-nil, zero value otherwise.

### GetSubRefIdOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetSubRefIdOk() (*int64, bool)`

GetSubRefIdOk returns a tuple with the SubRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRefId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetSubRefId(v int64)`

SetSubRefId sets SubRefId field to given value.

### HasSubRefId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasSubRefId() bool`

HasSubRefId returns a boolean if a field has been set.

### SetSubRefIdNil

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetSubRefIdNil(b bool)`

 SetSubRefIdNil sets the value for SubRefId to be an explicit nil

### UnsetSubRefId
`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) UnsetSubRefId()`

UnsetSubRefId ensures that no value is present for SubRefId, not even an explicit nil
### GetIsPublic

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetSystemImage

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetSystemImage() bool`

GetSystemImage returns the SystemImage field if non-nil, zero value otherwise.

### GetSystemImageOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetSystemImageOk() (*bool, bool)`

GetSystemImageOk returns a tuple with the SystemImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemImage

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetSystemImage(v bool)`

SetSystemImage sets SystemImage field to given value.

### HasSystemImage

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasSystemImage() bool`

HasSystemImage returns a boolean if a field has been set.

### GetDiskIndex

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetDiskIndex() int64`

GetDiskIndex returns the DiskIndex field if non-nil, zero value otherwise.

### GetDiskIndexOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) GetDiskIndexOk() (*int64, bool)`

GetDiskIndexOk returns a tuple with the DiskIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIndex

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) SetDiskIndex(v int64)`

SetDiskIndex sets DiskIndex field to given value.

### HasDiskIndex

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInnerLocationsInner) HasDiskIndex() bool`

HasDiskIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


