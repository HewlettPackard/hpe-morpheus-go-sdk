# GetImageBuild200ResponseImageBuildExecutionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ImageBuild** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**BuildNumber** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**StatusPercent** | Pointer to **int64** |  | [optional] 
**StatusEta** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy**](GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy.md) |  | [optional] 
**TempInstance** | Pointer to **string** |  | [optional] 
**VirtualImages** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetImageBuild200ResponseImageBuildExecutionsInner

`func NewGetImageBuild200ResponseImageBuildExecutionsInner() *GetImageBuild200ResponseImageBuildExecutionsInner`

NewGetImageBuild200ResponseImageBuildExecutionsInner instantiates a new GetImageBuild200ResponseImageBuildExecutionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImageBuild200ResponseImageBuildExecutionsInnerWithDefaults

`func NewGetImageBuild200ResponseImageBuildExecutionsInnerWithDefaults() *GetImageBuild200ResponseImageBuildExecutionsInner`

NewGetImageBuild200ResponseImageBuildExecutionsInnerWithDefaults instantiates a new GetImageBuild200ResponseImageBuildExecutionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageBuild

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetImageBuild() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetImageBuildOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) SetImageBuild(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetImageBuild sets ImageBuild field to given value.

### HasImageBuild

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) HasImageBuild() bool`

HasImageBuild returns a boolean if a field has been set.

### GetBuildNumber

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetBuildNumber() int64`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetBuildNumberOk() (*int64, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) SetBuildNumber(v int64)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetStartDate

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStatusPercent

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetStatusPercent() int64`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetStatusPercentOk() (*int64, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) SetStatusPercent(v int64)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetStatus

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetCreatedBy() GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetCreatedByOk() (*GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) SetCreatedBy(v GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTempInstance

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetTempInstance() string`

GetTempInstance returns the TempInstance field if non-nil, zero value otherwise.

### GetTempInstanceOk

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetTempInstanceOk() (*string, bool)`

GetTempInstanceOk returns a tuple with the TempInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempInstance

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) SetTempInstance(v string)`

SetTempInstance sets TempInstance field to given value.

### HasTempInstance

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) HasTempInstance() bool`

HasTempInstance returns a boolean if a field has been set.

### GetVirtualImages

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetVirtualImages() []map[string]interface{}`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) GetVirtualImagesOk() (*[]map[string]interface{}, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) SetVirtualImages(v []map[string]interface{})`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *GetImageBuild200ResponseImageBuildExecutionsInner) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


