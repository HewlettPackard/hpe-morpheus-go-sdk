# AttachServerVolumeRequestMountPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controller** | [**AttachServerVolumeRequestMountPointController**](AttachServerVolumeRequestMountPointController.md) |  | 
**UnitNumber** | **string** | The unit number for the disk (e.g., \&quot;3\&quot;) | 

## Methods

### NewAttachServerVolumeRequestMountPoint

`func NewAttachServerVolumeRequestMountPoint(controller AttachServerVolumeRequestMountPointController, unitNumber string, ) *AttachServerVolumeRequestMountPoint`

NewAttachServerVolumeRequestMountPoint instantiates a new AttachServerVolumeRequestMountPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachServerVolumeRequestMountPointWithDefaults

`func NewAttachServerVolumeRequestMountPointWithDefaults() *AttachServerVolumeRequestMountPoint`

NewAttachServerVolumeRequestMountPointWithDefaults instantiates a new AttachServerVolumeRequestMountPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetController

`func (o *AttachServerVolumeRequestMountPoint) GetController() AttachServerVolumeRequestMountPointController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *AttachServerVolumeRequestMountPoint) GetControllerOk() (*AttachServerVolumeRequestMountPointController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *AttachServerVolumeRequestMountPoint) SetController(v AttachServerVolumeRequestMountPointController)`

SetController sets Controller field to given value.


### GetUnitNumber

`func (o *AttachServerVolumeRequestMountPoint) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *AttachServerVolumeRequestMountPoint) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *AttachServerVolumeRequestMountPoint) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


