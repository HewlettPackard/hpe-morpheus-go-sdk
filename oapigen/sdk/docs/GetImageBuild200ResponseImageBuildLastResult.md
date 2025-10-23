# GetImageBuild200ResponseImageBuildLastResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ImageBuild** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**BuildNumber** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**StatusPercent** | Pointer to **int64** |  | [optional] 
**StatusEta** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to [**ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy**](ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy.md) |  | [optional] 
**TempInstance** | Pointer to **NullableString** |  | [optional] 
**VirtualImages** | Pointer to [**[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 

## Methods

### NewGetImageBuild200ResponseImageBuildLastResult

`func NewGetImageBuild200ResponseImageBuildLastResult() *GetImageBuild200ResponseImageBuildLastResult`

NewGetImageBuild200ResponseImageBuildLastResult instantiates a new GetImageBuild200ResponseImageBuildLastResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImageBuild200ResponseImageBuildLastResultWithDefaults

`func NewGetImageBuild200ResponseImageBuildLastResultWithDefaults() *GetImageBuild200ResponseImageBuildLastResult`

NewGetImageBuild200ResponseImageBuildLastResultWithDefaults instantiates a new GetImageBuild200ResponseImageBuildLastResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetImageBuild200ResponseImageBuildLastResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageBuild

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetImageBuild() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetImageBuildOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetImageBuild(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetImageBuild sets ImageBuild field to given value.

### HasImageBuild

`func (o *GetImageBuild200ResponseImageBuildLastResult) HasImageBuild() bool`

HasImageBuild returns a boolean if a field has been set.

### GetBuildNumber

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetBuildNumber() int64`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetBuildNumberOk() (*int64, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetBuildNumber(v int64)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *GetImageBuild200ResponseImageBuildLastResult) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetStartDate

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetImageBuild200ResponseImageBuildLastResult) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetImageBuild200ResponseImageBuildLastResult) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *GetImageBuild200ResponseImageBuildLastResult) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetStatusMessage

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetImageBuild200ResponseImageBuildLastResult) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetImageBuild200ResponseImageBuildLastResult) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusPercent

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetStatusPercent() int64`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetStatusPercentOk() (*int64, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetStatusPercent(v int64)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *GetImageBuild200ResponseImageBuildLastResult) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *GetImageBuild200ResponseImageBuildLastResult) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### SetStatusEtaNil

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetStatusEtaNil(b bool)`

 SetStatusEtaNil sets the value for StatusEta to be an explicit nil

### UnsetStatusEta
`func (o *GetImageBuild200ResponseImageBuildLastResult) UnsetStatusEta()`

UnsetStatusEta ensures that no value is present for StatusEta, not even an explicit nil
### GetStatus

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetImageBuild200ResponseImageBuildLastResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetImageBuild200ResponseImageBuildLastResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *GetImageBuild200ResponseImageBuildLastResult) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetCreatedBy

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetCreatedBy() ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetCreatedByOk() (*ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetCreatedBy(v ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetImageBuild200ResponseImageBuildLastResult) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTempInstance

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetTempInstance() string`

GetTempInstance returns the TempInstance field if non-nil, zero value otherwise.

### GetTempInstanceOk

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetTempInstanceOk() (*string, bool)`

GetTempInstanceOk returns a tuple with the TempInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempInstance

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetTempInstance(v string)`

SetTempInstance sets TempInstance field to given value.

### HasTempInstance

`func (o *GetImageBuild200ResponseImageBuildLastResult) HasTempInstance() bool`

HasTempInstance returns a boolean if a field has been set.

### SetTempInstanceNil

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetTempInstanceNil(b bool)`

 SetTempInstanceNil sets the value for TempInstance to be an explicit nil

### UnsetTempInstance
`func (o *GetImageBuild200ResponseImageBuildLastResult) UnsetTempInstance()`

UnsetTempInstance ensures that no value is present for TempInstance, not even an explicit nil
### GetVirtualImages

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetVirtualImages() []GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *GetImageBuild200ResponseImageBuildLastResult) GetVirtualImagesOk() (*[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetVirtualImages(v []GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *GetImageBuild200ResponseImageBuildLastResult) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### SetVirtualImagesNil

`func (o *GetImageBuild200ResponseImageBuildLastResult) SetVirtualImagesNil(b bool)`

 SetVirtualImagesNil sets the value for VirtualImages to be an explicit nil

### UnsetVirtualImages
`func (o *GetImageBuild200ResponseImageBuildLastResult) UnsetVirtualImages()`

UnsetVirtualImages ensures that no value is present for VirtualImages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


