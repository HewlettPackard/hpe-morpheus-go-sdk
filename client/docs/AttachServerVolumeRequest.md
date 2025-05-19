# AttachServerVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountPoint** | [**AttachServerVolumeRequestMountPoint**](AttachServerVolumeRequestMountPoint.md) |  | 

## Methods

### NewAttachServerVolumeRequest

`func NewAttachServerVolumeRequest(mountPoint AttachServerVolumeRequestMountPoint, ) *AttachServerVolumeRequest`

NewAttachServerVolumeRequest instantiates a new AttachServerVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachServerVolumeRequestWithDefaults

`func NewAttachServerVolumeRequestWithDefaults() *AttachServerVolumeRequest`

NewAttachServerVolumeRequestWithDefaults instantiates a new AttachServerVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountPoint

`func (o *AttachServerVolumeRequest) GetMountPoint() AttachServerVolumeRequestMountPoint`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *AttachServerVolumeRequest) GetMountPointOk() (*AttachServerVolumeRequestMountPoint, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *AttachServerVolumeRequest) SetMountPoint(v AttachServerVolumeRequestMountPoint)`

SetMountPoint sets MountPoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


