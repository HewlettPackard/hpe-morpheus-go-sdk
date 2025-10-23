# GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner

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
**VirtualImages** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner

`func NewGetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner() *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner`

NewGetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner instantiates a new GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInnerWithDefaults

`func NewGetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInnerWithDefaults() *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner`

NewGetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInnerWithDefaults instantiates a new GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageBuild

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetImageBuild() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetImageBuildOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetImageBuild(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetImageBuild sets ImageBuild field to given value.

### HasImageBuild

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) HasImageBuild() bool`

HasImageBuild returns a boolean if a field has been set.

### GetBuildNumber

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetBuildNumber() int64`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetBuildNumberOk() (*int64, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetBuildNumber(v int64)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetStartDate

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetStatusMessage

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusPercent

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetStatusPercent() int64`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetStatusPercentOk() (*int64, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetStatusPercent(v int64)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### SetStatusEtaNil

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetStatusEtaNil(b bool)`

 SetStatusEtaNil sets the value for StatusEta to be an explicit nil

### UnsetStatusEta
`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) UnsetStatusEta()`

UnsetStatusEta ensures that no value is present for StatusEta, not even an explicit nil
### GetStatus

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetCreatedBy

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetCreatedBy() ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetCreatedByOk() (*ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetCreatedBy(v ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTempInstance

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetTempInstance() string`

GetTempInstance returns the TempInstance field if non-nil, zero value otherwise.

### GetTempInstanceOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetTempInstanceOk() (*string, bool)`

GetTempInstanceOk returns a tuple with the TempInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempInstance

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetTempInstance(v string)`

SetTempInstance sets TempInstance field to given value.

### HasTempInstance

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) HasTempInstance() bool`

HasTempInstance returns a boolean if a field has been set.

### SetTempInstanceNil

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetTempInstanceNil(b bool)`

 SetTempInstanceNil sets the value for TempInstance to be an explicit nil

### UnsetTempInstance
`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) UnsetTempInstance()`

UnsetTempInstance ensures that no value is present for TempInstance, not even an explicit nil
### GetVirtualImages

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetVirtualImages() []map[string]interface{}`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) GetVirtualImagesOk() (*[]map[string]interface{}, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetVirtualImages(v []map[string]interface{})`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### SetVirtualImagesNil

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) SetVirtualImagesNil(b bool)`

 SetVirtualImagesNil sets the value for VirtualImages to be an explicit nil

### UnsetVirtualImages
`func (o *GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner) UnsetVirtualImages()`

UnsetVirtualImages ensures that no value is present for VirtualImages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


