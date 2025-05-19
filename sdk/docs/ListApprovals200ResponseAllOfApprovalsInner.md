# ListApprovals200ResponseAllOfApprovalsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ExternalName** | Pointer to **string** |  | [optional] 
**RequestType** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Approver** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**AccountIntegration** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**RequestBy** | Pointer to **string** |  | [optional] 

## Methods

### NewListApprovals200ResponseAllOfApprovalsInner

`func NewListApprovals200ResponseAllOfApprovalsInner() *ListApprovals200ResponseAllOfApprovalsInner`

NewListApprovals200ResponseAllOfApprovalsInner instantiates a new ListApprovals200ResponseAllOfApprovalsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApprovals200ResponseAllOfApprovalsInnerWithDefaults

`func NewListApprovals200ResponseAllOfApprovalsInnerWithDefaults() *ListApprovals200ResponseAllOfApprovalsInner`

NewListApprovals200ResponseAllOfApprovalsInnerWithDefaults instantiates a new ListApprovals200ResponseAllOfApprovalsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListApprovals200ResponseAllOfApprovalsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListApprovals200ResponseAllOfApprovalsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListApprovals200ResponseAllOfApprovalsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListApprovals200ResponseAllOfApprovalsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalId

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListApprovals200ResponseAllOfApprovalsInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListApprovals200ResponseAllOfApprovalsInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListApprovals200ResponseAllOfApprovalsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListApprovals200ResponseAllOfApprovalsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalName

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *ListApprovals200ResponseAllOfApprovalsInner) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *ListApprovals200ResponseAllOfApprovalsInner) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetRequestType

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *ListApprovals200ResponseAllOfApprovalsInner) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *ListApprovals200ResponseAllOfApprovalsInner) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetAccount

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListApprovals200ResponseAllOfApprovalsInner) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListApprovals200ResponseAllOfApprovalsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetApprover

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetApprover() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetApproverOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *ListApprovals200ResponseAllOfApprovalsInner) SetApprover(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *ListApprovals200ResponseAllOfApprovalsInner) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### GetAccountIntegration

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetAccountIntegration() string`

GetAccountIntegration returns the AccountIntegration field if non-nil, zero value otherwise.

### GetAccountIntegrationOk

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetAccountIntegrationOk() (*string, bool)`

GetAccountIntegrationOk returns a tuple with the AccountIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIntegration

`func (o *ListApprovals200ResponseAllOfApprovalsInner) SetAccountIntegration(v string)`

SetAccountIntegration sets AccountIntegration field to given value.

### HasAccountIntegration

`func (o *ListApprovals200ResponseAllOfApprovalsInner) HasAccountIntegration() bool`

HasAccountIntegration returns a boolean if a field has been set.

### GetStatus

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListApprovals200ResponseAllOfApprovalsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListApprovals200ResponseAllOfApprovalsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListApprovals200ResponseAllOfApprovalsInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ListApprovals200ResponseAllOfApprovalsInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListApprovals200ResponseAllOfApprovalsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListApprovals200ResponseAllOfApprovalsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListApprovals200ResponseAllOfApprovalsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListApprovals200ResponseAllOfApprovalsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRequestBy

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetRequestBy() string`

GetRequestBy returns the RequestBy field if non-nil, zero value otherwise.

### GetRequestByOk

`func (o *ListApprovals200ResponseAllOfApprovalsInner) GetRequestByOk() (*string, bool)`

GetRequestByOk returns a tuple with the RequestBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBy

`func (o *ListApprovals200ResponseAllOfApprovalsInner) SetRequestBy(v string)`

SetRequestBy sets RequestBy field to given value.

### HasRequestBy

`func (o *ListApprovals200ResponseAllOfApprovalsInner) HasRequestBy() bool`

HasRequestBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


