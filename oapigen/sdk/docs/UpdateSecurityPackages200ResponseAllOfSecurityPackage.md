# UpdateSecurityPackages200ResponseAllOfSecurityPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**UpdateSecurityPackages200ResponseAllOfSecurityPackageType**](UpdateSecurityPackages200ResponseAllOfSecurityPackageType.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateSecurityPackages200ResponseAllOfSecurityPackage

`func NewUpdateSecurityPackages200ResponseAllOfSecurityPackage() *UpdateSecurityPackages200ResponseAllOfSecurityPackage`

NewUpdateSecurityPackages200ResponseAllOfSecurityPackage instantiates a new UpdateSecurityPackages200ResponseAllOfSecurityPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetDescription

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetType() UpdateSecurityPackages200ResponseAllOfSecurityPackageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetTypeOk() (*UpdateSecurityPackages200ResponseAllOfSecurityPackageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) SetType(v UpdateSecurityPackages200ResponseAllOfSecurityPackageType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUuid

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateSecurityPackages200ResponseAllOfSecurityPackage) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


